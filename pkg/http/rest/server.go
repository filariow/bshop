package rest

import (
	"net/http"

	"github.com/filariow/bshop/pkg/storage"
)

//Server REST server structure
type Server struct {
	mux http.Handler

	BeerRepo storage.BeerRepository
}

func (s *Server) Configure() {
	s.registerRoutes()
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
