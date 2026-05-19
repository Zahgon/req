package req

import (
	"errors"
	"io"
	"mime/multipart"
	"net/textproto"
	"time"
)

type (
	// RequestMiddleware type is for request middleware, called before a request is sent
	RequestMiddleware func(client *Client, req *Request) error

	// ResponseMiddleware type is for response middleware, called after a response has been received
	ResponseMiddleware func(client *Client, resp *Response) error
)

func createMultipartHeader(file *FileUpload, contentType string) textproto.MIMEHeader {
	_ = "STUB: not implemented"
	return *new(textproto.MIMEHeader)
}

func closeq(v any) { _ = "STUB: not implemented"; return }

func writeMultipartFormFile(w *multipart.Writer, file *FileUpload, r *Request) error {
	_ = "STUB: not implemented"
	return nil
}

// reset file reader when retry a multipart file upload

// Auto detect actual multipart content type

func writeMultiPart(r *Request, w *multipart.Writer) {
	_ = "STUB: not implemented"
	// close multipart to write tailer boundary
	return
}

func handleMultiPart(c *Client, r *Request) (err error) { _ = "STUB: not implemented"; return nil }

// close pipe writer so that pipe reader could get EOF, and stop upload

func handleFormData(r *Request) { _ = "STUB: not implemented"; return }

var errBadOrderedFormData = errors.New("bad ordered form data, the number of key-value pairs should be an even number")

func handleOrderedFormData(r *Request) { _ = "STUB: not implemented"; return }

func handleMarshalBody(c *Client, r *Request) error { _ = "STUB: not implemented"; return nil }

func parseRequestBody(c *Client, r *Request) (err error) { _ = "STUB: not implemented"; return nil }

// handle multipart

// handle form data

// handle marshal body

// body is in-memory []byte, so we can guess content type

// ignore if content type set at client-level

// ignore if content-type set at request-level

func unmarshalBody(c *Client, r *Response, v any) (err error) {
	_ = "STUB: not implemented"
	return nil
	// in case req.SetResult or req.SetError with client.DisableAutoReadResponse(true)
}

func defaultResultStateChecker(resp *Response) ResultState {
	_ = "STUB: not implemented"
	return *new(ResultState)
}

func parseResponseBody(c *Client, r *Response) (err error) { _ = "STUB: not implemented"; return nil }

type callbackWriter struct {
	io.Writer
	written   int64
	totalSize int64
	lastTime  time.Time
	interval  time.Duration
	callback  func(written int64)
}

func (w *callbackWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type callbackReader struct {
	io.ReadCloser
	read     int64
	lastRead int64
	callback func(read int64)
	lastTime time.Time
	interval time.Duration
}

func (r *callbackReader) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func handleDownload(c *Client, r *Response) (err error) { _ = "STUB: not implemented"; return nil }

// already read

// must not nil

// generate URL
func parseRequestURL(c *Client, r *Request) error { _ = "STUB: not implemented"; return nil }

// Parsing request URL

// set scheme if missing

// If RawURL is relative path then added c.BaseURL into
// the request URL otherwise Request.URL will be used as-is

// Adding Query Param

// remove query param from client level by key
// since overrides happens for that key in the request

// Preserve query string order partially.
// Since not feasible in `SetQuery*` resty methods, because
// standard package `url.Encode(...)` sorts the query params
// alphabetically

func parseRequestHeader(c *Client, r *Request) error { _ = "STUB: not implemented"; return nil }

func parseRequestCookie(c *Client, r *Request) error { _ = "STUB: not implemented"; return nil }
