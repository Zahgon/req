// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"errors"
	"io"
	"sync"
)

// pipe is a goroutine-safe io.Reader/io.Writer pair. It's like
// io.Pipe except there are no PipeReader/PipeWriter halves, and the
// underlying buffer is an interface. (io.Pipe is always unbuffered)
type pipe struct {
	mu       sync.Mutex
	c        sync.Cond     // c.L lazily initialized to &p.mu
	b        pipeBuffer    // nil when done reading
	unread   int           // bytes unread when done
	err      error         // read error once empty. non-nil means closed.
	breakErr error         // immediate read error (caller doesn't see rest of b)
	donec    chan struct{} // closed on error
	readFn   func()        // optional code to run in Read before error
}

type pipeBuffer interface {
	Len() int
	io.Writer
	io.Reader
}

// setBuffer initializes the pipe buffer.
// It has no effect if the pipe is already closed.
func (p *pipe) setBuffer(b pipeBuffer) { _ = "STUB: not implemented"; return }

func (p *pipe) Len() int { _ = "STUB: not implemented"; return 0 }

// Read waits until data is available and copies bytes
// from the buffer into p.
func (p *pipe) Read(d []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// e.g. copy trailers
// not sticky like p.err

var (
	errClosedPipeWrite        = errors.New("write on closed buffer")
	errUninitializedPipeWrite = errors.New("write on uninitialized buffer")
)

// Write copies bytes from p into the buffer and wakes a reader.
// It is an error to write more data than the buffer can hold.
func (p *pipe) Write(d []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// pipe.setBuffer is never invoked, leaving the buffer uninitialized.
// We shouldn't try to write to an uninitialized pipe,
// but returning an error is better than panicking.

// CloseWithError causes the next Read (waking up a current blocked
// Read if needed) to return the provided err after all data has been
// read.
//
// The error must be non-nil.
func (p *pipe) CloseWithError(err error) { _ = "STUB: not implemented"; return }

// BreakWithError causes the next Read (waking up a current blocked
// Read if needed) to return the provided err immediately, without
// waiting for unread data.
func (p *pipe) BreakWithError(err error) { _ = "STUB: not implemented"; return }

// closeWithErrorAndCode is like CloseWithError but also sets some code to run
// in the caller's goroutine before returning the error.
func (p *pipe) closeWithErrorAndCode(err error, fn func()) { _ = "STUB: not implemented"; return }

func (p *pipe) closeWithError(dst *error, err error, fn func()) { _ = "STUB: not implemented"; return }

// Already been done.

// requires p.mu be held.
func (p *pipe) closeDoneLocked() { _ = "STUB: not implemented"; return }

// Close if unclosed. This isn't racy since we always
// hold p.mu while closing.

// Err returns the error (if any) first set by BreakWithError or CloseWithError.
func (p *pipe) Err() error { _ = "STUB: not implemented"; return nil }

// Done returns a channel which is closed if and when this pipe is closed
// with CloseWithError.
func (p *pipe) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Already hit an error.
