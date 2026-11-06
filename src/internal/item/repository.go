package item

import (
	"context"

	"virtual.economy/src/internal/adapters/postgresql/sqlc/helper"
)

type Repository interface {
	helper.RepoTx
	CreateItem(ctx context.Context, it Item) error
	GetItem(ctx context.Context, id string) (Item, error)
	DeleteItem(ctx context.Context, id string) error
	UpdateItem(ctx context.Context, itdto UpdateItemDTO) error
}
