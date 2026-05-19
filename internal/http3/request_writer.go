package http3

import (
	"bytes"
	"io"
	"net/http"
	"sync"

	"github.com/imroc/req/v3/internal/dump"
	"github.com/quic-go/qpack"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3/qlog"
	"github.com/quic-go/quic-go/qlogwriter"
)

const bodyCopyBufferSize = 8 * 1024

type requestWriter struct {
	mutex     sync.Mutex
	encoder   *qpack.Encoder
	headerBuf *bytes.Buffer
}

func newRequestWriter() *requestWriter { _ = "STUB: not implemented"; return nil }

func (w *requestWriter) WriteRequestHeader(wr io.Writer, req *http.Request, gzip bool, streamID quic.StreamID, qlogger qlogwriter.Recorder, dumps []*dump.Dumper) error {
	_ = "STUB: not implemented"
	// TODO: figure out how to add support for trailers
	return nil
}

func (w *requestWriter) writeHeaders(wr io.Writer, req *http.Request, gzip bool, streamID quic.StreamID, qlogger qlogwriter.Recorder, dumps []*dump.Dumper) error {
	_ = "STUB: not implemented"
	return nil
}

func isExtendedConnectRequest(req *http.Request) bool { _ = "STUB: not implemented"; return false }

// copied from net/transport.go
// Modified to support Extended CONNECT:
// Contrary to what the godoc for the http.Request says,
// we do respect the Proto field if the method is CONNECT.
//
// The returned header fields are only set if doQlog is true.
func (w *requestWriter) encodeHeaders(req *http.Request, addGzipHeader bool, trailers string, contentLength int64, doQlog bool, dumps []*dump.Dumper) ([]qlog.HeaderField, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// http.NewRequest sets this field to HTTP/1.1

// Check for any invalid headers and return an error before we
// potentially pollute our hpack state. (We want to be able to
// continue to reuse the hpack encoder for future requests)

// 8.1.2.3 Request Pseudo-Header Fields
// The :path pseudo-header field includes the path and query parts of the
// target URI (the path-absolute production and optionally a '?' character
// followed by the query production (see Sections 3.3 and 3.4 of
// [RFC3986]).

// Match Go's http1 behavior: at most one
// User-Agent. If set to nil or empty string,
// then omit it. Otherwise if not mentioned,
// include the default (below).

// Do a first pass over the headers counting bytes to ensure
// we don't exceed cc.peerMaxHeaderListSize. This is done as a
// separate pass before encoding the headers to prevent
// modifying the hpack state.

// TODO: check maximum header list size
// if hlSize > cc.peerMaxHeaderListSize {
// 	return errRequestHeaderListSize
// }

// Header list size is ok. Write the headers.

// Header list size is ok. Write the headers.

// authorityAddr returns a given authority (a host/IP, or host:port / ip:port)
// and returns a host:port. The port 443 is added if needed.
func authorityAddr(authority string) (addr string) { _ = "STUB: not implemented"; return "" }

// authority didn't have a port

// IPv6 address literal, without a port:

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

// actualContentLength returns a sanitized version of
// req.ContentLength, where 0 actually means zero (not unknown) and -1
// means unknown.
func actualContentLength(req *http.Request) int64 { _ = "STUB: not implemented"; return 0 }

// shouldSendReqContentLength reports whether the http2.Transport should send
// a "content-length" request header. This logic is basically a copy of the net/http
// transferWriter.shouldSendContentLength.
// The contentLength is the corrected contentLength (so 0 means actually 0, not unknown).
// -1 means unknown.
func shouldSendReqContentLength(method string, contentLength int64) bool {
	_ = "STUB: not implemented"
	return false
}

// For zero bodies, whether we send a content-length depends on the method.
// It also kinda doesn't matter for http2 either way, with END_STREAM.
