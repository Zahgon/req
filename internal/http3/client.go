package http3

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/imroc/req/v3/internal/dump"
	"github.com/imroc/req/v3/internal/transport"
	"github.com/quic-go/quic-go"

	"github.com/quic-go/qpack"
)

const (
	// MethodGet0RTT allows a GET request to be sent using 0-RTT.
	// Note that 0-RTT doesn't provide replay protection and should only be used for idempotent requests.
	MethodGet0RTT = "GET_0RTT"
	// MethodHead0RTT allows a HEAD request to be sent using 0-RTT.
	// Note that 0-RTT doesn't provide replay protection and should only be used for idempotent requests.
	MethodHead0RTT = "HEAD_0RTT"
)

const (
	defaultUserAgent              = "quic-go HTTP/3"
	defaultMaxResponseHeaderBytes = 10 * 1 << 20 // 10 MB
)

type errConnUnusable struct{ e error }

func (e *errConnUnusable) Unwrap() error { _ = "STUB: not implemented"; return nil }
func (e *errConnUnusable) Error() string { _ = "STUB: not implemented"; return "" }

const max1xxResponses = 5 // arbitrary bound on number of informational responses

var defaultQuicConfig = &quic.Config{
	MaxIncomingStreams: -1, // don't allow the server to create bidirectional streams
	KeepAlivePeriod:    10 * time.Second,
}

// ClientConn is an HTTP/3 client doing requests to a single remote server.
type ClientConn struct {
	*transport.Options
	conn *Conn

	// Enable support for HTTP/3 datagrams (RFC 9297).
	// If a QUICConfig is set, datagram support also needs to be enabled on the QUIC layer by setting enableDatagrams.
	enableDatagrams bool

	// Additional HTTP/3 settings.
	// It is invalid to specify any settings defined by RFC 9114 (HTTP/3) and RFC 9297 (HTTP Datagrams).
	additionalSettings map[uint64]uint64

	// maxResponseHeaderBytes specifies a limit on how many response bytes are
	// allowed in the server's response header.
	maxResponseHeaderBytes int

	// disableCompression, if true, prevents the Transport from requesting compression with an
	// "Accept-Encoding: gzip" request header when the Request contains no existing Accept-Encoding value.
	// If the Transport requests gzip on its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body.
	// However, if the user explicitly requested gzip it is not automatically uncompressed.
	disableCompression bool

	logger *slog.Logger

	requestWriter *requestWriter
	decoder       *qpack.Decoder
}

var _ http.RoundTripper = &ClientConn{}

func newClientConn(
	opts *transport.Options,
	conn *quic.Conn,
	enableDatagrams bool,
	additionalSettings map[uint64]uint64,
	streamHijacker func(FrameType, quic.ConnectionTracingID, *quic.Stream, error) (hijacked bool, err error),
	uniStreamHijacker func(StreamType, quic.ConnectionTracingID, *quic.ReceiveStream, error) (hijacked bool),
	maxResponseHeaderBytes int,
	disableCompression bool,
	logger *slog.Logger,
) *ClientConn {
	_ = "STUB: not implemented"
	return nil
}

// client

// send the SETTINGs frame, using 0-RTT data, if possible

// OpenRequestStream opens a new request stream on the HTTP/3 connection.
func (c *ClientConn) OpenRequestStream(ctx context.Context) (*RequestStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ClientConn) setupConn() error {
	_ = "STUB: not implemented"
	// open the control stream
	return nil
}

// send the SETTINGS frame

func (c *ClientConn) handleBidirectionalStreams(streamHijacker func(FrameType, quic.ConnectionTracingID, *quic.Stream, error) (hijacked bool, err error)) {
	_ = "STUB: not implemented"
	return
}

// RoundTrip executes a request and returns a response
func (c *ClientConn) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if the context was canceled, return the context cancellation error

func (c *ClientConn) roundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	// Immediately send out this request, if this is a 0-RTT request.
	return nil, nil
}

// don't modify the original request

// don't modify the original request

// wait for the handshake to complete

// It is only possible to send an Extended CONNECT request once the SETTINGS were received.
// See section 3 of RFC 8441.

// wait for the server's SETTINGS frame to arrive

// Request Cancellation:
// This go routine keeps running even after RoundTripOpt() returns.
// It is shut down when the application is done processing the body.

// if any error occurred

// ReceivedSettings returns a channel that is closed once the server's HTTP/3 settings were received.
// Settings can be obtained from the Settings method after the channel was closed.
func (c *ClientConn) ReceivedSettings() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Settings returns the HTTP/3 settings for this connection.
// It is only valid to call this function after the channel returned by ReceivedSettings was closed.
func (c *ClientConn) Settings() *Settings { _ = "STUB: not implemented"; return nil }

// CloseWithError closes the connection with the given error code and message.
// It is invalid to call this function after the connection was closed.
func (c *ClientConn) CloseWithError(code ErrCode, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

// Context returns a context that is cancelled when the connection is closed.
func (c *ClientConn) Context() context.Context {
	_ = "STUB: not implemented"
	return *

	// cancelingReader reads from the io.Reader.
	// It cancels writing on the stream if any error other than io.EOF occurs.
	new(context.Context)
}

type cancelingReader struct {
	r   io.Reader
	str *RequestStream
}

func (r *cancelingReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *ClientConn) sendRequestBody(str *RequestStream, body io.ReadCloser, contentLength int64, dumps []*dump.Dumper) error {
	_ = "STUB: not implemented"
	return nil
}

// make sure we don't send more bytes than the content length

func (c *ClientConn) doRequest(req *http.Request, str *RequestStream) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// send the request body asynchronously

// According to the documentation for http.Request.ContentLength,
// a value of 0 with a non-nil Body is also treated as unknown content length.

// copy from net/http: support 1xx responses
// number of informational 1xx headers received

// treat 101 as a terminal status, see https://github.com/golang/go/issues/26161

// Conn returns the underlying HTTP/3 connection.
// This method is only useful for advanced use cases, such as when the application needs to
// open streams on the HTTP/3 connection (e.g. WebTransport).
func (c *ClientConn) Conn() *Conn { _ = "STUB: not implemented"; return nil }
