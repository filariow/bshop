package usecase

import (
	"context"

	"github.com/filariow/bshop"
)

type ReadBeerFunc func(context.Context, uint64) (bshop.Beer, error)

type beerProvider interface {
	Read(context.Context, uint64) (bshop.Beer, error)
}

func ReadBeer(bp beerProvider) ReadBeerFunc {
	return func(ctx context.Context, id uint64) (bshop.Beer, error) {
		return bp.Read(ctx, id)
	}
}
