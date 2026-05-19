package req

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)

// SetURL is a global wrapper methods which delegated
// to the default client, create a request and SetURL for request.
func SetURL(url string) *Request { _ = "STUB: not implemented"; return nil }

// SetFormDataFromValues is a global wrapper methods which delegated
// to the default client, create a request and SetFormDataFromValues for request.
func SetFormDataFromValues(data url.Values) *Request { _ = "STUB: not implemented"; return nil }

// SetFormData is a global wrapper methods which delegated
// to the default client, create a request and SetFormData for request.
func SetFormData(data map[string]string) *Request { _ = "STUB: not implemented"; return nil }

// SetOrderedFormData is a global wrapper methods which delegated
// to the default client, create a request and SetOrderedFormData for request.
func SetOrderedFormData(kvs ...string) *Request { _ = "STUB: not implemented"; return nil }

// SetFormDataAnyType is a global wrapper methods which delegated
// to the default client, create a request and SetFormDataAnyType for request.
func SetFormDataAnyType(data map[string]any) *Request { _ = "STUB: not implemented"; return nil }

// SetCookies is a global wrapper methods which delegated
// to the default client, create a request and SetCookies for request.
func SetCookies(cookies ...*http.Cookie) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryString is a global wrapper methods which delegated
// to the default client, create a request and SetQueryString for request.
func SetQueryString(query string) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryParamsFromValues is a global wrapper methods which delegated
// to the default client, create a request and SetQueryParamsFromValues for request.
func SetQueryParamsFromValues(params url.Values) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryParamsFromStruct is a global wrapper methods which delegated
// to the default client, create a request and SetQueryParamsFromStruct for request.
func SetQueryParamsFromStruct(v any) *Request { _ = "STUB: not implemented"; return nil }

// SetFileReader is a global wrapper methods which delegated
// to the default client, create a request and SetFileReader for request.
func SetFileReader(paramName, filePath string, reader io.Reader) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetFileBytes is a global wrapper methods which delegated
// to the default client, create a request and SetFileBytes for request.
func SetFileBytes(paramName, filename string, content []byte) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetFiles is a global wrapper methods which delegated
// to the default client, create a request and SetFiles for request.
func SetFiles(files map[string]string) *Request { _ = "STUB: not implemented"; return nil }

// SetFile is a global wrapper methods which delegated
// to the default client, create a request and SetFile for request.
func SetFile(paramName, filePath string) *Request { _ = "STUB: not implemented"; return nil }

// SetFileUpload is a global wrapper methods which delegated
// to the default client, create a request and SetFileUpload for request.
func SetFileUpload(f ...FileUpload) *Request { _ = "STUB: not implemented"; return nil }

// SetResult is a global wrapper methods which delegated
// to the default client, create a request and SetSuccessResult for request.
//
// Deprecated: Use SetSuccessResult instead.
func SetResult(result any) *Request { _ = "STUB: not implemented"; return nil }

// SetSuccessResult is a global wrapper methods which delegated
// to the default client, create a request and SetSuccessResult for request.
func SetSuccessResult(result any) *Request { _ = "STUB: not implemented"; return nil }

// SetError is a global wrapper methods which delegated
// to the default client, create a request and SetErrorResult for request.
//
// Deprecated: Use SetErrorResult instead.
func SetError(error any) *Request { _ = "STUB: not implemented"; return nil }

// SetErrorResult is a global wrapper methods which delegated
// to the default client, create a request and SetErrorResult for request.
func SetErrorResult(error any) *Request { _ = "STUB: not implemented"; return nil }

// SetBearerAuthToken is a global wrapper methods which delegated
// to the default client, create a request and SetBearerAuthToken for request.
func SetBearerAuthToken(token string) *Request { _ = "STUB: not implemented"; return nil }

// SetBasicAuth is a global wrapper methods which delegated
// to the default client, create a request and SetBasicAuth for request.
func SetBasicAuth(username, password string) *Request { _ = "STUB: not implemented"; return nil }

// SetDigestAuth is a global wrapper methods which delegated
// to the default client, create a request and SetDigestAuth for request.
func SetDigestAuth(username, password string) *Request { _ = "STUB: not implemented"; return nil }

// SetHeaders is a global wrapper methods which delegated
// to the default client, create a request and SetHeaders for request.
func SetHeaders(hdrs map[string]string) *Request { _ = "STUB: not implemented"; return nil }

// SetHeader is a global wrapper methods which delegated
// to the default client, create a request and SetHeader for request.
func SetHeader(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// SetHeaderOrder is a global wrapper methods which delegated
// to the default client, create a request and SetHeaderOrder for request.
func SetHeaderOrder(keys ...string) *Request { _ = "STUB: not implemented"; return nil }

// SetPseudoHeaderOrder is a global wrapper methods which delegated
// to the default client, create a request and SetPseudoHeaderOrder for request.
func SetPseudoHeaderOrder(keys ...string) *Request { _ = "STUB: not implemented"; return nil }

// SetOutputFile is a global wrapper methods which delegated
// to the default client, create a request and SetOutputFile for request.
func SetOutputFile(file string) *Request { _ = "STUB: not implemented"; return nil }

// SetOutput is a global wrapper methods which delegated
// to the default client, create a request and SetOutput for request.
func SetOutput(output io.Writer) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryParams is a global wrapper methods which delegated
// to the default client, create a request and SetQueryParams for request.
func SetQueryParams(params map[string]string) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryParamsAnyType is a global wrapper methods which delegated
// to the default client, create a request and SetQueryParamsAnyType for request.
func SetQueryParamsAnyType(params map[string]any) *Request { _ = "STUB: not implemented"; return nil }

// SetQueryParam is a global wrapper methods which delegated
// to the default client, create a request and SetQueryParam for request.
func SetQueryParam(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// AddQueryParam is a global wrapper methods which delegated
// to the default client, create a request and AddQueryParam for request.
func AddQueryParam(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// AddQueryParams is a global wrapper methods which delegated
// to the default client, create a request and AddQueryParams for request.
func AddQueryParams(key string, values ...string) *Request { _ = "STUB: not implemented"; return nil }

// SetPathParams is a global wrapper methods which delegated
// to the default client, create a request and SetPathParams for request.
func SetPathParams(params map[string]string) *Request { _ = "STUB: not implemented"; return nil }

// SetPathParam is a global wrapper methods which delegated
// to the default client, create a request and SetPathParam for request.
func SetPathParam(key, value string) *Request { _ = "STUB: not implemented"; return nil }

// MustGet is a global wrapper methods which delegated
// to the default client, create a request and MustGet for request.
func MustGet(url string) *Response { _ = "STUB: not implemented"; return nil }

// Get is a global wrapper methods which delegated
// to the default client, create a request and Get for request.
func Get(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustPost is a global wrapper methods which delegated
// to the default client, create a request and Get for request.
func MustPost(url string) *Response { _ = "STUB: not implemented"; return nil }

// Post is a global wrapper methods which delegated
// to the default client, create a request and Post for request.
func Post(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustPut is a global wrapper methods which delegated
// to the default client, create a request and MustPut for request.
func MustPut(url string) *Response { _ = "STUB: not implemented"; return nil }

// Put is a global wrapper methods which delegated
// to the default client, create a request and Put for request.
func Put(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustPatch is a global wrapper methods which delegated
// to the default client, create a request and MustPatch for request.
func MustPatch(url string) *Response { _ = "STUB: not implemented"; return nil }

// Patch is a global wrapper methods which delegated
// to the default client, create a request and Patch for request.
func Patch(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustDelete is a global wrapper methods which delegated
// to the default client, create a request and MustDelete for request.
func MustDelete(url string) *Response { _ = "STUB: not implemented"; return nil }

// Delete is a global wrapper methods which delegated
// to the default client, create a request and Delete for request.
func Delete(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustOptions is a global wrapper methods which delegated
// to the default client, create a request and MustOptions for request.
func MustOptions(url string) *Response { _ = "STUB: not implemented"; return nil }

// Options is a global wrapper methods which delegated
// to the default client, create a request and Options for request.
func Options(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// MustHead is a global wrapper methods which delegated
// to the default client, create a request and MustHead for request.
func MustHead(url string) *Response { _ = "STUB: not implemented"; return nil }

// Head is a global wrapper methods which delegated
// to the default client, create a request and Head for request.
func Head(url string) (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// SetBody is a global wrapper methods which delegated
// to the default client, create a request and SetBody for request.
func SetBody(body any) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyBytes is a global wrapper methods which delegated
// to the default client, create a request and SetBodyBytes for request.
func SetBodyBytes(body []byte) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyString is a global wrapper methods which delegated
// to the default client, create a request and SetBodyString for request.
func SetBodyString(body string) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyJsonString is a global wrapper methods which delegated
// to the default client, create a request and SetBodyJsonString for request.
func SetBodyJsonString(body string) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyJsonBytes is a global wrapper methods which delegated
// to the default client, create a request and SetBodyJsonBytes for request.
func SetBodyJsonBytes(body []byte) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyJsonMarshal is a global wrapper methods which delegated
// to the default client, create a request and SetBodyJsonMarshal for request.
func SetBodyJsonMarshal(v any) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyXmlString is a global wrapper methods which delegated
// to the default client, create a request and SetBodyXmlString for request.
func SetBodyXmlString(body string) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyXmlBytes is a global wrapper methods which delegated
// to the default client, create a request and SetBodyXmlBytes for request.
func SetBodyXmlBytes(body []byte) *Request { _ = "STUB: not implemented"; return nil }

// SetBodyXmlMarshal is a global wrapper methods which delegated
// to the default client, create a request and SetBodyXmlMarshal for request.
func SetBodyXmlMarshal(v any) *Request { _ = "STUB: not implemented"; return nil }

// SetContentType is a global wrapper methods which delegated
// to the default client, create a request and SetContentType for request.
func SetContentType(contentType string) *Request { _ = "STUB: not implemented"; return nil }

// SetContext is a global wrapper methods which delegated
// to the default client, create a request and SetContext for request.
func SetContext(ctx context.Context) *Request { _ = "STUB: not implemented"; return nil }

// DisableTrace is a global wrapper methods which delegated
// to the default client, create a request and DisableTrace for request.
func DisableTrace() *Request { _ = "STUB: not implemented"; return nil }

// EnableTrace is a global wrapper methods which delegated
// to the default client, create a request and EnableTrace for request.
func EnableTrace() *Request { _ = "STUB: not implemented"; return nil }

// EnableForceChunkedEncoding is a global wrapper methods which delegated
// to the default client, create a request and EnableForceChunkedEncoding for request.
func EnableForceChunkedEncoding() *Request { _ = "STUB: not implemented"; return nil }

// DisableForceChunkedEncoding is a global wrapper methods which delegated
// to the default client, create a request and DisableForceChunkedEncoding for request.
func DisableForceChunkedEncoding() *Request { _ = "STUB: not implemented"; return nil }

// EnableForceMultipart is a global wrapper methods which delegated
// to the default client, create a request and EnableForceMultipart for request.
func EnableForceMultipart() *Request { _ = "STUB: not implemented"; return nil }

// DisableForceMultipart is a global wrapper methods which delegated
// to the default client, create a request and DisableForceMultipart for request.
func DisableForceMultipart() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpTo is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpTo for request.
func EnableDumpTo(output io.Writer) *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpToFile is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpToFile for request.
func EnableDumpToFile(filename string) *Request { _ = "STUB: not implemented"; return nil }

// SetDumpOptions is a global wrapper methods which delegated
// to the default client, create a request and SetDumpOptions for request.
func SetDumpOptions(opt *DumpOptions) *Request { _ = "STUB: not implemented"; return nil }

// EnableDump is a global wrapper methods which delegated
// to the default client, create a request and EnableDump for request.
func EnableDump() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutBody is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpWithoutBody for request.
func EnableDumpWithoutBody() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutHeader is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpWithoutHeader for request.
func EnableDumpWithoutHeader() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutResponse is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpWithoutResponse for request.
func EnableDumpWithoutResponse() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutRequest is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpWithoutRequest for request.
func EnableDumpWithoutRequest() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutRequestBody is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpWithoutRequestBody for request.
func EnableDumpWithoutRequestBody() *Request { _ = "STUB: not implemented"; return nil }

// EnableDumpWithoutResponseBody is a global wrapper methods which delegated
// to the default client, create a request and EnableDumpWithoutResponseBody for request.
func EnableDumpWithoutResponseBody() *Request { _ = "STUB: not implemented"; return nil }

// SetRetryCount is a global wrapper methods which delegated
// to the default client, create a request and SetRetryCount for request.
func SetRetryCount(count int) *Request { _ = "STUB: not implemented"; return nil }

// SetRetryInterval is a global wrapper methods which delegated
// to the default client, create a request and SetRetryInterval for request.
func SetRetryInterval(getRetryIntervalFunc GetRetryIntervalFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetRetryFixedInterval is a global wrapper methods which delegated
// to the default client, create a request and SetRetryFixedInterval for request.
func SetRetryFixedInterval(interval time.Duration) *Request { _ = "STUB: not implemented"; return nil }

// SetRetryBackoffInterval is a global wrapper methods which delegated
// to the default client, create a request and SetRetryBackoffInterval for request.
func SetRetryBackoffInterval(min, max time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetRetryHook is a global wrapper methods which delegated
// to the default client, create a request and SetRetryHook for request.
func SetRetryHook(hook RetryHookFunc) *Request { _ = "STUB: not implemented"; return nil }

// AddRetryHook is a global wrapper methods which delegated
// to the default client, create a request and AddRetryHook for request.
func AddRetryHook(hook RetryHookFunc) *Request { _ = "STUB: not implemented"; return nil }

// SetRetryCondition is a global wrapper methods which delegated
// to the default client, create a request and SetRetryCondition for request.
func SetRetryCondition(condition RetryConditionFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

// AddRetryCondition is a global wrapper methods which delegated
// to the default client, create a request and AddRetryCondition for request.
func AddRetryCondition(condition RetryConditionFunc) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetUploadCallback is a global wrapper methods which delegated
// to the default client, create a request and SetUploadCallback for request.
func SetUploadCallback(callback UploadCallback) *Request { _ = "STUB: not implemented"; return nil }

// SetUploadCallbackWithInterval is a global wrapper methods which delegated
// to the default client, create a request and SetUploadCallbackWithInterval for request.
func SetUploadCallbackWithInterval(callback UploadCallback, minInterval time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetDownloadCallback is a global wrapper methods which delegated
// to the default client, create a request and SetDownloadCallback for request.
func SetDownloadCallback(callback DownloadCallback) *Request { _ = "STUB: not implemented"; return nil }

// SetDownloadCallbackWithInterval is a global wrapper methods which delegated
// to the default client, create a request and SetDownloadCallbackWithInterval for request.
func SetDownloadCallbackWithInterval(callback DownloadCallback, minInterval time.Duration) *Request {
	_ = "STUB: not implemented"
	return nil
}

// EnableCloseConnection is a global wrapper methods which delegated
// to the default client, create a request and EnableCloseConnection for request.
func EnableCloseConnection() *Request { _ = "STUB: not implemented"; return nil }
