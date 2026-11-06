package wallet

import (
	"context"
	"time"
)

type Repository interface {
	GetLastSnapshot(ctx context.Context, profileId string) (WalletSnapshot, error)
	SaveSnapshot(ctx context.Context, snapshot WalletSnapshot) error
	GetProfileLastLedgesFromDate(ctx context.Context, profileId string, from time.Time) ([]UserLedge, error)
}
