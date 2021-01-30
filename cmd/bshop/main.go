package main

import (
	"log"
	"net/http"
	"os"

	"github.com/filariow/bshop/pkg/http/rest"
	"github.com/filariow/bshop/pkg/storage/inmem"
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
	s := rest.Server{BeerRepo: r}
	s.Configure()

	log.Println("Start listening...")
	if err := http.ListenAndServe(addr, &s); err != nil {
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
