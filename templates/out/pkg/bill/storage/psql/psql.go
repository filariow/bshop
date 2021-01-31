package psql

import (
	"context"
	"errors"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/bill/storage"
	"github.com/jackc/pgx/v4"
)

type Database struct {
	db *pgx.Conn
}

func New(conn *pgx.Conn) storage.BillRepository {
	return &Database{db: conn}
}

func (d *Database) Create(ctx context.Context, b bshop.Bill) (uint64, error) {
	sql := `
INSERT INTO bills(name, price, cost, size, vol, brewer)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id`
	id := uint64(0)
	err := d.db.QueryRow(ctx, sql, b.Name, b.Price, b.Cost, b.Size, b.Vol, b.Brewer).Scan(&id)
	return id, err
}

func (d *Database) Read(ctx context.Context, id uint64) (bshop.Bill, error) {
	sql := `
SELECT id, name, price, cost, size, vol, brewer
FROM bills
WHERE id = $1 and is_deleted=false`
	b := bshop.Bill{}
	err := d.db.QueryRow(ctx, sql, id).
		Scan(&b.ID, &b.Name, &b.Price, &b.Cost, &b.Size, &b.Vol, &b.Brewer)
	if errors.Is(err, pgx.ErrNoRows) {
		return b, storage.ErrorNotFound
	}
	return b, err
}

func (d *Database) Update(ctx context.Context, b bshop.Bill) error {
	sql := `
UPDATE bills
SET name=$2, price=$3, cost=$4, size=$5, vol=$6, brewer=$7
WHERE id = $1 and is_deleted=false`
	ct, err := d.db.Exec(ctx, sql, b.ID, b.Name, b.Price, b.Cost, b.Size, b.Vol, b.Brewer)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return storage.ErrorNotFound
	}
	return nil
}

func (d *Database) Delete(ctx context.Context, id uint64) error {
	sql := `
UPDATE bills
SET is_deleted=true
WHERE id=$1 and is_deleted=false`
	ct, err := d.db.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	if ct.RowsAffected() == 0 {
		return storage.ErrorNotFound
	}
	return nil
}

func (d *Database) List(ctx context.Context) ([]bshop.Bill, error) {
	sql := `
SELECT id, name, price, cost, size, vol, brewer
FROM bills
WHERE is_deleted=false`
	rr, err := d.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	bb := []bshop.Bill{}
	for rr.Next() {
		b := bshop.Bill{}
		rr.Scan(&b.ID, &b.Name, &b.Price, &b.Cost, &b.Size, &b.Vol, &b.Brewer)
		bb = append(bb, b)
	}
	return bb, nil
}
