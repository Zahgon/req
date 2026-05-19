// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"errors"
	"sync"
)

// Buffer chunks are allocated from a pool to reduce pressure on GC.
// The maximum wasted space per dataBuffer is 2x the largest size class,
// which happens when the dataBuffer has multiple chunks and there is
// one unread byte in both the first and last chunks. We use a few size
// classes to minimize overheads for servers that typically receive very
// small request bodies.
//
// TODO: Benchmark to determine if the pools are necessary. The GC may have
// improved enough that we can instead allocate chunks like this:
// make([]byte, max(16<<10, expectedBytesRemaining))
var dataChunkPools = [...]sync.Pool{
	{New: func() any { return new([1 << 10]byte) }},
	{New: func() any { return new([2 << 10]byte) }},
	{New: func() any { return new([4 << 10]byte) }},
	{New: func() any { return new([8 << 10]byte) }},
	{New: func() any { return new([16 << 10]byte) }},
}

func getDataBufferChunk(size int64) []byte { _ = "STUB: not implemented"; return nil }

func putDataBufferChunk(p []byte) { _ = "STUB: not implemented"; return }

// dataBuffer is an io.ReadWriter backed by a list of data chunks.
// Each dataBuffer is used to read DATA frames on a single stream.
// The buffer is divided into chunks so the server can limit the
// total memory used by a single connection without limiting the
// request body size on any single stream.
type dataBuffer struct {
	chunks   [][]byte
	r        int   // next byte to read is chunks[0][r]
	w        int   // next byte to write is chunks[len(chunks)-1][w]
	size     int   // total buffered bytes
	expected int64 // we expect at least this many bytes in future Write calls (ignored if <= 0)
}

var errReadEmpty = errors.New("read from empty dataBuffer")

// Read copies bytes from the buffer into p.
// It is an error to read when no data is available.
func (b *dataBuffer) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// If the first chunk has been consumed, advance to the next chunk.

func (b *dataBuffer) bytesFromFirstChunk() []byte { _ = "STUB: not implemented"; return nil }

// Len returns the number of bytes of the unread portion of the buffer.
func (b *dataBuffer) Len() int {
	_ = "STUB: not implemented"

	// Write appends p to the buffer.
	return 0
}

func (b *dataBuffer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// If the last chunk is empty, allocate a new chunk. Try to allocate
// enough to fully copy p plus any additional bytes we expect to
// receive. However, this may allocate less than len(p).

func (b *dataBuffer) lastChunkOrAlloc(want int64) []byte { _ = "STUB: not implemented"; return nil }
