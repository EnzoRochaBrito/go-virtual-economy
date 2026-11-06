package exchange

import (
	"context"
	"errors"
	"time"

	"virtual.economy/src/internal/wallet"
	"virtual.economy/src/pkg/uuid"
	veUtils "virtual.economy/src/shared/utils"
)

type Service interface {
	CreateExchange(ctx context.Context, id *string) (string, error)
	Pay(ctx context.Context, input PayInput) (string, error)
	GiveAmount(ctx context.Context, input GiveAmountDTO) error
	DrawAmount(ctx context.Context, input DrawAmountDTO) error
	GetProfileSystemExchangeId(profileId string) string
}

func NewService(repo Repository, wallet wallet.Service, policy Policy) Service {
	return &svc{
		repo:   repo,
		wallet: wallet,
		policy: policy,
	}
}

type svc struct {
	repo   Repository
	wallet wallet.Service
	policy Policy
}

// Pay implements [Service].
func (s *svc) Pay(ctx context.Context, input PayInput) (string, error) {
	if input.Exchange == nil {
		exchangeId, err := s.CreateExchange(ctx, nil)
		if err != nil {
			return "", err
		}
		input.Exchange = new(exchangeId)
	}

	balance, err := s.wallet.GetBalance(ctx, input.From)

	if err != nil {
		return "", err
	}

	if input.Amount > balance.Balance {
		return "", errors.New("Insuficient balance")
	}

	ledges := s.genPayDoubleEntry(input)

	return *input.Exchange, s.repo.SaveLedges(ctx, ledges)
}

func (s *svc) CreateExchange(ctx context.Context, id *string) (string, error) {
	var exchangeId string
	if id == nil {
		exchangeId = uuid.Generate().String()
	} else {
		exchangeId = *id
	}
	err := s.repo.CreateExchange(ctx, exchangeId)
	return exchangeId, err
}

func (s *svc) genPayDoubleEntry(i PayInput) []Ledge {
	calculatedFee := s.policy.CalculateFee(i.Amount)
	return []Ledge{
		{
			ProfileId:  i.From, // payer
			Amount:     i.Amount,
			ExchangeId: *i.Exchange,
			Movement:   veUtils.MovementSend,
			CreatedAt:  time.Now(),
		},
		{
			ProfileId:  "SYSTEM_ID", // system,
			Amount:     calculatedFee.Fee,
			ExchangeId: *i.Exchange,
			Movement:   veUtils.MovementFee,
			CreatedAt:  time.Now(),
		},
		{
			ProfileId:  i.To, // reciever
			Amount:     calculatedFee.Net,
			ExchangeId: *i.Exchange,
			Movement:   veUtils.MovementReceive,
			CreatedAt:  time.Now(),
		},
	}
}

func (s *svc) GiveAmount(ctx context.Context, input GiveAmountDTO) error {
	if input.Amount.Int() <= 0 {
		return errors.New("Value must be greater than zero")
	}
	err := s.repo.SaveLedges(ctx, []Ledge{
		{
			ProfileId:  input.ProfileId,
			Amount:     input.Amount,
			ExchangeId: s.policy.ProfileSystemExchangeId(input.ProfileId),
			Movement:   veUtils.MovementExternalReceive,
			CreatedAt:  time.Now(),
		},
	})
	return err
}

func (s *svc) DrawAmount(ctx context.Context, input DrawAmountDTO) error {
	if input.Amount.Int() <= 0 {
		return errors.New("Value must be greater than zero")
	}
	return s.repo.SaveLedges(ctx, []Ledge{
		{
			ProfileId:  input.ProfileId,
			Amount:     input.Amount,
			ExchangeId: s.policy.ProfileSystemExchangeId(input.ProfileId),
			Movement:   veUtils.MovementExternalSend,
			CreatedAt:  time.Now(),
		},
	})
}

func (s *svc) GetProfileSystemExchangeId(profileId string) string {
	return s.policy.ProfileSystemExchangeId(profileId)
}
