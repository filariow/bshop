package storage

import (
	"context"
	"fmt"

	"github.com/filariow/bshop"
)

type BeerRepository interface {
	Create(context.Context, bshop.Beer) (int64, error)
	Read(context.Context, int64) (bshop.Beer, error)
	Update(context.Context, bshop.Beer) error
	Delete(context.Context, int64) error
	List(context.Context) ([]bshop.Beer, error)
}

var (
	ErrorNotFound = fmt.Errorf("requested entry not found in db")
)
