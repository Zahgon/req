package compress

import (
	"compress/gzip"
	"io"
)

// GzipReader wraps a response body so it can lazily
// call gzip.NewReader on the first call to Read
type GzipReader struct {
	Body io.ReadCloser // underlying Response.Body
	zr   *gzip.Reader  // lazily-initialized gzip reader
	zerr error         // sticky error
}

func NewGzipReader(body io.ReadCloser) *GzipReader { _ = "STUB: not implemented"; return nil }

func (gz *GzipReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (gz *GzipReader) Close() error { _ = "STUB: not implemented"; return nil }

func (gz *GzipReader) GetUnderlyingBody() io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (gz *GzipReader) SetUnderlyingBody(body io.ReadCloser) { _ = "STUB: not implemented"; return }
