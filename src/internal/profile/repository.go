package profile

import (
	"context"
)

type Repository interface {
	CreateProfile(ctx context.Context, p CreateProfileDTO) error
	GetProfile(ctx context.Context, id string) (Profile, error)
	DeleteProfile(ctx context.Context, id string) error
	UpdateProfile(ctx context.Context, pdto UpdateProfileDTO) error
	AddItem(ctx context.Context, dto AddItemDTO) error
	HasItem(ctx context.Context, dto HasItemDTO) bool
	GetInventory(ctx context.Context, profileId string) ([]InventoryItem, error)
}
