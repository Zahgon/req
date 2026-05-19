package http3

import (
	"errors"
	"net/http"

	"github.com/imroc/req/v3/internal/dump"
	"github.com/quic-go/qpack"
)

type qpackError struct{ err error }

func (e *qpackError) Error() string { _ = "STUB: not implemented"; return "" }
func (e *qpackError) Unwrap() error { _ = "STUB: not implemented"; return nil }

var errHeaderTooLarge = errors.New("http3: headers too large")

type header struct {
	// Pseudo header fields defined in RFC 9114
	Path      string
	Method    string
	Authority string
	Scheme    string
	Status    string
	// for Extended connect
	Protocol string
	// parsed and deduplicated. -1 if no Content-Length header is sent
	ContentLength int64
	// all non-pseudo headers
	Headers http.Header
}

// connection-specific header fields must not be sent on HTTP/3
var invalidHeaderFields = [...]string{
	"connection",
	"keep-alive",
	"proxy-connection",
	"transfer-encoding",
	"upgrade",
}

func parseHeaders(decodeFn qpack.DecodeFunc, isRequest bool, sizeLimit int, headerFields *[]qpack.HeaderField, ds dump.Dumpers) (header, error) {
	_ = "STUB: not implemented"
	return *new(header), nil
}

// RFC 9114, section 4.2.2:
// The size of a field list is calculated based on the uncompressed size of fields,
// including the length of the name and value in bytes plus an overhead of 32 bytes for each field.

// field names need to be lowercase, see section 4.2 of RFC 9114

// all pseudo headers must appear before regular header fields, see section 4.3 of RFC 9114

// pseudo headers are either valid for requests or for responses
// pseudo headers are allowed to appear exactly once

// Ignore duplicate Content-Length headers.
// Fail if the duplicates differ.

// use ParseUint instead of ParseInt, so that parsing fails on negative values

func parseTrailers(decodeFn qpack.DecodeFunc, headerFields *[]qpack.HeaderField) (http.Header, error) {
	_ = "STUB: not implemented"
	return *new(http.Header), nil
}

// updateResponseFromHeaders sets up http.Response as an HTTP/3 response,
// using the decoded qpack header filed.
// It is only called for the HTTP header (and not the HTTP trailer).
// It takes an http.Response as an argument to allow the caller to set the trailer later on.
func updateResponseFromHeaders(rsp *http.Response, decodeFn qpack.DecodeFunc, sizeLimit int, headerFields *[]qpack.HeaderField, ds dump.Dumpers) error {
	_ = "STUB: not implemented"
	return nil
}

// processTrailers initializes the rsp.Trailer map, and adds keys for every announced header value.
// The Trailer header is removed from the http.Response.Header map.
// It handles both duplicate as well as comma-separated values for the Trailer header.
// For example:
//
//	Trailer: Trailer1, Trailer2
//	Trailer: Trailer3
//
// Will result in a http.Response.Trailer map containing the keys "Trailer1", "Trailer2", "Trailer3".
func processTrailers(rsp *http.Response) { _ = "STUB: not implemented"; return }
