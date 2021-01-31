package server

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (s *server) registerRoutes() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// Health
	r.Get("/health", s.handleHealth())

	// Controllers
	for _, c := range s.ControllerRegistrations {
		log.Println("Registering handler on path:", c.PathPrefix)
		r.Route(c.PathPrefix, func(r chi.Router) {
			c.Controller.RegisterRoutes(r)
		})
	}
	s.mux = r
}
