package item

import (
	"context"
	"errors"

	"virtual.economy/src/internal/exchange"
)

type Service interface {
	CreateItem(ctx context.Context, dto CreateItemInput) (Item, error)
	GetItem(ctx context.Context, id string) (Item, error)
	UpdateItem(ctx context.Context) (Item, error)
	DeleteItem(ctx context.Context, dto DeleteItemInput) error
}

type svc struct {
	repo     Repository
	exchange exchange.Service
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}

// CreateItem implements [Service].
func (s *svc) CreateItem(ctx context.Context, dto CreateItemInput) (Item, error) {

	item := NewItem(dto.CreatorId, dto.Name, dto.Value)
	err := s.repo.CreateItem(ctx, item)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

// DeleteItem implements [Service].
func (s *svc) DeleteItem(ctx context.Context, dto DeleteItemInput) error {
	txQuery, tx, err := s.repo.InitTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Commit(ctx)
	item, err := txQuery.GetItem(ctx, dto.ItemId)
	if err != nil {
		return err
	}
	if item.CreatorID != dto.ProfileId {
		return errors.New("Profile cannot delete item")
	}
	return txQuery.DeleteItem(ctx, dto.ItemId)
}

// GetItem implements [Service].
func (s *svc) GetItem(ctx context.Context, id string) (Item, error) {
	// needs caching
	return s.repo.GetItem(ctx, id)
}

// UpdateItem implements [Service].
func (s *svc) UpdateItem(ctx context.Context) (Item, error) {
	panic("unimplemented")
}
