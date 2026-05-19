package http3

import (
	"errors"
	"io"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/qlogwriter"
	"github.com/quic-go/quic-go/quicvarint"
)

// FrameType is the frame type of a HTTP/3 frame
type FrameType uint64

type unknownFrameHandlerFunc func(FrameType, error) (processed bool, err error)

type frame any

var errHijacked = errors.New("hijacked")

type countingByteReader struct {
	quicvarint.Reader
	NumRead int
}

func (r *countingByteReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *countingByteReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *countingByteReader) Reset() { _ = "STUB: not implemented"; return }

type frameParser struct {
	r                   io.Reader
	streamID            quic.StreamID
	closeConn           func(quic.ApplicationErrorCode, string) error
	unknownFrameHandler unknownFrameHandlerFunc
}

func (p *frameParser) ParseNext(qlogger qlogwriter.Recorder) (frame, error) {
	_ = "STUB: not implemented"
	return *new(frame), nil
}

// Call the unknownFrameHandler for frames not defined in the HTTP/3 spec

// If the unknownFrameHandler didn't process the frame, it is our responsibility to skip it.

// DATA

// HEADERS

// SETTINGS

// unsupported: CANCEL_PUSH

// unsupported: PUSH_PROMISE

// GOAWAY

// unsupported: MAX_PUSH_ID

// reserved frame types

// unknown frame types

// skip over the payload

type dataFrame struct {
	Length uint64
}

func (f *dataFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }

type headersFrame struct {
	Length    uint64
	headerLen int // number of bytes read for type and length field
}

func (f *headersFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }

const (
	// SETTINGS_MAX_FIELD_SECTION_SIZE
	settingMaxFieldSectionSize = 0x6
	// Extended CONNECT, RFC 9220
	settingExtendedConnect = 0x8
	// HTTP Datagrams, RFC 9297
	settingDatagram = 0x33
)

type settingsFrame struct {
	MaxFieldSectionSize int64 // SETTINGS_MAX_FIELD_SECTION_SIZE, -1 if not set

	Datagram        bool              // HTTP Datagrams, RFC 9297
	ExtendedConnect bool              // Extended CONNECT, RFC 9220
	Other           map[uint64]uint64 // all settings that we don't explicitly recognize
}

func pointer[T any](v T) *T { _ = "STUB: not implemented"; return nil }

func parseSettingsFrame(r *countingByteReader, l uint64, streamID quic.StreamID, qlogger qlogwriter.Recorder) (*settingsFrame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// should not happen. We allocated the whole frame already.

// should not happen. We allocated the whole frame already.

func (f *settingsFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }

type goAwayFrame struct {
	StreamID quic.StreamID
}

func parseGoAwayFrame(r *countingByteReader, l uint64, streamID quic.StreamID, qlogger qlogwriter.Recorder) (*goAwayFrame, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *goAwayFrame) Append(b []byte) []byte { _ = "STUB: not implemented"; return nil }
