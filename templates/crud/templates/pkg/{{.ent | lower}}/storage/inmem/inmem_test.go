package inmem_test

import (
	"context"
	"testing"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/{{.ent | lower }}/storage"
	"github.com/filariow/bshop/pkg/{{.ent | lower }}/storage/inmem"
	"github.com/matryer/is"
)

func Test_Create(t *testing.T) {
	cc := []struct {
		b bshop.{{.ent}}
	}{
		{
			b: bshop.{{.ent}}{
				Product: bshop.Product{
					Name:  "Test1",
					Price: 4.5,
					Cost:  2.5,
				},
				Size:   0.33,
				Vol:    4.5,
				Brewer: "Test1 Brewer",
			},
		},
	}

	for _, c := range cc {
		t.Run(c.b.Name, func(t *testing.T) {
			is := is.New(t)
			r := inmem.New()

			ctx := context.Background()
			id, err := r.Create(ctx, c.b)
			is.NoErr(err)
			is.True(id > 0)
			c.b.ID = id

			b, err := r.Read(ctx, id)
			is.NoErr(err)
			is.True(b == c.b)
		})
	}
}

func Test_CRUD(t *testing.T) {
	cc := []struct {
		b bshop.{{.ent}}
	}{
		{
			b: bshop.{{.ent}}{
				Product: bshop.Product{
					Name:  "Test1",
					Price: 4.5,
					Cost:  2.5,
				},
				Size:   0.33,
				Vol:    4.5,
				Brewer: "Test1 Brewer",
			},
		},
	}

	for _, c := range cc {
		t.Run(c.b.Name, func(t *testing.T) {
			is := is.New(t)
			r := inmem.New()

			ctx := context.Background()
			id, err := r.Create(ctx, c.b)
			is.NoErr(err)
			is.True(id > 0)
			c.b.ID = id

			c.b.Name = "Test1_Upd"
			err = r.Update(ctx, c.b)
			is.NoErr(err)

			b, err := r.Read(ctx, id)
			is.NoErr(err)
			is.True(b == c.b)

			err = r.Delete(ctx, id)
			is.NoErr(err)

			b, err = r.Read(ctx, id)
			is.True(err == storage.ErrorNotFound)
			is.True(b == bshop.{{.ent}}{})

		})
	}
}

func Test_List(t *testing.T) {
	cc := []struct {
		bb []bshop.{{.ent}}
	}{
		{

			bb: []bshop.{{.ent}}{
				{
					Product: bshop.Product{
						Name:  "Test1",
						Price: 4.5,
						Cost:  2.5,
					},
					Size:   0.33,
					Vol:    4.5,
					Brewer: "Test1 Brewer",
				},
				{
					Product: bshop.Product{
						Name:  "Test2",
						Price: 5.5,
						Cost:  3.5,
					},
					Size:   0.44,
					Vol:    5.5,
					Brewer: "Test2 Brewer",
				},
			},
		},
	}

	for _, c := range cc {
		t.Run("List", func(t *testing.T) {
			is := is.New(t)
			r := inmem.New()

			ctx := context.Background()
			for _, b := range c.bb {
				id, err := r.Create(ctx, b)
				is.NoErr(err)
				is.True(id > 0)
			}

			bb, err := r.List(ctx)
			is.True(len(c.bb) == len(bb))
			is.NoErr(err)
			for i, b := range c.bb {
				b.ID = uint64(i + 1)
				fb := func(id uint64) *bshop.{{.ent}} {
					for _, cb := range bb {
						if cb.ID == id {
							return &cb
						}
					}
					return nil
				}

				ub := fb(b.ID)
				is.True(b == *ub)
			}
		})
	}
}
