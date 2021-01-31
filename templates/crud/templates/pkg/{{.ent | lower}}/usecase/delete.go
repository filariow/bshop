package usecase

import "context"

type Delete{{.ent}}Func func(context.Context, uint64) error

type {{.ent | lower }}Eraser interface {
	Delete(context.Context, uint64) error
}

func Delete{{.ent}}(bp {{.ent | lower }}Eraser) Delete{{.ent}}Func {
	return func(ctx context.Context, id uint64) error {
		return bp.Delete(ctx, id)
	}
}
