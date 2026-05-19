package http3

import (
	"context"
	"errors"
	"io"
	"sync"

	"github.com/quic-go/quic-go"
)

// A Hijacker allows hijacking of the stream creating part of a quic.Conn from a http.ResponseWriter.
// It is used by WebTransport to create WebTransport streams after a session has been established.
type Hijacker interface {
	Connection() *Conn
}

var errTooMuchData = errors.New("peer sent too much data")

// The body is used in the requestBody (for a http.Request) and the responseBody (for a http.Response).
type body struct {
	str *Stream

	remainingContentLength int64
	violatedContentLength  bool
	hasContentLength       bool
}

func newBody(str *Stream, contentLength int64) *body { _ = "STUB: not implemented"; return nil }

func (r *body) StreamID() quic.StreamID { _ = "STUB: not implemented"; return *new(quic.StreamID) }

func (r *body) checkContentLengthViolation() error { _ = "STUB: not implemented"; return nil }

func (r *body) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *body) Close() error { _ = "STUB: not implemented"; return nil }

type requestBody struct {
	body
	connCtx      context.Context
	rcvdSettings <-chan struct{}
	getSettings  func() *Settings
}

var _ io.ReadCloser = &requestBody{}

type hijackableBody struct {
	body body

	// only set for the http.Response
	// The channel is closed when the user is done with this response:
	// either when Read() errors, or when Close() is called.
	reqDone     chan<- struct{}
	reqDoneOnce sync.Once
}

var _ io.ReadCloser = &hijackableBody{}

func newResponseBody(str *Stream, contentLength int64, done chan<- struct{}) *hijackableBody {
	_ = "STUB: not implemented"
	return nil
}

func (r *hijackableBody) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *hijackableBody) requestDone() { _ = "STUB: not implemented"; return }

func (r *hijackableBody) Close() error {
	_ = "STUB: not implemented"

	// If the EOF was read, CancelRead() is a no-op.
	return nil
}
