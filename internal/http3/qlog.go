package http3

import (
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3/qlog"
	"github.com/quic-go/quic-go/qlogwriter"

	"github.com/quic-go/qpack"
)

func maybeQlogInvalidHeadersFrame(qlogger qlogwriter.Recorder, streamID quic.StreamID, l uint64) {
	_ = "STUB: not implemented"
	return
}

func qlogParsedHeadersFrame(qlogger qlogwriter.Recorder, streamID quic.StreamID, hf *headersFrame, hfs []qpack.HeaderField) {
	_ = "STUB: not implemented"
	return
}

func qlogCreatedHeadersFrame(qlogger qlogwriter.Recorder, streamID quic.StreamID, length, payloadLength int, hfs []qlog.HeaderField) {
	_ = "STUB: not implemented"
	return
}
