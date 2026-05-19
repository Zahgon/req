// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ascii

// EqualFold is strings.EqualFold, ASCII only. It reports whether s and t
// are equal, ASCII-case-insensitively.
func EqualFold(s, t string) bool { _ = "STUB: not implemented"; return false }

// lower returns the ASCII lowercase version of b.
func lower(b byte) byte { _ = "STUB: not implemented"; return 0 }

// IsPrint returns whether s is ASCII and printable according to
// https://tools.ietf.org/html/rfc20#section-4.2.
func IsPrint(s string) bool { _ = "STUB: not implemented"; return false }

// Is returns whether s is ASCII.
func Is(s string) bool { _ = "STUB: not implemented"; return false }

// ToLower returns the lowercase version of s if s is ASCII and printable.
func ToLower(s string) (lower string, ok bool) { _ = "STUB: not implemented"; return "", false }
