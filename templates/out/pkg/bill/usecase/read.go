package usecase

import (
	"context"

	"github.com/filariow/bshop"
)

type ReadBillFunc func(context.Context, uint64) (bshop.Bill, error)

type billProvider interface {
	Read(context.Context, uint64) (bshop.Bill, error)
}

func ReadBill(bp billProvider) ReadBillFunc {
	return func(ctx context.Context, id uint64) (bshop.Bill, error) {
		return bp.Read(ctx, id)
	}
}
