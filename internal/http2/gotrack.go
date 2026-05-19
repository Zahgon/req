// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Defensive debug-only utility to track that functions run on the
// goroutine that they're supposed to.

package http2

import (
	"os"
	"sync"
)

var DebugGoroutines = os.Getenv("DEBUG_HTTP2_GOROUTINES") == "1"

type goroutineLock uint64

func newGoroutineLock() goroutineLock { _ = "STUB: not implemented"; return *new(goroutineLock) }

func (g goroutineLock) check() { _ = "STUB: not implemented"; return }

func (g goroutineLock) checkNotOn() { _ = "STUB: not implemented"; return }

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 { _ = "STUB: not implemented"; return 0 }

// Parse the 4707 out of "goroutine 4707 ["

var littleBuf = sync.Pool{
	New: func() any {
		buf := make([]byte, 64)
		return &buf
	},
}

// parseUintBytes is like strconv.ParseUint, but using a []byte.
func parseUintBytes(s []byte, base int, bitSize int) (n uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// valid base; nothing to do

// Look for octal, hex prefix.

// n*base overflows

// n+v overflows

// Return the first number n such that n*base >= 1<<64.
func cutoff64(base int) uint64 { _ = "STUB: not implemented"; return 0 }
