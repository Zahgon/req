package req

import (
	"context"
	"net"
	"net/http/httptrace"
	"time"
)

const (
	traceFmt = `TotalTime         : %v
DNSLookupTime     : %v
TCPConnectTime    : %v
TLSHandshakeTime  : %v
FirstResponseTime : %v
ResponseTime      : %v
IsConnReused:     : false
RemoteAddr        : %v
LocalAddr         : %v`
	traceReusedFmt = `TotalTime         : %v
FirstResponseTime : %v
ResponseTime      : %v
IsConnReused:     : true
RemoteAddr        : %v
LocalAddr         : %v`
)

// Blame return the human-readable reason of why request is slowing.
func (t TraceInfo) Blame() string { _ = "STUB: not implemented"; return "" }

// String return the details of trace information.
func (t TraceInfo) String() string { _ = "STUB: not implemented"; return "" }

// TraceInfo represents the trace information.
type TraceInfo struct {
	// DNSLookupTime is a duration that transport took to perform
	// DNS lookup.
	DNSLookupTime time.Duration

	// ConnectTime is a duration that took to obtain a successful connection.
	ConnectTime time.Duration

	// TCPConnectTime is a duration that took to obtain the TCP connection.
	TCPConnectTime time.Duration

	// TLSHandshakeTime is a duration that TLS handshake took place.
	TLSHandshakeTime time.Duration

	// FirstResponseTime is a duration that server took to respond first byte since
	// connection ready (after tls handshake if it's tls and not a reused connection).
	FirstResponseTime time.Duration

	// ResponseTime is a duration since first response byte from server to
	// request completion.
	ResponseTime time.Duration

	// TotalTime is a duration that total request took end-to-end.
	TotalTime time.Duration

	// IsConnReused is whether this connection has been previously
	// used for another HTTP request.
	IsConnReused bool

	// IsConnWasIdle is whether this connection was obtained from an
	// idle pool.
	IsConnWasIdle bool

	// ConnIdleTime is a duration how long the connection was previously
	// idle, if IsConnWasIdle is true.
	ConnIdleTime time.Duration

	// RemoteAddr returns the remote network address.
	RemoteAddr net.Addr

	// LocalAddr returns the local network address.
	LocalAddr net.Addr
}

type clientTrace struct {
	getConn              time.Time
	dnsStart             time.Time
	dnsDone              time.Time
	connectDone          time.Time
	tlsHandshakeStart    time.Time
	tlsHandshakeDone     time.Time
	gotConn              time.Time
	gotFirstResponseByte time.Time
	endTime              time.Time
	gotConnInfo          httptrace.GotConnInfo
}

func (t *clientTrace) createContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}
