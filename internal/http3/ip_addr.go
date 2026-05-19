package http3

import (
	"net"
)

// An addrList represents a list of network endpoint addresses.
// Copy from [net.addrList] and change type from [net.Addr] to [net.IPAddr]
type addrList []net.IPAddr

// isIPv4 reports whether addr contains an IPv4 address.
func isIPv4(addr net.IPAddr) bool { _ = "STUB: not implemented"; return false }

// isNotIPv4 reports whether addr does not contain an IPv4 address.
func isNotIPv4(addr net.IPAddr) bool { _ = "STUB: not implemented"; return false }

// forResolve returns the most appropriate address in address for
// a call to ResolveTCPAddr, ResolveUDPAddr, or ResolveIPAddr.
// IPv4 is preferred, unless addr contains an IPv6 literal.
func (addrs addrList) forResolve(network, addr string) net.IPAddr {
	_ = "STUB: not implemented"
	return *new(net.IPAddr)
}

// IPv6 literal (addr does NOT contain a port)

// IPv6 literal. (addr contains a port, so look for '[')

// first returns the first address which satisfies strategy, or if
// none do, then the first address of any kind.
func (addrs addrList) first(strategy func(net.IPAddr) bool) net.IPAddr {
	_ = "STUB: not implemented"
	return *new(net.IPAddr)
}
