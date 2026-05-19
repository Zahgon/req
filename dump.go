package req

import (
	"io"

	"github.com/imroc/req/v3/internal/dump"
)

// DumpOptions controls the dump behavior.
type DumpOptions struct {
	Output               io.Writer
	RequestOutput        io.Writer
	ResponseOutput       io.Writer
	RequestHeaderOutput  io.Writer
	RequestBodyOutput    io.Writer
	ResponseHeaderOutput io.Writer
	ResponseBodyOutput   io.Writer
	RequestHeader        bool
	RequestBody          bool
	ResponseHeader       bool
	ResponseBody         bool
	Async                bool
}

// Clone return a copy of DumpOptions
func (do *DumpOptions) Clone() *DumpOptions { _ = "STUB: not implemented"; return nil }

type dumpOptions struct {
	*DumpOptions
}

func (o dumpOptions) Output() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (o dumpOptions) RequestHeaderOutput() io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (o dumpOptions) RequestBodyOutput() io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (o dumpOptions) ResponseHeaderOutput() io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (o dumpOptions) ResponseBodyOutput() io.Writer {
	_ = "STUB: not implemented"
	return *new(io.Writer)
}

func (o dumpOptions) RequestHeader() bool { _ = "STUB: not implemented"; return false }

func (o dumpOptions) RequestBody() bool { _ = "STUB: not implemented"; return false }

func (o dumpOptions) ResponseHeader() bool { _ = "STUB: not implemented"; return false }

func (o dumpOptions) ResponseBody() bool { _ = "STUB: not implemented"; return false }

func (o dumpOptions) Async() bool { _ = "STUB: not implemented"; return false }

func (o dumpOptions) Clone() dump.Options { _ = "STUB: not implemented"; return *new(dump.Options) }

func newDefaultDumpOptions() *DumpOptions { _ = "STUB: not implemented"; return nil }

func newDumper(opt *DumpOptions) *dump.Dumper { _ = "STUB: not implemented"; return nil }
