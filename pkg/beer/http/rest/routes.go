package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (c *controller) registerRoutes(pathPrefix string) {
	r := mux.NewRouter()
	r.
		HandleFunc(pathPrefix, c.handlerCreateBeer()).
		Methods(http.MethodPost)

	r.
		HandleFunc(fmt.Sprintf("%s/{id}", pathPrefix), c.handlerReadBeer()).
		Methods(http.MethodGet)

	r.
		HandleFunc(fmt.Sprintf("%s/{id}", pathPrefix), c.handlerUpdateBeer()).
		Methods(http.MethodPut)

	r.
		HandleFunc(fmt.Sprintf("%s/{id}", pathPrefix), c.handlerDeleteBeer()).
		Methods(http.MethodDelete)

	r.
		HandleFunc(pathPrefix, c.handlerListBeer()).
		Methods(http.MethodGet)

	c.mux = r
}
