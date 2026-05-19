// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"errors"
)

// An ErrCode is an unsigned 32-bit error code as defined in the HTTP/2 spec.
type ErrCode uint32

const (
	ErrCodeNo                 ErrCode = 0x0
	ErrCodeProtocol           ErrCode = 0x1
	ErrCodeInternal           ErrCode = 0x2
	ErrCodeFlowControl        ErrCode = 0x3
	ErrCodeSettingsTimeout    ErrCode = 0x4
	ErrCodeStreamClosed       ErrCode = 0x5
	ErrCodeFrameSize          ErrCode = 0x6
	ErrCodeRefusedStream      ErrCode = 0x7
	ErrCodeCancel             ErrCode = 0x8
	ErrCodeCompression        ErrCode = 0x9
	ErrCodeConnect            ErrCode = 0xa
	ErrCodeEnhanceYourCalm    ErrCode = 0xb
	ErrCodeInadequateSecurity ErrCode = 0xc
	ErrCodeHTTP11Required     ErrCode = 0xd
)

var errCodeName = map[ErrCode]string{
	ErrCodeNo:                 "NO_ERROR",
	ErrCodeProtocol:           "PROTOCOL_ERROR",
	ErrCodeInternal:           "INTERNAL_ERROR",
	ErrCodeFlowControl:        "FLOW_CONTROL_ERROR",
	ErrCodeSettingsTimeout:    "SETTINGS_TIMEOUT",
	ErrCodeStreamClosed:       "STREAM_CLOSED",
	ErrCodeFrameSize:          "FRAME_SIZE_ERROR",
	ErrCodeRefusedStream:      "REFUSED_STREAM",
	ErrCodeCancel:             "CANCEL",
	ErrCodeCompression:        "COMPRESSION_ERROR",
	ErrCodeConnect:            "CONNECT_ERROR",
	ErrCodeEnhanceYourCalm:    "ENHANCE_YOUR_CALM",
	ErrCodeInadequateSecurity: "INADEQUATE_SECURITY",
	ErrCodeHTTP11Required:     "HTTP_1_1_REQUIRED",
}

func (e ErrCode) String() string { _ = "STUB: not implemented"; return "" }

func (e ErrCode) stringToken() string { _ = "STUB: not implemented"; return "" }

// ConnectionError is an error that results in the termination of the
// entire connection.
type ConnectionError ErrCode

func (e ConnectionError) Error() string { _ = "STUB: not implemented"; return "" }

// StreamError is an error that only affects one stream within an
// HTTP/2 connection.
type StreamError struct {
	StreamID uint32
	Code     ErrCode
	Cause    error // optional additional detail
}

// errFromPeer is a sentinel error value for StreamError.Cause to
// indicate that the StreamError was sent from the peer over the wire
// and wasn't locally generated in the Transport.
var errFromPeer = errors.New("received from peer")

func streamError(id uint32, code ErrCode) StreamError {
	_ = "STUB: not implemented"
	return *new(StreamError)
}

func (e StreamError) Error() string { _ = "STUB: not implemented"; return "" }

// connError represents an HTTP/2 ConnectionError error code, along
// with a string (for debugging) explaining why.
//
// Errors of this type are only returned by the frame parser functions
// and converted into ConnectionError(Code), after stashing away
// the Reason into the Framer's errDetail field, accessible via
// the (*Framer).ErrorDetail method.
type connError struct {
	Code   ErrCode // the ConnectionError error code
	Reason string  // additional reason
}

func (e connError) Error() string { _ = "STUB: not implemented"; return "" }

type pseudoHeaderError string

func (e pseudoHeaderError) Error() string { _ = "STUB: not implemented"; return "" }

type duplicatePseudoHeaderError string

func (e duplicatePseudoHeaderError) Error() string { _ = "STUB: not implemented"; return "" }

type headerFieldNameError string

func (e headerFieldNameError) Error() string { _ = "STUB: not implemented"; return "" }

type headerFieldValueError string

func (e headerFieldValueError) Error() string { _ = "STUB: not implemented"; return "" }

var (
	errMixPseudoHeaderTypes = errors.New("mix of request and response pseudo headers")
	errPseudoAfterRegular   = errors.New("pseudo header field after regular")
)
