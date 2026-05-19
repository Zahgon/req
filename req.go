package req

import (
	"net/http"
	"net/url"
)

type kv struct {
	Key   string
	Value string
}

// ContentDisposition represents parameters in `Content-Disposition`
// MIME header of multipart request.
type ContentDisposition struct {
	kv []kv
}

// Add adds a new key-value pair of Content-Disposition
func (c *ContentDisposition) Add(key, value string) *ContentDisposition {
	_ = "STUB: not implemented"
	return nil
}

func (c *ContentDisposition) string() string { _ = "STUB: not implemented"; return "" }

// FileUpload represents a "form-data" multipart
type FileUpload struct {
	// "name" parameter in `Content-Disposition`
	ParamName string
	// "filename" parameter in `Content-Disposition`
	FileName string
	// The file to be uploaded.
	GetFileContent GetContentFunc
	// Optional file length in bytes.
	FileSize int64
	// Optional Content-Type
	ContentType string

	// Optional extra ContentDisposition parameters.
	// According to the HTTP specification, this should be nil,
	// but some servers may not follow the specification and
	// requires `Content-Disposition` parameters more than just
	// "name" and "filename".
	ExtraContentDisposition *ContentDisposition
}

// UploadInfo is the information for each UploadCallback call.
type UploadInfo struct {
	// parameter name in multipart upload
	ParamName string
	// filename in multipart upload
	FileName string
	// total file length in bytes.
	FileSize int64
	// uploaded file length in bytes.
	UploadedSize int64
}

// UploadCallback is the callback which will be invoked during
// multipart upload.
type UploadCallback func(info UploadInfo)

// DownloadInfo is the information for each DownloadCallback call.
type DownloadInfo struct {
	// Response is the corresponding Response during download.
	Response *Response
	// downloaded body length in bytes.
	DownloadedSize int64
}

// DownloadCallback is the callback which will be invoked during
// response body download.
type DownloadCallback func(info DownloadInfo)

func cloneSlice[T any](s []T) []T { _ = "STUB: not implemented"; return nil }

func cloneUrlValues(v url.Values) url.Values { _ = "STUB: not implemented"; return *new(url.Values) }

func cloneMap(h map[string]string) map[string]string { _ = "STUB: not implemented"; return nil }

// convertHeaderToString converts http header to a string.
func convertHeaderToString(h http.Header) string { _ = "STUB: not implemented"; return "" }
