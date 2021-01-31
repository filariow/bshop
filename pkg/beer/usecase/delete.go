package usecase

import "context"

type DeleteBeerFunc func(context.Context, uint64) error

type beerEraser interface {
	Delete(context.Context, uint64) error
}

func DeleteBeer(bp beerEraser) DeleteBeerFunc {
	return func(ctx context.Context, id uint64) error {
		return bp.Delete(ctx, id)
	}
}
