package http3

import (
	"context"
	"sync"

	"github.com/quic-go/quic-go"
)

const streamDatagramQueueLen = 32

// stateTrackingStream is an implementation of quic.Stream that delegates
// to an underlying stream
// it takes care of proxying send and receive errors onto an implementation of
// the errorSetter interface (intended to be occupied by a datagrammer)
// it is also responsible for clearing the stream based on its ID from its
// parent connection, this is done through the streamClearer interface when
// both the send and receive sides are closed
type stateTrackingStream struct {
	*quic.Stream

	sendDatagram func([]byte) error
	hasData      chan struct{}
	queue        [][]byte // TODO: use a ring buffer

	mx      sync.Mutex
	sendErr error
	recvErr error

	clearer streamClearer
}

var _ datagramStream = &stateTrackingStream{}

type streamClearer interface {
	clearStream(quic.StreamID)
}

func newStateTrackingStream(s *quic.Stream, clearer streamClearer, sendDatagram func([]byte) error) *stateTrackingStream {
	_ = "STUB: not implemented"
	return nil
}

func (s *stateTrackingStream) closeSend(e error) { _ = "STUB: not implemented"; return }

// clear the stream the first time both the send
// and receive are finished

func (s *stateTrackingStream) closeReceive(e error) { _ = "STUB: not implemented"; return }

// clear the stream the first time both the send
// and receive are finished

func (s *stateTrackingStream) Close() error { _ = "STUB: not implemented"; return nil }

func (s *stateTrackingStream) CancelWrite(e quic.StreamErrorCode) {
	_ = "STUB: not implemented"
	return
}

func (s *stateTrackingStream) Write(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *stateTrackingStream) CancelRead(e quic.StreamErrorCode) { _ = "STUB: not implemented"; return }

func (s *stateTrackingStream) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *stateTrackingStream) SendDatagram(b []byte) error { _ = "STUB: not implemented"; return nil }

func (s *stateTrackingStream) signalHasDatagram() { _ = "STUB: not implemented"; return }

func (s *stateTrackingStream) enqueueDatagram(data []byte) { _ = "STUB: not implemented"; return }

func (s *stateTrackingStream) ReceiveDatagram(ctx context.Context) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *stateTrackingStream) QUICStream() *quic.Stream { _ = "STUB: not implemented"; return nil }
