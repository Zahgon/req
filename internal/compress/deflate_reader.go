package compress

import (
	"io"
)

type DeflateReader struct {
	Body io.ReadCloser // underlying Response.Body
	dr   io.ReadCloser // lazily-initialized deflate reader
	derr error         // sticky error
}

func NewDeflateReader(body io.ReadCloser) *DeflateReader { _ = "STUB: not implemented"; return nil }

func (df *DeflateReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (df *DeflateReader) Close() error { _ = "STUB: not implemented"; return nil }

func (df *DeflateReader) GetUnderlyingBody() io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (df *DeflateReader) SetUnderlyingBody(body io.ReadCloser) { _ = "STUB: not implemented"; return }
