package netutil

import (
	"net/url"
)

func AuthorityKey(u *url.URL) string { _ = "STUB: not implemented"; return "" }

// AuthorityAddr returns a given authority (a host/IP, or host:port / ip:port)
// and returns a host:port. The port 443 is added if needed.
func AuthorityAddr(scheme, authority string) (addr string) { _ = "STUB: not implemented"; return "" }

// IPv6 address literal, without a port:

func AuthorityHostPort(scheme, authority string) (host, port string) {
	_ = "STUB: not implemented"
	return "", ""
}

// authority didn't have a port
