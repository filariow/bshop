package inmem

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/{{.ent | lower }}/storage"
)

type {{.ent}}Repository struct {
	{{.ent | lower }}s map[uint64]bshop.{{.ent}}
	mux   sync.RWMutex

	id_counter uint64
}

func New() storage.{{.ent}}Repository {
	return &{{.ent}}Repository{
		{{.ent | lower }}s:      map[uint64]bshop.{{.ent}}{},
		id_counter: 0,
	}
}

func (r *{{.ent}}Repository) Create(ctx context.Context, {{.ent | lower }} bshop.{{.ent}}) (uint64, error) {
	r.mux.Lock()
	defer r.mux.Unlock()

	i := atomic.AddUint64(&r.id_counter, 1)
	{{.ent | lower }}.ID = i
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	r.{{.ent | lower }}s[i] = {{.ent | lower }}
	return i, nil
}

func (r *{{.ent}}Repository) Read(ctx context.Context, id uint64) (bshop.{{.ent}}, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	b, ok := r.{{.ent | lower }}s[id]
	if !ok {
		return bshop.{{.ent}}{}, storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return bshop.{{.ent}}{}, err
	}
	return b, nil
}

func (r *{{.ent}}Repository) Update(ctx context.Context, {{.ent | lower }} bshop.{{.ent}}) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.{{.ent | lower }}s[{{.ent | lower }}.ID]; !ok {
		return storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return err
	}
	r.{{.ent | lower }}s[{{.ent | lower }}.ID] = {{.ent | lower }}
	return nil
}

func (r *{{.ent}}Repository) Delete(ctx context.Context, id uint64) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	if _, ok := r.{{.ent | lower }}s[id]; !ok {
		return storage.ErrorNotFound
	}

	if err := ctx.Err(); err != nil {
		return err
	}
	delete(r.{{.ent | lower }}s, id)
	return nil
}

func (r *{{.ent}}Repository) List(ctx context.Context) ([]bshop.{{.ent}}, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	bb := make([]bshop.{{.ent}}, len(r.{{.ent | lower }}s))
	i := uint64(0)
	for _, v := range r.{{.ent | lower }}s {
		bb[i] = v
		i++
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return bb, nil
}
