package http3

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"

	"github.com/imroc/req/v3/internal/transport"
	"github.com/quic-go/quic-go"
)

// Settings are HTTP/3 settings that apply to the underlying connection.
type Settings struct {
	// Support for HTTP/3 datagrams (RFC 9297)
	EnableDatagrams bool
	// Extended CONNECT, RFC 9220
	EnableExtendedConnect bool
	// Other settings, defined by the application
	Other map[uint64]uint64
}

// RoundTripOpt are options for the Transport.RoundTripOpt method.
type RoundTripOpt struct {
	// OnlyCachedConn controls whether the Transport may create a new QUIC connection.
	// If set true and no cached connection is available, RoundTripOpt will return ErrNoCachedConn.
	OnlyCachedConn bool
}

type clientConn interface {
	OpenRequestStream(context.Context) (*RequestStream, error)
	RoundTrip(*http.Request) (*http.Response, error)
}

type roundTripperWithCount struct {
	cancel     context.CancelFunc
	dialing    chan struct{} // closed as soon as quic.Dial(Early) returned
	dialErr    error
	conn       *quic.Conn
	clientConn clientConn

	useCount atomic.Int64
}

func (r *roundTripperWithCount) Close() error { _ = "STUB: not implemented"; return nil }

// Transport implements the http.RoundTripper interface
type Transport struct {
	*transport.Options
	// TLSClientConfig specifies the TLS configuration to use with
	// tls.Client. If nil, the default configuration is used.
	TLSClientConfig *tls.Config

	// QUICConfig is the quic.Config used for dialing new connections.
	// If nil, reasonable default values will be used.
	QUICConfig *quic.Config

	// Dial specifies an optional dial function for creating QUIC
	// connections for requests.
	// If Dial is nil, a UDPConn will be created at the first request
	// and will be reused for subsequent connections to other servers.
	Dial func(ctx context.Context, addr string, tlsCfg *tls.Config, cfg *quic.Config) (*quic.Conn, error)

	// Enable support for HTTP/3 datagrams (RFC 9297).
	// If a QUICConfig is set, datagram support also needs to be enabled on the QUIC layer by setting EnableDatagrams.
	EnableDatagrams bool

	// Additional HTTP/3 settings.
	// It is invalid to specify any settings defined by RFC 9114 (HTTP/3) and RFC 9297 (HTTP Datagrams).
	AdditionalSettings map[uint64]uint64

	// MaxResponseHeaderBytes specifies a limit on how many response bytes are
	// allowed in the server's response header.
	// Zero means to use a default limit.
	MaxResponseHeaderBytes int

	// DisableCompression, if true, prevents the Transport from requesting compression with an
	// "Accept-Encoding: gzip" request header when the Request contains no existing Accept-Encoding value.
	// If the Transport requests gzip on its own and gets a gzipped response, it's transparently
	// decoded in the Response.Body.
	// However, if the user explicitly requested gzip it is not automatically uncompressed.
	DisableCompression bool

	StreamHijacker    func(FrameType, quic.ConnectionTracingID, *quic.Stream, error) (hijacked bool, err error)
	UniStreamHijacker func(StreamType, quic.ConnectionTracingID, *quic.ReceiveStream, error) (hijacked bool)

	Logger *slog.Logger

	mutex sync.Mutex

	initOnce sync.Once
	initErr  error

	newClientConn func(*quic.Conn) clientConn

	clients   map[string]*roundTripperWithCount
	transport *quic.Transport
	closed    bool
}

var (
	_ http.RoundTripper = &Transport{}
	_ io.Closer         = &Transport{}
)

var (
	// ErrNoCachedConn is returned when Transport.OnlyCachedConn is set
	ErrNoCachedConn = errors.New("http3: no cached connection was available")
	// ErrTransportClosed is returned when attempting to use a closed Transport
	ErrTransportClosed = errors.New("http3: transport is closed")
)

func (t *Transport) init() error {
	if t.newClientConn == nil {
		t.newClientConn = func(conn *quic.Conn) clientConn {
			return newClientConn(
				t.Options,
				conn,
				t.EnableDatagrams,
				t.AdditionalSettings,
				t.StreamHijacker,
				t.UniStreamHijacker,
				t.MaxResponseHeaderBytes,
				t.DisableCompression,
				t.Logger,
			)
		}
	}
	if t.QUICConfig == nil {
		t.QUICConfig = defaultQuicConfig.Clone()
		t.QUICConfig.EnableDatagrams = t.EnableDatagrams
	}
	if t.EnableDatagrams && !t.QUICConfig.EnableDatagrams {
		return errors.New("HTTP Datagrams enabled, but QUIC Datagrams disabled")
	}
	if len(t.QUICConfig.Versions) == 0 {
		t.QUICConfig = t.QUICConfig.Clone()
		t.QUICConfig.Versions = []quic.Version{quic.SupportedVersions()[0]}
	}
	if len(t.QUICConfig.Versions) != 1 {
		return errors.New("can only use a single QUIC version for dialing a HTTP/3 connection")
	}
	if t.QUICConfig.MaxIncomingStreams == 0 {
		t.QUICConfig.MaxIncomingStreams = -1 // don't allow any bidirectional streams
	}
	if t.Dial == nil {
		udpConn, err := net.ListenUDP("udp", nil)
		if err != nil {
			return err
		}
		t.transport = &quic.Transport{Conn: udpConn}
	}
	return nil
}

// RoundTripOpt is like RoundTrip, but takes options.
func (t *Transport) RoundTripOpt(req *http.Request, opt RoundTripOpt) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) roundTripOpt(req *http.Request, opt RoundTripOpt) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) doRoundTripOpt(req *http.Request, opt RoundTripOpt, isRetried bool) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// request aborted due to context cancellation

func canRetryRequest(err error, req *http.Request) (*http.Request, error) {
	_ = "STUB: not implemented"
	// error occurred while opening the stream, we can be sure that the request wasn't sent out
	return nil, nil
}

// If the request stream is reset, we can only be sure that the request wasn't processed
// if the error code is H3_REQUEST_REJECTED.

// if the body is nil (or http.NoBody), it's safe to reuse this request and its body

// if the request body can be reset back to its original state via req.GetBody, do that

// RoundTrip does a round trip.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RoundTripOnlyCachedConn round trip only cached conn.
func (t *Transport) RoundTripOnlyCachedConn(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AddConn add a http3 connection, dial new conn if not exists.
func (t *Transport) AddConn(ctx context.Context, addr string) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *Transport) getClient(ctx context.Context, hostname string, onlyCached bool) (rtc *roundTripperWithCount, isReused bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (t *Transport) dial(ctx context.Context, hostname string) (*quic.Conn, clientConn, error) {
	_ = "STUB: not implemented"
	return nil, *new(clientConn), nil
}

// It's ok if net.SplitHostPort returns an error - it could be a hostname/IP address without a port.

// Replace existing ALPNs by H3

func (t *Transport) resolveUDPAddr(ctx context.Context, network, addr string) (*net.UDPAddr, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *Transport) removeClient(hostname string) { _ = "STUB: not implemented"; return }

// NewClientConn creates a new HTTP/3 client connection on top of a QUIC connection.
// Most users should use RoundTrip instead of creating a connection directly.
// Specifically, it is not needed to perform GET, POST, HEAD and CONNECT requests.
//
// Obtaining a ClientConn is only needed for more advanced use cases, such as
// using Extended CONNECT for WebTransport or the various MASQUE protocols.
func (t *Transport) NewClientConn(conn *quic.Conn) *ClientConn {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the QUIC connections that this Transport has used.
// A Transport cannot be used after it has been closed.
func (t *Transport) Close() error { _ = "STUB: not implemented"; return nil }

func hostnameFromURL(url *url.URL) string { _ = "STUB: not implemented"; return "" }

func validMethod(method string) bool {
	_ = "STUB: not implemented"
	/*
				     Method         = "OPTIONS"                ; Section 9.2
		   		                    | "GET"                    ; Section 9.3
		   		                    | "HEAD"                   ; Section 9.4
		   		                    | "POST"                   ; Section 9.5
		   		                    | "PUT"                    ; Section 9.6
		   		                    | "DELETE"                 ; Section 9.7
		   		                    | "TRACE"                  ; Section 9.8
		   		                    | "CONNECT"                ; Section 9.9
		   		                    | extension-method
		   		   extension-method = token
		   		     token          = 1*<any CHAR except CTLs or separators>
	*/return false
}

// copied from net/http/http.go
func isNotToken(r rune) bool { _ = "STUB: not implemented"; return false }

// CloseIdleConnections closes any QUIC connections in the transport's pool that are currently idle.
// An idle connection is one that was previously used for requests but is now sitting unused.
// This method does not interrupt any connections currently in use.
// It also does not affect connections obtained via NewClientConn.
func (t *Transport) CloseIdleConnections() { _ = "STUB: not implemented"; return }
