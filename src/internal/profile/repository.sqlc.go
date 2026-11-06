package profile

import (
	"context"

	sqlcRepo "virtual.economy/src/internal/adapters/postgresql/sqlc/gen"
	veUtils "virtual.economy/src/shared/utils"
)

// CreateProfile(ctx context.Context, p CreateProfileDTO) error
// GetProfile(ctx context.Context, id string) (Profile, error)
// DeleteProfile(ctx context.Context, id string) error
// UpdateProfile(ctx context.Context, pdto UpdateProfileDTO) error

type SqlcRepository struct {
	db *sqlcRepo.Queries
}

func NewSqlcRepository(db *sqlcRepo.Queries) Repository {
	return &SqlcRepository{
		db: db,
	}
}

func (r *SqlcRepository) CreateProfile(ctx context.Context, p CreateProfileDTO) error {
	return r.db.CreateProfile(ctx, sqlcRepo.CreateProfileParams{
		ID:   p.Id,
		Name: p.Name,
	})
}

func (r *SqlcRepository) GetProfile(ctx context.Context, id string) (Profile, error) {
	profile, err := r.db.GetProfile(ctx, id)
	if err != nil {
		return Profile{}, err
	}
	return Profile{
		Id:        id,
		Name:      profile.Name,
		CreatedAt: profile.CreatedAt.Time,
	}, nil
}

func (r *SqlcRepository) DeleteProfile(ctx context.Context, id string) error {
	return r.db.DeleteProfile(ctx, id)
}

func (r *SqlcRepository) UpdateProfile(ctx context.Context, pdto UpdateProfileDTO) error {
	return r.db.UpdateProfile(ctx, sqlcRepo.UpdateProfileParams{
		ID:   pdto.Id,
		Name: pdto.Name,
	})
}

func (r *SqlcRepository) AddItem(ctx context.Context, dto AddItemDTO) error {
	return r.db.CreateInventory(ctx, sqlcRepo.CreateInventoryParams{
		ProfileID: dto.ProfileId,
		ItemID:    dto.ItemId,
	})
}

func (r *SqlcRepository) HasItem(ctx context.Context, dto HasItemDTO) bool {
	_, err := r.db.GetInventoryItem(ctx, sqlcRepo.GetInventoryItemParams{
		ProfileID: dto.ProfileId,
		ItemID:    dto.ItemId,
	})
	if err != nil {
		return false
	}
	return true
}

func (r *SqlcRepository) GetInventory(ctx context.Context, profileId string) ([]InventoryItem, error) {
	inv, err := r.db.GetInventory(ctx, profileId)
	if err != nil {
		return []InventoryItem{}, err
	}
	var items []InventoryItem
	for _, i := range inv {
		items = append(items, InventoryItem{
			ItemId:      i.ID,
			ItemName:    i.Name,
			PurchasedAt: i.PurchasedAt.Time,
			Value:       veUtils.Currency(i.Value.Int32),
		})
	}
	return items, err
}
