package usecase

import (
	"context"

	"github.com/filariow/bshop"
)

type ListBillFunc func(context.Context) ([]bshop.Bill, error)

type billsProvider interface {
	List(context.Context) ([]bshop.Bill, error)
}

func ListBill(bp billsProvider) ListBillFunc {
	return func(ctx context.Context) ([]bshop.Bill, error) {
		return bp.List(ctx)
	}
}
