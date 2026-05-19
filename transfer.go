// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package req

import (
	"bufio"
	"errors"
	"io"
	"net/http"
	"reflect"
	"sync"

	"github.com/imroc/req/v3/internal/dump"
	"github.com/imroc/req/v3/internal/godebug"
)

type errorReader struct {
	err error
}

func (r errorReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

type byteReader struct {
	b    byte
	done bool
}

func (br *byteReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// transferWriter inspects the fields of a user-supplied Request or Response,
// sanitizes them without changing the user object and provides methods for
// writing the respective header, body and trailer in wire format.
type transferWriter struct {
	Method           string
	Body             io.Reader
	BodyCloser       io.Closer
	ContentLength    int64 // -1 means unknown, 0 means exactly none
	Close            bool
	TransferEncoding []string
	Header           http.Header
	Trailer          http.Header
	bodyReadError    error // any non-EOF error from reading Body

	FlushHeaders bool            // flush headers to network before body
	ByteReadCh   chan readResult // non-nil if probeRequestBody called
}

func newTransferWriter(r *http.Request) (t *transferWriter, err error) {
	_ = "STUB: not implemented"
	return nil,

		// Extract relevant fields
		nil
}

// If there's a body, conservatively flush the headers
// to any bufio.Writer we're writing to, just in case
// the server needs the headers early, before we copy
// the body and possibly block. We make an exception
// for the common standard library in-memory types,
// though, to avoid unnecessary TCP packets on the
// wire. (Issue 22088.)

// Transport requests are always 1.1 or 2.0

// Sanitize Body,ContentLength,TransferEncoding

// no chunking, no body

// Sanitize Trailer

// shouldSendChunkedRequestBody reports whether we should try to send a
// chunked request body to the server. In particular, the case we really
// want to prevent is sending a GET or other typically-bodyless request to a
// server with a chunked body when the body has zero bytes, since GETs with
// bodies (while acceptable according to specs), even zero-byte chunked
// bodies, are approximately never seen in the wild and confuse most
// servers. See Issue 18257, as one example.
//
// The only reason we'd send such a request is if the user set the Body to a
// non-nil value (say, io.NopCloser(bytes.NewReader(nil))) and didn't
// set ContentLength, or NewRequest set it to -1 (unknown), so then we assume
// there's bytes to send.
//
// This code tries to read a byte from the Request.Body in such cases to see
// whether the body actually has content (super rare) or is actually just
// a non-nil content-less ReadCloser (the more common case). In that more
// common case, we act as if their Body were nil instead, and don't send
// a body.
func (t *transferWriter) shouldSendChunkedRequestBody() bool {
	_ = "STUB: not implemented"
	// Note that t.ContentLength is the corrected content length
	// from rr.outgoingLength, so 0 actually means zero, not unknown.
	return false
}

// redundant checks; caller did them

// Only probe the Request.Body for GET/HEAD/DELETE/etc
// requests, because it's only those types of requests
// that confuse servers.
// adjusts t.Body, t.ContentLength

// For all other request types (PUT, POST, PATCH, or anything
// made-up we've never heard of), assume it's normal and the server
// can deal with a chunked request body. Maybe we'll adjust this
// later.

// probeRequestBody reads a byte from t.Body to see whether it's empty
// (returns io.EOF right away).
//
// But because we've had problems with this blocking users in the past
// (issue 17480) when the body is a pipe (perhaps waiting on the response
// headers before the pipe is fed data), we need to be careful and bound how
// long we wait for it. This delay will only affect users if all the following
// are true:
//   - the request body blocks
//   - the content length is not set (or set to -1)
//   - the method doesn't usually have a body (GET, HEAD, DELETE, ...)
//   - there is no transfer-encoding=chunked already set.
//
// In other words, this delay will not normally affect anybody, and there
// are workarounds if it does.
func (t *transferWriter) probeRequestBody() { _ = "STUB: not implemented"; return }

// It was empty.

// Too slow. Don't wait. Read it later, and keep
// assuming that this is ContentLength == -1
// (unknown), which means we'll send a
// "Transfer-Encoding: chunked" header.

// Request that Request.Write flush the headers to the
// network before writing the body, since our body may not
// become readable until it's seen the response headers.

func noResponseBodyExpected(requestMethod string) bool { _ = "STUB: not implemented"; return false }

func (t *transferWriter) shouldSendContentLength() bool { _ = "STUB: not implemented"; return false }

// Many servers expect a Content-Length for these methods

func (t *transferWriter) writeHeader(writeHeader func(key string, values ...string) error) error {
	_ = "STUB: not implemented"
	return nil
}

// Write Content-Length and/or Transfer-Encoding whose values are a
// function of the sanitized field triple (Body, ContentLength,
// TransferEncoding)

// Write Trailer header

// TODO: could do better allocation-wise here, but trailers are rare,
// so being lazy for now.

// always closes t.BodyCloser
func (t *transferWriter) writeBody(w io.Writer, dumps []*dump.Dumper) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// raw writer

// Write body. We "unwrap" the body first if it was wrapped in a
// nopCloser or readTrackingBody. This is to ensure that we can take advantage of
// OS-level optimizations in the event that the body is an
// *os.File.

// Write Trailer header

// Last chunk, empty trailer

// doBodyCopy wraps a copy operation, with any resulting error also
// being saved in bodyReadError.
//
// This function is only intended for use in writeBody.
func (t *transferWriter) doBodyCopy(dst io.Writer, src io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// unwrapBody unwraps the body's inner reader if it's a
// nopCloser. This is to ensure that body writes sourced from local
// files (*os.File types) are properly optimized.
//
// This function is only intended for use in writeBody.
func (t *transferWriter) unwrapBody() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

type transferReader struct {
	// Input
	Header        http.Header
	StatusCode    int
	RequestMethod string
	ProtoMajor    int
	ProtoMinor    int
	// Output
	Body          io.ReadCloser
	ContentLength int64
	Chunked       bool
	Close         bool
	Trailer       http.Header
}

func (t *transferReader) protoAtLeast(m, n int) bool { _ = "STUB: not implemented"; return false }

// bodyAllowedForStatus reports whether a given response status code
// permits a body. See RFC 7230, section 3.3.
func bodyAllowedForStatus(status int) bool { _ = "STUB: not implemented"; return false }

// msg is *http.Request or *http.Response.
func readTransfer(msg any, r *bufio.Reader) (err error) { _ = "STUB: not implemented"; return nil }

// Unify input

// Default to HTTP/1.1

// Transfer-Encoding: chunked, and overriding Content-Length.

// Trailer

// If there is no Content-Length or chunked Transfer-Encoding on a *Response
// and the status is not 1xx, 204 or 304, then the body is unbounded.
// See RFC 7230, section 3.3.

// Unbounded body.

// Prepare body reader. ContentLength < 0 means chunked encoding
// or close connection when finished, since multipart is not supported yet

// realLength < 0, i.e. "Content-Length" not mentioned in header

// Close semantics (i.e. HTTP/1.0)

// Persistent connection (i.e. HTTP/1.1)

// Unify output

// Checks whether chunked is part of the encodings stack.
func chunked(te []string) bool { _ = "STUB: not implemented"; return false }

// Checks whether the encoding is explicitly "identity".
func isIdentity(te []string) bool { _ = "STUB: not implemented"; return false }

// unsupportedTEError reports unsupported transfer-encodings.
type unsupportedTEError struct {
	err string
}

func (uste *unsupportedTEError) Error() string {
	_ = "STUB: not implemented"

	// parseTransferEncoding sets t.Chunked based on the Transfer-Encoding header.
	return ""
}

func (t *transferReader) parseTransferEncoding() error { _ = "STUB: not implemented"; return nil }

// Issue 12785; ignore Transfer-Encoding on HTTP/1.0 requests.

// Like nginx, we only support a single Transfer-Encoding header field, and
// only if set to "chunked". This is one of the most security sensitive
// surfaces in HTTP/1.1 due to the risk of request smuggling, so we keep it
// strict and simple.

// Determine the expected body length, using RFC 7230 Section 3.3. This
// function is not a method, because ultimately it should be shared by
// ReadResponse and ReadRequest.
func fixLength(isResponse bool, status int, requestMethod string, header http.Header, chunked bool) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Hardening against HTTP request smuggling

// Per RFC 7230 Section 3.3.2, prevent multiple
// Content-Length headers if they differ in value.
// If there are dups of the value, remove the dups.
// See Issue 16490.

// deduplicate Content-Length

// Reject requests with invalid Content-Length headers.

// Logic based on response type or status

// According to RFC 9112, "If a message is received with both a
// Transfer-Encoding and a Content-Length header field, the Transfer-Encoding
// overrides the Content-Length. Such a message might indicate an attempt to
// perform request smuggling (Section 11.2) or response splitting (Section 11.1)
// and ought to be handled as an error. An intermediary that chooses to forward
// the message MUST first remove the received Content-Length field and process
// the Transfer-Encoding (as described below) prior to forwarding the message downstream."
//
// Chunked-encoding requests with either valid Content-Length
// headers or no Content-Length headers are accepted after removing
// the Content-Length field from header.
//
// Logic based on Transfer-Encoding

// Logic based on Content-Length

// RFC 7230 neither explicitly permits nor forbids an
// entity-body on a GET request so we permit one if
// declared, but we default to 0 here (not -1 below)
// if there's no mention of a body.
// Likewise, all other request methods are assumed to have
// no body if neither Transfer-Encoding chunked nor a
// Content-Length are set.

// Body-EOF logic based on other methods (like closing, or chunked coding)

// Determine whether to hang up after sending a request and body, or
// receiving a response and body
// 'header' is the request headers.
func shouldClose(major, minor int, header http.Header, removeCloseHeader bool) bool {
	_ = "STUB: not implemented"
	return false
}

// Parse the trailer header.
func fixTrailer(header http.Header, chunked bool) (http.Header, error) {
	_ = "STUB: not implemented"
	return *new(http.Header), nil
}

// Trailer and no chunking:
// this is an invalid use case for trailer header.
// Nevertheless, no error will be returned and we
// let users decide if this is a valid HTTP message.
// The Trailer header will be kept in Response.Header
// but not populate Response.Trailer.
// See issue #27197.

// body turns a textprotoReader into a ReadCloser.
// Close ensures that the body has been fully read
// and then reads the trailer if necessary.
type body struct {
	src          io.Reader
	hdr          any           // non-nil (Response or Request) value means read trailer
	r            *bufio.Reader // underlying wire-format reader for the trailer
	closing      bool          // is the connection to be closed after reading body?
	doEarlyClose bool          // whether Close should stop early

	mu         sync.Mutex // guards following, and calls to Read and Close
	sawEOF     bool
	closed     bool
	earlyClose bool   // Close called and we didn't read to the end of src
	onHitEOF   func() // if non-nil, func to call when EOF is Read
}

func (b *body) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Must hold b.mu.
func (b *body) readLocked(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Chunked case. Read the trailer.

// Something went wrong in the trailer, we must not allow any
// further reads of any kind to succeed from body, nor any
// subsequent requests on the server connection. See
// golang.org/issue/12027

// If the server declared the Content-Length, our body is a LimitedReader
// and we need to check whether this EOF arrived early.

// If we can return an EOF here along with the read data, do
// so. This is optional per the io.textprotoReader contract, but doing
// so helps the HTTP transport code recycle its connection
// earlier (since it will see this EOF itself), even if the
// client doesn't do future reads or Close.

var (
	singleCRLF = []byte("\r\n")
	doubleCRLF = []byte("\r\n\r\n")
)

func seeUpcomingDoubleCRLF(r *bufio.Reader) bool { _ = "STUB: not implemented"; return false }

// This loop stops when Peek returns an error,
// which it does when r's buffer has been filled.

var errTrailerEOF = errors.New("http: unexpected EOF reading trailer")

func (b *body) readTrailer() error {
	_ = "STUB: not implemented"
	// The common case, since nobody uses trailers.
	return nil
}

// Make sure there's a header terminator coming up, to prevent
// a DoS with an unbounded size Trailer. It's not easy to
// slip in a LimitReader here, as textproto.NewReader requires
// a concrete *bufio.textprotoReader. Also, we can't get all the way
// back up to our conn's LimitedReader that *might* be backing
// this bufio.textprotoReader. Instead, a hack: we iteratively Peek up
// to the bufio.textprotoReader's max size, looking for a double CRLF.
// This limits the trailer to the underlying buffer size, typically 4kB.

func mergeSetHeader(dst *http.Header, src http.Header) { _ = "STUB: not implemented"; return }

func (b *body) Close() error { _ = "STUB: not implemented"; return nil }

// Already saw EOF, so no need going to look for it.

// no trailer and closing the connection next.
// no point in reading to EOF.

// Read up to maxPostHandlerReadBytes bytes of the body, looking
// for EOF (and trailers), so we can reuse this connection.

// There was a declared Content-Length, and we have more bytes remaining
// than our maxPostHandlerReadBytes tolerance. So, give up.

// Consume the body, or, which will also lead to us reading
// the trailer headers after the body, if present.

// Fully consume the body, which will also lead to us reading
// the trailer headers after the body, if present.

// bodyLocked is an io.Reader reading from a *body when its mutex is
// already held.
type bodyLocked struct {
	b *body
}

func (bl bodyLocked) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

var httplaxContentLength = godebug.New("httplaxcontentlength")

// parseContentLength checks that the header is valid and then trims
// whitespace. It returns -1 if no value is set otherwise the value
// if it's >= 0.
func parseContentLength(clHeaders []string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// The Content-Length must be a valid numeric value.
// See: https://datatracker.ietf.org/doc/html/rfc2616/#section-14.13

// finishAsyncByteRead finishes reading the 1-byte sniff
// from the ContentLength==0, Body!=nil case.
type finishAsyncByteRead struct {
	tw *transferWriter
}

func (fr finishAsyncByteRead) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var (
	nopCloserType         = reflect.TypeOf(io.NopCloser(nil))
	nopCloserWriterToType = reflect.TypeOf(io.NopCloser(struct {
		io.Reader
		io.WriterTo
	}{}))
)

// unwrapNopCloser return the underlying reader and true if r is a NopCloser
// else it return false.
func unwrapNopCloser(r io.Reader) (underlyingReader io.Reader, isNopCloser bool) {
	_ = "STUB: not implemented"
	return *new(io.Reader), false
}

// isKnownInMemoryReader reports whether r is a type known to not
// block on Read. Its caller uses this as an optional optimization to
// send fewer TCP packets.
func isKnownInMemoryReader(r io.Reader) bool { _ = "STUB: not implemented"; return false }

// bufioFlushWriter is an io.Writer wrapper that flushes all writes
// on its wrapped writer if it's a *bufio.Writer.
type bufioFlushWriter struct{ w io.Writer }

func (fw bufioFlushWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
