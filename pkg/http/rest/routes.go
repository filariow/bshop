package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerRoutes() {
	r := mux.NewRouter()
	// Health
	{
		r.HandleFunc("/health", s.handleHealth()).Methods(http.MethodGet)
	}

	// Beers
	{
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

		r.
			HandleFunc("/beers", s.handlerListBeer()).
			Methods(http.MethodGet)
	}
	s.mux = r
}
