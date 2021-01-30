package inmem

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/storage"
)

type BeerRepository struct {
	beers map[int64]bshop.Beer
	mux   sync.RWMutex

	id_counter int64
}

func New() storage.BeerRepository {
	return &BeerRepository{
		beers:      map[int64]bshop.Beer{},
		id_counter: 0,
	}
}

func (r *BeerRepository) Create(ctx context.Context, beer bshop.Beer) (int64, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	i := atomic.AddInt64(&r.id_counter, 1)
	beer.ID = i
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	r.beers[i] = beer
	return i, nil
}

func (r *BeerRepository) Read(ctx context.Context, id int64) (bshop.Beer, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	b, ok := r.beers[id]
	if !ok {
		return bshop.Beer{}, storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return bshop.Beer{}, err
	}
	return b, nil
}

func (r *BeerRepository) Update(ctx context.Context, beer bshop.Beer) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.beers[beer.ID]; !ok {
		return storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return err
	}
	r.beers[beer.ID] = beer
	return nil
}

func (r *BeerRepository) Delete(ctx context.Context, id int64) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.beers[id]; !ok {
		return storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return err
	}
	delete(r.beers, id)
	return nil
}

func (r *BeerRepository) List(ctx context.Context) ([]bshop.Beer, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	bb := make([]bshop.Beer, len(r.beers))
	i := int64(0)
	for _, v := range r.beers {
		bb[i] = v
		i++
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return bb, nil
}
