package req

import (
	"io"
)

var textContentTypes = []string{"text", "json", "xml", "html", "java"}

var autoDecodeText = autoDecodeContentTypeFunc(textContentTypes...)

func autoDecodeContentTypeFunc(contentTypes ...string) func(contentType string) bool {
	_ = "STUB: not implemented"
	return nil
}

type decodeReaderCloser struct {
	io.ReadCloser
	decodeReader io.Reader
}

func (d *decodeReaderCloser) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func newAutoDecodeReadCloser(input io.ReadCloser, t *Transport) *autoDecodeReadCloser {
	_ = "STUB: not implemented"
	return nil
}

type autoDecodeReadCloser struct {
	io.ReadCloser
	t            *Transport
	decodeReader io.Reader
	detected     bool
	peek         []byte
}

func (a *autoDecodeReadCloser) peekRead(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *autoDecodeReadCloser) peekDrain(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (a *autoDecodeReadCloser) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// can not determine charset, not decode
