package usecase

import (
	"context"
	"log"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/errs"
)

type Create{{.ent}}Func func(context.Context, bshop.{{.ent}}) (uint64, error)

type {{.ent | lower }}Creator interface {
	Create(context.Context, bshop.{{.ent}}) (uint64, error)
}

func Create{{.ent}}(bp {{.ent | lower }}Creator) Create{{.ent}}Func {
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
		log.Printf("Validation errors: %v", ee)
		return len(ee) == 0, ee
	}

	return func(ctx context.Context, b bshop.{{.ent}}) (uint64, error) {
		if ok, ee := isValid(b); !ok {
			return 0, &errs.ErrValidation{Errors: ee}
		}
		return bp.Create(ctx, b)
	}
}
