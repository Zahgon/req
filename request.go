package req

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	urlpkg "net/url"
	"time"
)

// Request struct is used to compose and fire individual request from
// req client. Request provides lots of chainable settings which can
// override client level settings.
type Request struct {
	PathParams      map[string]string
	QueryParams     urlpkg.Values
	FormData        urlpkg.Values
	OrderedFormData []string
	Headers         http.Header
	Cookies         []*http.Cookie
	Result          any
	Error           any
	RawRequest      *http.Request
	StartTime       time.Time
	RetryAttempt    int
	RawURL          string // read only
	Method          string
	Body            []byte
	GetBody         GetContentFunc
	// URL is an auto-generated field, and is nil in request middleware (OnBeforeRequest),
	// consider using RawURL if you want, it's not nil in client middleware (WrapRoundTripFunc)
	URL *urlpkg.URL

	isMultiPart              bool
	disableAutoReadResponse  bool
	forceChunkedEncoding     bool
	isSaveResponse           bool
	close                    bool
	error                    error
	client                   *Client
	uploadCallback           UploadCallback
	uploadCallbackInterval   time.Duration
	downloadCallback         DownloadCallback
	downloadCallbackInterval time.Duration
	unReplayableBody         io.ReadCloser
	retryOption              *retryOption
	bodyReadCloser           io.ReadCloser
	dumpOptions              *DumpOptions
	marshalBody              any
	ctx                      context.Context
	uploadFiles              []*FileUpload
	uploadReader             []io.ReadCloser
	outputFile               string
	output                   io.Writer
	trace                    *clientTrace
	dumpBuffer               *bytes.Buffer
	responseReturnTime       time.Time
	afterResponse            []ResponseMiddleware
}

type GetContentFunc func() (io.ReadCloser, error)

func (r *Request) getHeader(key string) string { _ = "STUB: not implemented"; return "" }

// TraceInfo returns the trace information, only available if trace is enabled
// (see Request.EnableTrace and Client.EnableTraceAll).
func (r *Request) TraceInfo() TraceInfo { _ = "STUB: not implemented"; return *new(TraceInfo) }

// in case timeout

// Only calculate on successful connections

// Only calculate on successful connections

// Only calculate on successful connections

// Capture remote address info when connection is non-nil

// HeaderToString get all header as string.
func (r *Request) HeaderToString() string { _ = "STUB: not implemented"; return "" }

// SetURL set the url for request.
func (r *Request) SetURL(url string) *Request { _ = "STUB: not implemented"; return nil }

// SetFormDataFromValues set the form data from url.Values, will not
// been used if request method does not allow payload.
func (r *Request) SetFormDataFromValues(data urlpkg.Values) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetFormData set the form data from a map, will not been used
// if request method does not allow payload.
func (r *Request) SetFormData(data map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetOrderedFormData set the ordered form data from key-values pairs.
func (r *Request) SetOrderedFormData(kvs ...string) *Request { _ = "STUB: not implemented"; return nil }

// SetFormDataAnyType set the form data from a map, which value could be any type,
// will convert to string automatically.
// It will not been used if request method does not allow payload.
func (r *Request) SetFormDataAnyType(data map[string]any) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetCookies set http cookies for the request.
func (r *Request) SetCookies(cookies ...*http.Cookie) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetQueryString set URL query parameters for the request using
// raw query string.
func (r *Request) SetQueryString(query string) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryParamsFromValues sets query parameters from a url.Values map.
// This method allows direct configuration of query parameters from url.Values,
// which is commonly used with libraries like go-querystring.
func (r *Request) SetQueryParamsFromValues(params urlpkg.Values) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetQueryParamsFromStruct sets query parameters from a struct using go-querystring.
// This method provides a higher-level abstraction by allowing users to directly pass
// a struct to configure query parameters. The struct should use `url` tags to specify
// parameter names.
func (r *Request) SetQueryParamsFromStruct(v any) *Request { _ = "STUB: not implemented"; return nil }

// SetFileReader set up a multipart form with a reader to upload file.
func (r *Request) SetFileReader(paramName, filename string, reader io.Reader) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetFileBytes set up a multipart form with given []byte to upload.
func (r *Request) SetFileBytes(paramName, filename string, content []byte) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetFiles set up a multipart form from a map to upload, which
// key is the parameter name, and value is the file path.
func (r *Request) SetFiles(files map[string]string) *Request { _ = "STUB: not implemented"; return nil }

// SetFile set up a multipart form from file path to upload,
// which read file from filePath automatically to upload.
func (r *Request) SetFile(paramName, filePath string) *Request {
	_ = "STUB: not implemented"
	return nil
}

var (
	errMissingParamName   = errors.New("missing param name in multipart file upload")
	errMissingFileName    = errors.New("missing filename in multipart file upload")
	errMissingFileContent = errors.New("missing file content in multipart file upload")
)

// SetFileUpload set the fully customized multipart file upload options.
func (r *Request) SetFileUpload(uploads ...FileUpload) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetUploadCallback set the UploadCallback which will be invoked at least
// every 200ms during file upload, usually used to show upload progress.
func (r *Request) SetUploadCallback(callback UploadCallback) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetUploadCallbackWithInterval set the UploadCallback which will be invoked at least
// every `minInterval` during file upload, usually used to show upload progress.
func (r *Request) SetUploadCallbackWithInterval(callback UploadCallback, minInterval time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetDownloadCallback set the DownloadCallback which will be invoked at least
// every 200ms during file upload, usually used to show download progress.
func (r *Request) SetDownloadCallback(callback DownloadCallback) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetDownloadCallbackWithInterval set the DownloadCallback which will be invoked at least
// every `minInterval` during file upload, usually used to show download progress.
func (r *Request) SetDownloadCallbackWithInterval(callback DownloadCallback, minInterval time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetResult set the result that response Body will be unmarshalled to if
// no error occurs and Response.ResultState() returns SuccessState, by default
// it requires HTTP status `code >= 200 && code <= 299`, you can also use
// Client.SetResultStateCheckFunc to customize the result state check logic.
//
// Deprecated: Use SetSuccessResult instead.
func (r *Request) SetResult(result any) *Request { _ = "STUB: not implemented"; return nil }

// SetSuccessResult set the result that response Body will be unmarshalled to if
// no error occurs and Response.ResultState() returns SuccessState, by default
// it requires HTTP status `code >= 200 && code <= 299`, you can also use
// Client.SetResultStateCheckFunc to customize the result state check logic.
func (r *Request) SetSuccessResult(result any) *Request { _ = "STUB: not implemented"; return nil }

// SetError set the result that response body will be unmarshalled to if
// no error occurs and Response.ResultState() returns ErrorState, by default
// it requires HTTP status `code >= 400`, you can also use
// Client.SetResultStateCheckFunc to customize the result state check logic.
//
// Deprecated: Use SetErrorResult result.
func (r *Request) SetError(err any) *Request { _ = "STUB: not implemented"; return nil }

// SetErrorResult set the result that response body will be unmarshalled to if
// no error occurs and Response.ResultState() returns ErrorState, by default
// it requires HTTP status `code >= 400`, you can also
// use Client.SetResultStateCheckFunc to customize the result state check logic.
func (r *Request) SetErrorResult(err any) *Request { _ = "STUB: not implemented"; return nil }

// SetBearerAuthToken set bearer auth token for the request.
func (r *Request) SetBearerAuthToken(token string) *Request { _ = "STUB: not implemented"; return nil }

// SetBasicAuth set basic auth for the request.
func (r *Request) SetBasicAuth(username, password string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetDigestAuth sets the Digest Access auth scheme for the HTTP request. If a server responds with 401 and sends a
// Digest challenge in the WWW-Authenticate Header, the request will be resent with the appropriate Authorization Header.
//
// For Example: To set the Digest scheme with username "roc" and password "123456"
//
//	client.R().SetDigestAuth("roc", "123456")
//
// Information about Digest Access Authentication can be found in RFC7616:
//
//	https://datatracker.ietf.org/doc/html/rfc7616
//
// Deprecated: Use Client.SetCommonDigestAuth instead. Request level digest auth is not recommended,
func (r *Request) SetDigestAuth(username, password string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// OnAfterResponse add a response middleware which hooks after response received.
func (r *Request) OnAfterResponse(m ResponseMiddleware) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetHeaders set headers from a map for the request.
func (r *Request) SetHeaders(hdrs map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetHeader set a header for the request.
func (r *Request) SetHeader(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// SetHeadersNonCanonical set headers from a map for the request which key is a
// non-canonical key (keep case unchanged), only valid for HTTP/1.1.
func (r *Request) SetHeadersNonCanonical(hdrs map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetHeaderNonCanonical set a header for the request which key is a
// non-canonical key (keep case unchanged), only valid for HTTP/1.1.
func (r *Request) SetHeaderNonCanonical(key, value string) *Request {
	_ = "STUB: not implemented"
	return nil
}

const (
	// HeaderOderKey is the key of header order, which specifies the order
	// of the http header.
	HeaderOderKey = "__header_order__"
	// PseudoHeaderOderKey is the key of pseudo header order, which specifies
	// the order of the http2 and http3 pseudo header.
	PseudoHeaderOderKey = "__pseudo_header_order__"
)

// SetHeaderOrder set the order of the http header (case-insensitive).
// For example:
//
//	client.R().SetHeaderOrder(
//	    "custom-header",
//	    "cookie",
//	    "user-agent",
//	    "accept-encoding",
//	)
func (r *Request) SetHeaderOrder(keys ...string) *Request { _ = "STUB: not implemented"; return nil }

// SetPseudoHeaderOrder set the order of the pseudo http header (case-insensitive).
// Note this is only valid for http2 and http3.
// For example:
//
//	client.R().SetPseudoHeaderOrder(
//	    ":scheme",
//	    ":authority",
//	    ":path",
//	    ":method",
//	)
func (r *Request) SetPseudoHeaderOrder(keys ...string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetOutputFile set the file that response Body will be downloaded to.
func (r *Request) SetOutputFile(file string) *Request { _ = "STUB: not implemented"; return nil }

// SetOutput set the io.Writer that response Body will be downloaded to.
func (r *Request) SetOutput(output io.Writer) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryParams set URL query parameters from a map for the request.
func (r *Request) SetQueryParams(params map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetQueryParamsAnyType set URL query parameters from a map for the request.
// The value of map is any type, will be convert to string automatically.
func (r *Request) SetQueryParamsAnyType(params map[string]any) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetQueryParam set an URL query parameter for the request.
func (r *Request) SetQueryParam(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// AddQueryParam add a URL query parameter for the request.
func (r *Request) AddQueryParam(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// AddQueryParams add one or more values of specified URL query parameter for the request.
func (r *Request) AddQueryParams(key string, values ...string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetPathParams set URL path parameters from a map for the request.
func (r *Request) SetPathParams(params map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetPathParam set a URL path parameter for the request.
func (r *Request) SetPathParam(key, value string) *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) appendError(err error) { _ = "STUB: not implemented"; return }

var errRetryableWithUnReplayableBody = errors.New("retryable request should not have unreplayable Body (io.Reader)")

func (r *Request) newErrorResponse(err error) *Response { _ = "STUB: not implemented"; return nil }

// Do fires http request, 0 or 1 context is allowed, and returns the *Response which
// is always not nil, and Response.Err is not nil if error occurs.
func (r *Request) Do(ctx ...context.Context) *Response { _ = "STUB: not implemented"; return nil }

// retryable request should not have unreplayable Body

func (r *Request) do() (resp *Response, err error) { _ = "STUB: not implemented"; return nil, nil }

// Determine if the error is from a canceled context.
// Store it here so it doesn't get lost when processing the AfterResponse middleware.

// absolutely cannot retry.

// check retry whether is needed.
// default behaviour: retry if error occurs
// override default behaviour if custom RetryConditions has been set.

// no retry is needed.

// need retry, attempt to retry

// run retry hooks in reverse order

// clean up before retry

// Send fires http request with specified method and url, returns the
// *Response which is always not nil, and the error is not nil if error occurs.
func (r *Request) Send(method, url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustGet like Get, panic if error happens, should only be used to
// test without error handling.
func (r *Request) MustGet(url string) *Response { _ = "STUB: not implemented"; return nil }

// Get fires http request with GET method and the specified URL.
func (r *Request) Get(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustPost like Post, panic if error happens. should only be used to
// test without error handling.
func (r *Request) MustPost(url string) *Response { _ = "STUB: not implemented"; return nil }

// Post fires http request with POST method and the specified URL.
func (r *Request) Post(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustPut like Put, panic if error happens, should only be used to
// test without error handling.
func (r *Request) MustPut(url string) *Response { _ = "STUB: not implemented"; return nil }

// Put fires http request with PUT method and the specified URL.
func (r *Request) Put(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustPatch like Patch, panic if error happens, should only be used
// to test without error handling.
func (r *Request) MustPatch(url string) *Response { _ = "STUB: not implemented"; return nil }

// Patch fires http request with PATCH method and the specified URL.
func (r *Request) Patch(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustDelete like Delete, panic if error happens, should only be used
// to test without error handling.
func (r *Request) MustDelete(url string) *Response { _ = "STUB: not implemented"; return nil }

// Delete fires http request with DELETE method and the specified URL.
func (r *Request) Delete(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustOptions like Options, panic if error happens, should only be
// used to test without error handling.
func (r *Request) MustOptions(url string) *Response { _ = "STUB: not implemented"; return nil }

// Options fires http request with OPTIONS method and the specified URL.
func (r *Request) Options(url string) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustHead like Head, panic if error happens, should only be used
// to test without error handling.
func (r *Request) MustHead(url string) *Response { _ = "STUB: not implemented"; return nil }

// Head fires http request with HEAD method and the specified URL.
func (r *Request) Head(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// SetBody set the request Body, accepts string, []byte, io.Reader, map and struct.
func (r *Request) SetBody(body any) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyBytes set the request Body as []byte.
func (r *Request) SetBodyBytes(body []byte) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyString set the request Body as string.
func (r *Request) SetBodyString(body string) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyJsonString set the request Body as string and set Content-Type header
// as "application/json; charset=utf-8"
func (r *Request) SetBodyJsonString(body string) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyJsonBytes set the request Body as []byte and set Content-Type header
// as "application/json; charset=utf-8"
func (r *Request) SetBodyJsonBytes(body []byte) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyJsonMarshal set the request Body that marshaled from object, and
// set Content-Type header as "application/json; charset=utf-8"
func (r *Request) SetBodyJsonMarshal(v any) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyXmlString set the request Body as string and set Content-Type header
// as "text/xml; charset=utf-8"
func (r *Request) SetBodyXmlString(body string) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyXmlBytes set the request Body as []byte and set Content-Type header
// as "text/xml; charset=utf-8"
func (r *Request) SetBodyXmlBytes(body []byte) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyXmlMarshal set the request Body that marshaled from object, and
// set Content-Type header as "text/xml; charset=utf-8"
func (r *Request) SetBodyXmlMarshal(v any) *Request { _ = "STUB: not implemented"; return nil }

// SetContentType set the `Content-Type` for the request.
func (r *Request) SetContentType(contentType string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// Context method returns the Context if its already set in request
// otherwise it creates new one using `context.Background()`.
func (r *Request) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SetContext method sets the context.Context for current Request. It allows
// to interrupt the request execution if ctx.Done() channel is closed.
// See https://blog.golang.org/context article and the "context" package
// documentation.
//
// Attention: make sure call SetContext before EnableDumpXXX if you want to
// dump at the request level.
func (r *Request) SetContext(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }

// SetContextData sets the key-value pair data for current Request, so you
// can access some extra context info for current Request in hook or middleware.
func (r *Request) SetContextData(key, val any) *Request { _ = "STUB: not implemented"; return nil }

// GetContextData returns the context data of specified key, which set by SetContextData.
func (r *Request) GetContextData(key any) any { _ = "STUB: not implemented"; return *new(any) }

// DisableAutoReadResponse disable read response body automatically (enabled by default).
func (r *Request) DisableAutoReadResponse() *Request { _ = "STUB: not implemented"; return nil }

// EnableAutoReadResponse enable read response body automatically (enabled by default).
func (r *Request) EnableAutoReadResponse() *Request { _ = "STUB: not implemented"; return nil }

// DisableTrace disables trace.
func (r *Request) DisableTrace() *Request { _ = "STUB: not implemented"; return nil }

// EnableTrace enables trace (http3 currently does not support trace).
func (r *Request) EnableTrace() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) getDumpBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func (r *Request) getDumpOptions() *DumpOptions { _ = "STUB: not implemented"; return nil }

// EnableDumpTo enables dump and save to the specified io.Writer.
func (r *Request) EnableDumpTo(output io.Writer) *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpToFile enables dump and save to the specified filename.
func (r *Request) EnableDumpToFile(filename string) *Request { _ = "STUB: not implemented"; return nil }

// SetDumpOptions sets DumpOptions at request level.
func (r *Request) SetDumpOptions(opt *DumpOptions) *Request { _ = "STUB: not implemented"; return nil }

// EnableDump enables dump, including all content for the request and response by default.
func (r *Request) EnableDump() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutBody enables dump only header for the request and response.
func (r *Request) EnableDumpWithoutBody() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutHeader enables dump only Body for the request and response.
func (r *Request) EnableDumpWithoutHeader() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutResponse enables dump only request.
func (r *Request) EnableDumpWithoutResponse() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutRequest enables dump only response.
func (r *Request) EnableDumpWithoutRequest() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutRequestBody enables dump with request Body excluded,
// can be used in upload request to avoid dump the unreadable binary content.
func (r *Request) EnableDumpWithoutRequestBody() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutResponseBody enables dump with response Body excluded,
// can be used in download request to avoid dump the unreadable binary content.
func (r *Request) EnableDumpWithoutResponseBody() *Request { _ = "STUB: not implemented"; return nil }

// EnableForceChunkedEncoding enables force using chunked encoding when uploading.
func (r *Request) EnableForceChunkedEncoding() *Request { _ = "STUB: not implemented"; return nil }

// DisableForceChunkedEncoding disables force using chunked encoding when uploading.
func (r *Request) DisableForceChunkedEncoding() *Request { _ = "STUB: not implemented"; return nil }

// EnableForceMultipart enables force using multipart to upload form data.
func (r *Request) EnableForceMultipart() *Request { _ = "STUB: not implemented"; return nil }

// DisableForceMultipart disables force using multipart to upload form data.
func (r *Request) DisableForceMultipart() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) getRetryOption() *retryOption { _ = "STUB: not implemented"; return nil }

// SetRetryCount enables retry and set the maximum retry count.
// It will retry infinitely if count is negative.
func (r *Request) SetRetryCount(count int) *Request { _ = "STUB: not implemented"; return nil }

// SetRetryInterval sets the custom GetRetryIntervalFunc, you can use this to
// implement your own backoff retry algorithm.
// For example:
//
//	req.SetRetryInterval(func(resp *req.Response, attempt int) time.Duration {
//	    sleep := 0.01 * math.Exp2(float64(attempt))
//	    return time.Duration(math.Min(2, sleep)) * time.Second
//	})
func (r *Request) SetRetryInterval(getRetryIntervalFunc GetRetryIntervalFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetRetryFixedInterval set retry to use a fixed interval.
func (r *Request) SetRetryFixedInterval(interval time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetRetryBackoffInterval set retry to use a capped exponential backoff with jitter.
// https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
func (r *Request) SetRetryBackoffInterval(min, max time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetRetryHook set the retry hook which will be executed before a retry.
// It will override other retry hooks if any been added before (including
// client-level retry hooks).
func (r *Request) SetRetryHook(hook RetryHookFunc) *Request { _ = "STUB: not implemented"; return nil }

// AddRetryHook adds a retry hook which will be executed before a retry.
func (r *Request) AddRetryHook(hook RetryHookFunc) *Request { _ = "STUB: not implemented"; return nil }

// SetRetryCondition sets the retry condition, which determines whether the
// request should retry.
// It will override other retry conditions if any been added before (including
// client-level retry conditions).
func (r *Request) SetRetryCondition(condition RetryConditionFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

// AddRetryCondition adds a retry condition, which determines whether the
// request should retry.
func (r *Request) AddRetryCondition(condition RetryConditionFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetClient change the client of request dynamically.
func (r *Request) SetClient(client *Client) *Request { _ = "STUB: not implemented"; return nil }

// GetClient returns the current client used by request.
func (r *Request) GetClient() *Client {
	_ = "STUB: not implemented"

	// EnableCloseConnection closes the connection after sending this
	// request and reading its response if set to true in HTTP/1.1 and
	// HTTP/2.
	//
	// Setting this field prevents reuse of TCP connections between
	// requests to the same hosts event if EnableKeepAlives() were called.
	return nil
}

func (r *Request) EnableCloseConnection() *Request { _ = "STUB: not implemented"; return nil }
