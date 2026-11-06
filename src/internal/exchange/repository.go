package exchange

import (
	"context"
)

type Repository interface {
	CreateExchange(ctx context.Context, id string) error
	GetExchange(ctx context.Context, id string) (Exchange, error)
	SaveLedges(ctx context.Context, ledges []Ledge) error
}
