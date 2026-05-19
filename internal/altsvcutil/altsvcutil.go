package altsvcutil

import (
	"bytes"
	"net/url"
	"time"

	"github.com/imroc/req/v3/pkg/altsvc"
)

type altAvcParser struct {
	*bytes.Buffer
}

// validOptionalPort reports whether port is either an empty string
// or matches /^:\d*$/
func validOptionalPort(port string) bool { _ = "STUB: not implemented"; return false }

// splitHostPort separates host and port. If the port is not valid, it returns
// the entire input as host, and it doesn't check the validity of the host.
// Unlike net.SplitHostPort, but per RFC 3986, it requires ports to be numeric.
func splitHostPort(hostPort string) (host, port string) { _ = "STUB: not implemented"; return "", "" }

// ParseHeader parses the AltSvc from header value.
func ParseHeader(value string) ([]*altsvc.AltSvc, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newAltSvcParser(value string) *altAvcParser { _ = "STUB: not implemented"; return nil }

var endOfTime = time.Date(9999, 12, 31, 23, 59, 59, 0, time.UTC)

func (p *altAvcParser) Parse() (as []*altsvc.AltSvc, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *altAvcParser) parseKv() (key, value string, haveNextField bool, err error) {
	_ = "STUB: not implemented"
	return "", "", false, nil
}

func (p *altAvcParser) parseOne() (as *altsvc.AltSvc, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// drain useless fields

// ConvertURL converts the raw request url to expected alt-svc's url.
func ConvertURL(a *altsvc.AltSvc, u *url.URL) *url.URL { _ = "STUB: not implemented"; return nil }
