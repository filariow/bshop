package main

import (
	"log"
	"net/http"
	"os"

	"github.com/filariow/bshop/pkg/beer/http/rest"
	"github.com/filariow/bshop/pkg/beer/storage/inmem"
	"github.com/filariow/bshop/pkg/http/rest/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

const addrDefault = ":8080"

func run() error {
	addr := address()
	log.Println("Configuring server on address", addr)
	r := inmem.New()

	cc := []server.ControllerRegistration{
		{
			PathPrefix: "/beers",
			Controller: rest.New(r),
		},
	}
	s := server.New(cc)

	log.Println("Start listening...")
	if err := http.ListenAndServe(addr, s); err != nil {
		return err
	}
	return nil
}

func address() string {
	if address := os.Getenv("ADDRESS"); address != "" {
		return address
	}
	return addrDefault
}
