// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Transport code.

package http2

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/textproto"
	"sync"
	"time"

	"golang.org/x/net/http2/hpack"

	"github.com/imroc/req/v3/http2"
	"github.com/imroc/req/v3/internal/dump"
	"github.com/imroc/req/v3/internal/transport"
	reqtls "github.com/imroc/req/v3/pkg/tls"
)

const (
	// transportDefaultConnFlow is how many connection-level flow control
	// tokens we give the server at start-up, past the default 64k.
	transportDefaultConnFlow = 1 << 30

	// transportDefaultStreamFlow is how many stream-level flow
	// control tokens we announce to the peer, and how many bytes
	// we buffer per stream.
	transportDefaultStreamFlow = 4 << 20

	// initialMaxConcurrentStreams is a connections maxConcurrentStreams until
	// it's received servers initial SETTINGS frame, which corresponds with the
	// spec's minimum recommended value.
	initialMaxConcurrentStreams = 100

	// defaultMaxConcurrentStreams is a connections default maxConcurrentStreams
	// if the server doesn't include one in its initial SETTINGS frame.
	defaultMaxConcurrentStreams = 1000
)

// Transport is an HTTP/2 Transport.
//
// A Transport internally caches connections to servers. It is safe
// for concurrent use by multiple goroutines.
type Transport struct {
	*transport.Options

	// DialTLS specifies an optional dial function for creating
	// TLS connections for requests.
	//
	// If DialTLS is nil, tls.Dial is used.
	//
	// If the returned net.Conn has a ConnectionState method like tls.Conn,
	// it will be used to set http.Response.TLS.
	DialTLS func(network, addr string, cfg *tls.Config) (net.Conn, error)

	// ConnPool optionally specifies an alternate connection pool to use.
	// If nil, the default is used.
	ConnPool ClientConnPool

	// AllowHTTP, if true, permits HTTP/2 requests using the insecure,
	// plain-text "http" scheme. Note that this does not enable h2c support.
	AllowHTTP bool

	// MaxHeaderListSize is the http2 SETTINGS_MAX_HEADER_LIST_SIZE to
	// send in the initial settings frame. It is how many bytes
	// of response headers are allowed. Unlike the http2 spec, zero here
	// means to use a default limit (currently 10MB). If you actually
	// want to advertise an unlimited value to the peer, Transport
	// interprets the highest possible value here (0xffffffff or 1<<32-1)
	// to mean no limit.
	MaxHeaderListSize uint32

	// StrictMaxConcurrentStreams controls whether the server's
	// SETTINGS_MAX_CONCURRENT_STREAMS should be respected
	// globally. If false, new TCP connections are created to the
	// server as needed to keep each under the per-connection
	// SETTINGS_MAX_CONCURRENT_STREAMS limit. If true, the
	// server's SETTINGS_MAX_CONCURRENT_STREAMS is interpreted as
	// a global limit and callers of RoundTrip block when needed,
	// waiting for their turn.
	StrictMaxConcurrentStreams bool

	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout time.Duration

	// ReadIdleTimeout is the timeout after which a health check using ping
	// frame will be carried out if no frame is received on the connection.
	// Note that a ping response will is considered a received frame, so if
	// there is no other traffic on the connection, the health check will
	// be performed every ReadIdleTimeout interval.
	// If zero, no health check is performed.
	ReadIdleTimeout time.Duration

	// PingTimeout is the timeout after which the connection will be closed
	// if a response to Ping is not received.
	// Defaults to 15s.
	PingTimeout time.Duration

	// WriteByteTimeout is the timeout after which the connection will be
	// closed no data can be written to it. The timeout begins when data is
	// available to write, and is extended whenever any bytes are written.
	WriteByteTimeout time.Duration

	// CountError, if non-nil, is called on HTTP/2 transport errors.
	// It's intended to increment a metric for monitoring, such
	// as an expvar or Prometheus metric.
	// The errType consists of only ASCII word characters.
	CountError func(errType string)

	Settings []http2.Setting

	ConnectionFlow uint32
	HeaderPriority http2.PriorityParam
	PriorityFrames []http2.PriorityFrame

	connPoolOnce  sync.Once
	connPoolOrDef ClientConnPool // non-nil version of ConnPool
}

// newTimer creates a new time.Timer, or a synthetic timer in tests.
func (t *Transport) newTimer(d time.Duration) timer { _ = "STUB: not implemented"; return *new(timer) }

// afterFunc creates a new time.AfterFunc timer, or a synthetic timer in tests.
func (t *Transport) afterFunc(d time.Duration, f func()) timer {
	_ = "STUB: not implemented"
	return *new(timer)
}

func (t *Transport) contextWithTimeout(ctx context.Context, d time.Duration) (context.Context, context.CancelFunc) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc)
}

func (t *Transport) maxHeaderListSize() uint32 { _ = "STUB: not implemented"; return 0 }

func (t *Transport) pingTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (t *Transport) connPool() ClientConnPool {
	_ = "STUB: not implemented"
	return *new(ClientConnPool)
}

func (t *Transport) initConnPool() { _ = "STUB: not implemented"; return }

// ClientConn is the state of a single HTTP/2 client connection to an
// HTTP/2 server.
type ClientConn struct {
	t             *Transport
	tconn         net.Conn             // usually TLSConn, except specialized impls
	tlsState      *tls.ConnectionState // nil only for specialized impls
	reused        uint32               // whether conn is being reused; atomic
	singleUse     bool                 // whether being used for a single http.Request
	getConnCalled bool                 // used by clientConnPool

	// readLoop goroutine fields:
	readerDone chan struct{} // closed on error
	readerErr  error         // set before readerDone is closed

	idleTimeout time.Duration // or 0 for never
	idleTimer   timer

	mu              sync.Mutex // guards following
	cond            *sync.Cond // hold mu; broadcast on flow/closed changes
	flow            outflow    // our conn-level flow control quota (cs.outflow is per stream)
	inflow          inflow     // peer's conn-level flow control
	doNotReuse      bool       // whether conn is marked to not be reused for any future requests
	closing         bool
	closed          bool
	seenSettings    bool                     // true if we've seen a settings frame, false otherwise
	wantSettingsAck bool                     // we sent a SETTINGS frame and haven't heard back
	goAway          *GoAwayFrame             // if non-nil, the GoAwayFrame we received
	goAwayDebug     string                   // goAway frame's debug data, retained as a string
	streams         map[uint32]*clientStream // client-initiated
	streamsReserved int                      // incr by ReserveNewRequest; decr on RoundTrip
	nextStreamID    uint32
	pendingRequests int                       // requests blocked and waiting to be sent because len(streams) == maxConcurrentStreams
	pings           map[[8]byte]chan struct{} // in flight ping data to notification channel
	br              *bufio.Reader
	lastActive      time.Time
	lastIdle        time.Time // time last idle
	// Settings from peer: (also guarded by wmu)
	maxFrameSize          uint32
	maxConcurrentStreams  uint32
	peerMaxHeaderListSize uint64
	initialWindowSize     uint32

	// reqHeaderMu is a 1-element semaphore channel controlling access to sending new requests.
	// Write to reqHeaderMu to lock it, read from it to unlock.
	// Lock reqmu BEFORE mu or wmu.
	reqHeaderMu chan struct{}

	// wmu is held while writing.
	// Acquire BEFORE mu when holding both, to avoid blocking mu on network writes.
	// Only acquire both at the same time when changing peer settings.
	wmu  sync.Mutex
	bw   *bufio.Writer
	fr   *Framer
	werr error        // first write error that has occurred
	hbuf bytes.Buffer // HPACK encoder writes into this
	henc *hpack.Encoder
}

// clientStream is the state for a single HTTP/2 stream. One of these
// is created for each Transport.RoundTrip call.
type clientStream struct {
	currentRequest *http.Request
	cc             *ClientConn

	// Fields of Request that we may access even after the response body is closed.
	ctx       context.Context
	reqCancel <-chan struct{}

	trace         *httptrace.ClientTrace // or nil
	ID            uint32
	bufPipe       pipe // buffered pipe with the flow-controlled response payload
	requestedGzip bool
	isHead        bool

	abortOnce sync.Once
	abort     chan struct{} // closed to signal stream should end immediately
	abortErr  error         // set if abort is closed

	peerClosed chan struct{} // closed when the peer sends an END_STREAM flag
	donec      chan struct{} // closed after the stream is in the closed state
	on100      chan struct{} // buffered; written to if a 100 is received

	respHeaderRecv chan struct{}  // closed when headers are received
	res            *http.Response // set if respHeaderRecv is closed

	flow        outflow // guarded by cc.mu
	inflow      inflow  // guarded by cc.mu
	bytesRemain int64   // -1 means unknown; owned by transportResponseBody.Read
	readErr     error   // sticky read error; owned by transportResponseBody.Read

	reqBody              io.ReadCloser
	reqBodyContentLength int64         // -1 means unknown
	reqBodyClosed        chan struct{} // guarded by cc.mu; non-nil on Close, closed when done

	// owned by writeRequest:
	sentEndStream bool // sent an END_STREAM flag to the peer
	sentHeaders   bool

	// owned by clientConnReadLoop:
	firstByte    bool  // got the first response byte
	pastHeaders  bool  // got first MetaHeadersFrame (actual headers)
	pastTrailers bool  // got optional second MetaHeadersFrame (trailers)
	num1xx       uint8 // number of 1xx responses seen
	readClosed   bool  // peer sent an END_STREAM flag
	readAborted  bool  // read loop reset the stream

	trailer    http.Header  // accumulated trailers
	resTrailer *http.Header // client's Response.Trailer
}

var got1xxFuncForTests func(int, textproto.MIMEHeader) error

// get1xxTraceFunc returns the value of request's httptrace.ClientTrace.Got1xxResponse func,
// if any. It returns nil if not set or if the Go version is too old.
func (cs *clientStream) get1xxTraceFunc() func(int, textproto.MIMEHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func (cs *clientStream) abortStream(err error) { _ = "STUB: not implemented"; return }

func (cs *clientStream) abortStreamLocked(err error) { _ = "STUB: not implemented"; return }

// TODO(dneil): Clean up tests where cs.cc.cond is nil.

// Wake up writeRequestBody if it is waiting on flow control.

func (cs *clientStream) abortRequestBodyWrite() { _ = "STUB: not implemented"; return }

func (cs *clientStream) closeReqBodyLocked() { _ = "STUB: not implemented"; return }

type stickyErrWriter struct {
	conn    net.Conn
	timeout time.Duration
	err     *error
}

func (sew stickyErrWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Keep extending the deadline so long as we're making progress.

// noCachedConnError is the concrete type of ErrNoCachedConn, which
// needs to be detected by net/http regardless of whether it's its
// bundled version (in h2_bundle.go with a rewritten type name) or
// from a user's x/net/http2. As such, as it has a unique method name
// (IsHTTP2NoCachedConnError) that net/http sniffs for via func
// IsNoCachedConnError.
type noCachedConnError struct{}

func (noCachedConnError) IsHTTP2NoCachedConnError() { _ = "STUB: not implemented"; return }

func (noCachedConnError) Error() string { _ = "STUB: not implemented"; return "" }

// IsNoCachedConnError reports whether err is of type noCachedConnError
// or its equivalent renamed type in net/http2's h2_bundle.go. Both types
// may coexist in the same running program.
func IsNoCachedConnError(err error) bool { _ = "STUB: not implemented"; return false }

var ErrNoCachedConn error = noCachedConnError{}

// RoundTripOpt are options for the Transport.RoundTripOpt method.
type RoundTripOpt struct {
	// OnlyCachedConn controls whether RoundTripOpt may
	// create a new TCP connection. If set true and
	// no cached connection is available, RoundTripOpt
	// will return ErrNoCachedConn.
	OnlyCachedConn bool
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) RoundTripOnlyCachedConn(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// authorityAddr returns a given authority (a host/IP, or host:port / ip:port)
// and returns a host:port. The port 443 is added if needed.
func authorityAddr(scheme string, authority string) (addr string) {
	_ = "STUB: not implemented"
	return ""
}

// authority didn't have a port

// authority's port was empty

// IPv6 address literal, without a port:

func (t *Transport) AddConn(conn net.Conn, addr string) (used bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// RoundTripOpt is like RoundTrip, but takes options.
func (t *Transport) RoundTripOpt(req *http.Request, opt RoundTripOpt) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// After the first retry, do exponential backoff with 10% jitter.

// CloseIdleConnections closes any connections which were previously
// connected from previous requests but are now sitting idle.
// It does not interrupt any connections currently in use.
func (t *Transport) CloseIdleConnections() { _ = "STUB: not implemented"; return }

var (
	errClientConnClosed    = errors.New("http2: client conn is closed")
	errClientConnUnusable  = errors.New("http2: client conn not usable")
	errClientConnGotGoAway = errors.New("http2: Transport received Server's graceful shutdown GOAWAY")
)

// shouldRetryRequest is called by RoundTrip when a request fails to get
// response headers. It is always called with a non-nil error.
// It returns either a request to retry (either the same request, or a
// modified clone), or an error if the request can't be replayed.
func shouldRetryRequest(req *http.Request, err error) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If the Body is nil (or http.NoBody), it's safe to reuse
// this request and its Body.

// If the request body can be reset back to its original
// state via the optional req.GetBody, do that.

// The Request.Body can't reset back to the beginning, but we
// don't seem to have started to read from it yet, so reuse
// the request directly.

func canRetryError(err error) bool { _ = "STUB: not implemented"; return false }

// See golang/go#47635, golang/go#42777

func (t *Transport) dialClientConn(ctx context.Context, addr string, singleUse bool) (*ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) newTLSConfig(host string) *tls.Config { _ = "STUB: not implemented"; return nil }

var zeroDialer net.Dialer

type tlsHandshakeTimeoutError struct{}

func (tlsHandshakeTimeoutError) Timeout() bool   { _ = "STUB: not implemented"; return false }
func (tlsHandshakeTimeoutError) Temporary() bool { _ = "STUB: not implemented"; return false }
func (tlsHandshakeTimeoutError) Error() string   { _ = "STUB: not implemented"; return "" }

// dialTLSWithContext uses tls.Dialer, added in Go 1.15, to open a TLS
// connection.
func (t *Transport) dialTLSWithContext(ctx context.Context, network, addr string, cfg *tls.Config) (reqtls.Conn, error) {
	_ = "STUB: not implemented"
	return *new(reqtls.Conn), nil
}

// for canceling TLS handshake

func (t *Transport) dialTLS(ctx context.Context) func(string, string, *tls.Config) (net.Conn, error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) NewClientConn(c net.Conn) (*ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) newClientConn(c net.Conn, singleUse bool) (*ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// spec default
// spec default
// "infinite", per spec. Use a smaller value until we have received server settings.
// "infinite", per spec. Use 2^64-1 instead.

// TODO: adjust this writer size to account for frame size +
// MTU + crypto/tls record padding.

// TODO: SetMaxDynamicTableSize, SetMaxDynamicTableSizeLimit on
// henc in response to SETTINGS frames?

// Start the idle timer after the connection is fully initialized.

func (cc *ClientConn) healthCheck() { _ = "STUB: not implemented"; return }

// We don't need to periodically ping in the health check, because the readLoop of ClientConn will
// trigger the healthCheck again if there is no frame received.

// SetDoNotReuse marks cc as not reusable for future HTTP requests.
func (cc *ClientConn) SetDoNotReuse() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) setGoAway(f *GoAwayFrame) { _ = "STUB: not implemented"; return }

// Merge the previous and current GoAway error frames.

// The server's GOAWAY indicates that it received this stream.
// It will either finish processing it, or close the connection
// without doing so. Either way, leave the stream alone for now.

// Don't retry the first stream on a connection if we get a non-NO error.
// If the server is sending an error on a new connection,
// retrying the request on a new one probably isn't going to work.

// Aborting the stream with errClentConnGotGoAway indicates that
// the request should be retried on a new connection.

// CanTakeNewRequest reports whether the connection can take a new request,
// meaning it has not been closed or received or sent a GOAWAY.
//
// If the caller is going to immediately make a new request on this
// connection, use ReserveNewRequest instead.
func (cc *ClientConn) CanTakeNewRequest() bool { _ = "STUB: not implemented"; return false }

// ReserveNewRequest is like CanTakeNewRequest but also reserves a
// concurrent stream in cc. The reservation is decremented on the
// next call to RoundTrip.
func (cc *ClientConn) ReserveNewRequest() bool { _ = "STUB: not implemented"; return false }

// ClientConnState describes the state of a ClientConn.
type ClientConnState struct {
	// Closed is whether the connection is closed.
	Closed bool

	// Closing is whether the connection is in the process of
	// closing. It may be closing due to shutdown, being a
	// single-use connection, being marked as DoNotReuse, or
	// having received a GOAWAY frame.
	Closing bool

	// StreamsActive is how many streams are active.
	StreamsActive int

	// StreamsReserved is how many streams have been reserved via
	// ClientConn.ReserveNewRequest.
	StreamsReserved int

	// StreamsPending is how many requests have been sent in excess
	// of the peer's advertised MaxConcurrentStreams setting and
	// are waiting for other streams to complete.
	StreamsPending int

	// MaxConcurrentStreams is how many concurrent streams the
	// peer advertised as acceptable. Zero means no SETTINGS
	// frame has been received yet.
	MaxConcurrentStreams uint32

	// LastIdle, if non-zero, is when the connection last
	// transitioned to idle state.
	LastIdle time.Time
}

// clientConnIdleState describes the suitability of a client
// connection to initiate a new RoundTrip request.
type clientConnIdleState struct {
	canTakeNewRequest bool
}

func (cc *ClientConn) idleState() clientConnIdleState {
	_ = "STUB: not implemented"
	return *new(clientConnIdleState)
}

func (cc *ClientConn) idleStateLocked() (st clientConnIdleState) {
	_ = "STUB: not implemented"
	return *new(clientConnIdleState)
}

// We'll tell the caller we can take a new request to
// prevent the caller from dialing a new TCP
// connection, but then we'll block later before
// writing it.

func (cc *ClientConn) canTakeNewRequestLocked() bool { _ = "STUB: not implemented"; return false }

// tooIdleLocked reports whether this connection has been sitting idle
// for too much wall time.
func (cc *ClientConn) tooIdleLocked() bool {
	_ = "STUB: not implemented"
	// The Round(0) strips the monotonic clock reading so the
	// times are compared based on their wall time. We don't want
	// to reuse a connection that's been sitting idle during
	// VM/laptop suspend if monotonic time was also frozen.
	return false
}

// onIdleTimeout is called from a time.AfterFunc goroutine. It will
// only be called when we're idle, but because we're coming from a new
// goroutine, there could be a new request coming in at the same time,
// so this simply calls the synchronized closeIfIdle to shut down this
// connection. The timer could just call closeIfIdle, but this is more
// clear.
func (cc *ClientConn) onIdleTimeout() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) closeConn() { _ = "STUB: not implemented"; return }

// A tls.Conn.Close can hang for a long time if the peer is unresponsive.
// Try to shut it down more aggressively.
func (cc *ClientConn) forceCloseConn() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) closeIfIdle() { _ = "STUB: not implemented"; return }

// TODO: do clients send GOAWAY too? maybe? Just Close:

func (cc *ClientConn) isDoNotReuseAndIdle() bool { _ = "STUB: not implemented"; return false }

var shutdownEnterWaitStateHook = func() {}

// Shutdown gracefully closes the client connection, waiting for running streams to complete.
func (cc *ClientConn) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Wait for all in-flight streams to complete or connection to close

// guarded by cc.mu

// Free the goroutine above

func (cc *ClientConn) sendGoAway() error { _ = "STUB: not implemented"; return nil }

// GOAWAY sent already

// Send a graceful shutdown frame to server

// Prevent new requests

// closes the client connection immediately. In-flight requests are interrupted.
// err is sent to streams.
func (cc *ClientConn) closeForError(err error) { _ = "STUB: not implemented"; return }

// Close closes the client connection immediately.
//
// In-flight requests are interrupted. For a graceful shutdown, use Shutdown instead.
func (cc *ClientConn) Close() error { _ = "STUB: not implemented"; return nil }

// closes the client connection immediately. In-flight requests are interrupted.
func (cc *ClientConn) closeForLostPing() { _ = "STUB: not implemented"; return }

func commaSeparatedTrailers(req *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (cc *ClientConn) responseHeaderTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// checkConnHeaders checks whether req has any invalid connection-level headers.
// per RFC 7540 section 8.1.2.2: Connection-Specific Header Fields.
// Certain headers are special-cased as okay but not transmitted later.
func checkConnHeaders(req *http.Request) error { _ = "STUB: not implemented"; return nil }

// actualContentLength returns a sanitized version of
// req.ContentLength, where 0 actually means zero (not unknown) and -1
// means unknown.
func actualContentLength(req *http.Request) int64 { _ = "STUB: not implemented"; return 0 }

func (cc *ClientConn) decrStreamReservations() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) decrStreamReservationsLocked() { _ = "STUB: not implemented"; return }

func (cc *ClientConn) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cc *ClientConn) roundTrip(req *http.Request, streamf func(*clientStream)) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(bradfitz): this is a copy of the logic in net/http. Unify somewhere?

// Request gzip only, not deflate. Deflate is ambiguous and
// not as universally supported anyway.
// See: https://zlib.net/zlib_faq.html#faq39
//
// Note that we don't request this for HEAD requests,
// due to a bug in nginx:
//   http://trac.nginx.org/nginx/ticket/358
//   https://golang.org/issue/5522
//
// We don't request gzip if the request is for a range, since
// auto-decoding a portion of a gzipped document will just fail
// anyway. See https://golang.org/issue/8923

// On error or status code 3xx, 4xx, 5xx, etc abort any
// ongoing write, assuming that the server doesn't care
// about our request body. If the server replied with 1xx or
// 2xx, however, then assume the server DOES potentially
// want our body (e.g. full-duplex streaming:
// golang.org/issue/13444). If it turns out the server
// doesn't, they'll RST_STREAM us soon enough. This is a
// heuristic to avoid adding knobs to Transport. Hopefully
// we can keep it.

// If there isn't a request or response body still being
// written, then wait for the stream to be closed before
// RoundTrip returns.

// Wait for the request body to be closed.
//
// If nothing closed the body before now, abortStreamLocked
// will have started a goroutine to close it.
//
// Closing the body before returning avoids a race condition
// with net/http checking its readTrackingBody to see if the
// body was read from or closed. See golang/go#60041.
//
// The body is closed in a separate goroutine without the
// connection mutex held, but dropping the mutex before waiting
// will keep us from holding it indefinitely if the body
// close is slow for some reason.

// If both cs.respHeaderRecv and cs.abort are signaling,
// pick respHeaderRecv. The server probably wrote the
// response and immediately reset the stream.
// golang.org/issue/49645

// doRequest runs for the duration of the request lifetime.
//
// It sends the request and performs post-request cleanup (closing Request.Body, etc.).
func (cs *clientStream) doRequest(req *http.Request, streamf func(*clientStream)) {
	_ = "STUB: not implemented"
	return
}

// writeRequest sends a request.
//
// It returns nil after the request is written, the response read,
// and the request stream is half-closed by the peer.
//
// It returns non-nil if the request ends otherwise.
// If the returned error is StreamError, the error Code may be used in resetting the stream.
func (cs *clientStream) writeRequest(req *http.Request, streamf func(*clientStream)) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Acquire the new-request lock by writing to reqHeaderMu.
// This lock guards the critical section covering allocating a new stream ID
// (requires mu) and creating the stream (requires wmu).

// for tests

// assigns stream ID

// Past this point (where we send request headers), it is possible for
// RoundTrip to return successfully. Since the RoundTrip contract permits
// the caller to "mutate or reuse" the Request after closing the Response's Body,
// we must take care when referencing the Request from here on.

// Wait until the peer half-closes its end of the stream,
// or until the request is aborted (via context, error, or otherwise),
// whichever comes first.

// keep waiting for END_STREAM

func (cs *clientStream) encodeAndWriteHeaders(req *http.Request, dumps []*dump.Dumper) error {
	_ = "STUB: not implemented"
	return nil
}

// If the request was canceled while waiting for cc.mu, just quit.

// Encode headers.
//
// we send: HEADERS{1}, CONTINUATION{0,} + DATA{0,} (DATA is
// sent by writeRequestBody below, along with any Trailers,
// again in form HEADERS{1}, CONTINUATION{0,})

// Write the request.

// cleanupWriteRequest performs post-request tasks.
//
// If err (the result of writeRequest) is non-nil and the stream is not closed,
// cleanupWriteRequest will send a reset to the peer.
func (cs *clientStream) cleanupWriteRequest(err error) { _ = "STUB: not implemented"; return }

// We were canceled before creating the stream, so return our reservation.

// TODO: write h12Compare test showing whether
// Request.Body is closed by the Transport,
// and in multiple cases: server replies <=299 and >299
// while still writing request body

// If the connection is closed immediately after the response is read,
// we may be aborted before finishing up here. If the stream was closed
// cleanly on both sides, there is no error.

// possibly redundant, but harmless

// no-op if already closed

// awaitOpenSlotForStreamLocked waits until len(streams) < maxConcurrentStreams.
// Must hold cc.mu.
func (cc *ClientConn) awaitOpenSlotForStreamLocked(cs *clientStream) error {
	_ = "STUB: not implemented"
	return nil
}

// requires cc.wmu be held
func (cc *ClientConn) writeHeaders(streamID uint32, endStream bool, maxFrameSize int, hdrs []byte) error {
	_ = "STUB: not implemented"
	// first frame written (HEADERS is first, then CONTINUATION)
	return nil
}

// internal error values; they don't escape to callers
var (
	// abort request body write; don't send cancel
	errStopReqBodyWrite = errors.New("http2: aborting request body write")

	// abort request body write, but send stream reset of cancel.
	errStopReqBodyWriteAndCancel = errors.New("http2: canceling request")

	errReqBodyTooLong = errors.New("http2: request body larger than specified content length")
)

// frameScratchBufferLen returns the length of a buffer to use for
// outgoing request bodies to read/write to/from.
//
// It returns max(1, min(peer's advertised max frame size,
// Request.ContentLength+1, 512KB)).
func (cs *clientStream) frameScratchBufferLen(maxFrameSize int) int {
	_ = "STUB: not implemented"
	return 0
}

// Add an extra byte past the declared content-length to
// give the caller's Request.Body io.textprotoReader a chance to
// give us more bytes than they declared, so we can catch it
// early.

// doesn't truncate; max is 512K

// Seven bufPools manage different frame sizes. This helps to avoid scenarios where long-running
// streaming requests using small frame sizes occupy large buffers initially allocated for prior
// requests needing big buffers. The size ranges are as follows:
// {0 KB, 16 KB], {16 KB, 32 KB], {32 KB, 64 KB], {64 KB, 128 KB], {128 KB, 256 KB],
// {256 KB, 512 KB], {512 KB, infinity}
// In practice, the maximum scratch buffer size should not exceed 512 KB due to
// frameScratchBufferLen(maxFrameSize), thus the "infinity pool" should never be used.
// It exists mainly as a safety measure, for potential future increases in max buffer size.
var bufPools [7]sync.Pool       // of *[]byte
func bufPoolIndex(size int) int { _ = "STUB: not implemented"; return 0 }

func (cs *clientStream) writeRequestBody(req *http.Request, dumps []*dump.Dumper) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// whether we sent the final DATA frame w/ END_STREAM

// Scratch buffer for reading into & writing from.

// The request body's Content-Length was predeclared and
// we just finished reading it all, but the underlying io.textprotoReader
// returned the final chunk with a nil error (which is one of
// the two valid things a textprotoReader can do at EOF). Because we'd prefer
// to send the END_STREAM bit early, double-check that we're actually
// at EOF. Subsequent reads should return (0, EOF) at this point.
// If either value is different, we return an error in one of two ways below.

// TODO(bradfitz): this flush is for latency, not bandwidth.
// Most requests won't need this. Make this opt-in or
// opt-out?  Use some heuristic on the body type? Nagel-like
// timers?  Based on 'n'? Only last chunk of this for loop,
// unless flow control tokens are low? For now, always.
// If we change this, see comment below.

// Already sent END_STREAM (which implies we have no
// trailers) and flushed, because currently all
// WriteData frames above get a flush. So we're done.

// Since the RoundTrip contract permits the caller to "mutate or reuse"
// a request after the Response's Body is closed, verify that this hasn't
// happened before accessing the trailers.

// Two ways to send END_STREAM: either with trailers, or
// with an empty DATA frame.

// awaitFlowControl waits for [1, min(maxBytes, cc.cs.maxFrameSize)] flow
// control tokens from the server.
// It returns either the non-zero number of tokens taken or an error
// if the stream is dead.
func (cs *clientStream) awaitFlowControl(maxBytes int) (taken int32, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// can't truncate int; take is int32

func validateHeaders(hdrs http.Header) string { _ = "STUB: not implemented"; return "" }

// Don't include the value in the error,
// because it may be sensitive.

var errNilRequestURL = errors.New("http2: Request.URI is nil")

// requires cc.wmu be held.
func (cc *ClientConn) encodeHeaders(req *http.Request, addGzipHeader bool, trailers string, contentLength int64, dumps []*dump.Dumper) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check for any invalid headers+trailers and return an error before we
// potentially pollute our hpack state. (We want to be able to
// continue to reuse the hpack encoder for future requests)

// 8.1.2.3 Request Pseudo-Header Fields
// The :path pseudo-header field includes the path and query parts of the
// target URI (the path-absolute production and optionally a '?' character
// followed by the query production, see Sections 3.3 and 3.4 of
// [RFC3986]).

// Match Go's http1 behavior: at most one
// User-Agent. If set to nil or empty string,
// then omit it. Otherwise if not mentioned,
// include the default (below).

// Per 8.1.2.5 To allow for better compression efficiency, the
// Cookie header field MAY be split into separate header fields,
// each with one or more cookie-pairs.

// writeHeader("cookie", v[:p])

// strip space after semicolon if any.

// writeHeader("cookie", v)

// Do a first pass over the headers counting bytes to ensure
// we don't exceed cc.peerMaxHeaderListSize. This is done as a
// separate pass before encoding the headers to prevent
// modifying the hpack state.

// Header list size is ok. Write the headers.

// Skip writing invalid headers. Per RFC 7540, Section 8.1.2, header
// field names have to be ASCII characters (just as in HTTP/1.x).

// shouldSendReqContentLength reports whether the http2.Transport should send
// a "content-length" request header. This logic is basically a copy of the net/http
// transferWriter.shouldSendContentLength.
// The contentLength is the corrected contentLength (so 0 means actually 0, not unknown).
// -1 means unknown.
func shouldSendReqContentLength(method string, contentLength int64) bool {
	_ = "STUB: not implemented"
	return false
}

// For zero bodies, whether we send a content-length depends on the method.
// It also kinda doesn't matter for http2 either way, with END_STREAM.

// requires cc.wmu be held.
func (cc *ClientConn) encodeTrailers(trailer http.Header, dumps []*dump.Dumper) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip writing invalid headers. Per RFC 7540, Section 8.1.2, header
// field names have to be ASCII characters (just as in HTTP/1.x).

// Transfer-Encoding, etc.. have already been filtered at the
// start of RoundTrip

func (cc *ClientConn) writeHeader(name, value string) { _ = "STUB: not implemented"; return }

type resAndError struct {
	_   incomparable
	res *http.Response
	err error
}

// requires cc.mu be held.
func (cc *ClientConn) addStreamLocked(cs *clientStream) { _ = "STUB: not implemented"; return }

func (cc *ClientConn) forgetStreamID(id uint32) { _ = "STUB: not implemented"; return }

// Wake up writeRequestBody via clientStream.awaitFlowControl and
// wake up RoundTrip if there is a pending request.

// clientConnReadLoop is the state owned by the clientConn's frame-reading readLoop.
type clientConnReadLoop struct {
	_  incomparable
	cc *ClientConn
}

// readLoop runs in its own goroutine and reads and dispatches frames.
func (cc *ClientConn) readLoop() { _ = "STUB: not implemented"; return }

// GoAwayError is returned by the Transport when the server closes the
// TCP connection after sending a GOAWAY frame.
type GoAwayError struct {
	LastStreamID uint32
	ErrCode      ErrCode
	DebugData    string
}

func (e GoAwayError) Error() string { _ = "STUB: not implemented"; return "" }

func isEOFOrNetReadError(err error) bool { _ = "STUB: not implemented"; return false }

func (rl *clientConnReadLoop) cleanup() { _ = "STUB: not implemented"; return }

// Close any response bodies if the server closes prematurely.
// TODO: also do this if we've written the headers but not
// gotten a response yet.

// The server closed the stream before closing the conn,
// so no need to interrupt it.

// countReadFrameError calls Transport.CountError with a string
// representing err.
func (cc *ClientConn) countReadFrameError(err error) { _ = "STUB: not implemented"; return }

func (rl *clientConnReadLoop) run() error { _ = "STUB: not implemented"; return nil }

func (rl *clientConnReadLoop) processHeaders(f *MetaHeadersFrame) error {
	_ = "STUB: not implemented"
	return nil
}

// We'd get here if we canceled a request while the
// server had its response still in flight. So if this
// was just something we canceled, ignore it.

// TODO(bradfitz): move first response byte earlier,
// when we first read the 9 byte header, not waiting
// until all the HEADERS+CONTINUATION frames have been
// merged. This works for now.

// Any other error type is a stream error.

// return nil from process* funcs to keep conn alive

// (nil, nil) special case. See handleResponse docs.

// foreachHeaderElement splits v according to the "#rule" construction
// in RFC 7230 section 7 and calls fn for each non-empty element.
func foreachHeaderElement(v string, fn func(string)) { _ = "STUB: not implemented"; return }

// may return error types nil, or ConnectionError. Any other error value
// is a StreamError of type ErrCodeProtocol. The returned error in that case
// is the detail.
//
// As a special case, handleResponse may return (nil, nil) to skip the
// frame (currently only used for 1xx responses).
func (rl *clientConnReadLoop) handleResponse(cs *clientStream, f *MetaHeadersFrame) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// More than likely this will be a single-element key.
// Most headers aren't multi-valued.
// Set the capacity on strs[0] to 1, so any future append
// won't extend the slice into the other strings.

// arbitrary bound on number of informational responses, same as net/http

// do it all again

// TODO: care? unlike http/1, it won't mess up our framing, so it's
// more safe smuggling-wise to ignore.

// TODO: care? unlike http/1, it won't mess up our framing, so it's
// more safe smuggling-wise to ignore.

func (rl *clientConnReadLoop) processTrailers(cs *clientStream, f *MetaHeadersFrame) error {
	_ = "STUB: not implemented"
	return nil

	// Too many HEADERS frames for this stream.
}

// We expect that any headers for trailers also
// has END_STREAM.

// No pseudo header fields are defined for trailers.
// TODO: ConnectionError might be overly harsh? Check.

// transportResponseBody is the concrete type of Transport.RoundTrip's
// Response.Body. It is an io.ReadCloser.
type transportResponseBody struct {
	cs *clientStream
}

func (b transportResponseBody) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// No flow control tokens to send back.

// No need to refresh if the stream is over or failed.

var errClosedResponseBody = errors.New("http2: response body closed")

func (b transportResponseBody) Close() error { _ = "STUB: not implemented"; return nil }

// Return connection-level flow control.

// TODO(dneil): Acquiring this mutex can block indefinitely.
// Move flow control return to a goroutine?

// Return connection-level flow control.

// See golang/go#49366: The net/http package can cancel the
// request context after the response body is fully read.
// Don't treat this as an error.

func (rl *clientConnReadLoop) processData(f *DataFrame) error {
	_ = "STUB: not implemented"
	return nil
}

// We never asked for this.

// We probably did ask for this, but canceled. Just ignore it.
// TODO: be stricter here? only silently ignore things which
// we canceled, but not things which were closed normally
// by the peer? Tough without accumulating too much state.

// But at least return their flow control:

// Check connection-level flow control.

// Return any padded flow control now, since we won't
// refund it later on body reads.

// Return len(data) now if the stream is already closed,
// since data will never be read.

func (rl *clientConnReadLoop) endStream(cs *clientStream) {
	_ = "STUB: not implemented"
	// TODO: check that any declared content-length matches, like
	// server.go's (*stream).endStream method.
	return
}

// Close cs.bufPipe and cs.peerClosed with cc.mu held to avoid a
// race condition: The caller can read io.EOF from Response.Body
// and close the body before we close cs.peerClosed, causing
// cleanupWriteRequest to send a RST_STREAM.

func (rl *clientConnReadLoop) endStreamError(cs *clientStream, err error) {
	_ = "STUB: not implemented"
	return
}

func (rl *clientConnReadLoop) streamByID(id uint32) *clientStream {
	_ = "STUB: not implemented"
	return nil
}

func (cs *clientStream) copyTrailers() { _ = "STUB: not implemented"; return }

func (rl *clientConnReadLoop) processGoAway(f *GoAwayFrame) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: deal with GOAWAY more. particularly the error code

func (rl *clientConnReadLoop) processSettings(f *SettingsFrame) error {
	_ = "STUB: not implemented"

	// Locking both mu and wmu here allows frame encoding to read settings with only wmu held.
	// Acquiring wmu when f.IsAck() is unnecessary, but convenient and mostly harmless.
	return nil
}

func (rl *clientConnReadLoop) processSettingsNoWrite(f *SettingsFrame) error {
	_ = "STUB: not implemented"
	return nil
}

// Values above the maximum flow-control
// window size of 2^31-1 MUST be treated as a
// connection error (Section 5.4.1) of type
// FLOW_CONTROL_ERROR.

// Adjust flow control of currently-open
// frames by the difference of the old initial
// window size and this one.

// TODO(bradfitz): handle more settings? SETTINGS_HEADER_TABLE_SIZE probably.

// This was the servers initial SETTINGS frame and it
// didn't contain a MAX_CONCURRENT_STREAMS field so
// increase the number of concurrent streams this
// connection can establish to our default.

func (rl *clientConnReadLoop) processWindowUpdate(f *WindowUpdateFrame) error {
	_ = "STUB: not implemented"
	return nil
}

// For stream, the sender sends RST_STREAM with an error code of FLOW_CONTROL_ERROR

func (rl *clientConnReadLoop) processResetStream(f *RSTStreamFrame) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: return error if server tries to RST_STREAM an idle stream

// Ping sends a PING frame to the server and waits for the ack.
func (cc *ClientConn) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Generate a random payload

// check for dup before insert

// connection closed

func (rl *clientConnReadLoop) processPing(f *PingFrame) error {
	_ = "STUB: not implemented"
	return nil
}

// If ack, notify listener if any

func (rl *clientConnReadLoop) processPushPromise(f *PushPromiseFrame) error {
	_ = "STUB: not implemented"
	// We told the peer we don't want them.
	// Spec says:
	// "PUSH_PROMISE MUST NOT be sent if the SETTINGS_ENABLE_PUSH
	// setting of the peer endpoint is set to 0. An endpoint that
	// has set this setting and has received acknowledgement MUST
	// treat the receipt of a PUSH_PROMISE frame as a connection
	// error (Section 5.4.1) of type PROTOCOL_ERROR."
	return nil
}

func (cc *ClientConn) writeStreamReset(streamID uint32, code ErrCode, err error) {
	_ = "STUB: not implemented"
	// TODO: map err to more interesting error codes, once the
	// HTTP community comes up with some. But currently for
	// RST_STREAM there's no equivalent to GOAWAY frame's debug
	// data, and the error codes are all pretty vague ("cancel").
	return
}

var (
	errResponseHeaderListSize = errors.New("http2: response header list larger than advertised limit")
	errRequestHeaderListSize  = errors.New("http2: request header list larger than peer's advertised limit")
)

func (cc *ClientConn) logf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (cc *ClientConn) vlogf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (t *Transport) vlogf(format string, args ...any) { _ = "STUB: not implemented"; return }

func (t *Transport) logf(format string, args ...any) { _ = "STUB: not implemented"; return }

var noBody io.ReadCloser = noBodyReader{}

type noBodyReader struct{}

func (noBodyReader) Close() error             { _ = "STUB: not implemented"; return nil }
func (noBodyReader) Read([]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type missingBody struct{}

func (missingBody) Close() error { _ = "STUB: not implemented"; return nil }

func (missingBody) Read([]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func strSliceContains(ss []string, s string) bool { _ = "STUB: not implemented"; return false }

type erringRoundTripper struct{ err error }

func (rt erringRoundTripper) RoundTripErr() error { _ = "STUB: not implemented"; return nil }

func (rt erringRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil,

		// isConnectionCloseRequest reports whether req should use its own
		// connection for a single request and then close the connection.
		nil
}

func isConnectionCloseRequest(req *http.Request) bool { _ = "STUB: not implemented"; return false }

// noDialH2RoundTripper is a RoundTripper which only tries to complete the request
// if there's already has a cached connection to the host.
// (The field is exported so it can be accessed via reflect from net/http; tested
// by TestNoDialH2RoundTripperType)
type noDialH2RoundTripper struct{ *Transport }

func (rt noDialH2RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
