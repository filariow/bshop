package usecase

import (
	"context"

	"github.com/filariow/bshop"
)

type Read{{.ent}}Func func(context.Context, uint64) (bshop.{{.ent}}, error)

type {{.ent | lower }}Provider interface {
	Read(context.Context, uint64) (bshop.{{.ent}}, error)
}

func Read{{.ent}}(bp {{.ent | lower }}Provider) Read{{.ent}}Func {
	return func(ctx context.Context, id uint64) (bshop.{{.ent}}, error) {
		return bp.Read(ctx, id)
	}
}
