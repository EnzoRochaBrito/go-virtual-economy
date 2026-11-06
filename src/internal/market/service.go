package market

import (
	"context"
	"errors"

	"virtual.economy/src/internal/exchange"
	"virtual.economy/src/internal/item"
	"virtual.economy/src/internal/profile"
)

type Service interface {
	BuyItem(ctx context.Context, input BuyItemDTO) (BuyItemOutput, error)
}

type svc struct {
	item     item.Service
	exchange exchange.Service
	profile  profile.Service
}

func NewService(item item.Service, exchange exchange.Service, profile profile.Service) Service {
	return &svc{
		item,
		exchange,
		profile,
	}
}

// BuyItem implements [Service].
func (s *svc) BuyItem(ctx context.Context, input BuyItemDTO) (BuyItemOutput, error) {
	item, err := s.item.GetItem(ctx, input.ItemId)

	if err != nil {
		return BuyItemOutput{}, err
	}

	hasItem := s.profile.HasItem(ctx, profile.HasItemDTO{
		ProfileId: input.ProfileId,
		ItemId:    input.ItemId,
	})

	if hasItem {
		return BuyItemOutput{}, errors.New("Profile already have item")
	}

	exchangeId, err := s.exchange.Pay(ctx, exchange.PayInput{
		From:     input.ItemId,
		To:       item.CreatorId,
		Amount:   item.Value,
		Exchange: nil,
	})

	if err != nil {
		return BuyItemOutput{}, err
	}

	s.profile.AddItem(ctx, profile.AddItemDTO{
		ProfileId: input.ProfileId,
		ItemId:    input.ItemId,
	})

	return BuyItemOutput{
		ExchangeId: exchangeId,
	}, nil
}
