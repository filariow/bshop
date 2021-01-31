package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func (s *server) registerRoutes() {
	r := mux.NewRouter()
	// Health
	r.HandleFunc("/health", s.handleHealth()).Methods(http.MethodGet)
	// Controllers
	for _, c := range s.Controllers {
		log.Println("Registering handler on path:", c.Path)
		r.PathPrefix(c.Path).Handler(c.Handler)
	}
	lr := handlers.LoggingHandler(os.Stdout, r)
	s.mux = lr
}
