// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"context"
	"net"
	"net/http"
	"sync"
)

// ClientConnPool manages a pool of HTTP/2 client connections.
type ClientConnPool interface {
	// GetClientConn returns a specific HTTP/2 connection (usually
	// a TLS-TCP connection) to an HTTP/2 server. On success, the
	// returned ClientConn accounts for the upcoming RoundTrip
	// call, so the caller should not omit it. If the caller needs
	// to, ClientConn.RoundTrip can be called with a bogus
	// new(http.Request) to release the stream reservation.
	GetClientConn(req *http.Request, addr string, dialOnMiss bool) (*ClientConn, error)
	MarkDead(*ClientConn)
	CloseIdleConnections()
	AddConnIfNeeded(key string, t *Transport, c net.Conn) (used bool, err error)
}

// TODO: use singleflight for dialing and addConnCalls?
type clientConnPool struct {
	t *Transport

	mu sync.Mutex // TODO: maybe switch to RWMutex
	// TODO: add support for sharing conns based on cert names
	// (e.g. share conn for googleapis.com and appspot.com)
	conns        map[string][]*ClientConn // key is host:port
	dialing      map[string]*dialCall     // currently in-flight dials
	keys         map[*ClientConn][]string
	addConnCalls map[string]*addConnCall // in-flight addConnIfNeeded calls
}

func (p *clientConnPool) GetClientConn(req *http.Request, addr string, dialOnMiss bool) (*ClientConn, error) {
	_ = "STUB: not implemented"
	// TODO(dneil): Dial a new connection when t.DisableKeepAlives is set?
	return nil, nil
}

// It gets its own connection.

// When a connection is presented to us by the net/http package,
// the GetConn hook has already been called.
// Don't call it a second time here.

// dialCall is an in-flight Transport dial call to a host.
type dialCall struct {
	_ incomparable
	p *clientConnPool
	// the context associated with the request
	// that created this dialCall
	ctx  context.Context
	done chan struct{} // closed when done
	res  *ClientConn   // valid after done is closed
	err  error         // valid after done is closed
}

// requires p.mu is held.
func (p *clientConnPool) getStartDialLocked(ctx context.Context, addr string) *dialCall {
	_ = "STUB: not implemented"
	return nil
}

// A dial is already in-flight. Don't start another.

// run in its own goroutine.
func (c *dialCall) dial(ctx context.Context, addr string) { _ = "STUB: not implemented"; return }

// shared conn

// addConnIfNeeded makes a NewClientConn out of c if a connection for key doesn't
// already exist. It coalesces concurrent calls with the same key.
// This is used by the http1 Transport code when it creates a new connection. Because
// the http1 Transport doesn't de-dup TCP dials to outbound hosts (because it doesn't know
// the protocol), it can get into a situation where it has multiple TLS connections.
// This code decides which ones live or die.
// The return value used is whether c was used.
// c is never closed.
func (p *clientConnPool) AddConnIfNeeded(key string, t *Transport, c net.Conn) (used bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

type addConnCall struct {
	_    incomparable
	p    *clientConnPool
	done chan struct{} // closed when done
	err  error
}

func (c *addConnCall) run(t *Transport, key string, tc net.Conn) { _ = "STUB: not implemented"; return }

// already called by the net/http package

// p.mu must be held
func (p *clientConnPool) addConnLocked(key string, cc *ClientConn) {
	_ = "STUB: not implemented"
	return
}

func (p *clientConnPool) MarkDead(cc *ClientConn) { _ = "STUB: not implemented"; return }

func (p *clientConnPool) CloseIdleConnections() { _ = "STUB: not implemented"; return }

// TODO: don't close a cc if it was just added to the pool
// milliseconds ago and has never been used. There's currently
// a small race window with the HTTP/1 Transport's integration
// where it can add an idle conn just before using it, and
// somebody else can concurrently call CloseIdleConns and
// break some caller's RoundTrip.

func filterOutClientConn(in []*ClientConn, exclude *ClientConn) []*ClientConn {
	_ = "STUB: not implemented"
	return nil
}

// If we filtered it out, zero out the last item to prevent
// the GC from seeing it.

// shouldRetryDial reports whether the current request should
// retry dialing after the call finished unsuccessfully, for example
// if the dial was canceled because of a context cancellation or
// deadline expiry.
func shouldRetryDial(call *dialCall, req *http.Request) bool {
	_ = "STUB: not implemented"
	return false

	// No error, no need to retry
}

// If the call has the same context as the request, the dial
// should not be retried, since any cancellation will have come
// from this request.

// If the call error is not because of a context cancellation or a deadline expiry,
// the dial should not be retried.

// Only retry if the error is a context cancellation error or deadline expiry
// and the context associated with the call was canceled or expired.
