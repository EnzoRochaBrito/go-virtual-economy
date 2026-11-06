package helper

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	sqlcRepo "virtual.economy/src/internal/adapters/postgresql/sqlc/gen"
)

type RepoTxS struct {
	poosl *pgxpool.Pool
}

type RepoTx interface {
	InitTx(ctx context.Context) (*sqlcRepo.Queries, pgx.Tx, error)
}
