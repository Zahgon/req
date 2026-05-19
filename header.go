package req

import (
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/imroc/req/v3/internal/header"
)

var headerNewlineToSpace = strings.NewReplacer("\n", " ", "\r", " ")

// stringWriter implements WriteString on a Writer.
type stringWriter struct {
	w io.Writer
}

func (w stringWriter) WriteString(s string) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil

	// A headerSorter implements sort.Interface by sorting a []keyValues
	// by key. It's used as a pointer, so it can fit in a sort.Interface
	// interface value without allocation.
}

type headerSorter struct {
	kvs []header.KeyValues
}

func (s *headerSorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s *headerSorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s *headerSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

var headerSorterPool = sync.Pool{
	New: func() any { return new(headerSorter) },
}

// get is like Get, but key must already be in CanonicalHeaderKey form.
func headerGet(h http.Header, key string) string { _ = "STUB: not implemented"; return "" }

// has reports whether h has the provided key defined, even if it's
// set to 0-length slice.
func headerHas(h http.Header, key string) bool { _ = "STUB: not implemented"; return false }

// sortedKeyValues returns h's keys sorted in the returned kvs
// slice. The headerSorter used to sort is also returned, for possible
// return to headerSorterCache.
func headerSortedKeyValues(h http.Header, exclude map[string]bool) (kvs []header.KeyValues, hs *headerSorter) {
	_ = "STUB: not implemented"
	return nil, nil
}

func headerWrite(h http.Header, writeHeader func(key string, values ...string) error, sort bool) error {
	_ = "STUB: not implemented"
	return nil
}

func headerWriteSubset(h http.Header, exclude map[string]bool, writeHeader func(key string, values ...string) error, sort bool) error {
	_ = "STUB: not implemented"
	return nil
}

// This could be an error. In the common case of
// writing response headers, however, we have no good
// way to provide the error back to the server
// handler, so just drop invalid headers instead.
