package dump

import (
	"context"
	"io"
	"net/http"
)

// Options controls the dump behavior.
type Options interface {
	Output() io.Writer
	RequestHeaderOutput() io.Writer
	RequestBodyOutput() io.Writer
	ResponseHeaderOutput() io.Writer
	ResponseBodyOutput() io.Writer
	RequestHeader() bool
	RequestBody() bool
	ResponseHeader() bool
	ResponseBody() bool
	Async() bool
	Clone() Options
}

func (d *Dumper) WrapResponseBodyReadCloser(rc io.ReadCloser) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

type dumpResponseBodyReadCloser struct {
	io.ReadCloser
	dump *Dumper
}

func (r *dumpResponseBodyReadCloser) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Dumper) WrapRequestBodyWriteCloser(rc io.WriteCloser) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

type dumpRequestBodyWriteCloser struct {
	io.WriteCloser
	dump *Dumper
}

func (w *dumpRequestBodyWriteCloser) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type dumpRequestHeaderWriter struct {
	w    io.Writer
	dump *Dumper
}

func (w *dumpRequestHeaderWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Dumper) WrapRequestHeaderWriter(w io.Writer) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

type dumpRequestBodyWriter struct {
	w    io.Writer
	dump *Dumper
}

func (w *dumpRequestBodyWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (d *Dumper) WrapRequestBodyWriter(w io.Writer) io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

// GetResponseHeaderDumpers return Dumpers which need dump response header.
func GetResponseHeaderDumpers(ctx context.Context, dump *Dumper) Dumpers {
	_ = "STUB: not implemented"
	return *new(Dumpers)
}

// Dumpers is an array of Dumpper
type Dumpers []*Dumper

// ShouldDump is true if Dumper is not empty.
func (ds Dumpers) ShouldDump() bool { _ = "STUB: not implemented"; return false }

func (ds Dumpers) DumpResponseHeader(p []byte) { _ = "STUB: not implemented"; return }

// Dumper is the dump tool.
type Dumper struct {
	Options
	ch chan *dumpTask
}

type dumpTask struct {
	Data   []byte
	Output io.Writer
}

// NewDumper create a new Dumper.
func NewDumper(opt Options) *Dumper { _ = "STUB: not implemented"; return nil }

func (d *Dumper) SetOptions(opt Options) { _ = "STUB: not implemented"; return }

func (d *Dumper) Clone() *Dumper { _ = "STUB: not implemented"; return nil }

func (d *Dumper) DumpTo(p []byte, output io.Writer) { _ = "STUB: not implemented"; return }

func (d *Dumper) DumpDefault(p []byte) { _ = "STUB: not implemented"; return }

func (d *Dumper) DumpRequestHeader(p []byte) { _ = "STUB: not implemented"; return }

func (d *Dumper) DumpRequestBody(p []byte) { _ = "STUB: not implemented"; return }

func (d *Dumper) DumpResponseHeader(p []byte) { _ = "STUB: not implemented"; return }

func (d *Dumper) DumpResponseBody(p []byte) { _ = "STUB: not implemented"; return }

func (d *Dumper) Stop() { _ = "STUB: not implemented"; return }

func (d *Dumper) Start() { _ = "STUB: not implemented"; return }

type dumperKeyType int

const DumperKey dumperKeyType = iota

func GetDumpers(ctx context.Context, dump *Dumper) []*Dumper { _ = "STUB: not implemented"; return nil }

func WrapResponseBodyIfNeeded(res *http.Response, req *http.Request, dump *Dumper) {
	_ = "STUB: not implemented"
	return
}
