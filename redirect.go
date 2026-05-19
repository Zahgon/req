package req

import (
	"net/http"
)

// RedirectPolicy represents the redirect policy for Client.
type RedirectPolicy func(req *http.Request, via []*http.Request) error

// MaxRedirectPolicy specifies the max number of redirect
func MaxRedirectPolicy(noOfRedirect int) RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

// DefaultRedirectPolicy allows up to 10 redirects
func DefaultRedirectPolicy() RedirectPolicy { _ = "STUB: not implemented"; return *new(RedirectPolicy) }

// NoRedirectPolicy disable redirect behaviour
func NoRedirectPolicy() RedirectPolicy { _ = "STUB: not implemented"; return *new(RedirectPolicy) }

// SameDomainRedirectPolicy allows redirect only if the redirected domain
// is the same as original domain, e.g. redirect to "www.imroc.cc" from
// "imroc.cc" is allowed, but redirect to "google.com" is not allowed.
func SameDomainRedirectPolicy() RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

// SameHostRedirectPolicy allows redirect only if the redirected host
// is the same as original host, e.g. redirect to "www.imroc.cc" from
// "imroc.cc" is not the allowed.
func SameHostRedirectPolicy() RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

// AllowedHostRedirectPolicy allows redirect only if the redirected host
// match one of the host that specified.
func AllowedHostRedirectPolicy(hosts ...string) RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

// AllowedDomainRedirectPolicy allows redirect only if the redirected domain
// match one of the domain that specified.
func AllowedDomainRedirectPolicy(hosts ...string) RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}

func getHostname(host string) (hostname string) { _ = "STUB: not implemented"; return "" }

func getDomain(host string) string { _ = "STUB: not implemented"; return "" }

// AlwaysCopyHeaderRedirectPolicy ensures that the given sensitive headers will
// always be copied on redirect.
// By default, golang will copy all of the original request's headers on redirect,
// unless they're sensitive, like "Authorization" or "Www-Authenticate". Only send
// sensitive ones to the same origin, or subdomains thereof (https://go-review.googlesource.com/c/go/+/28930/)
// Check discussion: https://github.com/golang/go/issues/4800
// For example:
//
//	client.SetRedirectPolicy(req.AlwaysCopyHeaderRedirectPolicy("Authorization"))
func AlwaysCopyHeaderRedirectPolicy(headers ...string) RedirectPolicy {
	_ = "STUB: not implemented"
	return *new(RedirectPolicy)
}
