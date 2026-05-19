// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package req

import (
	"errors"
	"net/http"
	"strings"
	"syscall/js"
)

var uint8Array = js.Global().Get("Uint8Array")

// jsFetchMode is a Request.Header map key that, if present,
// signals that the map entry is actually an option to the Fetch API mode setting.
// Valid values are: "cors", "no-cors", "same-origin", "navigate"
// The default is "same-origin".
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters
const jsFetchMode = "js.fetch:mode"

// jsFetchCreds is a Request.Header map key that, if present,
// signals that the map entry is actually an option to the Fetch API credentials setting.
// Valid values are: "omit", "same-origin", "include"
// The default is "same-origin".
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters
const jsFetchCreds = "js.fetch:credentials"

// jsFetchRedirect is a Request.Header map key that, if present,
// signals that the map entry is actually an option to the Fetch API redirect setting.
// Valid values are: "follow", "error", "manual"
// The default is "follow".
//
// Reference: https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch#Parameters
const jsFetchRedirect = "js.fetch:redirect"

// jsFetchMissing will be true if the Fetch API is not present in
// the browser globals.
var jsFetchMissing = js.Global().Get("fetch").IsUndefined()

// jsFetchDisabled controls whether the use of Fetch API is disabled.
// It's set to true when we detect we're running in Node.js, so that
// RoundTrip ends up talking over the same fake network the HTTP servers
// currently use in various tests and examples. See go.dev/issue/57613.
//
// TODO(go.dev/issue/60810): See if it's viable to test the Fetch API
// code path.
var jsFetchDisabled = js.Global().Get("process").Type() == js.TypeObject &&
	strings.HasPrefix(js.Global().Get("process").Get("argv0").String(), "node")

// RoundTrip implements the [RoundTripper] interface using the WHATWG Fetch API.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	// The Transport has a documented contract that states that if the DialContext or
	// DialTLSContext functions are set, they will be used to set up the connections.
	// If they aren't set then the documented contract is to use Dial or DialTLS, even
	// though they are deprecated. Therefore, if any of these are set, we should obey
	// the contract and dial using the regular round-trip instead. Otherwise, we'll try
	// to fall back on the Fetch API, unless it's not available.
	return nil, nil
}

// Some browsers that support WASM don't necessarily support
// the AbortController. See
// https://developer.mozilla.org/en-US/docs/Web/API/AbortController#Browser_compatibility.

// See https://developer.mozilla.org/en-US/docs/Web/API/WindowOrWorkerGlobalScope/fetch
// for options available.

// TODO(johanbrandhorst): Stream request body when possible.
// See https://bugs.chromium.org/p/chromium/issues/detail?id=688906 for Blink issue.
// See https://bugzilla.mozilla.org/show_bug.cgi?id=1387483 for Firefox issue.
// See https://github.com/web-platform-tests/wpt/issues/7693 for WHATWG tests issue.
// See https://developer.mozilla.org/en-US/docs/Web/API/Streams_API for more details on the Streams API
// and browser support.
// NOTE(haruyama480): Ensure HTTP/1 fallback exists.
// See https://go.dev/issue/61889 for discussion.

// RoundTrip must always close the body, including on errors.

// https://developer.mozilla.org/en-US/docs/Web/API/Headers/entries

// Content-Length values less than 0 are invalid.
// See: https://datatracker.ietf.org/doc/html/rfc2616/#section-14.13

// If the response length is not declared, set it to -1.

// The body is undefined when the browser does not support streaming response bodies (Firefox),
// and null in certain error cases, i.e. when the request is blocked because of CORS settings.

// Fall back to using ArrayBuffer
// https://developer.mozilla.org/en-US/docs/Web/API/Body/arrayBuffer

// The fetch api will decode the gzip, but Content-Encoding not be deleted.

// The error is a JS Error type
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Error
// We can use the toString() method to get a string representation of the error.

// Errors can optionally contain a cause.

// The exact type of the cause is not defined,
// but if it's another error, we can call toString() on it too.

// Abort the Fetch request.

var errClosed = errors.New("net/http: reader is closed")

// streamReader implements an io.ReadCloser wrapper for ReadableStream.
// See https://fetch.spec.whatwg.org/#readablestream for more information.
type streamReader struct {
	pending []byte
	stream  js.Value
	err     error // sticky read error
}

func (r *streamReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Assumes it's a TypeError. See
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/TypeError
// for more information on this type. See
// https://streams.spec.whatwg.org/#byob-reader-read for the spec on
// the read method.

func (r *streamReader) Close() error {
	_ = "STUB: not implemented"
	// This ignores any error returned from cancel method. So far, I did not encounter any concrete
	// situation where reporting the error is meaningful. Most users ignore error from resp.Body.Close().
	// If there's a need to report error here, it can be implemented and tested when that need comes up.
	return nil
}

// arrayReader implements an io.ReadCloser wrapper for ArrayBuffer.
// https://developer.mozilla.org/en-US/docs/Web/API/Body/arrayBuffer.
type arrayReader struct {
	arrayPromise js.Value
	pending      []byte
	read         bool
	err          error // sticky read error
}

func (r *arrayReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Wrap the input ArrayBuffer with a Uint8Array

// Assumes it's a TypeError. See
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/TypeError
// for more information on this type.
// See https://fetch.spec.whatwg.org/#concept-body-consume-body for reasons this might error.

func (r *arrayReader) Close() error { _ = "STUB: not implemented"; return nil }
