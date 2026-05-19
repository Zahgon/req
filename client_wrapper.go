package req

import (
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/imroc/req/v3/http2"
	utls "github.com/refraction-networking/utls"
)

// WrapRoundTrip is a global wrapper methods which delegated
// to the default client's Client.WrapRoundTrip.
func WrapRoundTrip(wrappers ...RoundTripWrapper) *Client { _ = "STUB: not implemented"; return nil }

// WrapRoundTripFunc is a global wrapper methods which delegated
// to the default client's Client.WrapRoundTripFunc.
func WrapRoundTripFunc(funcs ...RoundTripWrapperFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonError is a global wrapper methods which delegated
// to the default client's Client.SetCommonErrorResult.
//
// Deprecated: Use SetCommonErrorResult instead.
func SetCommonError(err any) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonErrorResult is a global wrapper methods which delegated
// to the default client's Client.SetCommonError.
func SetCommonErrorResult(err any) *Client { _ = "STUB: not implemented"; return nil }

// SetResultStateCheckFunc is a global wrapper methods which delegated
// to the default client's Client.SetCommonResultStateCheckFunc.
func SetResultStateCheckFunc(fn func(resp *Response) ResultState) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonFormDataFromValues is a global wrapper methods which delegated
// to the default client's Client.SetCommonFormDataFromValues.
func SetCommonFormDataFromValues(data url.Values) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonFormData is a global wrapper methods which delegated
// to the default client's Client.SetCommonFormData.
func SetCommonFormData(data map[string]string) *Client { _ = "STUB: not implemented"; return nil }

// SetMultipartBoundaryFunc is a global wrapper methods which delegated
// to the default client's Client.SetMultipartBoundaryFunc.
func SetMultipartBoundaryFunc(fn func() string) *Client { _ = "STUB: not implemented"; return nil }

// SetBaseURL is a global wrapper methods which delegated
// to the default client's Client.SetBaseURL.
func SetBaseURL(u string) *Client { _ = "STUB: not implemented"; return nil }

// SetOutputDirectory is a global wrapper methods which delegated
// to the default client's Client.SetOutputDirectory.
func SetOutputDirectory(dir string) *Client { _ = "STUB: not implemented"; return nil }

// SetCertFromFile is a global wrapper methods which delegated
// to the default client's Client.SetCertFromFile.
func SetCertFromFile(certFile, keyFile string) *Client { _ = "STUB: not implemented"; return nil }

// SetCerts is a global wrapper methods which delegated
// to the default client's Client.SetCerts.
func SetCerts(certs ...tls.Certificate) *Client { _ = "STUB: not implemented"; return nil }

// SetRootCertFromString is a global wrapper methods which delegated
// to the default client's Client.SetRootCertFromString.
func SetRootCertFromString(pemContent string) *Client { _ = "STUB: not implemented"; return nil }

// SetRootCertsFromFile is a global wrapper methods which delegated
// to the default client's Client.SetRootCertsFromFile.
func SetRootCertsFromFile(pemFiles ...string) *Client { _ = "STUB: not implemented"; return nil }

// GetTLSClientConfig is a global wrapper methods which delegated
// to the default client's Client.GetTLSClientConfig.
func GetTLSClientConfig() *tls.Config { _ = "STUB: not implemented"; return nil }

// SetRedirectPolicy is a global wrapper methods which delegated
// to the default client's Client.SetRedirectPolicy.
func SetRedirectPolicy(policies ...RedirectPolicy) *Client { _ = "STUB: not implemented"; return nil }

// DisableKeepAlives is a global wrapper methods which delegated
// to the default client's Client.DisableKeepAlives.
func DisableKeepAlives() *Client { _ = "STUB: not implemented"; return nil }

// EnableKeepAlives is a global wrapper methods which delegated
// to the default client's Client.EnableKeepAlives.
func EnableKeepAlives() *Client { _ = "STUB: not implemented"; return nil }

// DisableCompression is a global wrapper methods which delegated
// to the default client's Client.DisableCompression.
func DisableCompression() *Client { _ = "STUB: not implemented"; return nil }

// EnableCompression is a global wrapper methods which delegated
// to the default client's Client.EnableCompression.
func EnableCompression() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSClientConfig is a global wrapper methods which delegated
// to the default client's Client.SetTLSClientConfig.
func SetTLSClientConfig(conf *tls.Config) *Client { _ = "STUB: not implemented"; return nil }

// EnableInsecureSkipVerify is a global wrapper methods which delegated
// to the default client's Client.EnableInsecureSkipVerify.
func EnableInsecureSkipVerify() *Client { _ = "STUB: not implemented"; return nil }

// DisableInsecureSkipVerify is a global wrapper methods which delegated
// to the default client's Client.DisableInsecureSkipVerify.
func DisableInsecureSkipVerify() *Client { _ = "STUB: not implemented"; return nil }

// SetCommonQueryParams is a global wrapper methods which delegated
// to the default client's Client.SetCommonQueryParams.
func SetCommonQueryParams(params map[string]string) *Client { _ = "STUB: not implemented"; return nil }

// AddCommonQueryParam is a global wrapper methods which delegated
// to the default client's Client.AddCommonQueryParam.
func AddCommonQueryParam(key, value string) *Client { _ = "STUB: not implemented"; return nil }

// AddCommonQueryParams is a global wrapper methods which delegated
// to the default client's Client.AddCommonQueryParams.
func AddCommonQueryParams(key string, values ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonPathParam is a global wrapper methods which delegated
// to the default client's Client.SetCommonPathParam.
func SetCommonPathParam(key, value string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonPathParams is a global wrapper methods which delegated
// to the default client's Client.SetCommonPathParams.
func SetCommonPathParams(pathParams map[string]string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonQueryParam is a global wrapper methods which delegated
// to the default client's Client.SetCommonQueryParam.
func SetCommonQueryParam(key, value string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonQueryString is a global wrapper methods which delegated
// to the default client's Client.SetCommonQueryString.
func SetCommonQueryString(query string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonQueryParamsFromValues is a global wrapper methods which delegated
// to the default client's Client.SetCommonQueryParamsFromValues.
func SetCommonQueryParamsFromValues(params url.Values) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonQueryParamsFromStruct is a global wrapper methods which delegated
// to the default client's Client.SetCommonQueryParamsFromStruct.
func SetCommonQueryParamsFromStruct(v any) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonCookies is a global wrapper methods which delegated
// to the default client's Client.SetCommonCookies.
func SetCommonCookies(cookies ...*http.Cookie) *Client { _ = "STUB: not implemented"; return nil }

// DisableDebugLog is a global wrapper methods which delegated
// to the default client's Client.DisableDebugLog.
func DisableDebugLog() *Client { _ = "STUB: not implemented"; return nil }

// EnableDebugLog is a global wrapper methods which delegated
// to the default client's Client.EnableDebugLog.
func EnableDebugLog() *Client { _ = "STUB: not implemented"; return nil }

// DevMode is a global wrapper methods which delegated
// to the default client's Client.DevMode.
func DevMode() *Client { _ = "STUB: not implemented"; return nil }

// SetScheme is a global wrapper methods which delegated
// to the default client's Client.SetScheme.
func SetScheme(scheme string) *Client { _ = "STUB: not implemented"; return nil }

// SetLogger is a global wrapper methods which delegated
// to the default client's Client.SetLogger.
func SetLogger(log Logger) *Client { _ = "STUB: not implemented"; return nil }

// SetTimeout is a global wrapper methods which delegated
// to the default client's Client.SetTimeout.
func SetTimeout(d time.Duration) *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAll is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAll.
func EnableDumpAll() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllToFile is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllToFile.
func EnableDumpAllToFile(filename string) *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllTo is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllTo.
func EnableDumpAllTo(output io.Writer) *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllAsync is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllAsync.
func EnableDumpAllAsync() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutRequestBody is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllWithoutRequestBody.
func EnableDumpAllWithoutRequestBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutResponseBody is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllWithoutResponseBody.
func EnableDumpAllWithoutResponseBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutResponse is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllWithoutResponse.
func EnableDumpAllWithoutResponse() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutRequest is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllWithoutRequest.
func EnableDumpAllWithoutRequest() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutHeader is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllWithoutHeader.
func EnableDumpAllWithoutHeader() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpAllWithoutBody is a global wrapper methods which delegated
// to the default client's Client.EnableDumpAllWithoutBody.
func EnableDumpAllWithoutBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequest is a global wrapper methods which delegated
// to the default client's Client.EnableDumpEachRequest.
func EnableDumpEachRequest() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequestWithoutBody is a global wrapper methods which delegated
// to the default client's Client.EnableDumpEachRequestWithoutBody.
func EnableDumpEachRequestWithoutBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequestWithoutHeader is a global wrapper methods which delegated
// to the default client's Client.EnableDumpEachRequestWithoutHeader.
func EnableDumpEachRequestWithoutHeader() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequestWithoutResponse is a global wrapper methods which delegated
// to the default client's Client.EnableDumpEachRequestWithoutResponse.
func EnableDumpEachRequestWithoutResponse() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequestWithoutRequest is a global wrapper methods which delegated
// to the default client's Client.EnableDumpEachRequestWithoutRequest.
func EnableDumpEachRequestWithoutRequest() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequestWithoutResponseBody is a global wrapper methods which delegated
// to the default client's Client.EnableDumpEachRequestWithoutResponseBody.
func EnableDumpEachRequestWithoutResponseBody() *Client { _ = "STUB: not implemented"; return nil }

// EnableDumpEachRequestWithoutRequestBody is a global wrapper methods which delegated
// to the default client's Client.EnableDumpEachRequestWithoutRequestBody.
func EnableDumpEachRequestWithoutRequestBody() *Client { _ = "STUB: not implemented"; return nil }

// DisableAutoReadResponse is a global wrapper methods which delegated
// to the default client's Client.DisableAutoReadResponse.
func DisableAutoReadResponse() *Client { _ = "STUB: not implemented"; return nil }

// EnableAutoReadResponse is a global wrapper methods which delegated
// to the default client's Client.EnableAutoReadResponse.
func EnableAutoReadResponse() *Client { _ = "STUB: not implemented"; return nil }

// SetAutoDecodeContentType is a global wrapper methods which delegated
// to the default client's Client.SetAutoDecodeContentType.
func SetAutoDecodeContentType(contentTypes ...string) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetAutoDecodeContentTypeFunc is a global wrapper methods which delegated
// to the default client's Client.SetAutoDecodeAllTypeFunc.
func SetAutoDecodeContentTypeFunc(fn func(contentType string) bool) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetAutoDecodeAllContentType is a global wrapper methods which delegated
// to the default client's Client.SetAutoDecodeAllContentType.
func SetAutoDecodeAllContentType() *Client { _ = "STUB: not implemented"; return nil }

// DisableAutoDecode is a global wrapper methods which delegated
// to the default client's Client.DisableAutoDecode.
func DisableAutoDecode() *Client { _ = "STUB: not implemented"; return nil }

// EnableAutoDecode is a global wrapper methods which delegated
// to the default client's Client.EnableAutoDecode.
func EnableAutoDecode() *Client { _ = "STUB: not implemented"; return nil }

// SetUserAgent is a global wrapper methods which delegated
// to the default client's Client.SetUserAgent.
func SetUserAgent(userAgent string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonBearerAuthToken is a global wrapper methods which delegated
// to the default client's Client.SetCommonBearerAuthToken.
func SetCommonBearerAuthToken(token string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonBasicAuth is a global wrapper methods which delegated
// to the default client's Client.SetCommonBasicAuth.
func SetCommonBasicAuth(username, password string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonDigestAuth is a global wrapper methods which delegated
// to the default client's Client.SetCommonDigestAuth.
func SetCommonDigestAuth(username, password string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonHeaders is a global wrapper methods which delegated
// to the default client's Client.SetCommonHeaders.
func SetCommonHeaders(hdrs map[string]string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonHeader is a global wrapper methods which delegated
// to the default client's Client.SetCommonHeader.
func SetCommonHeader(key, value string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonHeaderOrder is a global wrapper methods which delegated
// to the default client's Client.SetCommonHeaderOrder.
func SetCommonHeaderOrder(keys ...string) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonPseudoHeaderOder is a global wrapper methods which delegated
// to the default client's Client.SetCommonPseudoHeaderOder.
func SetCommonPseudoHeaderOder(keys ...string) *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2SettingsFrame is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2SettingsFrame.
func SetHTTP2SettingsFrame(settings ...http2.Setting) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2ConnectionFlow is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2ConnectionFlow.
func SetHTTP2ConnectionFlow(flow uint32) *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2HeaderPriority is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2HeaderPriority.
func SetHTTP2HeaderPriority(priority http2.PriorityParam) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2PriorityFrames is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2PriorityFrames.
func SetHTTP2PriorityFrames(frames ...http2.PriorityFrame) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetHTTP2MaxHeaderListSize is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2MaxHeaderListSize.
func SetHTTP2MaxHeaderListSize(max uint32) *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2StrictMaxConcurrentStreams is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2StrictMaxConcurrentStreams.
func SetHTTP2StrictMaxConcurrentStreams(strict bool) *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2ReadIdleTimeout is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2ReadIdleTimeout.
func SetHTTP2ReadIdleTimeout(timeout time.Duration) *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2PingTimeout is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2PingTimeout.
func SetHTTP2PingTimeout(timeout time.Duration) *Client { _ = "STUB: not implemented"; return nil }

// SetHTTP2WriteByteTimeout is a global wrapper methods which delegated
// to the default client's Client.SetHTTP2WriteByteTimeout.
func SetHTTP2WriteByteTimeout(timeout time.Duration) *Client { _ = "STUB: not implemented"; return nil }

// ImpersonateChrome is a global wrapper methods which delegated
// to the default client's Client.ImpersonateChrome.
func ImpersonateChrome() *Client { _ = "STUB: not implemented"; return nil }

// ImpersonateChrome is a global wrapper methods which delegated
// to the default client's Client.ImpersonateChrome.
func ImpersonateFirefox() *Client { _ = "STUB: not implemented"; return nil }

// ImpersonateChrome is a global wrapper methods which delegated
// to the default client's Client.ImpersonateChrome.
func ImpersonateSafari() *Client { _ = "STUB: not implemented"; return nil }

// SetCommonContentType is a global wrapper methods which delegated
// to the default client's Client.SetCommonContentType.
func SetCommonContentType(ct string) *Client { _ = "STUB: not implemented"; return nil }

// DisableDumpAll is a global wrapper methods which delegated
// to the default client's Client.DisableDumpAll.
func DisableDumpAll() *Client { _ = "STUB: not implemented"; return nil }

// SetCommonDumpOptions is a global wrapper methods which delegated
// to the default client's Client.SetCommonDumpOptions.
func SetCommonDumpOptions(opt *DumpOptions) *Client { _ = "STUB: not implemented"; return nil }

// SetProxy is a global wrapper methods which delegated
// to the default client's Client.SetProxy.
func SetProxy(proxy func(*http.Request) (*url.URL, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// OnBeforeRequest is a global wrapper methods which delegated
// to the default client's Client.OnBeforeRequest.
func OnBeforeRequest(m RequestMiddleware) *Client { _ = "STUB: not implemented"; return nil }

// OnAfterResponse is a global wrapper methods which delegated
// to the default client's Client.OnAfterResponse.
func OnAfterResponse(m ResponseMiddleware) *Client { _ = "STUB: not implemented"; return nil }

// SetProxyURL is a global wrapper methods which delegated
// to the default client's Client.SetProxyURL.
func SetProxyURL(proxyUrl string) *Client { _ = "STUB: not implemented"; return nil }

// DisableTraceAll is a global wrapper methods which delegated
// to the default client's Client.DisableTraceAll.
func DisableTraceAll() *Client { _ = "STUB: not implemented"; return nil }

// EnableTraceAll is a global wrapper methods which delegated
// to the default client's Client.EnableTraceAll.
func EnableTraceAll() *Client { _ = "STUB: not implemented"; return nil }

// SetCookieJar is a global wrapper methods which delegated
// to the default client's Client.SetCookieJar.
func SetCookieJar(jar http.CookieJar) *Client { _ = "STUB: not implemented"; return nil }

// GetCookies is a global wrapper methods which delegated
// to the default client's Client.GetCookies.
func GetCookies(url string) ([]*http.Cookie, error) { _ = "STUB: not implemented"; return nil, nil }

// ClearCookies is a global wrapper methods which delegated
// to the default client's Client.ClearCookies.
func ClearCookies() *Client { _ = "STUB: not implemented"; return nil }

// SetJsonMarshal is a global wrapper methods which delegated
// to the default client's Client.SetJsonMarshal.
func SetJsonMarshal(fn func(v any) ([]byte, error)) *Client { _ = "STUB: not implemented"; return nil }

// SetJsonUnmarshal is a global wrapper methods which delegated
// to the default client's Client.SetJsonUnmarshal.
func SetJsonUnmarshal(fn func(data []byte, v any) error) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetXmlMarshal is a global wrapper methods which delegated
// to the default client's Client.SetXmlMarshal.
func SetXmlMarshal(fn func(v any) ([]byte, error)) *Client { _ = "STUB: not implemented"; return nil }

// SetXmlUnmarshal is a global wrapper methods which delegated
// to the default client's Client.SetXmlUnmarshal.
func SetXmlUnmarshal(fn func(data []byte, v any) error) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetDialTLS is a global wrapper methods which delegated
// to the default client's Client.SetDialTLS.
func SetDialTLS(fn func(ctx context.Context, network, addr string) (net.Conn, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetDial is a global wrapper methods which delegated
// to the default client's Client.SetDial.
func SetDial(fn func(ctx context.Context, network, addr string) (net.Conn, error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSHandshakeTimeout is a global wrapper methods which delegated
// to the default client's Client.SetTLSHandshakeTimeout.
func SetTLSHandshakeTimeout(timeout time.Duration) *Client { _ = "STUB: not implemented"; return nil }

// EnableForceHTTP1 is a global wrapper methods which delegated
// to the default client's Client.EnableForceHTTP1.
func EnableForceHTTP1() *Client { _ = "STUB: not implemented"; return nil }

// EnableForceHTTP2 is a global wrapper methods which delegated
// to the default client's Client.EnableForceHTTP2.
func EnableForceHTTP2() *Client { _ = "STUB: not implemented"; return nil }

// EnableForceHTTP3 is a global wrapper methods which delegated
// to the default client's Client.EnableForceHTTP3.
func EnableForceHTTP3() *Client { _ = "STUB: not implemented"; return nil }

// EnableHTTP3 is a global wrapper methods which delegated
// to the default client's Client.EnableHTTP3.
func EnableHTTP3() *Client { _ = "STUB: not implemented"; return nil }

// DisableForceHttpVersion is a global wrapper methods which delegated
// to the default client's Client.DisableForceHttpVersion.
func DisableForceHttpVersion() *Client { _ = "STUB: not implemented"; return nil }

// EnableH2C is a global wrapper methods which delegated
// to the default client's Client.EnableH2C.
func EnableH2C() *Client { _ = "STUB: not implemented"; return nil }

// DisableH2C is a global wrapper methods which delegated
// to the default client's Client.DisableH2C.
func DisableH2C() *Client { _ = "STUB: not implemented"; return nil }

// DisableAllowGetMethodPayload is a global wrapper methods which delegated
// to the default client's Client.DisableAllowGetMethodPayload.
func DisableAllowGetMethodPayload() *Client { _ = "STUB: not implemented"; return nil }

// EnableAllowGetMethodPayload is a global wrapper methods which delegated
// to the default client's Client.EnableAllowGetMethodPayload.
func EnableAllowGetMethodPayload() *Client { _ = "STUB: not implemented"; return nil }

// SetCommonRetryCount is a global wrapper methods which delegated
// to the default client's Client.SetCommonRetryCount.
func SetCommonRetryCount(count int) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonRetryInterval is a global wrapper methods which delegated
// to the default client's Client.SetCommonRetryInterval.
func SetCommonRetryInterval(getRetryIntervalFunc GetRetryIntervalFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonRetryFixedInterval is a global wrapper methods which delegated
// to the default client's Client.SetCommonRetryFixedInterval.
func SetCommonRetryFixedInterval(interval time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonRetryBackoffInterval is a global wrapper methods which delegated
// to the default client's Client.SetCommonRetryBackoffInterval.
func SetCommonRetryBackoffInterval(min, max time.Duration) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetCommonRetryHook is a global wrapper methods which delegated
// to the default client's Client.SetCommonRetryHook.
func SetCommonRetryHook(hook RetryHookFunc) *Client { _ = "STUB: not implemented"; return nil }

// AddCommonRetryHook is a global wrapper methods which delegated
// to the default client's Client.AddCommonRetryHook.
func AddCommonRetryHook(hook RetryHookFunc) *Client { _ = "STUB: not implemented"; return nil }

// SetCommonRetryCondition is a global wrapper methods which delegated
// to the default client's Client.SetCommonRetryCondition.
func SetCommonRetryCondition(condition RetryConditionFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// AddCommonRetryCondition is a global wrapper methods which delegated
// to the default client's Client.AddCommonRetryCondition.
func AddCommonRetryCondition(condition RetryConditionFunc) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetResponseBodyTransformer is a global wrapper methods which delegated
// to the default client's Client.SetResponseBodyTransformer.
func SetResponseBodyTransformer(fn func(rawBody []byte, req *Request, resp *Response) (transformedBody []byte, err error)) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetUnixSocket is a global wrapper methods which delegated
// to the default client's Client.SetUnixSocket.
func SetUnixSocket(file string) *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprint is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprint.
func SetTLSFingerprint(clientHelloID utls.ClientHelloID) *Client {
	_ = "STUB: not implemented"
	return nil
}

// SetTLSFingerprintRandomized is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintRandomized.
func SetTLSFingerprintRandomized() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintChrome is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintChrome.
func SetTLSFingerprintChrome() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintAndroid is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintAndroid.
func SetTLSFingerprintAndroid() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprint360 is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprint360.
func SetTLSFingerprint360() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintEdge is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintEdge.
func SetTLSFingerprintEdge() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintFirefox is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintFirefox.
func SetTLSFingerprintFirefox() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintQQ is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintQQ.
func SetTLSFingerprintQQ() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintIOS is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintIOS.
func SetTLSFingerprintIOS() *Client { _ = "STUB: not implemented"; return nil }

// SetTLSFingerprintSafari is a global wrapper methods which delegated
// to the default client's Client.SetTLSFingerprintSafari.
func SetTLSFingerprintSafari() *Client { _ = "STUB: not implemented"; return nil }

// GetClient is a global wrapper methods which delegated
// to the default client's Client.GetClient.
func GetClient() *http.Client { _ = "STUB: not implemented"; return nil }

// NewRequest is a global wrapper methods which delegated
// to the default client's Client.NewRequest.
func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

// R is a global wrapper methods which delegated
// to the default client's Client.R().
func R() *Request { _ = "STUB: not implemented"; return nil }
