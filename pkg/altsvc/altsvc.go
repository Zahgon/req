package altsvc

import (
	"sync"
	"time"
)

// AltSvcJar is default implementation of Jar, which stores
// AltSvc in memory.
type AltSvcJar struct {
	entries map[string]*AltSvc
	mu      sync.Mutex
}

// NewAltSvcJar create a AltSvcJar which implements Jar.
func NewAltSvcJar() *AltSvcJar { _ = "STUB: not implemented"; return nil }

func (j *AltSvcJar) GetAltSvc(addr string) *AltSvc { _ = "STUB: not implemented"; return nil }

// expired

func (j *AltSvcJar) SetAltSvc(addr string, as *AltSvc) { _ = "STUB: not implemented"; return }

// AltSvc is the parsed alt-svc.
type AltSvc struct {
	// Protocol is the alt-svc proto, e.g. h3.
	Protocol string
	// Host is the alt-svc's host, could be empty if
	// it's the same host as the raw request.
	Host string
	// Port is the alt-svc's port.
	Port string
	// Expire is the time that the alt-svc should expire.
	Expire time.Time
}
