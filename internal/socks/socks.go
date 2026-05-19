// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package socks provides a SOCKS version 5 client implementation.
//
// SOCKS protocol version 5 is defined in RFC 1928.
// Username/Password authentication for SOCKS version 5 is defined in
// RFC 1929.
package socks

import (
	"context"
	"io"
	"net"
)

// A Command represents a SOCKS command.
type Command int

func (cmd Command) String() string { _ = "STUB: not implemented"; return "" }

// An AuthMethod represents a SOCKS authentication method.
type AuthMethod int

// A Reply represents a SOCKS command reply code.
type Reply int

func (code Reply) String() string { _ = "STUB: not implemented"; return "" }

// Wire protocol constants.
const (
	Version5 = 0x05

	AddrTypeIPv4 = 0x01
	AddrTypeFQDN = 0x03
	AddrTypeIPv6 = 0x04

	CmdConnect Command = 0x01 // establishes an active-open forward proxy connection
	cmdBind    Command = 0x02 // establishes a passive-open forward proxy connection

	AuthMethodNotRequired         AuthMethod = 0x00 // no authentication required
	AuthMethodUsernamePassword    AuthMethod = 0x02 // use username/password
	AuthMethodNoAcceptableMethods AuthMethod = 0xff // no acceptable authentication methods

	StatusSucceeded Reply = 0x00
)

// An Addr represents a SOCKS-specific address.
// Either Name or IP is used exclusively.
type Addr struct {
	Name string // fully-qualified domain name
	IP   net.IP
	Port int
}

// Network return "socks"
func (a *Addr) Network() string { _ = "STUB: not implemented"; return "" }

func (a *Addr) String() string { _ = "STUB: not implemented"; return "" }

// A Conn represents a forward proxy connection.
type Conn struct {
	net.Conn

	boundAddr net.Addr
}

// BoundAddr returns the address assigned by the proxy server for
// connecting to the command target address from the proxy server.
func (c *Conn) BoundAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

// A Dialer holds SOCKS-specific options.
type Dialer struct {
	cmd          Command // either CmdConnect or cmdBind
	proxyNetwork string  // network between a proxy server and a client
	proxyAddress string  // proxy server address

	// ProxyDial specifies the optional dial function for
	// establishing the transport connection.
	ProxyDial func(context.Context, string, string) (net.Conn, error)

	// AuthMethods specifies the list of request authentication
	// methods.
	// If empty, SOCKS client requests only AuthMethodNotRequired.
	AuthMethods []AuthMethod

	// Authenticate specifies the optional authentication
	// function. It must be non-nil when AuthMethods is not empty.
	// It must return an error when the authentication is failed.
	Authenticate func(context.Context, io.ReadWriter, AuthMethod) error
}

// DialContext connects to the provided address on the provided
// network.
//
// The returned error value may be a net.OpError. When the Op field of
// net.OpError contains "socks", the Source field contains a proxy
// server address and the Addr field contains a command target
// address.
//
// See func Dial of the net package of standard library for a
// description of the network and address parameters.
func (d *Dialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

// DialWithConn initiates a connection from SOCKS server to the target
// network and address using the connection c that is already
// connected to the SOCKS server.
//
// It returns the connection's local address assigned by the SOCKS
// server.
func (d *Dialer) DialWithConn(ctx context.Context, c net.Conn, network, address string) (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

func (d *Dialer) validateTarget(network, address string) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dialer) pathAddrs(address string) (proxy, dst net.Addr, err error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), *new(net.Addr), nil
}

// NewDialer returns a new Dialer that dials through the provided
// proxy server's network and address.
func NewDialer(network, address string) *Dialer { _ = "STUB: not implemented"; return nil }

const (
	authUsernamePasswordVersion = 0x01
	authStatusSucceeded         = 0x00
)

// UsernamePassword are the credentials for the username/password
// authentication method.
type UsernamePassword struct {
	Username string
	Password string
}

// Authenticate authenticates a pair of username and password with the
// proxy server.
func (up *UsernamePassword) Authenticate(ctx context.Context, rw io.ReadWriter, auth AuthMethod) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO(mikio): handle IO deadlines and cancellation if
// necessary
