package exchange

import (
	"context"
	"errors"

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

// CreateExchange implements [Repository].
func (s *SqlcRepository) CreateExchange(ctx context.Context, id string) error {
	return s.db.CreateExchange(ctx, id)
}

// SaveLedges implements [Repository].
func (s *SqlcRepository) SaveLedges(ctx context.Context, ledges []Ledge) error {
	if len(ledges) == 0 {
		return errors.New("No ledge passed to save")
	} else if len(ledges) == 1 {
		return s.db.CreateLedge(ctx, sqlcRepo.CreateLedgeParams{
			ProfileID:  ledges[0].ProfileId,
			ExchangeID: ledges[0].ExchangeId,
			Amount:     helper.Int4(ledges[0].Amount.Int()),
			Movement:   sqlcRepo.MovementType(ledges[0].Movement),
		})
	}
	var adaptedLedges []sqlcRepo.CreateManyLedgesParams = []sqlcRepo.CreateManyLedgesParams{}
	for _, l := range ledges {
		adaptedLedges = append(adaptedLedges, sqlcRepo.CreateManyLedgesParams{
			ProfileID: l.ProfileId,
			ExchangeID: l.ExchangeId,
			Amount: helper.Int4(l.Amount.Int()),
			Movement: sqlcRepo.MovementType(l.Movement),
		})
	}
	_, err := s.db.CreateManyLedges(ctx, adaptedLedges)
	return err
}

// GetExchange implements [Repository].
func (s *SqlcRepository) GetExchange(ctx context.Context, id string) (Exchange, error) {
	exc, err := s.db.GetExchangeLedges(ctx, id)

	if err != nil {
		return Exchange{}, err
	}

	var exchange Exchange = Exchange{
		Id:     id,
		Ledges: []ViewLedge{},
	}

	for _, e := range exc {
		exchange.Ledges = append(exchange.Ledges, ViewLedge{
			ProfileId: e.ProfileID,
			Amount:    veUtils.Currency(e.Amount.Int32),
			Movement:  veUtils.MovementType(e.Movement),
			CreatedAt: e.CreatedAt.Time,
		})
	}

	return exchange, nil
}
