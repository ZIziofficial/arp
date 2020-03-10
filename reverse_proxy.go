package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/google/uuid"
)

func NewReverseProxy(backend string) (*httputil.ReverseProxy, error) {
	var err error
	var target *url.URL
	var proxy *httputil.ReverseProxy

	if target, err = url.Parse(backend); err == nil {
		proxy = httputil.NewSingleHostReverseProxy(target)

		// Copy stdlib Director for reuse
		director := proxy.Director

		proxy.Director = func(req *http.Request) {
			req.Header.Set("X-Proxy-UUID", uuid.New().String())

			// for logging
			if req.URL.Scheme == "" {
				if req.TLS == nil {
					req.URL.Scheme = "http"
				} else {
					req.URL.Scheme = "https"
				}
			}

			log.Printf("uuid=%s on=request method=%s scheme=%s host=%s path=%s",
				req.Header.Get("X-Proxy-UUID"), req.Method, req.URL.Scheme, req.Host, req.URL.Path)

			if debug {
				log.Printf("req.Header: %v", req.Header)
			}

			// Exec stdlib director
			director(req)

			req.Header.Del("X-Forwarded-For")
			req.Host = target.Host

		}

		proxy.ModifyResponse = func(res *http.Response) error {
			req := res.Request
			log.Printf("uuid=%s on=proxied method=%s scheme=%s host=%s path=%s",
				req.Header.Get("X-Proxy-UUID"), req.Method, req.URL.Scheme, req.Host, req.URL.Path)

			if debug {
				log.Printf("res.Header: %v", res.Header)
			}

			return nil
		}
	}

	return proxy, err
}
