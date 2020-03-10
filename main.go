package main

import (
	"log"
	"os"

	"github.com/jmervine/arp/pkg/proxy"
)

var (
	backend string
	rp      *proxy.Proxy
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("app=arp ")
}

func main() {
	configure()
}

func configure() {
	var ok bool

	if backend, ok = os.LookupEnv("BACKEND"); !ok {
		log.Fatal("BACKEND is required")
	}

	rp = &proxy.Proxy{Backend: backend}
}
