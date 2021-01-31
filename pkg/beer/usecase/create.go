package usecase

import (
	"context"
	"log"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/errs"
)

type CreateBeerFunc func(context.Context, bshop.Beer) (uint64, error)

type beerCreator interface {
	Create(context.Context, bshop.Beer) (uint64, error)
}

func CreateBeer(bp beerCreator) CreateBeerFunc {
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
		log.Printf("Validation errors: %v", ee)
		return len(ee) == 0, ee
	}

	return func(ctx context.Context, b bshop.Beer) (uint64, error) {
		if ok, ee := isValid(b); !ok {
			return 0, &errs.ErrValidation{Errors: ee}
		}
		return bp.Create(ctx, b)
	}
}
