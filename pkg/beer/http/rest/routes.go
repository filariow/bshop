package rest

import (
	"github.com/go-chi/chi"
)

func (c *controller) RegisterRoutes(r chi.Router) error {
	r.Post("/", c.handlerCreateBeer())
	r.Get("/{id}", c.handlerReadBeer())
	r.Put("/{id}", c.handlerUpdateBeer())
	r.Delete("/{id}", c.handlerDeleteBeer())
	r.Get("/", c.handlerListBeer())
	return nil
}
