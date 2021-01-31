package storage

import (
	"context"
	"fmt"

	"github.com/filariow/bshop"
)

type {{.ent}}Repository interface {
	Create(context.Context, bshop.{{.ent}}) (uint64, error)
	Read(context.Context, uint64) (bshop.{{.ent}}, error)
	Update(context.Context, bshop.{{.ent}}) error
	Delete(context.Context, uint64) error
	List(context.Context) ([]bshop.{{.ent}}, error)
}

var (
	ErrorNotFound = fmt.Errorf("requested entry not found in db")
)
