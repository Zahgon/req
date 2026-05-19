package req

import "sync"

const copyBufPoolSize = 32 * 1024

var copyBufPool = sync.Pool{New: func() any { return new([copyBufPoolSize]byte) }}

func getCopyBuf() []byte { _ = "STUB: not implemented"; return nil }

func putCopyBuf(b []byte) { _ = "STUB: not implemented"; return }
