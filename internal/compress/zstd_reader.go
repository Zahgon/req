package compress

import (
	"io"

	"github.com/klauspost/compress/zstd"
)

type ZstdReader struct {
	Body io.ReadCloser // underlying Response.Body
	zr   *zstd.Decoder // lazily-initialized zstd reader
	zerr error         // sticky error
}

func NewZstdReader(body io.ReadCloser) *ZstdReader { _ = "STUB: not implemented"; return nil }

func (zr *ZstdReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (zr *ZstdReader) Close() error { _ = "STUB: not implemented"; return nil }

func (zr *ZstdReader) GetUnderlyingBody() io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (zr *ZstdReader) SetUnderlyingBody(body io.ReadCloser) { _ = "STUB: not implemented"; return }
