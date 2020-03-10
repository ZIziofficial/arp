package reverse_proxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	// black hole logging
	if os.Getenv("DEBUG") != "true" {
		log.SetOutput(ioutil.Discard)
	}
}

func TestNewReverseProxy(t *testing.T) {
	var backendCalled bool
	backendServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		backendCalled = true
	}))
	defer backendServer.Close()

	proxy, err := NewReverseProxy(backendServer.URL)

	if err != nil {
		t.Fatalf("Expected no errors (NewReverseProxy), got %v", err)
	}

	if proxy.ModifyResponse == nil {
		t.Fatal("Expected ModifyResponse to be set")
	}

	frontendProxy := httptest.NewServer(proxy)
	defer frontendProxy.Close()

	//resp, err := http.Get(frontendProxy.URL)
	_, err = http.Get(frontendProxy.URL)
	if err != nil {
		t.Fatalf("Expected no errors (http.Get), got %v", err)
	}

	if !backendCalled {
		t.Fatal("Expected a call to the backend, but it didn't happen")
	}
}

func TestNewReverseProxy_Director(t *testing.T) {
	proxy, err := NewReverseProxy("https://www.example.com")
	if err != nil {
		t.Fatalf("Expected no errors (NewReverseProxy), got %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/foo/bar", nil)
	if err != nil {
		t.Fatalf("Expected no errors (http.NewRequest), got %v", err)
	}

	proxy.Director(req)

	// Only testing modifications from our version of Director, not those that
	// come from stdlib Director
	//
	// By not testing stdlib Director, we lose assurance that it's actually being
	// executed, however, we avoid adding fragility should it change in the future.
	if req.Header.Get("X-Proxy-UUID") == "" {
		t.Fatal("Expected Header[X-Proxy-UUID] to be set")
	}

	if req.Header.Get("X-Forwarded-For") != "" {
		t.Fatalf("Expected Header[X-Proxy-UUID] to NOT be set, got: %s",
			req.Header.Get("X-Forwarded-For"))
	}

	if req.Host != "www.example.com" {
		t.Fatalf("Expected Host to be www.example.com, got: %s", req.Host)
	}
}
