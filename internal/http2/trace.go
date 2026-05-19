// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"net/http"
	"net/http/httptrace"
	"net/textproto"
)

func traceHasWroteHeaderField(trace *httptrace.ClientTrace) bool {
	_ = "STUB: not implemented"
	return false
}

func traceWroteHeaderField(trace *httptrace.ClientTrace, k, v string) {
	_ = "STUB: not implemented"
	return
}

func traceGot1xxResponseFunc(trace *httptrace.ClientTrace) func(int, textproto.MIMEHeader) error {
	_ = "STUB: not implemented"
	return nil
}

func traceGetConn(req *http.Request, hostPort string) { _ = "STUB: not implemented"; return }

func traceGotConn(req *http.Request, cc *ClientConn, reused bool) {
	_ = "STUB: not implemented"
	return
}

func traceWroteHeaders(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }

func traceGot100Continue(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }

func traceWait100Continue(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }

func traceWroteRequest(trace *httptrace.ClientTrace, err error) { _ = "STUB: not implemented"; return }

func traceFirstResponseByte(trace *httptrace.ClientTrace) { _ = "STUB: not implemented"; return }
