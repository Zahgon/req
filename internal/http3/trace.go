package http3

import (
	"crypto/tls"
	"net"
	"net/http/httptrace"
	"net/textproto"
	"time"

	"github.com/quic-go/quic-go"
)

func traceGetConn(trace *httptrace.ClientTrace, hostPort string) { _ = "STUB: not implemented"; return }

// fakeConn is a wrapper for quic.EarlyConnection
// because the quic connection does not implement net.Conn.
type fakeConn struct {
	conn *quic.Conn
}

func (c *fakeConn) Close() error                       { _ = "STUB: not implemented"; return nil }
func (c *fakeConn) Read(p []byte) (int, error)         { _ = "STUB: not implemented"; return 0, nil }
func (c *fakeConn) Write(p []byte) (int, error)        { _ = "STUB: not implemented"; return 0, nil }
func (c *fakeConn) SetDeadline(t time.Time) error      { _ = "STUB: not implemented"; return nil }
func (c *fakeConn) SetReadDeadline(t time.Time) error  { _ = "STUB: not implemented"; return nil }
func (c *fakeConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
func (c *fakeConn) RemoteAddr() net.Addr               { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *fakeConn) LocalAddr() net.Addr                { _ = "STUB: not implemented"; return *new(net.Addr) }

func traceGotConn(trace *httptrace.ClientTrace, conn *quic.Conn, reused bool) {
	_ = "STUB: not implemented"
	return
}

func traceGotFirstResponseByte(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }

func traceGot1xxResponse(trace *httptrace.ClientTrace, code int, header textproto.MIMEHeader) {
	_ = "STUB: not implemented"
	return
}

func traceGot100Continue(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }

func traceHasWroteHeaderField(trace *httptrace.ClientTrace) bool {
	_ = "STUB: not implemented"
	return false
}

func traceWroteHeaders(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }

func traceWroteRequest(trace *httptrace.ClientTrace, err error) { _ = "STUB: not implemented"; return }

func traceConnectStart(trace *httptrace.ClientTrace, network, addr string) {
	_ = "STUB: not implemented"
	return
}

func traceConnectDone(trace *httptrace.ClientTrace, network, addr string, err error) {
	_ = "STUB: not implemented"
	return
}

func traceTLSHandshakeStart(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }

func traceTLSHandshakeDone(trace *httptrace.ClientTrace, state tls.ConnectionState, err error) {
	_ = "STUB: not implemented"
	return
}
