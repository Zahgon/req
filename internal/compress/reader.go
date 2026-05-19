package compress

import "io"

type CompressReader interface {
	io.ReadCloser
	GetUnderlyingBody() io.ReadCloser
	SetUnderlyingBody(body io.ReadCloser)
}

func NewCompressReader(body io.ReadCloser, contentEncoding string) CompressReader {
	_ = "STUB: not implemented"
	return *new(CompressReader)
}
