package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerRoutes() {
	r := mux.NewRouter()

	r.
		HandleFunc("/beers", s.handlerCreateBeer()).
		Methods(http.MethodPost)

	r.
		HandleFunc("/beers/{id}", s.handlerReadBeer()).
		Methods(http.MethodGet)

	r.
		HandleFunc("/beers/{id}", s.handlerUpdateBeer()).
		Methods(http.MethodPut)

	r.
		HandleFunc("/beers/{id}", s.handlerDeleteBeer()).
		Methods(http.MethodDelete)

	s.mux = r
}
