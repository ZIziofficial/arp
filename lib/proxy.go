package proxy

import (
	"http/httputil"
)

type Proxy struct {
	Backend string

	proxy *httputil.ReverseProxy
}

func NewProxy(backend string) *Proxy {
	p := new(Proxy)
	p.Backend = backend
}
