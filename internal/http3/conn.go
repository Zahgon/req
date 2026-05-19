package http3

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/imroc/req/v3/internal/transport"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"

	"github.com/quic-go/qpack"
)

const maxQuarterStreamID = 1<<60 - 1

var errGoAway = errors.New("connection in graceful shutdown")

// invalidStreamID is a stream ID that is invalid. The first valid stream ID in QUIC is 0.
const invalidStreamID = quic.StreamID(-1)

// Conn is an HTTP/3 connection.
// It has all methods from the quic.Conn expect for AcceptStream, AcceptUniStream,
// SendDatagram and ReceiveDatagram.
type Conn struct {
	conn *quic.Conn
	*transport.Options

	ctx context.Context

	isServer bool
	logger   *slog.Logger

	enableDatagrams bool

	decoder *qpack.Decoder

	streamMx     sync.Mutex
	streams      map[quic.StreamID]*stateTrackingStream
	lastStreamID quic.StreamID
	maxStreamID  quic.StreamID

	settings         *Settings
	receivedSettings chan struct{}

	idleTimeout time.Duration
	idleTimer   *time.Timer

	qlogger qlogwriter.Recorder
}

func newConnection(
	ctx context.Context,
	quicConn *quic.Conn,
	enableDatagrams bool,
	isServer bool,
	logger *slog.Logger,
	idleTimeout time.Duration,
	options *transport.Options,
) *Conn {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) OpenStream() (*quic.Stream, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) OpenStreamSync(ctx context.Context) (*quic.Stream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) OpenUniStream() (*quic.SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) OpenUniStreamSync(ctx context.Context) (*quic.SendStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) HandshakeComplete() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (c *Conn) ConnectionState() quic.ConnectionState {
	_ = "STUB: not implemented"
	return *new(quic.ConnectionState)
}

func (c *Conn) onIdleTimer() { _ = "STUB: not implemented"; return }

func (c *Conn) clearStream(id quic.StreamID) { _ = "STUB: not implemented"; return }

// The server is performing a graceful shutdown.
// If no more streams are remaining, close the connection.

func (c *Conn) openRequestStream(
	ctx context.Context,
	requestWriter *requestWriter,
	reqDone chan<- struct{},
	disableCompression bool,
	maxHeaderBytes int,
) (*RequestStream, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Streams with stream ID equal to or greater than the stream ID carried in the GOAWAY frame
// will be rejected, see section 5.2 of RFC 9114.

func (c *Conn) decodeTrailers(r io.Reader, streamID quic.StreamID, hf *headersFrame, maxHeaderBytes int) (http.Header, error) {
	_ = "STUB: not implemented"
	return *new(http.Header), nil
}

func (c *Conn) CloseWithError(code quic.ApplicationErrorCode, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) handleUnidirectionalStreams(hijack func(StreamType, quic.ConnectionTracingID, *quic.ReceiveStream, error) (hijacked bool)) {
	_ = "STUB: not implemented"
	return
}

// We're only interested in the control stream here.

// Our QPACK implementation doesn't use the dynamic table yet.

// Our QPACK implementation doesn't use the dynamic table yet.

// only the server can push

// we never increased the Push ID, so we don't expect any push streams

// Only a single control stream is allowed.

func (c *Conn) handleControlStream(str *quic.ReceiveStream) { _ = "STUB: not implemented"; return }

// If datagram support was enabled on our side as well as on the server side,
// we can expect it to have been negotiated both on the transport and on the HTTP/3 layer.
// Note: ConnectionState() will block until the handshake is complete (relevant when using 0-RTT).

// we don't support server push, hence we don't expect any GOAWAY frames from the client

// GOAWAY is the only frame allowed at this point:
// * unexpected frames are ignored by the frame parser
// * we don't support any extension that might add support for more frames

// client-initiated, bidirectional streams

// immediately close the connection if there are currently no active requests

func (c *Conn) sendDatagram(streamID quic.StreamID, b []byte) error {
	_ = "STUB: not implemented"
	// TODO: this creates a lot of garbage and an additional copy
	return nil
}

func (c *Conn) receiveDatagrams() error { _ = "STUB: not implemented"; return nil }

// ReceivedSettings returns a channel that is closed once the peer's SETTINGS frame was received.
// Settings can be obtained from the Settings method after the channel was closed.
func (c *Conn) ReceivedSettings() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Settings returns the settings received on this connection.
// It is only valid to call this function after the channel returned by ReceivedSettings was closed.
func (c *Conn) Settings() *Settings {
	_ = "STUB: not implemented"

	// Context returns the context of the underlying QUIC connection.
	return nil
}

func (c *Conn) Context() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }
