package util

import (
	"reflect"
)

// IsJSONType method is to check JSON content type or not
func IsJSONType(ct string) bool { _ = "STUB: not implemented"; return false }

// IsXMLType method is to check XML content type or not
func IsXMLType(ct string) bool { _ = "STUB: not implemented"; return false }

// GetPointer return the pointer of the interface.
func GetPointer(v any) any { _ = "STUB: not implemented"; return *new(any) }

// pointer of pointer

// GetType return the underlying type.
func GetType(v any) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// CutString slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func CutString(s, sep string) (before, after string, found bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// CutBytes slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, nil, false.
//
// CutBytes returns slices of the original slice s, not copies.
func CutBytes(s, sep []byte) (before, after []byte, found bool) {
	_ = "STUB: not implemented"
	return nil, nil, false
}

// IsStringEmpty method tells whether given string is empty or not
func IsStringEmpty(str string) bool { _ = "STUB: not implemented"; return false }

// See 2 (end of page 4) https://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
func basicAuth(username, password string) string { _ = "STUB: not implemented"; return "" }

// BasicAuthHeaderValue return the header of basic auth.
func BasicAuthHeaderValue(username, password string) string { _ = "STUB: not implemented"; return "" }

// CreateDirectory create the directory.
func CreateDirectory(dir string) (err error) { _ = "STUB: not implemented"; return nil }
