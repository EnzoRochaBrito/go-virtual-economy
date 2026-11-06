package item

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlcRepo "virtual.economy/src/internal/adapters/postgresql/sqlc/gen"
	"virtual.economy/src/internal/adapters/postgresql/sqlc/helper"
	veUtils "virtual.economy/src/shared/utils"
)

type SqlcRepository struct {
	db   *sqlcRepo.Queries
	pool *pgxpool.Pool
}

func NewSqlcRepository(pool *pgxpool.Pool) Repository {
	return &SqlcRepository{
		pool: pool,
		db:   sqlcRepo.New(pool),
	}
}

// CreateItem implements [Repository].
func (r *SqlcRepository) CreateItem(ctx context.Context, it Item) error {
	r.db.CreateItem(ctx, sqlcRepo.CreateItemParams{
		ID:        it.Id,
		CreatorID: it.CreatorId,
		Name:      it.Name,
		Value:     helper.Int4(it.Value.Int()),
	})
	return nil
}

// DeleteItem implements [Repository].
func (r *SqlcRepository) DeleteItem(ctx context.Context, id string) error {
	return r.db.DeleteItem(ctx, id)
}

// GetItem implements [Repository].
func (r *SqlcRepository) GetItem(ctx context.Context, id string) (Item, error) {
	item, err := r.db.GetItem(ctx, id)
	if err != nil {
		return Item{}, err
	}
	return Item{
		Id:        item.ID,
		CreatorId: item.CreatorID,
		Name:      item.Name,
		Value:     veUtils.Currency(item.Value.Int32),
		CreatedAt: item.CreatedAt.Time,
	}, nil
}

// UpdateItem implements [Repository].
func (r *SqlcRepository) UpdateItem(ctx context.Context, itdto UpdateItemDTO) error {
	panic("unimplemented")
}

func (r *SqlcRepository) InitTx(ctx context.Context) (*sqlcRepo.Queries, pgx.Tx, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return &sqlcRepo.Queries{}, &pgxpool.Tx{}, err
	}
	return r.db.WithTx(tx), tx, nil
}
