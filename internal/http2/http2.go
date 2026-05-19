// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"bufio"
	"crypto/tls"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	VerboseLogs    bool
	logFrameWrites bool
	logFrameReads  bool
	inTests        bool
)

func init() {
	e := os.Getenv("GODEBUG")
	if strings.Contains(e, "http2debug=1") {
		VerboseLogs = true
	}
	if strings.Contains(e, "http2debug=2") {
		VerboseLogs = true
		logFrameWrites = true
		logFrameReads = true
	}
}

const (
	// ClientPreface is the string that must be sent by new
	// connections from clients.
	ClientPreface = "PRI * HTTP/2.0\r\n\r\nSM\r\n\r\n"

	// NextProtoTLS is the NPN/ALPN protocol negotiated during
	// HTTP/2's TLS setup.
	NextProtoTLS = "h2"

	// https://httpwg.org/specs/rfc7540.html#SettingValues
	initialHeaderTableSize = 4096

	initialWindowSize = 65535 // 6.9.2 Initial Flow Control Window Size
)

var clientPreface = []byte(ClientPreface)

// validWireHeaderFieldName reports whether v is a valid header field
// name (key). See httpguts.ValidHeaderName for the base rules.
//
// Further, http2 says:
//
//	"Just as in HTTP/1.x, header field names are strings of ASCII
//	characters that are compared in a case-insensitive
//	fashion. However, header field names MUST be converted to
//	lowercase prior to their encoding in HTTP/2. "
func validWireHeaderFieldName(v string) bool { _ = "STUB: not implemented"; return false }

func httpCodeString(code int) string { _ = "STUB: not implemented"; return "" }

// bufWriterPoolBufferSize is the size of bufio.Writer's
// buffers created using bufWriterPool.
//
// TODO: pick a less arbitrary value? this is a bit under
// (3 x typical 1500 byte MTU) at least. Other than that,
// not much thought went into it.
const bufWriterPoolBufferSize = 4 << 10

var bufWriterPool = sync.Pool{
	New: func() any {
		return bufio.NewWriterSize(nil, bufWriterPoolBufferSize)
	},
}

func mustUint31(v int32) uint32 { _ = "STUB: not implemented"; return 0 }

// bodyAllowedForStatus reports whether a given response status code
// permits a body. See RFC 7230, section 3.3.
func bodyAllowedForStatus(status int) bool { _ = "STUB: not implemented"; return false }

type httpError struct {
	_       incomparable
	msg     string
	timeout bool
}

func (e *httpError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *httpError) Timeout() bool { _ = "STUB: not implemented"; return false }

func (e *httpError) Temporary() bool { _ = "STUB: not implemented"; return false }

var errH2Timeout error = &httpError{msg: "http2: timeout awaiting response headers", timeout: true}

type connectionStater interface {
	ConnectionState() tls.ConnectionState
}

var sorterPool = sync.Pool{New: func() any { return new(sorter) }}

type sorter struct {
	v []string // owned by sorter
}

func (s *sorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (s *sorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s *sorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Keys returns the sorted keys of h.
//
// The returned slice is only valid until s used again or returned to
// its pool.
func (s *sorter) Keys(h http.Header) []string { _ = "STUB: not implemented"; return nil }

func (s *sorter) SortStrings(ss []string) {
	_ = "STUB: not implemented"
	// Our sorter works on s.v, which sorter owns, so
	// stash it away while we sort the user's buffer.
	return
}

// validPseudoPath reports whether v is a valid :path pseudo-header
// value. It must be either:
//
//	*) a non-empty string starting with '/'
//	*) the string '*', for OPTIONS requests.
//
// For now this is only used a quick check for deciding when to clean
// up Opaque URLs before sending requests from the Transport.
// See golang.org/issue/16847
//
// We used to enforce that the path also didn't start with "//", but
// Google's GFE accepts such paths and Chrome sends them, so ignore
// that part of the spec. See golang.org/issue/19103.
func validPseudoPath(v string) bool { _ = "STUB: not implemented"; return false }

// incomparable is a zero-width, non-comparable type. Adding it to a struct
// makes that struct also non-comparable, and generally doesn't add
// any size (as long as it's first).
type incomparable [0]func()
