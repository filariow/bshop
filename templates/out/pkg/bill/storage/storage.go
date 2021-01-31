package storage

import (
	"context"
	"fmt"

	"github.com/filariow/bshop"
)

type BillRepository interface {
	Create(context.Context, bshop.Bill) (uint64, error)
	Read(context.Context, uint64) (bshop.Bill, error)
	Update(context.Context, bshop.Bill) error
	Delete(context.Context, uint64) error
	List(context.Context) ([]bshop.Bill, error)
}

var (
	ErrorNotFound = fmt.Errorf("requested entry not found in db")
)
