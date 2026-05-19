package tests

import (
	"testing"
)

func AssertIsNil(t *testing.T, v any) { _ = "STUB: not implemented"; return }

func AssertAllNotNil(t *testing.T, vv ...any) { _ = "STUB: not implemented"; return }

func AssertNotNil(t *testing.T, v any) { _ = "STUB: not implemented"; return }

func AssertEqual(t *testing.T, e, g any) { _ = "STUB: not implemented"; return }

func AssertNoError(t *testing.T, err error) { _ = "STUB: not implemented"; return }

func AssertErrorContains(t *testing.T, err error, s string) { _ = "STUB: not implemented"; return }

func AssertContains(t *testing.T, s, substr string, shouldContain bool) {
	_ = "STUB: not implemented"
	return
}

func AssertClone(t *testing.T, e, g any) { _ = "STUB: not implemented"; return }

func equal(expected, got any) bool { _ = "STUB: not implemented"; return false }

func isNil(v any) bool { _ = "STUB: not implemented"; return false }
