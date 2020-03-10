package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jmervine/arp/reverse_proxy"
)

var (
	backend    string
	port, bind string
	debug      bool
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("app=arp ")
}

func main() {
	configure()

	proxy, err := reverse_proxy.NewReverseProxy(backend)
	if err != nil {
		log.Fatalf("at=main error=\"%v\"", err)
	}

	listen := bind + ":" + port

	log.Printf("at=main listener=%s backend=%s", listen, backend)
	log.Fatal(http.ListenAndServe(listen, proxy))
}

func configure() {
	var ok bool

	if backend, ok = os.LookupEnv("BACKEND"); !ok {
		log.Fatal("BACKEND is required")
	}

	if port, ok = os.LookupEnv("PORT"); !ok {
		port = "3000"
	}

	if bind, ok = os.LookupEnv("bind"); !ok {
		bind = "0.0.0.0"
	}

	if os.Getenv("DEBUG") == "true" {
		debug = true
	}
}
