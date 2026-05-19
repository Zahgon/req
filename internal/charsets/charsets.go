package charsets

import (
	"golang.org/x/text/encoding"
)

var boms = []struct {
	bom []byte
	enc string
}{
	{[]byte{0xfe, 0xff}, "utf-16be"},
	{[]byte{0xff, 0xfe}, "utf-16le"},
	{[]byte{0xef, 0xbb, 0xbf}, "utf-8"},
}

// FindEncoding sniff and find the encoding of the content.
func FindEncoding(content []byte) (enc encoding.Encoding, name string) {
	_ = "STUB: not implemented"
	return *new(encoding.Encoding), ""
}

func prescan(content []byte) (e encoding.Encoding, name string) {
	_ = "STUB: not implemented"
	return *new(encoding.Encoding), ""
}

func fromMetaElement(s string) string { _ = "STUB: not implemented"; return "" }
