package compress

import (
	"io"
)

type BrotliReader struct {
	Body io.ReadCloser // underlying Response.Body
	br   io.Reader     // lazily-initialized brotli reader
	berr error         // sticky error
}

func NewBrotliReader(body io.ReadCloser) *BrotliReader { _ = "STUB: not implemented"; return nil }

func (br *BrotliReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (br *BrotliReader) Close() error { _ = "STUB: not implemented"; return nil }

func (br *BrotliReader) GetUnderlyingBody() io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (br *BrotliReader) SetUnderlyingBody(body io.ReadCloser) { _ = "STUB: not implemented"; return }
