package wallet

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	sqlcRepo "virtual.economy/src/internal/adapters/postgresql/sqlc/gen"
	"virtual.economy/src/internal/adapters/postgresql/sqlc/helper"
	veUtils "virtual.economy/src/shared/utils"
)

type SqlcRepository struct {
	db *sqlcRepo.Queries
}

func NewSqlcRepository(db *sqlcRepo.Queries) Repository {
	return &SqlcRepository{
		db: db,
	}
}

// GetLastSnapshot implements [Repository].
func (repo *SqlcRepository) GetLastSnapshot(ctx context.Context, profileId string) (WalletSnapshot, error) {
	snapshot, err := repo.db.GetLastSnapshot(ctx, profileId)
	if err != nil {
		return WalletSnapshot{}, err
	}
	return WalletSnapshot{
		ProfileId:   snapshot.ProfileID,
		Balance:     veUtils.Currency(snapshot.Balance.Int32),
		EffectiveAt: snapshot.EffectiveAt.Time,
	}, err
}

// SaveSnapshot implements [Repository].
func (repo *SqlcRepository) SaveSnapshot(ctx context.Context, snapshot WalletSnapshot) error {
	return repo.db.CreateWalletSnapshot(ctx, sqlcRepo.CreateWalletSnapshotParams{
		ProfileID:   snapshot.ProfileId,
		Balance:     helper.Int4(snapshot.Balance.Int()),
		EffectiveAt: helper.Date(snapshot.EffectiveAt),
	})
}

func (repo *SqlcRepository) GetProfileLastLedgesFromDate(ctx context.Context, profileId string, from time.Time) ([]UserLedge, error) {
	fetchedLedges, err := repo.db.GetProfileLedgesFromDate(ctx, sqlcRepo.GetProfileLedgesFromDateParams{
		ProfileID: profileId,
		CreatedAt: pgtype.Timestamptz(helper.Date(from)),
	})

	if err != nil {
		return []UserLedge{}, err
	}

	var ledges []UserLedge = []UserLedge{}

	for _, l := range fetchedLedges {
		ledges = append(ledges, UserLedge{
			ExchangeId: l.ExchangeID,
			Amount:     veUtils.Currency(l.Amount.Int32),
			Movement:   veUtils.MovementType(l.Movement),
			CreatedAt:  l.CreatedAt.Time,
		})
	}

	return ledges, nil
}
