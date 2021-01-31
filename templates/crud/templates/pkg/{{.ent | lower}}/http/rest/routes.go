package rest

import (
	"github.com/go-chi/chi"
)

func (c *controller) RegisterRoutes(r chi.Router) error {
	r.Post("/", c.handlerCreate{{.ent}}())
	r.Get("/{id}", c.handlerRead{{.ent}}())
	r.Put("/{id}", c.handlerUpdate{{.ent}}())
	r.Delete("/{id}", c.handlerDelete{{.ent}}())
	r.Get("/", c.handlerList{{.ent}}())
	return nil
}
