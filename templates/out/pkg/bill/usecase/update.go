package usecase

import (
	"context"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/errs"
)

type UpdateBillFunc func(context.Context, bshop.Bill) error

type billUpdater interface {
	Update(context.Context, bshop.Bill) error
}

func UpdateBill(bp billUpdater) UpdateBillFunc {
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
		return len(ee) == 0, ee
	}

	return func(ctx context.Context, b bshop.Bill) error {
		if ok, ee := isValid(b); !ok {
			return &errs.ErrValidation{Errors: ee}
		}
		return bp.Update(ctx, b)
	}
}
