package req

import (
	"net/http"
	"time"
)

// Response is the http response.
type Response struct {
	// The underlying http.Response is embed into Response.
	*http.Response
	// Err is the underlying error, not nil if some error occurs.
	// Usually used in the ResponseMiddleware, you can skip logic in
	// ResponseMiddleware that doesn't need to be executed when err occurs.
	Err error
	// Request is the Response's related Request.
	Request    *Request
	body       []byte
	receivedAt time.Time
	error      any
	result     any
}

// IsSuccess method returns true if no error occurs and HTTP status `code >= 200 and <= 299`
// by default, you can also use Client.SetResultStateCheckFunc to customize the result
// state check logic.
//
// Deprecated: Use IsSuccessState instead.
func (r *Response) IsSuccess() bool { _ = "STUB: not implemented"; return false }

// IsSuccessState method returns true if no error occurs and HTTP status `code >= 200 and <= 299`
// by default, you can also use Client.SetResultStateCheckFunc to customize the result state
// check logic.
func (r *Response) IsSuccessState() bool { _ = "STUB: not implemented"; return false }

// IsError method returns true if no error occurs and HTTP status `code >= 400`
// by default, you can also use Client.SetResultStateCheckFunc to customize the result
// state check logic.
//
// Deprecated: Use IsErrorState instead.
func (r *Response) IsError() bool { _ = "STUB: not implemented"; return false }

// IsErrorState method returns true if no error occurs and HTTP status `code >= 400`
// by default, you can also use Client.SetResultStateCheckFunc to customize the result
// state check logic.
func (r *Response) IsErrorState() bool { _ = "STUB: not implemented"; return false }

// GetContentType return the `Content-Type` header value.
func (r *Response) GetContentType() string { _ = "STUB: not implemented"; return "" }

// ResultState returns the result state.
// By default, it returns SuccessState if HTTP status `code >= 200 && code <= 299`, and returns
// ErrorState if HTTP status `code >= 400`, otherwise returns UnknownState.
// You can also use Client.SetResultStateCheckFunc to customize the result
// state check logic.
func (r *Response) ResultState() ResultState { _ = "STUB: not implemented"; return *new(ResultState) }

// Result returns the automatically unmarshalled object if Request.SetSuccessResult
// is called and ResultState returns SuccessState.
// Otherwise, return nil.
//
// Deprecated: Use SuccessResult instead.
func (r *Response) Result() any {
	_ = "STUB: not implemented"
	return *

	// SuccessResult returns the automatically unmarshalled object if Request.SetSuccessResult
	// is called and ResultState returns SuccessState.
	// Otherwise, return nil.
	new(any)
}

func (r *Response) SuccessResult() any {
	_ = "STUB: not implemented"

	// Error returns the automatically unmarshalled object when Request.SetErrorResult
	// or Client.SetCommonErrorResult is called, and ResultState returns ErrorState.
	// Otherwise, return nil.
	//
	// Deprecated: Use ErrorResult instead.
	return *new(any)
}

func (r *Response) Error() any {
	_ = "STUB: not implemented"

	// ErrorResult returns the automatically unmarshalled object when Request.SetErrorResult
	// or Client.SetCommonErrorResult is called, and ResultState returns ErrorState.
	// Otherwise, return nil.
	return *new(any)
}

func (r *Response) ErrorResult() any {
	_ = "STUB: not implemented"

	// TraceInfo returns the TraceInfo from Request.
	return *new(any)
}

func (r *Response) TraceInfo() TraceInfo { _ = "STUB: not implemented"; return *new(TraceInfo) }

// TotalTime returns the total time of the request, from request we sent to response we received.
func (r *Response) TotalTime() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// ReceivedAt returns the timestamp that response we received.
func (r *Response) ReceivedAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (r *Response) setReceivedAt() { _ = "STUB: not implemented"; return }

// UnmarshalJson unmarshalls JSON response body into the specified object.
func (r *Response) UnmarshalJson(v any) error { _ = "STUB: not implemented"; return nil }

// UnmarshalXml unmarshalls XML response body into the specified object.
func (r *Response) UnmarshalXml(v any) error { _ = "STUB: not implemented"; return nil }

// Unmarshal unmarshalls response body into the specified object according
// to response `Content-Type`.
func (r *Response) Unmarshal(v any) error { _ = "STUB: not implemented"; return nil }

// Into unmarshalls response body into the specified object according
// to response `Content-Type`.
func (r *Response) Into(v any) error { _ = "STUB: not implemented"; return nil }

// Set response body with byte array content
func (r *Response) SetBody(body []byte) {
	_ = "STUB: not implemented"

	// Set response body with string content
	return
}

func (r *Response) SetBodyString(body string) { _ = "STUB: not implemented"; return }

// Bytes return the response body as []bytes that have already been read, could be
// nil if not read, the following cases are already read:
//  1. `Request.SetResult` or `Request.SetError` is called.
//  2. `Client.DisableAutoReadResponse` and `Request.DisableAutoReadResponse` is not
//     called, and also `Request.SetOutput` and `Request.SetOutputFile` is not called.
func (r *Response) Bytes() []byte {
	_ = "STUB: not implemented"

	// String returns the response body as string that have already been read, could be
	// nil if not read, the following cases are already read:
	//  1. `Request.SetResult` or `Request.SetError` is called.
	//  2. `Client.DisableAutoReadResponse` and `Request.DisableAutoReadResponse` is not
	//     called, and also `Request.SetOutput` and `Request.SetOutputFile` is not called.
	return nil
}

func (r *Response) String() string { _ = "STUB: not implemented"; return "" }

// ToString returns the response body as string, read body if not have been read.
func (r *Response) ToString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ToBytes returns the response body as []byte, read body if not have been read.
func (r *Response) ToBytes() (body []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

// Dump return the string content that have been dumped for the request.
// `Request.Dump` or `Request.DumpXXX` MUST have been called.
func (r *Response) Dump() string { _ = "STUB: not implemented"; return "" }

// GetStatus returns the response status.
func (r *Response) GetStatus() string { _ = "STUB: not implemented"; return "" }

// GetStatusCode returns the response status code.
func (r *Response) GetStatusCode() int { _ = "STUB: not implemented"; return 0 }

// GetHeader returns the response header value by key.
func (r *Response) GetHeader(key string) string { _ = "STUB: not implemented"; return "" }

// GetHeaderValues returns the response header values by key.
func (r *Response) GetHeaderValues(key string) []string { _ = "STUB: not implemented"; return nil }

// HeaderToString get all header as string.
func (r *Response) HeaderToString() string { _ = "STUB: not implemented"; return "" }
