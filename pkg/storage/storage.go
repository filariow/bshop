package storage

import (
	"context"

	"github.com/filariow/bshop"
)

type BeerRepository interface {
	CreateBeer(context.Context, bshop.Beer) (int64, error)
	ReadBeer(int64) (context.Context, bshop.Beer, error)
	UpdateBeer(context.Context, bshop.Beer) error
	DeleteBeer(context.Context, int64) error
	ListBeers(context.Context) ([]bshop.Beer, error)
}
