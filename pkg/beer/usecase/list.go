package usecase

import (
	"context"

	"github.com/filariow/bshop"
)

type ListBeerFunc func(context.Context) ([]bshop.Beer, error)

type beersProvider interface {
	List(context.Context) ([]bshop.Beer, error)
}

func ListBeer(bp beersProvider) ListBeerFunc {
	return func(ctx context.Context) ([]bshop.Beer, error) {
		return bp.List(ctx)
	}
}
