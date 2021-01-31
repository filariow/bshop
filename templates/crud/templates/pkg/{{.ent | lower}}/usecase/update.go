package usecase

import (
	"context"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/errs"
)

type Update{{.ent}}Func func(context.Context, bshop.{{.ent}}) error

type {{.ent | lower }}Updater interface {
	Update(context.Context, bshop.{{.ent}}) error
}

func Update{{.ent}}(bp {{.ent | lower }}Updater) Update{{.ent}}Func {
	isValid := func(b bshop.{{.ent}}) (bool, map[string]string) {
		ee := map[string]string{}
		if b.Name == "" {
			ee["Name"] = "A Name is required"
		}

		if b.Cost < .0 {
			ee["Cost"] = "Cost must be bigger than or equal to 0"
		}

		if b.Price < .0 {
			ee["Price"] = "Price must be bigger than or equal to 0"
		}
		return len(ee) == 0, ee
	}

	return func(ctx context.Context, b bshop.{{.ent}}) error {
		if ok, ee := isValid(b); !ok {
			return &errs.ErrValidation{Errors: ee}
		}
		return bp.Update(ctx, b)
	}
}
