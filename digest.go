package req

import (
	"net/http"
	"sync"

	"github.com/icholy/digest"
)

// cchal is a cached challenge and the number of times it's been used.
type cchal struct {
	c *digest.Challenge
	n int
}

type digestAuth struct {
	Username   string
	Password   string
	HttpClient *http.Client
	cache      map[string]*cchal
	cacheMu    sync.Mutex
}

func (da *digestAuth) digest(req *http.Request, chal *digest.Challenge, count int) (*digest.Credentials, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// challenge returns a cached challenge and count for the provided request
func (da *digestAuth) challenge(req *http.Request) (*digest.Challenge, int, bool) {
	_ = "STUB: not implemented"
	return nil, 0, false
}

// prepare attempts to find a cached challenge that matches the
// requested domain, and use it to set the Authorization header
func (da *digestAuth) prepare(req *http.Request) error {
	_ = "STUB: not implemented"
	// add cookies
	return nil
}

// add auth

func (da *digestAuth) HttpRoundTripWrapperFunc(rt http.RoundTripper) HttpRoundTripFunc {
	_ = "STUB: not implemented"
	return *new(HttpRoundTripFunc)
}

// make a copy of the request

// prepare the first request using a cached challenge

// the first request will either succeed or return a 401

// drain and close the first message body

// find and cache the challenge

// existing cached challenge didn't work, so remove it

// found new challenge, so cache it

// make a second copy of the request

// prepare the second request based on the new challenge

// create response middleware for http digest authentication.
func handleDigestAuthFunc(username, password string) ResponseMiddleware {
	_ = "STUB: not implemented"
	return *new(ResponseMiddleware)
}

// re-setup body

func createDigestAuth(req *http.Request, resp *http.Response, username, password string) (auth string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// cloner returns a function which makes clones of the provided request
func cloner(req *http.Request) (func() (*http.Request, error), error) {
	_ = "STUB: not implemented"
	return nil,

		// if there's no GetBody function set we have to copy the body
		// into memory to use for future clones
		nil
}
