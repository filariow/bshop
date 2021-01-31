package usecase

import (
	"context"
	"log"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/errs"
)

type CreateBillFunc func(context.Context, bshop.Bill) (uint64, error)

type billCreator interface {
	Create(context.Context, bshop.Bill) (uint64, error)
}

func CreateBill(bp billCreator) CreateBillFunc {
	isValid := func(b bshop.Bill) (bool, map[string]string) {
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

	return func(ctx context.Context, b bshop.Bill) (uint64, error) {
		if ok, ee := isValid(b); !ok {
			return 0, &errs.ErrValidation{Errors: ee}
		}
		return bp.Create(ctx, b)
	}
}
