// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The wire protocol for HTTP's "chunked" Transfer-Encoding.

// Package internal contains HTTP internals shared by net/http and
// net/http/httputil.
package internal

import (
	"bufio"
	"errors"
	"io"
)

const maxLineLength = 4096 // assumed <= bufio.defaultBufSize

// ErrLineTooLong is the error that header line too long.
var ErrLineTooLong = errors.New("header line too long")

// NewChunkedReader returns a new chunkedReader that translates the data read from r
// out of HTTP "chunked" format before returning it.
// The chunkedReader returns io.EOF when the final 0-length chunk is read.
//
// NewChunkedReader is not needed by normal applications. The http package
// automatically decodes chunking when reading response bodies.
func NewChunkedReader(r io.Reader) io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

type chunkedReader struct {
	r        *bufio.Reader
	n        uint64 // unread bytes in chunk
	err      error
	buf      [2]byte
	checkEnd bool // whether need to check for \r\n chunk footer
}

func (cr *chunkedReader) beginChunk() {
	_ = "STUB: not implemented"
	// chunk-size CRLF
	return
}

func (cr *chunkedReader) chunkHeaderAvailable() bool { _ = "STUB: not implemented"; return false }

func (cr *chunkedReader) Read(b []uint8) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// We have some data. Return early (per the io.Reader
// contract) instead of potentially blocking while
// reading more.

// We've read enough. Don't potentially block
// reading a new chunk header.

// If we're at the end of a chunk, read the next two
// bytes to verify they are "\r\n".

// Read a line of bytes (up to \n) from b.
// Give up if the line exceeds maxLineLength.
// The returned bytes are owned by the bufio.Reader
// so they are only valid until the next bufio read.
func readChunkLine(b *bufio.Reader) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// We always know when EOF is coming.
// If the caller asked for a line, there should be a line.

func trimTrailingWhitespace(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func isASCIISpace(b byte) bool { _ = "STUB: not implemented"; return false }

var semi = []byte(";")

// removeChunkExtension removes any chunk-extension from p.
// For example,
//
//	"0" => "0"
//	"0;token" => "0"
//	"0;token=val" => "0"
//	`0;token="quoted string"` => "0"
func removeChunkExtension(p []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// TODO: care about exact syntax of chunk extensions? We're
// ignoring and stripping them anyway. For now just never
// return an error.

// NewChunkedWriter returns a new chunkedWriter that translates writes into HTTP
// "chunked" format before writing them to w. Closing the returned chunkedWriter
// sends the final 0-length chunk that marks the end of the stream but does
// not send the final CRLF that appears after trailers; trailers and the last
// CRLF must be written separately.
//
// NewChunkedWriter is not needed by normal applications. The http
// package adds chunking automatically if handlers don't set a
// Content-Length header. Using newChunkedWriter inside a handler
// would result in double chunking or chunking with a Content-Length
// length, both of which are wrong.
func NewChunkedWriter(w io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *

	// Writing to chunkedWriter translates to writing in HTTP chunked Transfer
	// Encoding wire format to the underlying Wire chunkedWriter.
	new(io.WriteCloser)
}

type chunkedWriter struct {
	Wire io.Writer
}

// Write the contents of data as one chunk to Wire.
// NOTE: Note that the corresponding chunk-writing procedure in Conn.Write has
// a bug since it does not check for success of io.WriteString
func (cw *chunkedWriter) Write(data []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// Don't send 0-length data. It looks like EOF for chunked encoding.
	return 0, nil
}

func (cw *chunkedWriter) Close() error { _ = "STUB: not implemented"; return nil }

// FlushAfterChunkWriter signals from the caller of NewChunkedWriter
// that each chunk should be followed by a flush. It is used by the
// http.Transport code to keep the buffering behavior for headers and
// trailers, but flush out chunks aggressively in the middle for
// request bodies which may be generated slowly. See Issue 6574.
type FlushAfterChunkWriter struct {
	*bufio.Writer
}

func parseHexUint(v []byte) (n uint64, err error) { _ = "STUB: not implemented"; return 0, nil }
