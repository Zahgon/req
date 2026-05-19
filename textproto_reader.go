// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package req

import (
	"bufio"
	"errors"
	"net/textproto"
	"sync"

	"github.com/imroc/req/v3/internal/dump"
)

func isASCIILetter(b byte) bool {
	_ = "STUB: not implemented"
	// make lower case
	return false
}

// TODO: This should be a distinguishable error (ErrMessageTooLarge)
// to allow mime/multipart to detect it.
var errMessageTooLarge = errors.New("message too large")

// A textprotoReader implements convenience methods for reading requests
// or responses from a text protocol network connection.
type textprotoReader struct {
	R        *bufio.Reader
	buf      []byte // a reusable buffer for readContinuedLineSlice
	readLine func() (line []byte, isPrefix bool, err error)
}

// NewReader returns a new textprotoReader reading from r.
//
// To avoid denial of service attacks, the provided bufio.Reader
// should be reading from an io.LimitReader or similar textprotoReader to bound
// the size of responses.
func newTextprotoReader(r *bufio.Reader, ds dump.Dumpers) *textprotoReader {
	_ = "STUB: not implemented"
	return nil
}

// ReadLine reads a single line from r,
// eliding the final \n or \r\n from the returned string.
func (r *textprotoReader) ReadLine() (string, error) { _ = "STUB: not implemented"; return "", nil }

// readLineSlice reads a single line from r,
// up to lim bytes long (or unlimited if lim is less than 0),
// eliding the final \r or \r\n from the returned string.
func (r *textprotoReader) readLineSlice(lim int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Avoid the copy if the first call produced a full line.

// trim returns s with leading and trailing spaces and tabs removed.
// It does not assume Unicode or UTF-8.
func trim(s []byte) []byte { _ = "STUB: not implemented"; return nil }

// readContinuedLineSlice reads continued lines from the reader buffer,
// returning a byte slice with all lines. The validateFirstLine function
// is run on the first read line, and if it returns an error then this
// error is returned from readContinuedLineSlice.
// It reads up to lim bytes of data (or unlimited if lim is less than 0).
func (r *textprotoReader) readContinuedLineSlice(lim int64, validateFirstLine func([]byte) error) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read the first line.

// blank line - no continuation

// Optimistically assume that we have started to buffer the next line
// and it starts with an ASCII letter (the next header key), or a blank
// line, so we can avoid copying that buffered data around in memory
// and skipping over non-existent whitespace.

// ReadByte or the next readLineSlice will flush the read buffer;
// copy the slice into buf.

// Read continuation lines.

// skipSpace skips R over all spaces and returns the number of bytes skipped.
func (r *textprotoReader) skipSpace() int { _ = "STUB: not implemented"; return 0 }

// Bufio will keep err until next read.

// A protocolError describes a protocol violation such
// as an invalid response or a hung-up connection.
type protocolError string

func (p protocolError) Error() string { _ = "STUB: not implemented"; return "" }

var colon = []byte(":")

// ReadMIMEHeader reads a MIME-style header from r.
// The header is a sequence of possibly continued Key: Value lines
// ending in a blank line.
// The returned map m maps [CanonicalMIMEHeaderKey](key) to a
// sequence of values in the same order encountered in the input.
//
// For example, consider this input:
//
//	My-Key: Value 1
//	Long-Key: Even
//	       Longer Value
//	My-Key: Value 2
//
// Given that input, ReadMIMEHeader returns the map:
//
//	map[string][]string{
//		"My-Key": {"Value 1", "Value 2"},
//		"Long-Key": {"Even Longer Value"},
//	}
func (r *textprotoReader) ReadMIMEHeader() (textproto.MIMEHeader, error) {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader), nil
}

// readMIMEHeader is a version of ReadMIMEHeader which takes a limit on the header size.
// It is called by the mime/multipart package.
func (r *textprotoReader) readMIMEHeader(maxMemory, maxHeaders int64) (textproto.MIMEHeader, error) {
	_ = "STUB: not implemented"
	// Avoid lots of small slice allocations later by allocating one
	// large one ahead of time which we'll cut up into smaller
	// slices. If this isn't big enough later, we allocate small ones.
	return *new(textproto.MIMEHeader), nil
}

// set a cap to avoid overallocation

// Account for 400 bytes of overhead for the MIMEHeader, plus 200 bytes per entry.
// Benchmarking map creation as of go1.20, a one-entry MIMEHeader is 416 bytes and large
// MIMEHeaders average about 200 bytes per entry.

// The first line cannot start with a leading space.

// arbitrary limit on how much of the line we'll quote

// Key ends at first colon.

// Skip initial spaces in value.

// More than likely this will be a single-element key.
// Most headers aren't multi-valued.
// Set the capacity on strs[0] to 1, so any future append
// won't extend the slice into the other strings.

// mustHaveFieldNameColon ensures that, per RFC 7230, the
// field-name is on a single line, so the first line must
// contain a colon.
func mustHaveFieldNameColon(line []byte) error { _ = "STUB: not implemented"; return nil }

var nl = []byte("\n")

// upcomingHeaderKeys returns an approximation of the number of keys
// that will be in this header. If it gets confused, it returns 0.
func (r *textprotoReader) upcomingHeaderKeys() (n int) {
	_ = "STUB: not implemented"
	// Try to determine the 'hint' size.
	return 0
}

// force a buffer load if empty

// Blank line separating headers from the body.

// Folded continuation of the previous line.

const toLower = 'a' - 'A'

// validHeaderFieldByte reports whether b is a valid byte in a header
// field name. RFC 7230 says:
//
//	header-field   = field-name ":" OWS field-value OWS
//	field-name     = token
//	tchar = "!" / "#" / "$" / "%" / "&" / "'" / "*" / "+" / "-" / "." /
//	        "^" / "_" / "`" / "|" / "~" / DIGIT / ALPHA
//	token = 1*tchar
func validHeaderFieldByte(b byte) bool { _ = "STUB: not implemented"; return false }

// canonicalMIMEHeaderKey is like CanonicalMIMEHeaderKey but is
// allowed to mutate the provided byte slice before returning the
// string.
//
// For invalid inputs (if a contains spaces or non-token bytes), a
// is unchanged and a string copy is returned.
//
// ok is true if the header key contains only valid characters and spaces.
// ReadMIMEHeader accepts header keys containing spaces, but does not
// canonicalize them.
func canonicalMIMEHeaderKey(a []byte) (_ string, ok bool) {
	_ = "STUB: not implemented"
	return "", false
}

// See if a looks like a header key. If not, return it unchanged.

// Don't canonicalize.

// We accept invalid headers with a space before the
// colon, but must not canonicalize them.
// See https://go.dev/issue/34540.

// Canonicalize: first letter upper case
// and upper case after each dash.
// (Host, User-Agent, If-Modified-Since).
// MIME headers are ASCII only, so no Unicode issues.

// for next time

// The compiler recognizes m[string(byteSlice)] as a special
// case, so a copy of a's bytes into a new string does not
// happen in this map lookup:

// validHeaderValueByte reports whether c is a valid byte in a header
// field value. RFC 7230 says:
//
//	field-content  = field-vchar [ 1*( SP / HTAB ) field-vchar ]
//	field-vchar    = VCHAR / obs-text
//	obs-text       = %x80-FF
//
// RFC 5234 says:
//
//	HTAB           =  %x09
//	SP             =  %x20
//	VCHAR          =  %x21-7E
func validHeaderValueByte(c byte) bool {
	_ = "STUB: not implemented"
	// mask is a 128-bit bitmap with 1s for allowed bytes,
	// so that the byte c can be tested with a shift and an and.
	// If c >= 128, then 1<<c and 1<<(c-64) will both be zero.
	// Since this is the obs-text range, we invert the mask to
	// create a bitmap with 1s for disallowed bytes.
	return false
}

// VCHAR: %x21-7E
// SP: %x20
// HTAB: %x09

// commonHeader interns common header strings.
var commonHeader map[string]string

var commonHeaderOnce sync.Once

func initCommonHeader() { _ = "STUB: not implemented"; return }

// isTokenTable is a copy of net/http/lex.go's isTokenTable.
// See https://httpwg.github.io/specs/rfc7230.html#rule.token.separators
var isTokenTable = [127]bool{
	'!':  true,
	'#':  true,
	'$':  true,
	'%':  true,
	'&':  true,
	'\'': true,
	'*':  true,
	'+':  true,
	'-':  true,
	'.':  true,
	'0':  true,
	'1':  true,
	'2':  true,
	'3':  true,
	'4':  true,
	'5':  true,
	'6':  true,
	'7':  true,
	'8':  true,
	'9':  true,
	'A':  true,
	'B':  true,
	'C':  true,
	'D':  true,
	'E':  true,
	'F':  true,
	'G':  true,
	'H':  true,
	'I':  true,
	'J':  true,
	'K':  true,
	'L':  true,
	'M':  true,
	'N':  true,
	'O':  true,
	'P':  true,
	'Q':  true,
	'R':  true,
	'S':  true,
	'T':  true,
	'U':  true,
	'W':  true,
	'V':  true,
	'X':  true,
	'Y':  true,
	'Z':  true,
	'^':  true,
	'_':  true,
	'`':  true,
	'a':  true,
	'b':  true,
	'c':  true,
	'd':  true,
	'e':  true,
	'f':  true,
	'g':  true,
	'h':  true,
	'i':  true,
	'j':  true,
	'k':  true,
	'l':  true,
	'm':  true,
	'n':  true,
	'o':  true,
	'p':  true,
	'q':  true,
	'r':  true,
	's':  true,
	't':  true,
	'u':  true,
	'v':  true,
	'w':  true,
	'x':  true,
	'y':  true,
	'z':  true,
	'|':  true,
	'~':  true,
}
