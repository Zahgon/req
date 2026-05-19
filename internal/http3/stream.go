package http3

import (
	"context"
	"io"
	"net/http"
	"net/http/httptrace"
	"time"

	"github.com/imroc/req/v3/internal/transport"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"

	"github.com/quic-go/qpack"
)

type datagramStream interface {
	io.ReadWriteCloser
	CancelRead(quic.StreamErrorCode)
	CancelWrite(quic.StreamErrorCode)
	StreamID() quic.StreamID
	Context() context.Context
	SetDeadline(time.Time) error
	SetReadDeadline(time.Time) error
	SetWriteDeadline(time.Time) error
	SendDatagram(b []byte) error
	ReceiveDatagram(ctx context.Context) ([]byte, error)

	QUICStream() *quic.Stream
}

// A Stream is an HTTP/3 stream.
//
// When writing to and reading from the stream, data is framed in HTTP/3 DATA frames.
type Stream struct {
	datagramStream
	conn        *Conn
	frameParser *frameParser

	buf []byte // used as a temporary buffer when writing the HTTP/3 frame headers

	bytesRemainingInFrame uint64

	qlogger qlogwriter.Recorder

	parseTrailer  func(io.Reader, *headersFrame) error
	parsedTrailer bool
}

func newStream(
	str datagramStream,
	conn *Conn,
	trace *httptrace.ClientTrace,
	parseTrailer func(io.Reader, *headersFrame) error,
	qlogger qlogwriter.Recorder,
) *Stream {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stream) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// parseNextFrame skips over unknown frame types
// Therefore, this condition is only entered when we parsed another known frame type.

func (s *Stream) hasMoreData() bool { _ = "STUB: not implemented"; return false }

func (s *Stream) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Stream) StreamID() quic.StreamID { _ = "STUB: not implemented"; return *new(quic.StreamID) }

func (s *Stream) SendDatagram(b []byte) error {
	_ = "STUB: not implemented"
	// TODO: reject if datagrams are not negotiated (yet)
	return nil
}

func (s *Stream) ReceiveDatagram(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	// TODO: reject if datagrams are not negotiated (yet)
	return nil, nil
}

// A RequestStream is a low-level abstraction representing an HTTP/3 request stream.
// It decouples sending of the HTTP request from reading the HTTP response, allowing
// the application to optimistically use the stream (and, for example, send datagrams)
// before receiving the response.
//
// This is only needed for advanced use case, e.g. WebTransport and the various
// MASQUE proxying protocols.
type RequestStream struct {
	ctx context.Context
	*transport.Options
	str *Stream

	responseBody io.ReadCloser // set by ReadResponse

	decoder            *qpack.Decoder
	requestWriter      *requestWriter
	maxHeaderBytes     int
	reqDone            chan<- struct{}
	disableCompression bool
	response           *http.Response

	sentRequest   bool
	requestedGzip bool
	isConnect     bool
}

func newRequestStream(
	ctx context.Context,
	options *transport.Options,
	str *Stream,
	requestWriter *requestWriter,
	reqDone chan<- struct{},
	decoder *qpack.Decoder,
	disableCompression bool,
	maxHeaderBytes int,
	rsp *http.Response,
) *RequestStream {
	_ = "STUB: not implemented"
	return nil
}

// Read reads data from the underlying stream.
//
// It can only be used after the request has been sent (using SendRequestHeader)
// and the response has been consumed (using ReadResponse).
func (s *RequestStream) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// StreamID returns the QUIC stream ID of the underlying QUIC stream.
func (s *RequestStream) StreamID() quic.StreamID {
	_ = "STUB: not implemented"
	return *

	// Write writes data to the stream.
	//
	// It can only be used after the request has been sent (using SendRequestHeader).
	new(quic.StreamID)
}

func (s *RequestStream) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close closes the send-direction of the stream.
// It does not close the receive-direction of the stream.
func (s *RequestStream) Close() error { _ = "STUB: not implemented"; return nil }

// CancelRead aborts receiving on this stream.
// See [quic.Stream.CancelRead] for more details.
func (s *RequestStream) CancelRead(errorCode quic.StreamErrorCode) {
	_ = "STUB: not implemented"
	return
}

// CancelWrite aborts sending on this stream.
// See [quic.Stream.CancelWrite] for more details.
func (s *RequestStream) CancelWrite(errorCode quic.StreamErrorCode) {
	_ = "STUB: not implemented"
	return
}

// Context returns a context derived from the underlying QUIC stream's context.
// See [quic.Stream.Context] for more details.
func (s *RequestStream) Context() context.Context {
	_ = "STUB: not implemented"
	return *

	// SetReadDeadline sets the deadline for Read calls.
	new(context.Context)
}

func (s *RequestStream) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetWriteDeadline sets the deadline for Write calls.
func (s *RequestStream) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SetDeadline sets the read and write deadlines associated with the stream.
// It is equivalent to calling both SetReadDeadline and SetWriteDeadline.
func (s *RequestStream) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

// SendDatagrams send a new HTTP Datagram (RFC 9297).
//
// It is only possible to send datagrams if the server enabled support for this extension.
// It is recommended (though not required) to send the request before calling this method,
// as the server might drop datagrams which it can't associate with an existing request.
func (s *RequestStream) SendDatagram(b []byte) error { _ = "STUB: not implemented"; return nil }

// ReceiveDatagram receives HTTP Datagrams (RFC 9297).
//
// It is only possible if support for HTTP Datagrams was enabled, using the EnableDatagram
// option on the [Transport].
func (s *RequestStream) ReceiveDatagram(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SendRequestHeader sends the HTTP request.
//
// It can only used for requests that don't have a request body.
// It is invalid to call it more than once.
// It is invalid to call it after Write has been called.
func (s *RequestStream) SendRequestHeader(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *RequestStream) sendRequestHeader(req *http.Request) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadResponse reads the HTTP response from the stream.
//
// It must be called after sending the request (using SendRequestHeader).
// It is invalid to call it more than once.
// It doesn't set Response.Request and Response.TLS.
// It is invalid to call it after Read has been called.
func (s *RequestStream) ReadResponse() (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Check that the server doesn't send more data in DATA frames than indicated by the Content-Length header (if set).
// See section 4.1.2 of RFC 9114.

// Rules for when to set Content-Length are defined in https://tools.ietf.org/html/rfc7230#section-3.3.2.

type tracingReader struct {
	io.Reader
	readFirst bool
	trace     *httptrace.ClientTrace
}

func (r *tracingReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
