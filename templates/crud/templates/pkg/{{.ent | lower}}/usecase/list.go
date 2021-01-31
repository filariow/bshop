package usecase

import (
	"context"

	"github.com/filariow/bshop"
)

type List{{.ent}}Func func(context.Context) ([]bshop.{{.ent}}, error)

type {{.ent | lower }}sProvider interface {
	List(context.Context) ([]bshop.{{.ent}}, error)
}

func List{{.ent}}(bp {{.ent | lower }}sProvider) List{{.ent}}Func {
	return func(ctx context.Context) ([]bshop.{{.ent}}, error) {
		return bp.List(ctx)
	}
}
