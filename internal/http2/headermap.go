// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"sync"
)

var (
	commonBuildOnce   sync.Once
	commonLowerHeader map[string]string // Go-Canonical-Case -> lower-case
	commonCanonHeader map[string]string // lower-case -> Go-Canonical-Case
)

func buildCommonHeaderMapsOnce() { _ = "STUB: not implemented"; return }

func buildCommonHeaderMaps() { _ = "STUB: not implemented"; return }

func lowerHeader(v string) (lower string, isAscii bool) {
	_ = "STUB: not implemented"
	return "", false
}

func canonicalHeader(v string) string { _ = "STUB: not implemented"; return "" }
