package rest

import (
	"github.com/go-chi/chi"
)

func (c *controller) RegisterRoutes(r chi.Router) error {
	r.Post("/", c.handlerCreateBill())
	r.Get("/{id}", c.handlerReadBill())
	r.Put("/{id}", c.handlerUpdateBill())
	r.Delete("/{id}", c.handlerDeleteBill())
	r.Get("/", c.handlerListBill())
	return nil
}
