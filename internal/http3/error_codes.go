package http3

import (
	"github.com/quic-go/quic-go"
)

type ErrCode quic.ApplicationErrorCode

const (
	ErrCodeNoError                  ErrCode = 0x100
	ErrCodeGeneralProtocolError     ErrCode = 0x101
	ErrCodeInternalError            ErrCode = 0x102
	ErrCodeStreamCreationError      ErrCode = 0x103
	ErrCodeClosedCriticalStream     ErrCode = 0x104
	ErrCodeFrameUnexpected          ErrCode = 0x105
	ErrCodeFrameError               ErrCode = 0x106
	ErrCodeExcessiveLoad            ErrCode = 0x107
	ErrCodeIDError                  ErrCode = 0x108
	ErrCodeSettingsError            ErrCode = 0x109
	ErrCodeMissingSettings          ErrCode = 0x10a
	ErrCodeRequestRejected          ErrCode = 0x10b
	ErrCodeRequestCanceled          ErrCode = 0x10c
	ErrCodeRequestIncomplete        ErrCode = 0x10d
	ErrCodeMessageError             ErrCode = 0x10e
	ErrCodeConnectError             ErrCode = 0x10f
	ErrCodeVersionFallback          ErrCode = 0x110
	ErrCodeDatagramError            ErrCode = 0x33
	ErrCodeQPACKDecompressionFailed ErrCode = 0x200
)

func (e ErrCode) String() string { _ = "STUB: not implemented"; return "" }

func (e ErrCode) string() string { _ = "STUB: not implemented"; return "" }
