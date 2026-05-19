// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socks

import (
	"context"
	"net"
	"time"
)

var (
	noDeadline   = time.Time{}
	aLongTimeAgo = time.Unix(1, 0)
)

func (d *Dialer) connect(ctx context.Context, c net.Conn, address string) (_ net.Addr, ctxErr error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

// the size here is just an estimate

func splitHostPort(address string) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}
