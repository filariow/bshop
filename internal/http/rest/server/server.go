package server

import (
	"net/http"
)

type server struct {
	mux http.Handler

	Controllers []ControllerRegistration
}

type ControllerRegistration struct {
	Path    string
	Handler http.Handler
}

func New(controllers []ControllerRegistration) http.Handler {
	s := server{
		Controllers: controllers,
	}
	s.registerRoutes()
	return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
