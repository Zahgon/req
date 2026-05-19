// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"sync"

	"github.com/imroc/req/v3/http2"
	"github.com/imroc/req/v3/internal/dump"
	"golang.org/x/net/http2/hpack"
)

const frameHeaderLen = 9

var padZeros = make([]byte, 255) // zeros for padding

// A FrameType is a registered frame type as defined in
// https://httpwg.org/specs/rfc7540.html#rfc.section.11.2
type FrameType uint8

const (
	FrameData         FrameType = 0x0
	FrameHeaders      FrameType = 0x1
	FramePriority     FrameType = 0x2
	FrameRSTStream    FrameType = 0x3
	FrameSettings     FrameType = 0x4
	FramePushPromise  FrameType = 0x5
	FramePing         FrameType = 0x6
	FrameGoAway       FrameType = 0x7
	FrameWindowUpdate FrameType = 0x8
	FrameContinuation FrameType = 0x9
)

var frameName = map[FrameType]string{
	FrameData:         "DATA",
	FrameHeaders:      "HEADERS",
	FramePriority:     "PRIORITY",
	FrameRSTStream:    "RST_STREAM",
	FrameSettings:     "SETTINGS",
	FramePushPromise:  "PUSH_PROMISE",
	FramePing:         "PING",
	FrameGoAway:       "GOAWAY",
	FrameWindowUpdate: "WINDOW_UPDATE",
	FrameContinuation: "CONTINUATION",
}

func (t FrameType) String() string { _ = "STUB: not implemented"; return "" }

// Flags is a bitmask of HTTP/2 flags.
// The meaning of flags varies depending on the frame type.
type Flags uint8

// Has reports whether f contains all (0 or more) flags in v.
func (f Flags) Has(v Flags) bool { _ = "STUB: not implemented"; return false }

// Frame-specific FrameHeader flag bits.
const (
	// Data Frame
	FlagDataEndStream Flags = 0x1
	FlagDataPadded    Flags = 0x8

	// Headers Frame
	FlagHeadersEndStream  Flags = 0x1
	FlagHeadersEndHeaders Flags = 0x4
	FlagHeadersPadded     Flags = 0x8
	FlagHeadersPriority   Flags = 0x20

	// Settings Frame
	FlagSettingsAck Flags = 0x1

	// Ping Frame
	FlagPingAck Flags = 0x1

	// Continuation Frame
	FlagContinuationEndHeaders Flags = 0x4

	FlagPushPromiseEndHeaders Flags = 0x4
	FlagPushPromisePadded     Flags = 0x8
)

var flagName = map[FrameType]map[Flags]string{
	FrameData: {
		FlagDataEndStream: "END_STREAM",
		FlagDataPadded:    "PADDED",
	},
	FrameHeaders: {
		FlagHeadersEndStream:  "END_STREAM",
		FlagHeadersEndHeaders: "END_HEADERS",
		FlagHeadersPadded:     "PADDED",
		FlagHeadersPriority:   "PRIORITY",
	},
	FrameSettings: {
		FlagSettingsAck: "ACK",
	},
	FramePing: {
		FlagPingAck: "ACK",
	},
	FrameContinuation: {
		FlagContinuationEndHeaders: "END_HEADERS",
	},
	FramePushPromise: {
		FlagPushPromiseEndHeaders: "END_HEADERS",
		FlagPushPromisePadded:     "PADDED",
	},
}

// a frameParser parses a frame given its FrameHeader and payload
// bytes. The length of payload will always equal fh.Length (which
// might be 0).
type frameParser func(fc *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error)

var frameParsers = map[FrameType]frameParser{
	FrameData:         parseDataFrame,
	FrameHeaders:      parseHeadersFrame,
	FramePriority:     parsePriorityFrame,
	FrameRSTStream:    parseRSTStreamFrame,
	FrameSettings:     parseSettingsFrame,
	FramePushPromise:  parsePushPromise,
	FramePing:         parsePingFrame,
	FrameGoAway:       parseGoAwayFrame,
	FrameWindowUpdate: parseWindowUpdateFrame,
	FrameContinuation: parseContinuationFrame,
}

func typeFrameParser(t FrameType) frameParser { _ = "STUB: not implemented"; return *new(frameParser) }

// A FrameHeader is the 9 byte header of all HTTP/2 frames.
//
// See https://httpwg.org/specs/rfc7540.html#FrameHeader
type FrameHeader struct {
	valid bool // caller can access []byte fields in the Frame

	// Type is the 1 byte frame type. There are ten standard frame
	// types, but extension frame types may be written by WriteRawFrame
	// and will be returned by ReadFrame (as UnknownFrame).
	Type FrameType

	// Flags are the 1 byte of 8 potential bit flags per frame.
	// They are specific to the frame type.
	Flags Flags

	// Length is the length of the frame, not including the 9 byte header.
	// The maximum size is one byte less than 16MB (uint24), but only
	// frames up to 16KB are allowed without peer agreement.
	Length uint32

	// StreamID is which stream this frame is for. Certain frames
	// are not stream-specific, in which case this field is 0.
	StreamID uint32
}

// Header returns h. It exists so FrameHeaders can be embedded in other
// specific frame types and implement the Frame interface.
func (h FrameHeader) Header() FrameHeader { _ = "STUB: not implemented"; return *new(FrameHeader) }

func (h FrameHeader) String() string { _ = "STUB: not implemented"; return "" }

func (h FrameHeader) writeDebug(buf *bytes.Buffer) { _ = "STUB: not implemented"; return }

func (h *FrameHeader) checkValid() { _ = "STUB: not implemented"; return }

func (h *FrameHeader) invalidate() {
	_ = "STUB: not implemented"

	// frame header bytes.
	// Used only by ReadFrameHeader.
	return
}

var fhBytes = sync.Pool{
	New: func() any {
		buf := make([]byte, frameHeaderLen)
		return &buf
	},
}

// ReadFrameHeader reads 9 bytes from r and returns a FrameHeader.
// Most users should use Framer.ReadFrame instead.
func ReadFrameHeader(r io.Reader) (FrameHeader, error) {
	_ = "STUB: not implemented"
	return *new(FrameHeader), nil
}

func readFrameHeader(buf []byte, r io.Reader) (FrameHeader, error) {
	_ = "STUB: not implemented"
	return *new(FrameHeader), nil
}

// A Frame is the base interface implemented by all frame types.
// Callers will generally type-assert the specific frame type:
// *HeadersFrame, *SettingsFrame, *WindowUpdateFrame, etc.
//
// Frames are only valid until the next call to Framer.ReadFrame.
type Frame interface {
	Header() FrameHeader

	// invalidate is called by Framer.ReadFrame to make this
	// frame's buffers as being invalid, since the subsequent
	// frame will reuse them.
	invalidate()
}

// A Framer reads and writes Frames.
type Framer struct {
	cc        *ClientConn
	r         io.Reader
	lastFrame Frame
	errDetail error

	// countError is a non-nil func that's called on a frame parse
	// error with some unique error path token. It's initialized
	// from Transport.CountError or Server.CountError.
	countError func(errToken string)

	// lastHeaderStream is non-zero if the last frame was an
	// unfinished HEADERS/CONTINUATION.
	lastHeaderStream uint32

	maxReadSize uint32
	headerBuf   [frameHeaderLen]byte

	// TODO: let getReadBuf be configurable, and use a less memory-pinning
	// allocator in server.go to minimize memory pinned for many idle conns.
	// Will probably also need to make frame invalidation have a hook too.
	getReadBuf func(size uint32) []byte
	readBuf    []byte // cache for default getReadBuf

	maxWriteSize uint32 // zero means unlimited; TODO: implement

	w    io.Writer
	wbuf []byte

	// AllowIllegalWrites permits the Framer's Write methods to
	// write frames that do not conform to the HTTP/2 spec. This
	// permits using the Framer to test other HTTP/2
	// implementations' conformance to the spec.
	// If false, the Write methods will prefer to return an error
	// rather than comply.
	AllowIllegalWrites bool

	// AllowIllegalReads permits the Framer's ReadFrame method
	// to return non-compliant frames or frame orders.
	// This is for testing and permits using the Framer to test
	// other HTTP/2 implementations' conformance to the spec.
	// It is not compatible with ReadMetaHeaders.
	AllowIllegalReads bool

	// ReadMetaHeaders if non-nil causes ReadFrame to merge
	// HEADERS and CONTINUATION frames together and return
	// MetaHeadersFrame instead.
	ReadMetaHeaders *hpack.Decoder

	// MaxHeaderListSize is the http2 MAX_HEADER_LIST_SIZE.
	// It's used only if ReadMetaHeaders is set; 0 means a sane default
	// (currently 16MB)
	// If the limit is hit, MetaHeadersFrame.Truncated is set true.
	MaxHeaderListSize uint32

	// TODO: track which type of frame & with which flags was sent
	// last. Then return an error (unless AllowIllegalWrites) if
	// we're in the middle of a header block and a
	// non-Continuation or Continuation on a different stream is
	// attempted to be written.

	logReads, logWrites bool

	debugFramer       *Framer // only use for logging written writes
	debugFramerBuf    *bytes.Buffer
	debugReadLoggerf  func(string, ...any)
	debugWriteLoggerf func(string, ...any)

	frameCache *frameCache // nil if frames aren't reused (default)
}

func (h2f *Framer) maxHeaderListSize() uint32 { _ = "STUB: not implemented"; return 0 }

// sane default, per docs

func (h2f *Framer) startWrite(ftype FrameType, flags Flags, streamID uint32) {
	_ = "STUB: not implemented"
	// Write the FrameHeader.
	return
}

// 3 bytes of length, filled in in endWrite

func (h2f *Framer) endWrite() error {
	_ = "STUB: not implemented"
	// Now that we know the final size, fill in the FrameHeader in
	// the space previously reserved for it. Abuse append.
	return nil
}

func (h2f *Framer) logWrite() { _ = "STUB: not implemented"; return }

// we log it ourselves, saying "wrote" below
// Let us read anything, even if we accidentally wrote it
// in the wrong order:

func (h2f *Framer) writeByte(v byte) { _ = "STUB: not implemented"; return }

func (h2f *Framer) writeBytes(v []byte) { _ = "STUB: not implemented"; return }

func (h2f *Framer) writeUint16(v uint16) { _ = "STUB: not implemented"; return }

func (h2f *Framer) writeUint32(v uint32) { _ = "STUB: not implemented"; return }

const (
	minMaxFrameSize = 1 << 14
	maxFrameSize    = 1<<24 - 1
)

// SetReuseFrames allows the Framer to reuse Frames.
// If called on a Framer, Frames returned by calls to ReadFrame are only
// valid until the next call to ReadFrame.
func (h2f *Framer) SetReuseFrames() { _ = "STUB: not implemented"; return }

type frameCache struct {
	dataFrame DataFrame
}

func (fc *frameCache) getDataFrame() *DataFrame { _ = "STUB: not implemented"; return nil }

// NewFramer returns a Framer that writes frames to w and reads them from r.
func NewFramer(w io.Writer, r io.Reader) *Framer { _ = "STUB: not implemented"; return nil }

// SetMaxReadFrameSize sets the maximum size of a frame
// that will be read by a subsequent call to ReadFrame.
// It is the caller's responsibility to advertise this
// limit with a SETTINGS frame.
func (h2f *Framer) SetMaxReadFrameSize(v uint32) { _ = "STUB: not implemented"; return }

// ErrorDetail returns a more detailed error of the last error
// returned by Framer.ReadFrame. For instance, if ReadFrame
// returns a StreamError with code PROTOCOL_ERROR, ErrorDetail
// will say exactly what was invalid. ErrorDetail is not guaranteed
// to return a non-nil value and like the rest of the http2 package,
// its return value is not protected by an API compatibility promise.
// ErrorDetail is reset after the next call to ReadFrame.
func (h2f *Framer) ErrorDetail() error { _ = "STUB: not implemented"; return nil }

// errFrameTooLarge is returned from Framer.ReadFrame when the peer
// sends a frame that is larger than declared with SetMaxReadFrameSize.
var errFrameTooLarge = errors.New("http2: frame too large")

// terminalReadFrameError reports whether err is an unrecoverable
// error from ReadFrame and no other frames should be read.
func terminalReadFrameError(err error) bool { _ = "STUB: not implemented"; return false }

func (h2f *Framer) streamByID(id uint32) *clientStream { _ = "STUB: not implemented"; return nil }

func (h2f *Framer) currentRequest(id uint32) *http.Request { _ = "STUB: not implemented"; return nil }

// ReadFrame reads a single frame. The returned Frame is only valid
// until the next call to ReadFrame.
//
// If the frame is larger than previously set with SetMaxReadFrameSize, the
// returned error is errFrameTooLarge. Other errors may be of type
// ConnectionError, StreamError, or anything else from the underlying
// reader.
//
// If ReadFrame returns an error and a non-nil Frame, the Frame's StreamID
// indicates the stream responsible for the error.
func (h2f *Framer) ReadFrame() (Frame, error) { _ = "STUB: not implemented"; return *new(Frame), nil }

// connError returns ConnectionError(code) but first
// stashes away a public reason to the caller can optionally relay it
// to the peer before hanging up on them. This might help others debug
// their implementations.
func (h2f *Framer) connError(code ErrCode, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

// checkFrameOrder reports an error if f is an invalid frame to return
// next from ReadFrame. Mostly it checks whether HEADERS and
// CONTINUATION frames are contiguous.
func (h2f *Framer) checkFrameOrder(f Frame) error { _ = "STUB: not implemented"; return nil }

// A DataFrame conveys arbitrary, variable-length sequences of octets
// associated with a stream.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.1
type DataFrame struct {
	FrameHeader
	data []byte
}

func (f *DataFrame) StreamEnded() bool { _ = "STUB: not implemented"; return false }

// Data returns the frame's data octets, not including any padding
// size byte or padding suffix bytes.
// The caller must not retain the returned memory past the next
// call to ReadFrame.
func (f *DataFrame) Data() []byte { _ = "STUB: not implemented"; return nil }

func parseDataFrame(fc *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *

	// DATA frames MUST be associated with a stream. If a
	// DATA frame is received whose stream identifier
	// field is 0x0, the recipient MUST respond with a
	// connection error (Section 5.4.1) of type
	// PROTOCOL_ERROR.
	new(Frame), nil
}

// If the length of the padding is greater than the
// length of the frame payload, the recipient MUST
// treat this as a connection error.
// Filed: https://github.com/http2/http2-spec/issues/610

var (
	errStreamID    = errors.New("invalid stream ID")
	errDepStreamID = errors.New("invalid dependent stream ID")
	errPadLength   = errors.New("pad length too large")
	errPadBytes    = errors.New("padding bytes must all be zeros unless AllowIllegalWrites is enabled")
)

func validStreamIDOrZero(streamID uint32) bool { _ = "STUB: not implemented"; return false }

func validStreamID(streamID uint32) bool { _ = "STUB: not implemented"; return false }

// writeData writes a DATA frame.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility not to violate the maximum frame size
// and to not call other Write methods concurrently.
func (h2f *Framer) WriteData(streamID uint32, endStream bool, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteDataPadded writes a DATA frame with optional padding.
//
// If pad is nil, the padding bit is not sent.
// The length of pad must not exceed 255 bytes.
// The bytes of pad must all be zero, unless f.AllowIllegalWrites is set.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility not to violate the maximum frame size
// and to not call other Write methods concurrently.
func (h2f *Framer) WriteDataPadded(streamID uint32, endStream bool, data, pad []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// startWriteDataPadded is WriteDataPadded, but only writes the frame to the Framer's internal buffer.
// The caller should call endWrite to flush the frame to the underlying writer.
func (h2f *Framer) startWriteDataPadded(streamID uint32, endStream bool, data, pad []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// "Padding octets MUST be set to zero when sending."

// A SettingsFrame conveys configuration parameters that affect how
// endpoints communicate, such as preferences and constraints on peer
// behavior.
//
// See https://httpwg.org/specs/rfc7540.html#SETTINGS
type SettingsFrame struct {
	FrameHeader
	p []byte
}

func parseSettingsFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// When this (ACK 0x1) bit is set, the payload of the
// SETTINGS frame MUST be empty. Receipt of a
// SETTINGS frame with the ACK flag set and a length
// field value other than 0 MUST be treated as a
// connection error (Section 5.4.1) of type
// FRAME_SIZE_ERROR.

// SETTINGS frames always apply to a connection,
// never a single stream. The stream identifier for a
// SETTINGS frame MUST be zero (0x0).  If an endpoint
// receives a SETTINGS frame whose stream identifier
// field is anything other than 0x0, the endpoint MUST
// respond with a connection error (Section 5.4.1) of
// type PROTOCOL_ERROR.

// Expecting even number of 6 byte settings.

// Values above the maximum flow control window size of 2^31 - 1 MUST
// be treated as a connection error (Section 5.4.1) of type
// FLOW_CONTROL_ERROR.

func (f *SettingsFrame) IsAck() bool { _ = "STUB: not implemented"; return false }

func (f *SettingsFrame) Value(id http2.SettingID) (v uint32, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// Setting returns the setting from the frame at the given 0-based index.
// The index must be >= 0 and less than f.NumSettings().
func (f *SettingsFrame) Setting(i int) http2.Setting {
	_ = "STUB: not implemented"
	return *new(http2.Setting)
}

func (f *SettingsFrame) NumSettings() int {
	_ = "STUB: not implemented"

	// HasDuplicates reports whether f contains any duplicate setting IDs.
	return 0
}

func (f *SettingsFrame) HasDuplicates() bool { _ = "STUB: not implemented"; return false }

// If it's small enough (the common case), just do the n^2
// thing and avoid a map allocation.

// ForeachSetting runs fn for each setting.
// It stops and returns the first error.
func (f *SettingsFrame) ForeachSetting(fn func(http2.Setting) error) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteSettings writes a SETTINGS frame with zero or more settings
// specified and the ACK bit not set.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility to not call other Write methods concurrently.
func (h2f *Framer) WriteSettings(settings ...http2.Setting) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteSettingsAck writes an empty SETTINGS frame with the ACK bit set.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility to not call other Write methods concurrently.
func (h2f *Framer) WriteSettingsAck() error { _ = "STUB: not implemented"; return nil }

// A PingFrame is a mechanism for measuring a minimal round trip time
// from the sender, as well as determining whether an idle connection
// is still functional.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.7
type PingFrame struct {
	FrameHeader
	Data [8]byte
}

func (f *PingFrame) IsAck() bool { _ = "STUB: not implemented"; return false }

func parsePingFrame(_ *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

func (h2f *Framer) WritePing(ack bool, data [8]byte) error { _ = "STUB: not implemented"; return nil }

// A GoAwayFrame informs the remote peer to stop creating streams on this connection.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.8
type GoAwayFrame struct {
	FrameHeader
	LastStreamID uint32
	ErrCode      ErrCode
	debugData    []byte
}

// DebugData returns any debug data in the GOAWAY frame. Its contents
// are not defined.
// The caller must not retain the returned memory past the next
// call to ReadFrame.
func (f *GoAwayFrame) DebugData() []byte { _ = "STUB: not implemented"; return nil }

func parseGoAwayFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

func (h2f *Framer) WriteGoAway(maxStreamID uint32, code ErrCode, debugData []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// An UnknownFrame is the frame type returned when the frame type is unknown
// or no specific frame type parser exists.
type UnknownFrame struct {
	FrameHeader
	p []byte
}

// Payload returns the frame's payload (after the header).  It is not
// valid to call this method after a subsequent call to
// Framer.ReadFrame, nor is it valid to retain the returned slice.
// The memory is owned by the Framer and is invalidated when the next
// frame is read.
func (f *UnknownFrame) Payload() []byte { _ = "STUB: not implemented"; return nil }

func parseUnknownFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// A WindowUpdateFrame is used to implement flow control.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.9
type WindowUpdateFrame struct {
	FrameHeader
	Increment uint32 // never read with high bit set
}

func parseWindowUpdateFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// mask off high reserved bit

// A receiver MUST treat the receipt of a
// WINDOW_UPDATE frame with an flow control window
// increment of 0 as a stream error (Section 5.4.2) of
// type PROTOCOL_ERROR; errors on the connection flow
// control window MUST be treated as a connection
// error (Section 5.4.1).

// WriteWindowUpdate writes a WINDOW_UPDATE frame.
// The increment value must be between 1 and 2,147,483,647, inclusive.
// If the Stream ID is zero, the window update applies to the
// connection as a whole.
func (h2f *Framer) WriteWindowUpdate(streamID, incr uint32) error {
	_ = "STUB: not implemented"
	// "The legal range for the increment to the flow control window is 1 to 2^31-1 (2,147,483,647) octets."
	return nil
}

// A HeadersFrame is used to open a stream and additionally carries a
// header block fragment.
type HeadersFrame struct {
	FrameHeader

	// Priority is set if FlagHeadersPriority is set in the FrameHeader.
	Priority http2.PriorityParam

	headerFragBuf []byte // not owned
}

func (f *HeadersFrame) HeaderBlockFragment() []byte { _ = "STUB: not implemented"; return nil }

func (f *HeadersFrame) HeadersEnded() bool { _ = "STUB: not implemented"; return false }

func (f *HeadersFrame) StreamEnded() bool { _ = "STUB: not implemented"; return false }

func (f *HeadersFrame) HasPriority() bool { _ = "STUB: not implemented"; return false }

func parseHeadersFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (_ Frame, err error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// HEADERS frames MUST be associated with a stream. If a HEADERS frame
// is received whose stream identifier field is 0x0, the recipient MUST
// respond with a connection error (Section 5.4.1) of type
// PROTOCOL_ERROR.

// high bit was set

// HeadersFrameParam are the parameters for writing a HEADERS frame.
type HeadersFrameParam struct {
	// StreamID is the required Stream ID to initiate.
	StreamID uint32
	// BlockFragment is part (or all) of a Header Block.
	BlockFragment []byte

	// EndStream indicates that the header block is the last that
	// the endpoint will send for the identified stream. Setting
	// this flag causes the stream to enter one of "half closed"
	// states.
	EndStream bool

	// EndHeaders indicates that this frame contains an entire
	// header block and is not followed by any
	// CONTINUATION frames.
	EndHeaders bool

	// PadLength is the optional number of bytes of zeros to add
	// to this frame.
	PadLength uint8

	// Priority, if non-zero, includes stream priority information
	// in the HEADER frame.
	Priority http2.PriorityParam
}

// WriteHeaders writes a single HEADERS frame.
//
// This is a low-level header writing method. Encoding headers and
// splitting them into any necessary CONTINUATION frames is handled
// elsewhere.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility to not call other Write methods concurrently.
func (h2f *Framer) WriteHeaders(p HeadersFrameParam) error { _ = "STUB: not implemented"; return nil }

// A PriorityFrame specifies the sender-advised priority of a stream.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.3
type PriorityFrame struct {
	FrameHeader
	http2.PriorityParam
}

func parsePriorityFrame(_ *frameCache, fh FrameHeader, countError func(string), payload []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// mask off high bit

// was high bit set?

// WritePriority writes a PRIORITY frame.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility to not call other Write methods concurrently.
func (h2f *Framer) WritePriority(streamID uint32, p http2.PriorityParam) error {
	_ = "STUB: not implemented"
	return nil
}

// A RSTStreamFrame allows for abnormal termination of a stream.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.4
type RSTStreamFrame struct {
	FrameHeader
	ErrCode ErrCode
}

func parseRSTStreamFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// WriteRSTStream writes a RST_STREAM frame.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility to not call other Write methods concurrently.
func (h2f *Framer) WriteRSTStream(streamID uint32, code ErrCode) error {
	_ = "STUB: not implemented"
	return nil
}

// A ContinuationFrame is used to continue a sequence of header block fragments.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.10
type ContinuationFrame struct {
	FrameHeader
	headerFragBuf []byte
}

func parseContinuationFrame(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

func (f *ContinuationFrame) HeaderBlockFragment() []byte { _ = "STUB: not implemented"; return nil }

func (f *ContinuationFrame) HeadersEnded() bool { _ = "STUB: not implemented"; return false }

// WriteContinuation writes a CONTINUATION frame.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility to not call other Write methods concurrently.
func (h2f *Framer) WriteContinuation(streamID uint32, endHeaders bool, headerBlockFragment []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// A PushPromiseFrame is used to initiate a server stream.
// See https://httpwg.org/specs/rfc7540.html#rfc.section.6.6
type PushPromiseFrame struct {
	FrameHeader
	PromiseID     uint32
	headerFragBuf []byte // not owned
}

func (f *PushPromiseFrame) HeaderBlockFragment() []byte { _ = "STUB: not implemented"; return nil }

func (f *PushPromiseFrame) HeadersEnded() bool { _ = "STUB: not implemented"; return false }

func parsePushPromise(_ *frameCache, fh FrameHeader, countError func(string), p []byte) (_ Frame, err error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// PUSH_PROMISE frames MUST be associated with an existing,
// peer-initiated stream. The stream identifier of a
// PUSH_PROMISE frame indicates the stream it is associated
// with. If the stream identifier field specifies the value
// 0x0, a recipient MUST respond with a connection error
// (Section 5.4.1) of type PROTOCOL_ERROR.

// The PUSH_PROMISE frame includes optional padding.
// Padding fields and flags are identical to those defined for DATA frames

// like the DATA frame, error out if padding is longer than the body.

// PushPromiseParam are the parameters for writing a PUSH_PROMISE frame.
type PushPromiseParam struct {
	// StreamID is the required Stream ID to initiate.
	StreamID uint32

	// PromiseID is the required Stream ID which this
	// Push Promises
	PromiseID uint32

	// BlockFragment is part (or all) of a Header Block.
	BlockFragment []byte

	// EndHeaders indicates that this frame contains an entire
	// header block and is not followed by any
	// CONTINUATION frames.
	EndHeaders bool

	// PadLength is the optional number of bytes of zeros to add
	// to this frame.
	PadLength uint8
}

// WritePushPromise writes a single PushPromise Frame.
//
// As with Header Frames, This is the low level call for writing
// individual frames. Continuation frames are handled elsewhere.
//
// It will perform exactly one Write to the underlying Writer.
// It is the caller's responsibility to not call other Write methods concurrently.
func (h2f *Framer) WritePushPromise(p PushPromiseParam) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteRawFrame writes a raw frame. This can be used to write
// extension frames unknown to this package.
func (h2f *Framer) WriteRawFrame(t FrameType, flags Flags, streamID uint32, payload []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func readByte(p []byte) (remain []byte, b byte, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func readUint32(p []byte) (remain []byte, v uint32, err error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

type streamEnder interface {
	StreamEnded() bool
}

type headersEnder interface {
	HeadersEnded() bool
}

type headersOrContinuation interface {
	headersEnder
	HeaderBlockFragment() []byte
}

// A MetaHeadersFrame is the representation of one HEADERS frame and
// zero or more contiguous CONTINUATION frames and the decoding of
// their HPACK-encoded contents.
//
// This type of frame does not appear on the wire and is only returned
// by the Framer when Framer.ReadMetaHeaders is set.
type MetaHeadersFrame struct {
	*HeadersFrame

	// Fields are the fields contained in the HEADERS and
	// CONTINUATION frames. The underlying slice is owned by the
	// Framer and must not be retained after the next call to
	// ReadFrame.
	//
	// Fields are guaranteed to be in the correct http2 order and
	// not have unknown pseudo header fields or invalid header
	// field names or values. Required pseudo header fields may be
	// missing, however. Use the MetaHeadersFrame.Pseudo accessor
	// method access pseudo headers.
	Fields []hpack.HeaderField

	// Truncated is whether the max header list size limit was hit
	// and Fields is incomplete. The hpack decoder state is still
	// valid, however.
	Truncated bool
}

// PseudoValue returns the given pseudo header field's value.
// The provided pseudo field should not contain the leading colon.
func (mh *MetaHeadersFrame) PseudoValue(pseudo string) string { _ = "STUB: not implemented"; return "" }

// RegularFields returns the regular (non-pseudo) header fields of mh.
// The caller does not own the returned slice.
func (mh *MetaHeadersFrame) RegularFields() []hpack.HeaderField {
	_ = "STUB: not implemented"
	return nil
}

// PseudoFields returns the pseudo header fields of mh.
// The caller does not own the returned slice.
func (mh *MetaHeadersFrame) PseudoFields() []hpack.HeaderField {
	_ = "STUB: not implemented"
	return nil
}

func (mh *MetaHeadersFrame) checkPseudos() error { _ = "STUB: not implemented"; return nil }

// Check for duplicates.
// This would be a bad algorithm, but N is 4.
// And this doesn't allocate.

func (fr *Framer) maxHeaderStringLen() int { _ = "STUB: not implemented"; return 0 }

// If maxHeaderListSize overflows an int, use no limit (0).

// readMetaFrame returns 0 or more CONTINUATION frames from fr and
// merge them into the provided hf and returns a MetaHeadersFrame
// with the decoded hpack values.
func (h2f *Framer) readMetaFrame(hf *HeadersFrame, dumps []*dump.Dumper) (Frame, error) {
	_ = "STUB: not implemented"
	return *new(Frame), nil
}

// pseudo header field errors

// Don't include the value in the error, because it may be sensitive.

// Lose reference to MetaHeadersFrame:

// Avoid parsing large amounts of headers that we will then discard.
// If the sender exceeds the max header list size by too much,
// skip parsing the fragment and close the connection.
//
// "Too much" is either any CONTINUATION frame after we've already
// exceeded the max header list size (in which case remainSize is 0),
// or a frame whose encoded size is more than twice the remaining
// header list bytes we're willing to accept.

// It would be nice to send a RST_STREAM before sending the GOAWAY,
// but the structure of the server's frame writer makes this difficult.

// Also close the connection after any CONTINUATION frame following an
// invalid header, since we stop tracking the size of the headers after
// an invalid one.

// It would be nice to send a RST_STREAM before sending the GOAWAY,
// but the structure of the server's frame writer makes this difficult.

// guaranteed by checkFrameOrder

func summarizeFrame(f Frame) string { _ = "STUB: not implemented"; return "" }

// remove trailing comma
