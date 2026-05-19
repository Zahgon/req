// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP client implementation. See RFC 7230 through 7235.
//
// This is the low-level Transport implementation of http.RoundTripper.
// The high-level interface is in client.go.

package req

import (
	"bufio"
	"compress/gzip"
	"container/list"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"sync"
	"time"
	_ "unsafe"

	"github.com/imroc/req/v3/http2"
	h2internal "github.com/imroc/req/v3/internal/http2"
	"github.com/imroc/req/v3/internal/http3"
	"github.com/imroc/req/v3/internal/transport"
	"github.com/imroc/req/v3/pkg/altsvc"
)

// httpVersion represents http version.
type httpVersion string

const (
	// h1 represents "HTTP/1.1"
	h1 httpVersion = "1.1"
	// h2 represents "HTTP/2.0"
	h2 httpVersion = "2"
	// h3 represents "HTTP/3.0"
	h3 httpVersion = "3"
)

// defaultMaxIdleConnsPerHost is the default value of Transport's
// MaxIdleConnsPerHost.
const defaultMaxIdleConnsPerHost = 2

// Transport is an implementation of http.RoundTripper that supports HTTP,
// HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
//
// By default, Transport caches connections for future reuse.
// This may leave many open connections when accessing many hosts.
// This behavior can be managed using Transport's CloseIdleConnections method
// and the MaxIdleConnsPerHost and DisableKeepAlives fields.
//
// Transports should be reused instead of created as needed.
// Transports are safe for concurrent use by multiple goroutines.
//
// A Transport is a low-level primitive for making HTTP and HTTPS requests.
// For high-level functionality, such as cookies and redirects, see Client.
//
// Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2
// for HTTPS URLs, depending on whether the server supports HTTP/2,
// and how the Transport is configured. The DefaultTransport supports HTTP/2.
// To explicitly enable HTTP/2 on a transport, use golang.org/x/net/http2
// and call ConfigureTransport. See the package docs for more about HTTP/2.
//
// Responses with status codes in the 1xx range are either handled
// automatically (100 expect-continue) or ignored. The one
// exception is HTTP status code 101 (Switching Protocols), which is
// considered a terminal status and returned by RoundTrip. To see the
// ignored 1xx responses, use the httptrace trace package's
// ClientTrace.Got1xxResponse.
//
// Transport only retries a request upon encountering a network error
// if the request is idempotent and either has no body or has its
// Request.GetBody defined. HTTP requests are considered idempotent if
// they have HTTP methods GET, HEAD, OPTIONS, or TRACE; or if their
// Header map contains an "Idempotency-Key" or "X-Idempotency-Key"
// entry. If the idempotency key value is a zero-length slice, the
// request is treated as idempotent but the header is not sent on the
// wire.
type Transport struct {
	Headers http.Header
	Cookies []*http.Cookie

	idleMu       sync.Mutex
	closeIdle    bool                                // user has requested to close all idle conns
	idleConn     map[connectMethodKey][]*persistConn // most recently used at end
	idleConnWait map[connectMethodKey]wantConnQueue  // waiting getConns
	idleLRU      connLRU

	reqMu       sync.Mutex
	reqCanceler map[*http.Request]context.CancelCauseFunc

	connsPerHostMu   sync.Mutex
	connsPerHost     map[connectMethodKey]int
	connsPerHostWait map[connectMethodKey]wantConnQueue // waiting getConns
	dialsInProgress  wantConnQueue

	altSvcJar        altsvc.Jar
	pendingAltSvcs   map[string]*pendingAltSvc
	pendingAltSvcsMu sync.Mutex

	// Force using specific http version
	forceHttpVersion httpVersion

	transport.Options

	t2 *h2internal.Transport // non-nil if http2 wired up
	t3 *http3.Transport

	// disableAutoDecode, if true, prevents auto detect response
	// body's charset and decode it to utf-8
	disableAutoDecode bool

	// autoDecodeContentType specifies an optional function for determine
	// whether the response body should been auto decode to utf-8.
	// Only valid when DisableAutoDecode is true.
	autoDecodeContentType func(contentType string) bool
	wrappedRoundTrip      http.RoundTripper
	httpRoundTripWrappers []HttpRoundTripWrapper
}

// NewTransport is an alias of T
func NewTransport() *Transport {
	_ = "STUB: not implemented"

	// T create a Transport.
	return nil
}

func T() *Transport { _ = "STUB: not implemented"; return nil }

// HttpRoundTripFunc is a http.RoundTripper implementation, which is a simple function.
type HttpRoundTripFunc func(req *http.Request) (resp *http.Response, err error)

// RoundTrip implements http.RoundTripper.
func (fn HttpRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"

	// HttpRoundTripWrapper is transport middleware function.
	return nil, nil
}

type HttpRoundTripWrapper func(rt http.RoundTripper) http.RoundTripper

// HttpRoundTripWrapperFunc is transport middleware function, more convenient than HttpRoundTripWrapper.
type HttpRoundTripWrapperFunc func(rt http.RoundTripper) HttpRoundTripFunc

func (f HttpRoundTripWrapperFunc) wrapper() HttpRoundTripWrapper {
	_ = "STUB: not implemented"
	return *new(HttpRoundTripWrapper)
}

// WrapRoundTripFunc adds a transport middleware function that will give the caller
// an opportunity to wrap the underlying http.RoundTripper.
func (t *Transport) WrapRoundTripFunc(funcs ...HttpRoundTripWrapperFunc) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// WrapRoundTrip adds a transport middleware function that will give the caller
// an opportunity to wrap the underlying http.RoundTripper.
func (t *Transport) WrapRoundTrip(wrappers ...HttpRoundTripWrapper) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// DisableAutoDecode disable auto-detect charset and decode to utf-8
// (enabled by default).
func (t *Transport) DisableAutoDecode() *Transport { _ = "STUB: not implemented"; return nil }

// EnableAutoDecode enable auto-detect charset and decode to utf-8
// (enabled by default).
func (t *Transport) EnableAutoDecode() *Transport { _ = "STUB: not implemented"; return nil }

// SetAutoDecodeContentTypeFunc set the function that determines whether the
// specified `Content-Type` should be auto-detected and decode to utf-8.
func (t *Transport) SetAutoDecodeContentTypeFunc(fn func(contentType string) bool) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetAutoDecodeAllContentType enable try auto-detect charset and decode all
// content type to utf-8.
func (t *Transport) SetAutoDecodeAllContentType() *Transport { _ = "STUB: not implemented"; return nil }

// SetAutoDecodeContentType set the content types that will be auto-detected and decode
// to utf-8 (e.g. "json", "xml", "html", "text").
func (t *Transport) SetAutoDecodeContentType(contentTypes ...string) {
	_ = "STUB: not implemented"
	return
}

// GetMaxIdleConns returns MaxIdleConns.
func (t *Transport) GetMaxIdleConns() int { _ = "STUB: not implemented"; return 0 }

// SetMaxIdleConns set the MaxIdleConns, which controls the maximum number of idle (keep-alive)
// connections across all hosts. Zero means no limit.
func (t *Transport) SetMaxIdleConns(max int) *Transport { _ = "STUB: not implemented"; return nil }

// SetMaxConnsPerHost set the MaxConnsPerHost, optionally limits the
// total number of connections per host, including connections in the
// dialing, active, and idle states. On limit violation, dials will block.
//
// Zero means no limit.
func (t *Transport) SetMaxConnsPerHost(max int) *Transport { _ = "STUB: not implemented"; return nil }

// SetIdleConnTimeout set the IdleConnTimeout, which  is the maximum
// amount of time an idle (keep-alive) connection will remain idle before
// closing itself.
//
// Zero means no limit.
func (t *Transport) SetIdleConnTimeout(timeout time.Duration) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSHandshakeTimeout set the TLSHandshakeTimeout, which specifies the
// maximum amount of time waiting to wait for a TLS handshake.
//
// Zero means no timeout.
func (t *Transport) SetTLSHandshakeTimeout(timeout time.Duration) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetResponseHeaderTimeout set the ResponseHeaderTimeout, if non-zero, specifies
// the amount of time to wait for a server's response headers after fully writing
// the request (including its body, if any). This time does not include the time
// to read the response body.
func (t *Transport) SetResponseHeaderTimeout(timeout time.Duration) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetExpectContinueTimeout set the ExpectContinueTimeout, if non-zero, specifies
// the amount of time to wait for a server's first response headers after fully
// writing the request headers if the request has an "Expect: 100-continue" header.
// Zero means no timeout and causes the body to be sent immediately, without waiting
// for the server to approve.
// This time does not include the time to send the request header.
func (t *Transport) SetExpectContinueTimeout(timeout time.Duration) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetGetProxyConnectHeader set the GetProxyConnectHeader, which optionally specifies a func
// to return headers to send to proxyURL during a CONNECT request to the ip:port target.
// If it returns an error, the Transport's RoundTrip fails with that error. It can
// return (nil, nil) to not add headers.
// If GetProxyConnectHeader is non-nil, ProxyConnectHeader is ignored.
func (t *Transport) SetGetProxyConnectHeader(fn func(ctx context.Context, proxyURL *url.URL, target string) (http.Header, error)) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetProxyConnectHeader set the ProxyConnectHeader, which optionally specifies headers to
// send to proxies during CONNECT requests.
// To set the header dynamically, see SetGetProxyConnectHeader.
func (t *Transport) SetProxyConnectHeader(header http.Header) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetReadBufferSize set the ReadBufferSize, which specifies the size of the read buffer used
// when reading from the transport.
// If zero, a default (currently 4KB) is used.
func (t *Transport) SetReadBufferSize(size int) *Transport { _ = "STUB: not implemented"; return nil }

// SetWriteBufferSize set the WriteBufferSize, which specifies the size of the write buffer used
// when writing to the transport.
// If zero, a default (currently 4KB) is used.
func (t *Transport) SetWriteBufferSize(size int) *Transport { _ = "STUB: not implemented"; return nil }

// SetMaxResponseHeaderBytes set the MaxResponseHeaderBytes, which specifies a limit on how many
// response bytes are allowed in the server's response header.
//
// Zero means to use a default limit.
func (t *Transport) SetMaxResponseHeaderBytes(max int64) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2MaxHeaderListSize set the http2 MaxHeaderListSize,
// which is the http2 SETTINGS_MAX_HEADER_LIST_SIZE to
// send in the initial settings frame. It is how many bytes
// of response headers are allowed. Unlike the http2 spec, zero here
// means to use a default limit (currently 10MB). If you actually
// want to advertise an unlimited value to the peer, Transport
// interprets the highest possible value here (0xffffffff or 1<<32-1)
// to mean no limit.
func (t *Transport) SetHTTP2MaxHeaderListSize(max uint32) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2StrictMaxConcurrentStreams set the http2
// StrictMaxConcurrentStreams, which controls whether the
// server's SETTINGS_MAX_CONCURRENT_STREAMS should be respected
// globally. If false, new TCP connections are created to the
// server as needed to keep each under the per-connection
// SETTINGS_MAX_CONCURRENT_STREAMS limit. If true, the
// server's SETTINGS_MAX_CONCURRENT_STREAMS is interpreted as
// a global limit and callers of RoundTrip block when needed,
// waiting for their turn.
func (t *Transport) SetHTTP2StrictMaxConcurrentStreams(strict bool) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2ReadIdleTimeout set the http2 ReadIdleTimeout,
// which is the timeout after which a health check using ping
// frame will be carried out if no frame is received on the connection.
// Note that a ping response will is considered a received frame, so if
// there is no other traffic on the connection, the health check will
// be performed every ReadIdleTimeout interval.
// If zero, no health check is performed.
func (t *Transport) SetHTTP2ReadIdleTimeout(timeout time.Duration) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2PingTimeout set the http2 PingTimeout, which is the timeout
// after which the connection will be closed if a response to Ping is
// not received.
// Defaults to 15s
func (t *Transport) SetHTTP2PingTimeout(timeout time.Duration) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2WriteByteTimeout set the http2 WriteByteTimeout, which is the
// timeout after which the connection will be closed no data can be written
// to it. The timeout begins when data is available to write, and is
// extended whenever any bytes are written.
func (t *Transport) SetHTTP2WriteByteTimeout(timeout time.Duration) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2SettingsFrame set the ordered http2 settings frame.
func (t *Transport) SetHTTP2SettingsFrame(settings ...http2.Setting) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2ConnectionFlow set the default http2 connection flow, which is the increment
// value of initial WINDOW_UPDATE frame.
func (t *Transport) SetHTTP2ConnectionFlow(flow uint32) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2HeaderPriority set the header priority param.
func (t *Transport) SetHTTP2HeaderPriority(priority http2.PriorityParam) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2PriorityFrames set the ordered http2 priority frames.
func (t *Transport) SetHTTP2PriorityFrames(frames ...http2.PriorityFrame) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSClientConfig set the custom TLSClientConfig, which specifies the TLS configuration to
// use with tls.Client.
// If nil, the default configuration is used.
// If non-nil, HTTP/2 support may not be enabled by default.
func (t *Transport) SetTLSClientConfig(cfg *tls.Config) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetDebug set the optional debug function.
func (t *Transport) SetDebug(debugf func(format string, v ...any)) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetProxy set the http proxy, only valid for HTTP1 and HTTP2, which specifies a function
// to return a proxy for a given Request. If the function returns a non-nil error, the request
// is aborted with the provided error.
//
// The proxy type is determined by the URL scheme. "http",
// "https", and "socks5" are supported. If the scheme is empty,
// "http" is assumed.
//
// If Proxy is nil or returns a nil *URL, no proxy is used.
func (t *Transport) SetProxy(proxy func(*http.Request) (*url.URL, error)) *Transport {
	_ = "STUB: not implemented"
	return nil

	// SetDial set the custom DialContext function, only valid for HTTP1 and HTTP2, which specifies the
	// dial function for creating unencrypted TCP connections.
	// If it is nil, then the transport dials using package net.
	//
	// The dial function runs concurrently with calls to RoundTrip.
	// A RoundTrip call that initiates a dial may end up using a connection dialed previously when the
	// earlier connection becomes idle before the later dial function completes.
}

func (t *Transport) SetDial(fn func(ctx context.Context, network, addr string) (net.Conn, error)) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetDialTLS set the custom DialTLSContext function, only valid for HTTP1 and HTTP2, which specifies
// an optional dial function for creating TLS connections for non-proxied HTTPS requests (proxy will
// not work if set).
//
// If it is nil, DialContext and TLSClientConfig are used.
//
// If it is set, the function that set in SetDial is not used for HTTPS requests and the TLSClientConfig
// and TLSHandshakeTimeout are ignored. The returned net.Conn is assumed to already be past the TLS handshake.
func (t *Transport) SetDialTLS(fn func(ctx context.Context, network, addr string) (net.Conn, error)) *Transport {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSHandshake set the custom tls handshake function, only valid for HTTP1 and HTTP2, not HTTP3,
// it specifies an optional dial function for tls handshake, it works even if a proxy is set, can be
// used to customize the tls fingerprint.
func (t *Transport) SetTLSHandshake(fn func(ctx context.Context, addr string, plainConn net.Conn) (conn net.Conn, tlsState *tls.ConnectionState, err error)) *Transport {
	_ = "STUB: not implemented"
	return nil
}

type pendingAltSvc struct {
	CurrentIndex int
	Entries      []*altsvc.AltSvc
	Mu           sync.Mutex
	LastTime     time.Time
	Transport    http.RoundTripper
}

// EnableForceHTTP1 enable force using HTTP1 (disabled by default).
func (t *Transport) EnableForceHTTP1() *Transport { _ = "STUB: not implemented"; return nil }

// EnableForceHTTP2 enable force using HTTP2 for https requests
// (disabled by default).
func (t *Transport) EnableForceHTTP2() *Transport { _ = "STUB: not implemented"; return nil }

// EnableH2C enables HTTP2 over TCP without TLS.
func (t *Transport) EnableH2C() *Transport { _ = "STUB: not implemented"; return nil }

// DisableH2C disables HTTP2 over TCP without TLS.
func (t *Transport) DisableH2C() *Transport { _ = "STUB: not implemented"; return nil }

// EnableForceHTTP3 enable force using HTTP3 for https requests
// (disabled by default).
func (t *Transport) EnableForceHTTP3() *Transport { _ = "STUB: not implemented"; return nil }

// DisableForceHttpVersion disable force using specified http
// version (disabled by default).
func (t *Transport) DisableForceHttpVersion() *Transport { _ = "STUB: not implemented"; return nil }

func (t *Transport) DisableHTTP3() { _ = "STUB: not implemented"; return }

func (t *Transport) EnableHTTP3() { _ = "STUB: not implemented"; return }

type wrapResponseBodyKeyType int

const wrapResponseBodyKey wrapResponseBodyKeyType = iota

type wrapResponseBodyFunc func(rc io.ReadCloser) io.ReadCloser

func (t *Transport) handleResponseBody(res *http.Response, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

var allowedProtocols = map[string]bool{
	"h3": true,
}

func (t *Transport) handleAltSvc(req *http.Request, value string) {
	_ = "STUB: not implemented"
	return
}

func (t *Transport) handlePendingAltSvc(u *url.URL, pas *pendingAltSvc) {
	_ = "STUB: not implemented"
	return
}

// only support h3 in alt-svc for now

func (t *Transport) wrapResponseBody(res *http.Response, wrap wrapResponseBodyFunc) {
	_ = "STUB: not implemented"
	return
}

func (t *Transport) autoDecodeResponseBody(res *http.Response) { _ = "STUB: not implemented"; return }

// do not decode utf-8

func (t *Transport) writeBufferSize() int { _ = "STUB: not implemented"; return 0 }

func (t *Transport) readBufferSize() int { _ = "STUB: not implemented"; return 0 }

// Clone returns a deep copy of t's exported fields.
func (t *Transport) Clone() *Transport { _ = "STUB: not implemented"; return nil }

// clone transport middleware

// EnableDump enables the dump for all requests with specified dump options.
func (t *Transport) EnableDump(opt *DumpOptions) { _ = "STUB: not implemented"; return }

// DisableDump disables the dump.
func (t *Transport) DisableDump() { _ = "STUB: not implemented"; return }

func (t *Transport) hasCustomTLSDialer() bool { _ = "STUB: not implemented"; return false }

// transportRequest is a wrapper around a *Request that adds
// optional extra headers to write and stores any error to return
// from roundTrip.
type transportRequest struct {
	*http.Request                        // original request, not to be mutated
	extra         http.Header            // extra headers to write, or nil
	trace         *httptrace.ClientTrace // optional

	ctx    context.Context // canceled when we are done with the request
	cancel context.CancelCauseFunc

	mu  sync.Mutex // guards err
	err error      // first setError value for mapRoundTripError to consider
}

func (tr *transportRequest) extraHeaders() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (tr *transportRequest) setError(err error) { _ = "STUB: not implemented"; return }

func (t *Transport) roundTripAltSvc(req *http.Request, as *altsvc.AltSvc) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// impossible!

func (t *Transport) checkAltSvc(req *http.Request) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateHeaders(hdrs http.Header) string { _ = "STUB: not implemented"; return "" }

// Don't include the value in the error,
// because it may be sensitive.

// roundTrip implements a http.RoundTripper over HTTP.
func (t *Transport) roundTrip(req *http.Request) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Validate the outgoing headers.

// Validate the outgoing trailers too.

// Transport request context.
//
// If RoundTrip returns an error, it cancels this context before returning.
//
// If RoundTrip returns no error:
//   - For an HTTP/1 request, persistConn.readLoop cancels this context
//     after reading the request body.
//   - For an HTTP/2 request, RoundTrip cancels this context after the HTTP/2
//     RoundTripper returns.

// Convert Request.Cancel into context cancellation.

// Convert Transport.CancelRequest into context cancellation.
//
// This is lamentably expensive. CancelRequest has been deprecated for a long time
// and doesn't work on HTTP/2 requests. Perhaps we should drop support for it entirely.

// treq gets modified by roundTrip, so we need to recreate for each retry.

// Get the cached or newly-created connection to either the
// host (for http or https), the http proxy, or the http proxy
// pre-CONNECTed to https server. In any case, we'll be ready
// to send it requests.

// HTTP/2 path.

// HTTP/2 requests are not cancelable with CancelRequest,
// so we have no further need for the request context.
//
// On the HTTP/1 path, roundTrip takes responsibility for
// canceling the context after the response body is read.

// Failed. Clean up and determine whether to retry.

// Issue 16465: return underlying net.Conn.Read error from peek,
// as we've historically done.

// Issue 49621: Close the request body if pconn.roundTrip
// didn't do so already. This can happen if the pconn
// write loop exits without reading the write request.

// Rewind the body if we're able to.

func awaitLegacyCancel(ctx context.Context, cancel context.CancelCauseFunc, req *http.Request) {
	_ = "STUB: not implemented"
	return
}

var errCannotRewind = errors.New("net/http: cannot rewind body after connection loss")

type readTrackingBody struct {
	io.ReadCloser
	didRead  bool
	didClose bool
}

func (r *readTrackingBody) Read(data []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *readTrackingBody) Close() error { _ = "STUB: not implemented"; return nil }

// setupRewindBody returns a new request with a custom body wrapper
// that can report whether the body needs rewinding.
// This lets rewindBody avoid an error result when the request
// does not have GetBody but the body hasn't been read at all yet.
func setupRewindBody(req *http.Request) *http.Request { _ = "STUB: not implemented"; return nil }

// rewindBody returns a new request with the body rewound.
// It returns req unmodified if the body does not need rewinding.
// rewindBody takes care of closing req.Body when appropriate
// (in all cases except when rewindBody returns req unmodified).
func rewindBody(req *http.Request) (rewound *http.Request, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// nothing to rewind

// shouldRetryRequest reports whether we should retry sending a failed
// HTTP request on a new connection. The non-nil input error is the
// error from roundTrip.
func (pc *persistConn) shouldRetryRequest(req *http.Request, err error) bool {
	_ = "STUB: not implemented"
	return false
}

// Issue 16582: if the user started a bunch of
// requests at once, they can all pick the same conn
// and violate the server's max concurrent streams.
// Instead, match the HTTP/1 behavior for now and dial
// again to get a new TCP connection, rather than failing
// this request.

// User error.

// This was a fresh connection. There's no reason the server
// should've hung up on us.
//
// Also, if we retried now, we could loop forever
// creating new connections and retrying if the server
// is just hanging up on us because it doesn't like
// our request (as opposed to sending an error).

// We never wrote anything, so it's safe to retry, if there's no body or we
// can "rewind" the body with GetBody.

// Don't retry non-idempotent requests.

// We got some non-EOF net.Conn.Read failure reading
// the 1st response byte from the server.

// The server replied with io.EOF while we were trying to
// read the response. Probably an unfortunately keep-alive
// timeout, just as the client was writing a request.

// conservatively

// CloseIdleConnections closes any connections which were previously
// connected from previous requests but are now sitting idle in
// a "keep-alive" state. It does not interrupt any connections currently
// in use.
func (t *Transport) CloseIdleConnections() { _ = "STUB: not implemented"; return }

// close newly idle connections

// prepareTransportCancel sets up state to convert Transport.CancelRequest into context cancellation.
func (t *Transport) prepareTransportCancel(req *http.Request, origCancel context.CancelCauseFunc) context.CancelCauseFunc {
	_ = "STUB: not implemented"
	// Historically, RoundTrip has not modified the Request in any way.
	// We could avoid the need to keep a map of all in-flight requests by adding
	// a field to the Request containing its cancel func, and setting that field
	// while the request is in-flight. Callers aren't supposed to reuse a Request
	// until after the response body is closed, so this wouldn't violate any
	// concurrency guarantees.
	return *new(context.CancelCauseFunc)
}

// CancelRequest cancels an in-flight request by closing its connection.
// CancelRequest should only be called after [Transport.RoundTrip] has returned.
//
// Deprecated: Use [Request.WithContext] to create a request with a
// cancelable context instead. CancelRequest cannot cancel HTTP/2
// requests. This may become a no-op in a future release of Go.
func (t *Transport) CancelRequest(req *http.Request) { _ = "STUB: not implemented"; return }

// resetProxyConfig is used by tests.
func resetProxyConfig() { _ = "STUB: not implemented"; return }

func (t *Transport) connectMethodForRequest(treq *transportRequest) (cm connectMethod, err error) {
	_ = "STUB: not implemented"
	return *new(connectMethod), nil
}

// proxyAuth returns the Proxy-Authorization header to set
// on requests, if applicable.
func (cm *connectMethod) proxyAuth() string { _ = "STUB: not implemented"; return "" }

// error values for debugging and testing, not seen by users.
var (
	errKeepAlivesDisabled = errors.New("http: putIdleConn: keep alives disabled")
	errConnBroken         = errors.New("http: putIdleConn: connection is in bad state")
	errCloseIdle          = errors.New("http: putIdleConn: CloseIdleConnections was called")
	errTooManyIdle        = errors.New("http: putIdleConn: too many idle connections")
	errTooManyIdleHost    = errors.New("http: putIdleConn: too many idle connections for host")
	errCloseIdleConns     = errors.New("http: CloseIdleConnections called")
	errReadLoopExiting    = errors.New("http: persistConn.readLoop exiting")
	errIdleConnTimeout    = errors.New("http: idle connection timeout")

	// errServerClosedIdle is not seen by users for idempotent requests, but may be
	// seen by a user if the server shuts down an idle connection and sends its FIN
	// in flight with already-written POST body bytes from the client.
	// See https://github.com/golang/go/issues/19943#issuecomment-355607646
	errServerClosedIdle = errors.New("http: server closed idle connection")
)

// transportReadFromServerError is used by Transport.readLoop when the
// 1 byte peek read fails and we're actually anticipating a response.
// Usually this is just due to the inherent keep-alive shut down race,
// where the server closed the connection at the same time the client
// wrote. The underlying err field is usually io.EOF or some
// ECONNRESET sort of thing which varies by platform. But it might be
// the user's custom net.Conn.Read error too, so we carry it along for
// them to return from Transport.RoundTrip.
type transportReadFromServerError struct {
	err error
}

func (e transportReadFromServerError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e transportReadFromServerError) Error() string { _ = "STUB: not implemented"; return "" }

func (t *Transport) putOrCloseIdleConn(pconn *persistConn) { _ = "STUB: not implemented"; return }

func (t *Transport) maxIdleConnsPerHost() int { _ = "STUB: not implemented"; return 0 }

// tryPutIdleConn adds pconn to the list of idle persistent connections awaiting
// a new request.
// If pconn is no longer needed or not in a good state, tryPutIdleConn returns
// an error explaining why it wasn't registered.
// tryPutIdleConn does not close pconn. Use putOrCloseIdleConn instead for that.
func (t *Transport) tryPutIdleConn(pconn *persistConn) error { _ = "STUB: not implemented"; return nil }

// HTTP/2 (pconn.alt != nil) connections do not come out of the idle list,
// because multiple goroutines can use them simultaneously.
// If this is an HTTP/2 connection being “returned,” we're done.

// Deliver pconn to goroutine waiting for idle connection, if any.
// (They may be actively dialing, but this conn is ready first.
// Chrome calls this socket late binding.
// See https://www.chromium.org/developers/design-documents/network-stack#TOC-Connection-Management.)

// HTTP/1.
// Loop over the waiting list until we find a w that isn't done already, and hand it pconn.

// HTTP/2.
// Can hand the same pconn to everyone in the waiting list,
// and we still won't be done: we want to put it in the idle
// list unconditionally, for any future clients too.

// Set idle timer, but only for HTTP/1 (pconn.alt == nil).
// The HTTP/2 implementation manages the idle timer itself
// (see idleConnTimeout in h2_bundle.go).

// queueForIdleConn queues w to receive the next idle connection for w.cm.
// As an optimization hint to the caller, queueForIdleConn reports whether
// it successfully delivered an already-idle connection.
func (t *Transport) queueForIdleConn(w *wantConn) (delivered bool) {
	_ = "STUB: not implemented"
	return false
}

// Stop closing connections that become idle - we might want one.
// (That is, undo the effect of t.CloseIdleConnections.)

// Happens in test hook.

// If IdleConnTimeout is set, calculate the oldest
// persistConn.idleAt time we're willing to use a cached idle
// conn.

// Look for most recently-used idle connection.

// See whether this connection has been idle too long, considering
// only the wall time (the Round(0)), in case this is a laptop or VM
// coming out of suspend with previously cached idle connections.

// Async cleanup. Launch in its own goroutine (as if a
// time.AfterFunc called it); it acquires idleMu, which we're
// holding, and does a synchronous net.Conn.Close.

// If either persistConn.readLoop has marked the connection
// broken, but Transport.removeIdleConn has not yet removed it
// from the idle list, or if this persistConn is too old (it was
// idle too long), then ignore it and look for another. In both
// cases it's already in the process of being closed.

// HTTP/2: multiple clients can share pconn.
// Leave it in the list.

// HTTP/1: only one client can use pconn.
// Remove it from the list.

// Register to receive next connection that becomes idle.

// removeIdleConn marks pconn as dead.
func (t *Transport) removeIdleConn(pconn *persistConn) bool {
	_ = "STUB: not implemented"
	return false
}

// t.idleMu must be held.
func (t *Transport) removeIdleConnLocked(pconn *persistConn) bool {
	_ = "STUB: not implemented"
	return false
}

// Nothing

// Slide down, keeping most recently-used
// conns at the end.

var zeroDialer net.Dialer

func (t *Transport) dial(ctx context.Context, network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// A wantConn records state about a wanted connection
// (that is, an active call to getConn).
// The conn may be gotten by dialing or by finding an idle connection,
// or a cancellation may make the conn no longer wanted.
// These three options are racing against each other and use
// wantConn to coordinate and agree about the winning outcome.
type wantConn struct {
	cm  connectMethod
	key connectMethodKey // cm.key()

	// hooks for testing to know when dials are done
	// beforeDial is called in the getConn goroutine when the dial is queued.
	// afterDial is called when the dial is completed or canceled.
	beforeDial func()
	afterDial  func()

	mu        sync.Mutex      // protects ctx, done and sending of the result
	ctx       context.Context // context for dial, cleared after delivered or canceled
	cancelCtx context.CancelFunc
	done      bool             // true after delivered or canceled
	result    chan connOrError // channel to deliver connection or error
}

type connOrError struct {
	pc     *persistConn
	err    error
	idleAt time.Time
}

// waiting reports whether w is still waiting for an answer (connection or error).
func (w *wantConn) waiting() bool { _ = "STUB: not implemented"; return false }

// getCtxForDial returns context for dial or nil if connection was delivered or canceled.
func (w *wantConn) getCtxForDial() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// tryDeliver attempts to deliver pc, err to w and reports whether it succeeded.
func (w *wantConn) tryDeliver(pc *persistConn, err error, idleAt time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

// cancel marks w as no longer wanting a result (for example, due to cancellation).
// If a connection has been delivered already, cancel returns it with t.putOrCloseIdleConn.
func (w *wantConn) cancel(t *Transport, err error) { _ = "STUB: not implemented"; return }

// A wantConnQueue is a queue of wantConns.
type wantConnQueue struct {
	// This is a queue, not a deque.
	// It is split into two stages - head[headPos:] and tail.
	// popFront is trivial (headPos++) on the first stage, and
	// pushBack is trivial (append) on the second stage.
	// If the first stage is empty, popFront can swap the
	// first and second stages to remedy the situation.
	//
	// This two-stage split is analogous to the use of two lists
	// in Okasaki's purely functional queue but without the
	// overhead of reversing the list when swapping stages.
	head    []*wantConn
	headPos int
	tail    []*wantConn
}

// len returns the number of items in the queue.
func (q *wantConnQueue) len() int { _ = "STUB: not implemented"; return 0 }

// pushBack adds w to the back of the queue.
func (q *wantConnQueue) pushBack(w *wantConn) { _ = "STUB: not implemented"; return }

// popFront removes and returns the wantConn at the front of the queue.
func (q *wantConnQueue) popFront() *wantConn { _ = "STUB: not implemented"; return nil }

// Pick up tail as new head, clear tail.

// peekFront returns the wantConn at the front of the queue without removing it.
func (q *wantConnQueue) peekFront() *wantConn { _ = "STUB: not implemented"; return nil }

// cleanFrontNotWaiting pops any wantConns that are no longer waiting from the head of the
// queue, reporting whether any were popped.
func (q *wantConnQueue) cleanFrontNotWaiting() (cleaned bool) {
	_ = "STUB: not implemented"
	return false
}

// cleanFrontCanceled pops any wantConns with canceled dials from the head of the queue.
func (q *wantConnQueue) cleanFrontCanceled() { _ = "STUB: not implemented"; return }

// all iterates over all wantConns in the queue.
// The caller must not modify the queue while iterating.
func (q *wantConnQueue) all(f func(*wantConn)) { _ = "STUB: not implemented"; return }

func (t *Transport) customDialTLS(ctx context.Context, network, addr string) (conn net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// getConn dials and creates a new persistConn to the target as
// specified in the connectMethod. This includes doing a proxy CONNECT
// and/or setting up TLS.  If this doesn't return an error, the persistConn
// is ready to write requests to.
func (t *Transport) getConn(treq *transportRequest, cm connectMethod) (pc *persistConn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Detach from the request context's cancellation signal.
// The dial should proceed even if the request is canceled,
// because a future request may be able to make use of the connection.
//
// We retain the request context's values.

// Queue for idle connection.

// Wait for completion or cancellation.

// Trace success but only for HTTP/1.
// HTTP/2 calls trace.GotConn itself.

// If the request has been canceled, that's probably
// what caused r.err; if so, prefer to return the
// cancellation error (see golang.org/issue/16049).

// return below

// queueForDial queues w to wait for permission to begin dialing.
// Once w receives permission to dial, it will do so in a separate goroutine.
func (t *Transport) queueForDial(w *wantConn) { _ = "STUB: not implemented"; return }

// startDialConnFor calls dialConn in a new goroutine.
// t.connsPerHostMu must be held.
func (t *Transport) startDialConnForLocked(w *wantConn) { _ = "STUB: not implemented"; return }

// dialConnFor dials on behalf of w and delivers the result to w.
// dialConnFor has received permission to dial w.cm and is counted in t.connCount[w.cm.key()].
// If the dial is canceled or unsuccessful, dialConnFor decrements t.connCount[w.cm.key()].
func (t *Transport) dialConnFor(w *wantConn) { _ = "STUB: not implemented"; return }

// pconn was not passed to w,
// or it is HTTP/2 and can be shared.
// Add to the idle connection pool.

// decConnsPerHost decrements the per-host connection count for key,
// which may in turn give a different waiting goroutine permission to dial.
func (t *Transport) decConnsPerHost(key connectMethodKey) { _ = "STUB: not implemented"; return }

// Shouldn't happen, but if it does, the counting is buggy and could
// easily lead to a silent deadlock, so report the problem loudly.

// Can we hand this count to a goroutine still waiting to dial?
// (Some goroutines on the wait list may have timed out or
// gotten a connection another way. If they're all gone,
// we don't want to kick off any spurious dial operations.)

// q is a value (like a slice), so we have to store
// the updated q back into the map.

// Otherwise, decrement the recorded count.

// Add TLS to a persistent connection, i.e. negotiate a TLS session. If pconn is already a TLS
// tunnel, this function establishes a nested TLS session inside the encrypted channel.
// The remote endpoint's name may be overridden by TLSClientConfig.ServerName.
func (pc *persistConn) addTLS(ctx context.Context, name string, trace *httptrace.ClientTrace, forProxy bool) error {
	_ = "STUB: not implemented"
	// Initiate TLS and check remote host name against certificate.
	return nil
}

// for canceling TLS handshake

// Now that we have closed the connection,
// wait for the call to HandshakeContext to return.

func newHttp2NotSupportedError(negotiatedProtocol string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) customTlsHandshake(ctx context.Context, trace *httptrace.ClientTrace, addr string, pconn *persistConn) error {
	_ = "STUB: not implemented"
	return nil
}

// for canceling TLS handshake

var testHookProxyConnectTimeout = context.WithTimeout

func (t *Transport) dialConn(ctx context.Context, cm connectMethod) (pconn *persistConn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Return a typed error, per Issue 16997

// Handshake here, in case DialTLS didn't. TLSNextProto below
// depends on it for knowing the connection state.

// Proxy setup.

// Do nothing. Not using a proxy.

// Set a (long) timeout here to make sure we don't block forever
// and leak a goroutine if the connection stops replying after
// the TCP connect.

// closed after CONNECT write+read is done or fails

// write or read error

// Write the CONNECT request & read the response.

// Okay to use and discard buffered reader here, because
// TLS server will not speak until spoken to.

// resp or err now set

// persistConnWriter is the io.Writer written to by pc.bw.
// It accumulates the number of bytes written to the underlying conn,
// so the retry logic can determine whether any bytes made it across
// the wire.
// This is exactly 1 pointer field wide so it can go into an interface
// without allocation.
type persistConnWriter struct {
	pc *persistConn
}

func (w persistConnWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom exposes persistConnWriter's underlying Conn to io.Copy and if
// the Conn implements io.ReaderFrom, it can take advantage of optimizations
// such as sendfile.
func (w persistConnWriter) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var _ io.ReaderFrom = (*persistConnWriter)(nil)

// connectMethod is the map key (in its String form) for keeping persistent
// TCP connections alive for subsequent HTTP requests.
//
// A connect method may be of the following types:
//
//	connectMethod.key().String()      Description
//	------------------------------    -------------------------
//	|http|foo.com                     http directly to server, no proxy
//	|https|foo.com                    https directly to server, no proxy
//	|https,h1|foo.com                 https directly to server w/o HTTP/2, no proxy
//	http://proxy.com|https|foo.com    http to proxy, then CONNECT to foo.com
//	http://proxy.com|http             http to proxy, http to anywhere after that
//	socks5://proxy.com|http|foo.com   socks5 to proxy, then http to foo.com
//	socks5://proxy.com|https|foo.com  socks5 to proxy, then https to foo.com
//	https://proxy.com|https|foo.com   https to proxy, then CONNECT to foo.com
//	https://proxy.com|http            https to proxy, http to anywhere after that
type connectMethod struct {
	_            incomparable
	proxyURL     *url.URL // nil for no proxy, else full proxy URL
	targetScheme string   // "http" or "https"
	// If proxyURL specifies an http or https proxy, and targetScheme is http (not https),
	// then targetAddr is not included in the connect method key, because the socket can
	// be reused for different targetAddr values.
	targetAddr string
	onlyH1     bool // whether to disable HTTP/2 and force HTTP/1
}

func (cm *connectMethod) key() connectMethodKey {
	_ = "STUB: not implemented"
	return *new(connectMethodKey)
}

// scheme returns the first hop scheme: http, https, or socks5
func (cm *connectMethod) scheme() string { _ = "STUB: not implemented"; return "" }

// addr returns the first hop "host:port" to which we need to TCP connect.
func (cm *connectMethod) addr() string { _ = "STUB: not implemented"; return "" }

// tlsHost returns the host name to match against the peer's
// TLS certificate.
func (cm *connectMethod) tlsHost() string { _ = "STUB: not implemented"; return "" }

// connectMethodKey is the map key version of connectMethod, with a
// stringified proxy URL (or the empty string) instead of a pointer to
// a URL.
type connectMethodKey struct {
	proxy, scheme, addr string
	onlyH1              bool
}

func (k connectMethodKey) String() string {
	_ = "STUB: not implemented"
	// Only used by tests.
	return ""
}

// persistConn wraps a connection, usually a persistent one
// (but may be used for non-keep-alive requests as well)
type persistConn struct {
	// alt optionally specifies the TLS NextProto http.RoundTripper.
	// This is used for HTTP/2 today and future protocols later.
	// If it's non-nil, the rest of the fields are unused.
	alt http.RoundTripper

	t         *Transport
	cacheKey  connectMethodKey
	conn      net.Conn
	tlsState  *tls.ConnectionState
	br        *bufio.Reader       // from conn
	bw        *bufio.Writer       // to conn
	nwrite    int64               // bytes written
	reqch     chan requestAndChan // written by roundTrip; read by readLoop
	writech   chan writeRequest   // written by roundTrip; read by writeLoop
	closech   chan struct{}       // closed when conn closed
	isProxy   bool
	sawEOF    bool  // whether we've seen EOF from conn; owned by readLoop
	readLimit int64 // bytes allowed to be read; owned by readLoop
	// writeErrCh passes the request write error (usually nil)
	// from the writeLoop goroutine to the readLoop which passes
	// it off to the res.Body reader, which then uses it to decide
	// whether or not a connection can be reused. Issue 7569.
	writeErrCh chan error

	writeLoopDone chan struct{} // closed when write loop ends

	// Both guarded by Transport.idleMu:
	idleAt    time.Time   // time it last become idle
	idleTimer *time.Timer // holding an AfterFunc to close it

	mu                   sync.Mutex // guards following fields
	numExpectedResponses int
	closed               error // set non-nil when conn is closed, before closech is closed
	canceledErr          error // set non-nil if conn is canceled
	broken               bool  // an error has happened on this connection; marked broken so it's not reused.
	reused               bool  // whether conn has had successful request/response and is being reused.
	// mutateHeaderFunc is an optional func to modify extra
	// headers on each outbound request before it's written. (the
	// original Request given to RoundTrip is not modified)
	mutateHeaderFunc func(http.Header)
}

// RFC 7234, section 5.4: Should treat
//
//	Pragma: no-cache
//
// like
//
//	Cache-Control: no-cache
func fixPragmaCacheControl(header http.Header) { _ = "STUB: not implemented"; return }

// readResponse reads an HTTP response (or two, in the case of "Expect:
// 100-continue") from the server. It returns the final non-100 one.
// trace is optional.
func (pc *persistConn) _readResponse(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse the first line of the response.

// Parse the response headers.

func (pc *persistConn) maxHeaderResponseSize() int64 { _ = "STUB: not implemented"; return 0 }

// conservative default; same as http2

func (pc *persistConn) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// isBroken reports whether this connection is in a known broken state.
func (pc *persistConn) isBroken() bool { _ = "STUB: not implemented"; return false }

// canceled returns non-nil if the connection was closed due to
// CancelRequest or due to context cancellation.
func (pc *persistConn) canceled() error { _ = "STUB: not implemented"; return nil }

// isReused reports whether this connection has been used before.
func (pc *persistConn) isReused() bool { _ = "STUB: not implemented"; return false }

func (pc *persistConn) cancelRequest(err error) { _ = "STUB: not implemented"; return }

// closeConnIfStillIdle closes the connection if it's still sitting idle.
// This is what's called by the persistConn's idleTimer, and is run in its
// own goroutine.
func (pc *persistConn) closeConnIfStillIdle() { _ = "STUB: not implemented"; return }

// Not idle.

// mapRoundTripError returns the appropriate error value for
// persistConn.roundTrip.
//
// The provided err is the first error that (*persistConn).roundTrip
// happened to receive from its select statement.
//
// The startBytesWritten value should be the value of pc.nwrite before the roundTrip
// started writing the request.
func (pc *persistConn) mapRoundTripError(req *transportRequest, startBytesWritten int64, err error) error {
	_ = "STUB: not implemented"
	return nil
}

// Wait for the writeLoop goroutine to terminate to avoid data
// races on callers who mutate the request on failure.
//
// When resc in pc.roundTrip and hence rc.ch receives a responseAndError
// with a non-nil error it implies that the persistConn is either closed
// or closing. Waiting on pc.writeLoopDone is hence safe as all callers
// close closech which in turn ensures writeLoop returns.

// If the request was canceled, that's better than network
// failures that were likely the result of tearing down the
// connection.

// See if an error was set explicitly.

// Don't decorate

// Don't decorate

// errCallerOwnsConn is an internal sentinel error used when we hand
// off a writable response.Body to the caller. We use this to prevent
// closing a net.Conn that is now owned by the caller.
var errCallerOwnsConn = errors.New("read loop ending; caller owns writable underlying conn")

func (pc *persistConn) readLoop() { _ = "STUB: not implemented"; return }

// default value, if not changed below

// eofc is used to block caller goroutines reading from Response.Body
// at EOF until this goroutines has (potentially) added the connection
// back to the idle pool.

// unblock reader on errors

// Read this once, before loop starts. (to avoid races in tests)

// effectively no limit for response bodies

// Don't do keep-alive on error if either party requested a close
// or we get an unexpected informational (1xx) response.
// StatusCode 100 is already handled above.

// Put the idle conn back into the pool before we send the response
// so if they process it quickly and make another request, they'll
// get this same conn. But we use the unbuffered channel 'rc'
// to guarantee that persistConn.roundTrip got out of its select
// potentially waiting for this persistConn to close.

// Now that they've read from the unbuffered channel, they're safely
// out of the select that also waits on this goroutine to die, so
// we're allowed to exit now if needed (if alive is false)

// will be closed by deferred call at the end of the function

// see comment above eofc declaration

// Before looping back to the top of this function and peeking on
// the bufio.Reader, wait for the caller goroutine to finish
// reading the response body. (or for cancellation or death)

func (pc *persistConn) readLoopPeekFailLocked(peekErr error) { _ = "STUB: not implemented"; return }

// common case.

// is408Message reports whether buf has the prefix of an
// HTTP 408 Request Timeout response.
// See golang.org/issue/32310.
func is408Message(buf []byte) bool { _ = "STUB: not implemented"; return false }

// readResponse reads an HTTP response (or two, in the case of "Expect:
// 100-continue") from the server. It returns the final non-100 one.
// trace is optional.
func (pc *persistConn) readResponse(rc requestAndChan, trace *httptrace.ClientTrace) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// number of informational 1xx headers received
// arbitrary bound on number of informational responses

// treat 101 as a terminal status, see issue 26161

// reset the limit

// We send an "Expect: 100-continue" header, but the server
// responded with a terminal status and no 100 Continue.
//
// If we're going to keep using the connection, we need to send the request body.
// Tell writeLoop to skip sending the body if we're going to close the connection,
// or to send it otherwise.
//
// The case where we receive a 101 Switching Protocols response is a bit
// ambiguous, since we don't know what protocol we're switching to.
// Conceivably, it's one that doesn't need us to send the body.
// Given that we'll send the body if ExpectContinueTimeout expires,
// be consistent and always send it if we aren't closing the connection.

// don't send the body; the connection will close

// send the body

// waitForContinue returns the function to block until
// any response, timeout or connection close. After any of them,
// the function returns a bool which indicates if the body should be sent.
func (pc *persistConn) waitForContinue(continueCh <-chan struct{}) func() bool {
	_ = "STUB: not implemented"
	return nil
}

func newReadWriteCloserBody(br *bufio.Reader, rwc io.ReadWriteCloser) io.ReadWriteCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadWriteCloser)
}

// readWriteCloserBody is the Response.Body type used when we want to
// give users write access to the Body through the underlying
// connection (TCP, unless using custom dialers). This is then
// the concrete type for a Response.Body on the 101 Switching
// Protocols response, as used by WebSockets, h2c, etc.
type readWriteCloserBody struct {
	_  incomparable
	br *bufio.Reader // used until empty
	io.ReadWriteCloser
}

func (b *readWriteCloserBody) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// nothingWrittenError wraps a write errors which ended up writing zero bytes.
type nothingWrittenError struct {
	error
}

func (nwe nothingWrittenError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (pc *persistConn) writeLoop() { _ = "STUB: not implemented"; return }

// Errors reading from the user's
// Request.Body are high priority.
// Set it here before sending on the
// channels below or calling
// pc.close() which tears down
// connections and causes other
// errors.

// to the body reader, which might recycle us
// to the roundTrip function

// extraHeaders may be nil
// waitForContinue may be nil
// always closes body
func (pc *persistConn) writeRequest(r *http.Request, w io.Writer, usingProxy bool, extraHeaders http.Header, waitForContinue func() bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Find the target host. Prefer the Host: header, but if that
// is not given, use the host from the request URL.
//
// Clean the host, in case it arrives with unexpected stuff in it.

// Validate that the Host header is a valid header in general,
// but don't validate the host itself. This is sufficient to avoid
// header or request smuggling via the Host field.
// The server can (and will, if it's a net/http server) reject
// the request if it doesn't consider the host valid.

// Historically, we would truncate the Host header after '/' or ' '.
// Some users have relied on this truncation to convert a network
// address such as Unix domain socket path into a valid, ignored
// Host header (see https://go.dev/issue/61431).
//
// We don't preserve the truncation, because sending an altered
// header field opens a smuggling vector. Instead, zero out the
// Host header entirely if it isn't valid. (An empty Host is valid;
// see RFC 9112 Section 3.2.)
//
// Return an error if we're sending to a proxy, since the proxy
// probably can't do anything useful with an empty Host header.

// According to RFC 6874, an HTTP client, proxy, or other
// intermediary must remove any IPv6 zone identifier attached
// to an outgoing URI.

// CONNECT requests normally give just the host and port, not a full URL.

// TODO: validate r.Method too? At least it's less likely to
// come from an attacker (more likely to be a constant in
// code).

// Wrap the writer in a bufio Writer if it's not already buffered.
// Don't always call NewWriter, as that forces a bytes.Buffer
// and other small bufio Writers to have a minimum 4k buffer
// size.

// raw writer

// Header lines

// Use the defaultUserAgent unless the Header contains one, which
// may be blank to not send the header.

// Process Body,ContentLength,Close,Trailer

// sort and write headers

// Flush and wait for 100-continue if expected.

// Write body and trailer

// maxWriteWaitBeforeConnReuse is how long the a Transport RoundTrip
// will wait to see the Request's Body.Write result after getting a
// response from the server. See comments in (*persistConn).wroteRequest.
//
// In tests, we set this to a large value to avoid flakiness from inconsistent
// recycling of connections.
var maxWriteWaitBeforeConnReuse = 50 * time.Millisecond

// wroteRequest is a check before recycling a connection that the previous write
// (from writeLoop above) happened and was successful.
func (pc *persistConn) wroteRequest() bool { _ = "STUB: not implemented"; return false }

// Common case: the write happened well before the response, so
// avoid creating a timer.

// Rare case: the request was written in writeLoop above but
// before it could send to pc.writeErrCh, the reader read it
// all, processed it, and called us here. In this case, give the
// write goroutine a bit of time to finish its send.
//
// Less rare case: We also get here in the legitimate case of
// Issue 7569, where the writer is still writing (or stalled),
// but the server has already replied. In this case, we don't
// want to wait too long, and we want to return false so this
// connection isn't reused.

// responseAndError is how the goroutine reading from an HTTP/1 server
// communicates with the goroutine doing the RoundTrip.
type responseAndError struct {
	_   incomparable
	res *http.Response // else use this response (see res method)
	err error
}

type requestAndChan struct {
	_    incomparable
	treq *transportRequest
	ch   chan responseAndError // unbuffered; always send in select on callerGone

	// whether the Transport (as opposed to the user client code)
	// added the Accept-Encoding gzip header. If the Transport
	// set it, only then do we transparently decode the gzip.
	addedGzip bool

	// Optional blocking chan for Expect: 100-continue (for send).
	// If the request has an "Expect: 100-continue" header and
	// the server responds 100 Continue, readLoop send a value
	// to writeLoop via this chan.
	continueCh chan<- struct{}

	callerGone <-chan struct{} // closed when roundTrip caller has returned
}

// A writeRequest is sent by the caller's goroutine to the
// writeLoop's goroutine to write a request while the read loop
// concurrently waits on both the write response and the server's
// reply.
type writeRequest struct {
	req *transportRequest
	ch  chan<- error

	// Optional blocking chan for Expect: 100-continue (for receive).
	// If not nil, writeLoop blocks sending request body until
	// it receives from this chan.
	continueCh <-chan struct{}
}

// httpTimeoutError represents a timeout.
// It implements net.Error and wraps context.DeadlineExceeded.
type timeoutError struct {
	err string
}

func (e *timeoutError) Error() string     { _ = "STUB: not implemented"; return "" }
func (e *timeoutError) Timeout() bool     { _ = "STUB: not implemented"; return false }
func (e *timeoutError) Temporary() bool   { _ = "STUB: not implemented"; return false }
func (e *timeoutError) Is(err error) bool { _ = "STUB: not implemented"; return false }

var errTimeout error = &timeoutError{"net/http: timeout awaiting response headers"}

var errRequestCanceledConn = errors.New("net/http: request canceled while waiting for connection") // TODO: unify?

// errRequestDone is used to cancel the round trip Context after a request is successfully done.
// It should not be seen by the user.
var errRequestDone = errors.New("net/http: request completed")

func nop() {
	_ = "STUB: not implemented"

	// testHooks. Always non-nil.
	return
}

var (
	testHookEnterRoundTrip   = nop
	testHookWaitResLoop      = nop
	testHookRoundTripRetried = nop
	testHookPrePendingDial   = nop
	testHookPostPendingDial  = nop

	testHookMu                     sync.Locker = fakeLocker{} // guards following
	testHookReadLoopBeforeNextRead             = nop
)

func (pc *persistConn) roundTrip(req *transportRequest) (resp *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ask for a compressed version if the caller didn't set their
// own value for Accept-Encoding. We only attempt to
// uncompress the gzip stream if we were the layer that
// requested it.

// Request gzip only, not deflate. Deflate is ambiguous and
// not as universally supported anyway.
// See: https://zlib.net/zlib_faq.html#faq39
//
// Note that we don't request this for HEAD requests,
// due to a bug in nginx:
//   https://trac.nginx.org/nginx/ticket/358
//   https://golang.org/issue/5522
//
// We don't request gzip if the request is for a range, since
// auto-decoding a portion of a gzipped document will just fail
// anyway. See https://golang.org/issue/8923

// Write the request concurrently with waiting for a response,
// in case the server decides to reply before reading our full
// request body.

// prevent leaks

// The pconn closing raced with the response to the request,
// probably after the server wrote a response and immediately
// closed the connection. Use the response.

// readLoop is responsible for canceling req.ctx after
// it reads the response body. Check for a response racing
// the context close, and use the response if available.

// tLogKey is a context WithValue key for test debugging contexts containing
// a t.Logf func. See export_test.go's Request.WithT method.
type tLogKey struct{}

func (tr *transportRequest) logf(format string, args ...any) { _ = "STUB: not implemented"; return }

// markReused marks this connection as having been successfully used for a
// request and response.
func (pc *persistConn) markReused() { _ = "STUB: not implemented"; return }

// close closes the underlying TCP connection and closes
// the pc.closech channel.
//
// The provided err is only for testing and debugging; in normal
// circumstances it should never be seen by users.
func (pc *persistConn) close(err error) { _ = "STUB: not implemented"; return }

func (pc *persistConn) closeLocked(err error) { _ = "STUB: not implemented"; return }

// Close HTTP/1 (pc.alt == nil) connection.
// HTTP/2 closes its connection itself.

var portMap = map[string]string{
	"http":    "80",
	"https":   "443",
	"socks5":  "1080",
	"socks5h": "1080",
}

func idnaASCIIFromURL(url *url.URL) string { _ = "STUB: not implemented"; return "" }

// canonicalAddr returns url.Host but always with a ":port" suffix.
func canonicalAddr(url *url.URL) string { _ = "STUB: not implemented"; return "" }

// bodyEOFSignal is used by the HTTP/1 transport when reading response
// bodies to make sure we see the end of a response body before
// proceeding and reading on the connection again.
//
// It wraps a ReadCloser but runs fn (if non-nil) at most
// once, right before its final (error-producing) Read or Close call
// returns. fn should return the new error to return from Read or Close.
//
// If earlyCloseFn is non-nil and Close is called before io.EOF is
// seen, earlyCloseFn is called instead of fn, and its return value is
// the return value from Close.
type bodyEOFSignal struct {
	body         io.ReadCloser
	mu           sync.Mutex        // guards following 4 fields
	closed       bool              // whether Close has been called
	rerr         error             // sticky Read error
	fn           func(error) error // err will be nil on Read io.EOF
	earlyCloseFn func() error      // optional alt Close func used if io.EOF not seen
}

var errReadOnClosedResBody = errors.New("http: read on closed response body")

func (es *bodyEOFSignal) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (es *bodyEOFSignal) Close() error { _ = "STUB: not implemented"; return nil }

// caller must hold es.mu.
func (es *bodyEOFSignal) condfn(err error) error { _ = "STUB: not implemented"; return nil }

// gzipReader wraps a response body so it can lazily
// call gzip.NewReader on the first call to Read
type gzipReader struct {
	_    incomparable
	body *bodyEOFSignal // underlying HTTP/1 response body framing
	zr   *gzip.Reader   // lazily-initialized gzip reader
	zerr error          // any error from gzip.NewReader; sticky
}

func (gz *gzipReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (gz *gzipReader) Close() error { _ = "STUB: not implemented"; return nil }

type tlsHandshakeTimeoutError struct{}

func (tlsHandshakeTimeoutError) Timeout() bool   { _ = "STUB: not implemented"; return false }
func (tlsHandshakeTimeoutError) Temporary() bool { _ = "STUB: not implemented"; return false }
func (tlsHandshakeTimeoutError) Error() string   { _ = "STUB: not implemented"; return "" }

// fakeLocker is a sync.Locker which does nothing. It's used to guard
// test-only fields when not under test, to avoid runtime atomic
// overhead.
type fakeLocker struct{}

func (fakeLocker) Lock() { _ = "STUB: not implemented"; return }
func (fakeLocker) Unlock() {
	_ = "STUB: not implemented"

	// cloneTLSConfig returns a shallow clone of cfg, or a new zero tls.Config if
	// cfg is nil. This is safe to call even if cfg is in active use by a TLS
	// client or server.
	//
	// cloneTLSConfig should be an internal detail,
	// but widely used packages access it using linkname.
	// Notable members of the hall of shame include:
	//   - github.com/searKing/golang
	//
	// Do not remove or change the type signature.
	// See go.dev/issue/67401.
	//
	//go:linkname cloneTLSConfig
	return
}

func cloneTLSConfig(cfg *tls.Config) *tls.Config { _ = "STUB: not implemented"; return nil }

type connLRU struct {
	ll *list.List // list.Element.Value type is of *persistConn
	m  map[*persistConn]*list.Element
}

// add adds pc to the head of the linked list.
func (cl *connLRU) add(pc *persistConn) { _ = "STUB: not implemented"; return }

func (cl *connLRU) removeOldest() *persistConn { _ = "STUB: not implemented"; return nil }

// remove removes pc from cl.
func (cl *connLRU) remove(pc *persistConn) { _ = "STUB: not implemented"; return }

// len returns the number of items in the cache.
func (cl *connLRU) len() int { _ = "STUB: not implemented"; return 0 }
