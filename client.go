package req

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/http/cookiejar"
	urlpkg "net/url"
	"reflect"
	"time"

	utls "github.com/refraction-networking/utls"

	"github.com/imroc/req/v3/http2"
)

// DefaultClient returns the global default Client.
func DefaultClient() *Client { _ = "STUB: not implemented"; return nil }

// SetDefaultClient override the global default Client.
func SetDefaultClient(c *Client) { _ = "STUB: not implemented"; return }

var defaultClient = C()

// Client is the req's http client.
type Client struct {
	BaseURL               string
	PathParams            map[string]string
	QueryParams           urlpkg.Values
	FormData              urlpkg.Values
	DebugLog              bool
	AllowGetMethodPayload bool
	*Transport
	digestAuth              *digestAuth
	cookiejarFactory        func() *cookiejar.Jar
	trace                   bool
	disableAutoReadResponse bool
	commonErrorType         reflect.Type
	retryOption             *retryOption
	jsonMarshal             func(v any) ([]byte, error)
	jsonUnmarshal           func(data []byte, v any) error
	xmlMarshal              func(v any) ([]byte, error)
	xmlUnmarshal            func(data []byte, v any) error
	multipartBoundaryFunc   func() string
	outputDirectory         string
	scheme                  string
	log                     Logger
	dumpOptions             *DumpOptions
	httpClient              *http.Client
	beforeRequest           []RequestMiddleware
	udBeforeRequest         []RequestMiddleware
	afterResponse           []ResponseMiddleware
	wrappedRoundTrip        RoundTripper
	roundTripWrappers       []RoundTripWrapper
	responseBodyTransformer func(rawBody []byte, req *Request, resp *Response) (transformedBody []byte, err error)
	resultStateCheckFunc    func(resp *Response) ResultState
	onError                 ErrorHook
}

type ErrorHook func(client *Client, req *Request, resp *Response, err error)

// R create a new request.
func (c *Client) R() *Request { _ = "STUB: not implemented"; return nil }

// Get create a new GET request, accepts 0 or 1 url.
func (c *Client) Get(url ...string) *Request { _ = "STUB: not implemented"; return nil }

// Post create a new POST request.
func (c *Client) Post(url ...string) *Request { _ = "STUB: not implemented"; return nil }

// Patch create a new PATCH request.
func (c *Client) Patch(url ...string) *Request { _ = "STUB: not implemented"; return nil }

// Delete create a new DELETE request.
func (c *Client) Delete(url ...string) *Request { _ = "STUB: not implemented"; return nil }

// Put create a new PUT request.
func (c *Client) Put(url ...string) *Request { _ = "STUB: not implemented"; return nil }

// Head create a new HEAD request.
func (c *Client) Head(url ...string) *Request { _ = "STUB: not implemented"; return nil }

// Options create a new OPTIONS request.
func (c *Client) Options(url ...string) *Request { _ = "STUB: not implemented"; return nil }

// GetTransport return the underlying transport.
func (c *Client) GetTransport() *Transport {
	_ = "STUB: not implemented"

	// SetResponseBodyTransformer set the response body transformer, which can modify the
	// response body before unmarshalled if auto-read response body is not disabled.
	return nil
}

func (c *Client) SetResponseBodyTransformer(fn func(rawBody []byte, req *Request, resp *Response) (transformedBody []byte, err error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonError set the common result that response body will be unmarshalled to
// if no error occurs but Response.ResultState returns ErrorState, by default it
// is HTTP status `code >= 400`, you can also use SetCommonResultStateChecker
// to customize the result state check logic.
//
// Deprecated: Use SetCommonErrorResult instead.
func (c *Client) SetCommonError(err any) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonErrorResult set the common result that response body will be unmarshalled to
// if no error occurs but Response.ResultState returns ErrorState, by default it
// is HTTP status `code >= 400`, you can also use SetCommonResultStateChecker
// to customize the result state check logic.
func (c *Client) SetCommonErrorResult(err any) *Client { _ = "STUB: not implemented"; return nil }

// ResultState represents the state of the result.
type ResultState int

const (
	// SuccessState indicates the response is in success state,
	// and result will be unmarshalled if Request.SetSuccessResult
	// is called.
	SuccessState ResultState = iota
	// ErrorState indicates the response is in error state,
	// and result will be unmarshalled if Request.SetErrorResult
	// or Client.SetCommonErrorResult is called.
	ErrorState
	// UnknownState indicates the response is in unknown state,
	// and handler will be invoked if Request.SetUnknownResultHandlerFunc
	// or Client.SetCommonUnknownResultHandlerFunc is called.
	UnknownState
)

// SetResultStateCheckFunc overrides the default result state checker with customized one,
// which returns SuccessState when HTTP status `code >= 200 and <= 299`, and returns
// ErrorState when HTTP status `code >= 400`, otherwise returns UnknownState.
func (c *Client) SetResultStateCheckFunc(fn func(resp *Response) ResultState) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonFormDataFromValues set the form data from url.Values for requests
// fired from the client which request method allows payload.
func (c *Client) SetCommonFormDataFromValues(data urlpkg.Values) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonFormData set the form data from map for requests fired from the client
// which request method allows payload.
func (c *Client) SetCommonFormData(data map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetMultipartBoundaryFunc overrides the default function used to generate
// boundary delimiters for "multipart/form-data" requests with a customized one,
// which returns a boundary delimiter (without the two leading hyphens).
//
// Boundary delimiter may only contain certain ASCII characters, and must be
// non-empty and at most 70 bytes long (see RFC 2046, Section 5.1.1).
func (c *Client) SetMultipartBoundaryFunc(fn func() string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetBaseURL set the default base URL, will be used if request URL is
// a relative URL.
func (c *Client) SetBaseURL(u string) *Client { _ = "STUB: not implemented"; return nil }

// SetOutputDirectory set output directory that response will
// be downloaded to.
func (c *Client) SetOutputDirectory(dir string) *Client { _ = "STUB: not implemented"; return nil }

// SetCertFromFile helps to set client certificates from cert and key file.
func (c *Client) SetCertFromFile(certFile, keyFile string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCerts set client certificates.
func (c *Client) SetCerts(certs ...tls.Certificate) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) appendRootCertData(data []byte) { _ = "STUB: not implemented"; return }

// SetRootCertFromString set root certificates from string.
func (c *Client) SetRootCertFromString(pemContent string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetRootCertsFromFile set root certificates from files.
func (c *Client) SetRootCertsFromFile(pemFiles ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// GetTLSClientConfig return the underlying tls.Config.
func (c *Client) GetTLSClientConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

// SetRedirectPolicy set the RedirectPolicy which controls the behavior of receiving redirect
// responses (usually responses with 301 and 302 status code), see the predefined
// AllowedDomainRedirectPolicy, AllowedHostRedirectPolicy, DefaultRedirectPolicy, MaxRedirectPolicy,
// NoRedirectPolicy, SameDomainRedirectPolicy and SameHostRedirectPolicy.
func (c *Client) SetRedirectPolicy(policies ...RedirectPolicy) *Client {
	_ = "STUB: not implemented"
	return nil
}

// DisableKeepAlives disable the HTTP keep-alives (enabled by default)
// and will only use the connection to the server for a single
// HTTP request.
//
// This is unrelated to the similarly named TCP keep-alives.
func (c *Client) DisableKeepAlives() *Client { _ = "STUB: not implemented"; return nil }

// EnableKeepAlives enables HTTP keep-alives (enabled by default).
func (c *Client) EnableKeepAlives() *Client { _ = "STUB: not implemented"; return nil }

// DisableCompression disables the compression (enabled by default),
// which prevents the Transport from requesting compression
// with an "Accept-Encoding: gzip" request header when the
// Request contains no existing Accept-Encoding value. If
// the Transport requests gzip on its own and gets a gzipped
// response, it's transparently decoded in the Response.Body.
// However, if the user explicitly requested gzip it is not
// automatically uncompressed.
func (c *Client) DisableCompression() *Client { _ = "STUB: not implemented"; return nil }

// EnableCompression enables the compression (enabled by default).
func (c *Client) EnableCompression() *Client { _ = "STUB: not implemented"; return nil }

// EnableAutoDecompress enables the automatic decompression (disabled by default).
func (c *Client) EnableAutoDecompress() *Client { _ = "STUB: not implemented"; return nil }

// DisableAutoDecompress disables the automatic decompression (disabled by default).
func (c *Client) DisableAutoDecompress() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSClientConfig set the TLS client config. Be careful! Usually
// you don't need this, you can directly set the tls configuration with
// methods like EnableInsecureSkipVerify, SetCerts etc. Or you can call
// GetTLSClientConfig to get the current tls configuration to avoid
// overwriting some important configurations, such as not setting NextProtos
// will not use http2 by default.
func (c *Client) SetTLSClientConfig(conf *tls.Config) *Client {
	_ = "STUB: not implemented"
	return nil
}

// EnableInsecureSkipVerify enable send https without verifying
// the server's certificates (disabled by default).
func (c *Client) EnableInsecureSkipVerify() *Client { _ = "STUB: not implemented"; return nil }

// DisableInsecureSkipVerify disable send https without verifying
// the server's certificates (disabled by default).
func (c *Client) DisableInsecureSkipVerify() *Client { _ = "STUB: not implemented"; return nil }

// SetCommonQueryParams set URL query parameters with a map
// for requests fired from the client.
func (c *Client) SetCommonQueryParams(params map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// AddCommonQueryParam add a URL query parameter with a key-value
// pair for requests fired from the client.
func (c *Client) AddCommonQueryParam(key, value string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// AddCommonQueryParams add one or more values of specified URL query parameter
// for requests fired from the client.
func (c *Client) AddCommonQueryParams(key string, values ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) pathParams() map[string]string { _ = "STUB: not implemented"; return nil }

// SetCommonPathParam set a path parameter for requests fired from the client.
func (c *Client) SetCommonPathParam(key, value string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonPathParams set path parameters for requests fired from the client.
func (c *Client) SetCommonPathParams(pathParams map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonQueryParam set a URL query parameter with a key-value
// pair for requests fired from the client.
func (c *Client) SetCommonQueryParam(key, value string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonQueryString set URL query parameters with a raw query string
// for requests fired from the client.
func (c *Client) SetCommonQueryString(query string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonQueryParamsFromValues set URL query parameters from a url.Values map
// for requests fired from the client.
func (c *Client) SetCommonQueryParamsFromValues(params urlpkg.Values) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonQueryParamsFromStruct set URL query parameters from a struct using go-querystring
// for requests fired from the client.
func (c *Client) SetCommonQueryParamsFromStruct(v any) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonCookies set HTTP cookies for requests fired from the client.
func (c *Client) SetCommonCookies(cookies ...*http.Cookie) *Client {
	_ = "STUB: not implemented"
	return nil
}

// DisableDebugLog disable debug level log (disabled by default).
func (c *Client) DisableDebugLog() *Client { _ = "STUB: not implemented"; return nil }

// EnableDebugLog enable debug level log (disabled by default).
func (c *Client) EnableDebugLog() *Client { _ = "STUB: not implemented"; return nil }

// DevMode enables:
// 1. Dump content of all requests and responses to see details.
// 2. Output debug level log for deeper insights.
// 3. Trace all requests, so you can get trace info to analyze performance.
func (c *Client) DevMode() *Client { _ = "STUB: not implemented"; return nil }

// SetScheme set the default scheme for client, will be used when
// there is no scheme in the request URL (e.g. "github.com/imroc/req").
func (c *Client) SetScheme(scheme string) *Client { _ = "STUB: not implemented"; return nil }

// GetLogger return the internal logger, usually used in middleware.
func (c *Client) GetLogger() Logger { _ = "STUB: not implemented"; return *new(Logger) }

// SetLogger set the customized logger for client, will disable log if set to nil.
func (c *Client) SetLogger(log Logger) *Client { _ = "STUB: not implemented"; return nil }

// SetTimeout set timeout for requests fired from the client.
func (c *Client) SetTimeout(d time.Duration) *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) getDumpOptions() *DumpOptions { _ = "STUB: not implemented"; return nil }

// EnableDumpAll enable dump for requests fired from the client, including
// all content for the request and response by default.
func (c *Client) EnableDumpAll() *Client {
	_ = "STUB: not implemented"
	// dump already started
	return nil
}

// EnableDumpAllToFile enable dump for requests fired from the
// client and output to the specified file.
func (c *Client) EnableDumpAllToFile(filename string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// EnableDumpAllTo enable dump for requests fired from the
// client and output to the specified io.Writer.
func (c *Client) EnableDumpAllTo(output io.Writer) *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllAsync enable dump for requests fired from the
// client and output asynchronously, can be used for debugging
// in production environment without affecting performance.
func (c *Client) EnableDumpAllAsync() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutRequestBody enable dump for requests fired
// from the client without request body, can be used in the upload
// request to avoid dumping the unreadable binary content.
func (c *Client) EnableDumpAllWithoutRequestBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutResponseBody enable dump for requests fired
// from the client without response body, can be used in the download
// request to avoid dumping the unreadable binary content.
func (c *Client) EnableDumpAllWithoutResponseBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutResponse enable dump for requests fired from
// the client without response, can be used if you only care about
// the request.
func (c *Client) EnableDumpAllWithoutResponse() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutRequest enables dump for requests fired from
// the client without request, can be used if you only care about
// the response.
func (c *Client) EnableDumpAllWithoutRequest() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutHeader enable dump for requests fired from
// the client without header, can be used if you only care about
// the body.
func (c *Client) EnableDumpAllWithoutHeader() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutBody enable dump for requests fired from
// the client without body, can be used if you only care about
// the header.
func (c *Client) EnableDumpAllWithoutBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequest enable dump at the request-level for each request, and only
// temporarily stores the dump content in memory, call Response.Dump() to get the
// dump content when needed.
func (c *Client) EnableDumpEachRequest() *Client { _ = "STUB: not implemented"; return nil }

// Ignore on retry, no need to repeat enable dump.

// EnableDumpEachRequestWithoutBody enable dump without body at the request-level for
// each request, and only temporarily stores the dump content in memory, call
// Response.Dump() to get the dump content when needed.
func (c *Client) EnableDumpEachRequestWithoutBody() *Client { _ = "STUB: not implemented"; return nil }

// Ignore on retry, no need to repeat enable dump.

// EnableDumpEachRequestWithoutHeader enable dump without header at the request-level for
// each request, and only temporarily stores the dump content in memory, call
// Response.Dump() to get the dump content when needed.
func (c *Client) EnableDumpEachRequestWithoutHeader() *Client {
	_ = "STUB: not implemented"
	return nil
}

// Ignore on retry, no need to repeat enable dump.

// EnableDumpEachRequestWithoutRequest enable dump without request at the request-level for
// each request, and only temporarily stores the dump content in memory, call
// Response.Dump() to get the dump content when needed.
func (c *Client) EnableDumpEachRequestWithoutRequest() *Client {
	_ = "STUB: not implemented"
	return nil
}

// Ignore on retry, no need to repeat enable dump.

// EnableDumpEachRequestWithoutResponse enable dump without response at the request-level for
// each request, and only temporarily stores the dump content in memory, call
// Response.Dump() to get the dump content when needed.
func (c *Client) EnableDumpEachRequestWithoutResponse() *Client {
	_ = "STUB: not implemented"
	return nil
}

// Ignore on retry, no need to repeat enable dump.

// EnableDumpEachRequestWithoutResponseBody enable dump without response body at the
// request-level for each request, and only temporarily stores the dump content in memory,
// call Response.Dump() to get the dump content when needed.
func (c *Client) EnableDumpEachRequestWithoutResponseBody() *Client {
	_ = "STUB: not implemented"
	return nil
}

// Ignore on retry, no need to repeat enable dump.

// EnableDumpEachRequestWithoutRequestBody enable dump without request body at the
// request-level for each request, and only temporarily stores the dump content in memory,
// call Response.Dump() to get the dump content when needed.
func (c *Client) EnableDumpEachRequestWithoutRequestBody() *Client {
	_ = "STUB: not implemented"
	return nil
}

// Ignore on retry, no need to repeat enable dump.

// NewRequest is the alias of R()
func (c *Client) NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (c *Client) NewParallelDownload(url string) *ParallelDownload {
	_ = "STUB: not implemented"
	return nil
}

// DisableAutoReadResponse disable read response body automatically (enabled by default).
func (c *Client) DisableAutoReadResponse() *Client { _ = "STUB: not implemented"; return nil }

// EnableAutoReadResponse enable read response body automatically (enabled by default).
func (c *Client) EnableAutoReadResponse() *Client { _ = "STUB: not implemented"; return nil }

// SetAutoDecodeContentType set the content types that will be auto-detected and decode to utf-8
// (e.g. "json", "xml", "html", "text").
func (c *Client) SetAutoDecodeContentType(contentTypes ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetAutoDecodeContentTypeFunc set the function that determines whether the specified `Content-Type` should be auto-detected and decode to utf-8.
func (c *Client) SetAutoDecodeContentTypeFunc(fn func(contentType string) bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetAutoDecodeAllContentType enable try auto-detect charset and decode all content type to utf-8.
func (c *Client) SetAutoDecodeAllContentType() *Client { _ = "STUB: not implemented"; return nil }

// DisableAutoDecode disable auto-detect charset and decode to utf-8 (enabled by default).
func (c *Client) DisableAutoDecode() *Client { _ = "STUB: not implemented"; return nil }

// EnableAutoDecode enable auto-detect charset and decode to utf-8 (enabled by default).
func (c *Client) EnableAutoDecode() *Client { _ = "STUB: not implemented"; return nil }

// SetUserAgent set the "User-Agent" header for requests fired from the client.
func (c *Client) SetUserAgent(userAgent string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonBearerAuthToken set the bearer auth token for requests fired from the client.
func (c *Client) SetCommonBearerAuthToken(token string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonBasicAuth set the basic auth for requests fired from
// the client.
func (c *Client) SetCommonBasicAuth(username, password string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonDigestAuth sets the Digest Access auth scheme for requests fired from the client. If a server responds with
// 401 and sends a Digest challenge in the WWW-Authenticate Header, requests will be resent with the appropriate
// Authorization Header.
//
// For Example: To set the Digest scheme with user "roc" and password "123456"
//
//	client.SetCommonDigestAuth("roc", "123456")
//
// Information about Digest Access Authentication can be found in RFC7616:
//
//	https://datatracker.ietf.org/doc/html/rfc7616
func (c *Client) SetCommonDigestAuth(username, password string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonHeaders set headers for requests fired from the client.
func (c *Client) SetCommonHeaders(hdrs map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonHeader set a header for requests fired from the client.
func (c *Client) SetCommonHeader(key, value string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonHeaderNonCanonical set a header for requests fired from
// the client which key is a non-canonical key (keep case unchanged),
// only valid for HTTP/1.1.
func (c *Client) SetCommonHeaderNonCanonical(key, value string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonHeadersNonCanonical set headers for requests fired from the
// client which key is a non-canonical key (keep case unchanged), only
// valid for HTTP/1.1.
func (c *Client) SetCommonHeadersNonCanonical(hdrs map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonHeaderOrder set the order of the http header requests fired from the
// client (case-insensitive).
// For example:
//
//	client.R().SetCommonHeaderOrder(
//	    "custom-header",
//	    "cookie",
//	    "user-agent",
//	    "accept-encoding",
//	).Get(url
func (c *Client) SetCommonHeaderOrder(keys ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonPseudoHeaderOder set the order of the pseudo http header requests fired
// from the client (case-insensitive).
// Note this is only valid for http2 and http3.
// For example:
//
//	client.SetCommonPseudoHeaderOder(
//	    ":scheme",
//	    ":authority",
//	    ":path",
//	    ":method",
//	)
func (c *Client) SetCommonPseudoHeaderOder(keys ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2SettingsFrame set the ordered http2 settings frame.
func (c *Client) SetHTTP2SettingsFrame(settings ...http2.Setting) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2ConnectionFlow set the default http2 connection flow, which is the increment
// value of initial WINDOW_UPDATE frame.
func (c *Client) SetHTTP2ConnectionFlow(flow uint32) *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2HeaderPriority set the header priority param.
func (c *Client) SetHTTP2HeaderPriority(priority http2.PriorityParam) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2PriorityFrames set the ordered http2 priority frames.
func (c *Client) SetHTTP2PriorityFrames(frames ...http2.PriorityFrame) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonContentType set the `Content-Type` header for requests fired
// from the client.
func (c *Client) SetCommonContentType(ct string) *Client { _ = "STUB: not implemented"; return nil }

// DisableDumpAll disable dump for requests fired from the client.
func (c *Client) DisableDumpAll() *Client { _ = "STUB: not implemented"; return nil }

// SetCommonDumpOptions configures the underlying Transport's DumpOptions
// for requests fired from the client.
func (c *Client) SetCommonDumpOptions(opt *DumpOptions) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetProxy set the proxy function.
func (c *Client) SetProxy(proxy func(*http.Request) (*urlpkg.URL, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// OnError set the error hook which will be executed if any error returned,
// even if the occurs before request is sent (e.g. invalid URL).
func (c *Client) OnError(hook ErrorHook) *Client { _ = "STUB: not implemented"; return nil }

// OnBeforeRequest add a request middleware which hooks before request sent.
func (c *Client) OnBeforeRequest(m RequestMiddleware) *Client {
	_ = "STUB: not implemented"
	return nil
}

// OnAfterResponse add a response middleware which hooks after response received.
func (c *Client) OnAfterResponse(m ResponseMiddleware) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetProxyURL set proxy from the proxy URL.
func (c *Client) SetProxyURL(proxyUrl string) *Client { _ = "STUB: not implemented"; return nil }

// DisableTraceAll disable trace for requests fired from the client.
func (c *Client) DisableTraceAll() *Client { _ = "STUB: not implemented"; return nil }

// EnableTraceAll enable trace for requests fired from the client (http3
// currently does not support trace).
func (c *Client) EnableTraceAll() *Client { _ = "STUB: not implemented"; return nil }

// SetCookieJar set the cookie jar to the underlying `http.Client`, set to nil if you
// want to disable cookies.
// Note: If you use Client.Clone to clone a new Client, the new client will share the same
// cookie jar as the old Client after cloning. Use SetCookieJarFactory instead if you want
// to create a new CookieJar automatically when cloning a client.
func (c *Client) SetCookieJar(jar http.CookieJar) *Client { _ = "STUB: not implemented"; return nil }

// GetCookies get cookies from the underlying `http.Client`'s `CookieJar`.
func (c *Client) GetCookies(url string) ([]*http.Cookie, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ClearCookies clears all cookies if cookie is enabled, including
// cookies from cookie jar and cookies set by SetCommonCookies.
// Note: The cookie jar will not be cleared if you called SetCookieJar
// instead of SetCookieJarFactory.
func (c *Client) ClearCookies() *Client { _ = "STUB: not implemented"; return nil }

// SetJsonMarshal set the JSON marshal function which will be used
// to marshal request body.
func (c *Client) SetJsonMarshal(fn func(v any) ([]byte, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetJsonUnmarshal set the JSON unmarshal function which will be used
// to unmarshal response body.
func (c *Client) SetJsonUnmarshal(fn func(data []byte, v any) error) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetXmlMarshal set the XML marshal function which will be used
// to marshal request body.
func (c *Client) SetXmlMarshal(fn func(v any) ([]byte, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetXmlUnmarshal set the XML unmarshal function which will be used
// to unmarshal response body.
func (c *Client) SetXmlUnmarshal(fn func(data []byte, v any) error) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetDialTLS set the customized `DialTLSContext` function to Transport.
// Make sure the returned `conn` implements pkg/tls.Conn if you want your
// customized `conn` supports HTTP2.
func (c *Client) SetDialTLS(fn func(ctx context.Context, network, addr string) (net.Conn, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetDial set the customized `DialContext` function to Transport.
func (c *Client) SetDial(fn func(ctx context.Context, network, addr string) (net.Conn, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSFingerprintChrome uses tls fingerprint of Chrome browser.
func (c *Client) SetTLSFingerprintChrome() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintFirefox uses tls fingerprint of Firefox browser.
func (c *Client) SetTLSFingerprintFirefox() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintEdge uses tls fingerprint of Edge browser.
func (c *Client) SetTLSFingerprintEdge() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintQQ uses tls fingerprint of QQ browser.
func (c *Client) SetTLSFingerprintQQ() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintSafari uses tls fingerprint of Safari browser.
func (c *Client) SetTLSFingerprintSafari() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprint360 uses tls fingerprint of 360 browser.
func (c *Client) SetTLSFingerprint360() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintIOS uses tls fingerprint of IOS.
func (c *Client) SetTLSFingerprintIOS() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintAndroid uses tls fingerprint of Android.
func (c *Client) SetTLSFingerprintAndroid() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintRandomized uses randomized tls fingerprint.
func (c *Client) SetTLSFingerprintRandomized() *Client { _ = "STUB: not implemented"; return nil }

// uTLSConn is wrapper of UConn which implements the net.Conn interface.
type uTLSConn struct {
	*utls.UConn
}

func (conn *uTLSConn) ConnectionState() tls.ConnectionState {
	_ = "STUB: not implemented"
	return *new(tls.ConnectionState)
}

// SetTLSFingerprint set the tls fingerprint for tls handshake, will use utls
// (https://github.com/refraction-networking/utls) to perform the tls handshake,
// which uses the specified clientHelloID to simulate the tls fingerprint.
// Note this is valid for HTTP1 and HTTP2, not HTTP3.
func (c *Client) SetTLSFingerprint(clientHelloID utls.ClientHelloID) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSHandshake set the custom tls handshake function, only valid for HTTP1 and HTTP2, not HTTP3,
// it specifies an optional dial function for tls handshake, it works even if a proxy is set, can be
// used to customize the tls fingerprint.
func (c *Client) SetTLSHandshake(fn func(ctx context.Context, addr string, plainConn net.Conn) (conn net.Conn, tlsState *tls.ConnectionState, err error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSHandshakeTimeout set the TLS handshake timeout.
func (c *Client) SetTLSHandshakeTimeout(timeout time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// EnableForceHTTP1 enable force using HTTP1 (disabled by default).
//
// Attention: This method should not be called when ImpersonateXXX, SetTLSFingerPrint or
// SetTLSHandshake and other methods that will customize the tls handshake are called.
func (c *Client) EnableForceHTTP1() *Client { _ = "STUB: not implemented"; return nil }

// EnableForceHTTP2 enable force using HTTP2 for https requests (disabled by default).
//
// Attention: This method should not be called when ImpersonateXXX, SetTLSFingerPrint or
// SetTLSHandshake and other methods that will customize the tls handshake are called.
func (c *Client) EnableForceHTTP2() *Client { _ = "STUB: not implemented"; return nil }

// EnableForceHTTP3 enable force using HTTP3 for https requests (disabled by default).
//
// Attention: This method should not be called when ImpersonateXXX, SetTLSFingerPrint or
// SetTLSHandshake and other methods that will customize the tls handshake are called.
func (c *Client) EnableForceHTTP3() *Client { _ = "STUB: not implemented"; return nil }

// DisableForceHttpVersion disable force using specified http
// version (disabled by default).
func (c *Client) DisableForceHttpVersion() *Client { _ = "STUB: not implemented"; return nil }

// EnableH2C enables HTTP/2 over TCP without TLS.
func (c *Client) EnableH2C() *Client { _ = "STUB: not implemented"; return nil }

// DisableH2C disables HTTP/2 over TCP without TLS.
func (c *Client) DisableH2C() *Client { _ = "STUB: not implemented"; return nil }

// DisableAllowGetMethodPayload disable sending GET method requests with body.
func (c *Client) DisableAllowGetMethodPayload() *Client { _ = "STUB: not implemented"; return nil }

// EnableAllowGetMethodPayload allows sending GET method requests with body.
func (c *Client) EnableAllowGetMethodPayload() *Client { _ = "STUB: not implemented"; return nil }

func (c *Client) isPayloadForbid(m string) bool { _ = "STUB: not implemented"; return false }

// GetClient returns the underlying `http.Client`.
func (c *Client) GetClient() *http.Client { _ = "STUB: not implemented"; return nil }

func (c *Client) getRetryOption() *retryOption { _ = "STUB: not implemented"; return nil }

// SetCommonRetryCount enables retry and set the maximum retry count for requests
// fired from the client.
// It will retry infinitely if count is negative.
func (c *Client) SetCommonRetryCount(count int) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonRetryInterval sets the custom GetRetryIntervalFunc for requests fired
// from the client, you can use this to implement your own backoff retry algorithm.
// For example:
//
//		 req.SetCommonRetryInterval(func(resp *req.Response, attempt int) time.Duration {
//	     sleep := 0.01 * math.Exp2(float64(attempt))
//	     return time.Duration(math.Min(2, sleep)) * time.Second
//		 })
func (c *Client) SetCommonRetryInterval(getRetryIntervalFunc GetRetryIntervalFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonRetryFixedInterval set retry to use a fixed interval for requests
// fired from the client.
func (c *Client) SetCommonRetryFixedInterval(interval time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonRetryBackoffInterval set retry to use a capped exponential backoff
// with jitter for requests fired from the client.
// https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/
func (c *Client) SetCommonRetryBackoffInterval(min, max time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonRetryHook set the retry hook which will be executed before a retry.
// It will override other retry hooks if any been added before.
func (c *Client) SetCommonRetryHook(hook RetryHookFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// AddCommonRetryHook adds a retry hook for requests fired from the client,
// which will be executed before a retry.
func (c *Client) AddCommonRetryHook(hook RetryHookFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonRetryCondition sets the retry condition, which determines whether the
// request should retry.
// It will override other retry conditions if any been added before.
func (c *Client) SetCommonRetryCondition(condition RetryConditionFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// AddCommonRetryCondition adds a retry condition, which determines whether the
// request should retry.
func (c *Client) AddCommonRetryCondition(condition RetryConditionFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetUnixSocket set client to dial connection use unix socket.
// For example:
//
// client.SetUnixSocket("/var/run/custom.sock")
func (c *Client) SetUnixSocket(file string) *Client { _ = "STUB: not implemented"; return nil }

// DisableHTTP3 disables the http3 protocol.
func (c *Client) DisableHTTP3() *Client { _ = "STUB: not implemented"; return nil }

// EnableHTTP3 enables the http3 protocol.
func (c *Client) EnableHTTP3() *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2MaxHeaderListSize set the http2 MaxHeaderListSize,
// which is the http2 SETTINGS_MAX_HEADER_LIST_SIZE to
// send in the initial settings frame. It is how many bytes
// of response headers are allowed. Unlike the http2 spec, zero here
// means to use a default limit (currently 10MB). If you actually
// want to advertise an unlimited value to the peer, Transport
// interprets the highest possible value here (0xffffffff or 1<<32-1)
// to mean no limit.
func (c *Client) SetHTTP2MaxHeaderListSize(max uint32) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2StrictMaxConcurrentStreams set the http2
// StrictMaxConcurrentStreams, which controls whether the
// server's SETTINGS_MAX_CONCURRENT_STREAMS should be respected
// globally. If false, new TCP connections are created to the
// server as needed to keep each under the per-connection
// SETTINGS_MAX_CONCURRENT_STREAMS limit. If true, the
// server's SETTINGS_MAX_CONCURRENT_STREAMS is interpreted as
// a global limit and callers of RoundTrip block when needed,
// waiting for their turn.
func (c *Client) SetHTTP2StrictMaxConcurrentStreams(strict bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2ReadIdleTimeout set the http2 ReadIdleTimeout,
// which is the timeout after which a health check using ping
// frame will be carried out if no frame is received on the connection.
// Note that a ping response will is considered a received frame, so if
// there is no other traffic on the connection, the health check will
// be performed every ReadIdleTimeout interval.
// If zero, no health check is performed.
func (c *Client) SetHTTP2ReadIdleTimeout(timeout time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2PingTimeout set the http2 PingTimeout, which is the timeout
// after which the connection will be closed if a response to Ping is
// not received.
// Defaults to 15s
func (c *Client) SetHTTP2PingTimeout(timeout time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2WriteByteTimeout set the http2 WriteByteTimeout, which is the
// timeout after which the connection will be closed no data can be written
// to it. The timeout begins when data is available to write, and is
// extended whenever any bytes are written.
func (c *Client) SetHTTP2WriteByteTimeout(timeout time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// Do is compatible with http.Client.Do, which can make req integration easier
// in some scenarios. It should be noted that this will make some req features
// not work properly, such as automatic retry, client middleware, etc.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil,

		// NewClient is the alias of C
		nil
}

func NewClient() *Client {
	_ = "STUB: not implemented"

	// Clone copy and returns the Client
	return nil
}

func (c *Client) Clone() *Client {
	_ = "STUB: not implemented"

	// clone Transport
	return nil
}

// clone http.Client

// clone client middleware

// clone other fields that may need to be cloned

func memoryCookieJarFactory() *cookiejar.Jar { _ = "STUB: not implemented"; return nil }

// C create a new client.
func C() *Client { _ = "STUB: not implemented"; return nil }

// SetCookieJarFactory set the functional factory of cookie jar, which creates
// cookie jar that store cookies for underlying `http.Client`. After client clone,
// the cookie jar of the new client will also be regenerated using this factory
// function.
func (c *Client) SetCookieJarFactory(factory func() *cookiejar.Jar) *Client {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) initCookieJar() { _ = "STUB: not implemented"; return }

func (c *Client) initTransport() { _ = "STUB: not implemented"; return }

// RoundTripper is the interface of req's Client.
type RoundTripper interface {
	RoundTrip(*Request) (*Response, error)
}

// RoundTripFunc is a RoundTripper implementation, which is a simple function.
type RoundTripFunc func(req *Request) (resp *Response, err error)

// RoundTrip implements RoundTripper.
func (fn RoundTripFunc) RoundTrip(req *Request) (*Response, error) {
	_ = "STUB: not implemented"

	// RoundTripWrapper is client middleware function.
	return nil, nil
}

type RoundTripWrapper func(rt RoundTripper) RoundTripper

// RoundTripWrapperFunc is client middleware function, more convenient than RoundTripWrapper.
type RoundTripWrapperFunc func(rt RoundTripper) RoundTripFunc

func (f RoundTripWrapperFunc) wrapper() RoundTripWrapper {
	_ = "STUB: not implemented"
	return *new(RoundTripWrapper)
}

// WrapRoundTripFunc adds a client middleware function that will give the caller
// an opportunity to wrap the underlying http.RoundTripper.
func (c *Client) WrapRoundTripFunc(funcs ...RoundTripWrapperFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

type roundTripImpl struct {
	*Client
}

func (r roundTripImpl) RoundTrip(req *Request) (resp *Response, err error) {
	_ = "STUB: not implemented"
	return nil,

		// WrapRoundTrip adds a client middleware function that will give the caller
		// an opportunity to wrap the underlying http.RoundTripper.
		nil
}

func (c *Client) WrapRoundTrip(wrappers ...RoundTripWrapper) *Client {
	_ = "STUB: not implemented"
	return nil
}

// RoundTrip implements RoundTripper
func (c *Client) roundTrip(r *Request) (resp *Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// setup trace

// setup url and host

// Host header override

// setup header

// auto-read response body if possible

// restore body for re-reads
