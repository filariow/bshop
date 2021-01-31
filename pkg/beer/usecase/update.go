package usecase

import (
	"context"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/errs"
)

type UpdateBeerFunc func(context.Context, bshop.Beer) error

type beerUpdater interface {
	Update(context.Context, bshop.Beer) error
}

func UpdateBeer(bp beerUpdater) UpdateBeerFunc {
	isValid := func(b bshop.Beer) (bool, map[string]string) {
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

	return func(ctx context.Context, b bshop.Beer) error {
		if ok, ee := isValid(b); !ok {
			return &errs.ErrValidation{Errors: ee}
		}
		return bp.Update(ctx, b)
	}
}
