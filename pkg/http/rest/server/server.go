package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	mux http.Handler

	ControllerRegistrations []ControllerRegistration
}

type ControllerRegistration struct {
	PathPrefix string
	Controller Controller
}

type Controller interface {
	RegisterRoutes(r chi.Router) error
}

func New(controllersRegistrations []ControllerRegistration) http.Handler {
	s := server{
		ControllerRegistrations: controllersRegistrations,
	}
	s.registerRoutes()
	return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
