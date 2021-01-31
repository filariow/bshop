package usecase

import "context"

type DeleteBillFunc func(context.Context, uint64) error

type billEraser interface {
	Delete(context.Context, uint64) error
}

func DeleteBill(bp billEraser) DeleteBillFunc {
	return func(ctx context.Context, id uint64) error {
		return bp.Delete(ctx, id)
	}
}
