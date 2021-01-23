package rest

import (
	"github.com/gorilla/mux"
)

func (s *Server) registerRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", nil)

	s.mux = r
}
