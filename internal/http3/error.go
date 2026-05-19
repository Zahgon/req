package http3

// Error is returned from the round tripper (for HTTP clients)
// and inside the HTTP handler (for HTTP servers) if an HTTP/3 error occurs.
// See section 8 of RFC 9114.
type Error struct {
	Remote       bool
	ErrorCode    ErrCode
	ErrorMessage string
}

var _ error = &Error{}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// Usually errors are remote. Only make it explicit for local errors.

func (e *Error) Is(target error) bool { _ = "STUB: not implemented"; return false }

func maybeReplaceError(err error) error { _ = "STUB: not implemented"; return nil }
