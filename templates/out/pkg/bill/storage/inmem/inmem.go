package inmem

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/bill/storage"
)

type BillRepository struct {
	bills map[uint64]bshop.Bill
	mux   sync.RWMutex

	id_counter uint64
}

func New() storage.BillRepository {
	return &BillRepository{
		bills:      map[uint64]bshop.Bill{},
		id_counter: 0,
	}
}

func (r *BillRepository) Create(ctx context.Context, bill bshop.Bill) (uint64, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	i := atomic.AddUint64(&r.id_counter, 1)
	bill.ID = i
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	r.bills[i] = bill
	return i, nil
}

func (r *BillRepository) Read(ctx context.Context, id uint64) (bshop.Bill, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	b, ok := r.bills[id]
	if !ok {
		return bshop.Bill{}, storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return bshop.Bill{}, err
	}
	return b, nil
}

func (r *BillRepository) Update(ctx context.Context, bill bshop.Bill) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.bills[bill.ID]; !ok {
		return storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return err
	}
	r.bills[bill.ID] = bill
	return nil
}

func (r *BillRepository) Delete(ctx context.Context, id uint64) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.bills[id]; !ok {
		return storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return err
	}
	delete(r.bills, id)
	return nil
}

func (r *BillRepository) List(ctx context.Context) ([]bshop.Bill, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	bb := make([]bshop.Bill, len(r.bills))
	i := uint64(0)
	for _, v := range r.bills {
		bb[i] = v
		i++
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return bb, nil
}
