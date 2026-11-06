package profile

import (
	"context"
	"fmt"

	"virtual.economy/src/internal/exchange"
	"virtual.economy/src/pkg/uuid"
)

type Service interface {
	CreateProfile(ctx context.Context, p CreateProfileInput) (CreateProfileResponse, error)
	GetProfile(ctx context.Context, id string) (Profile, error)
	DeleteProfile(ctx context.Context, id string) error
	UpdateProfile(ctx context.Context, pdto UpdateProfileDTO) error
	AddItem(ctx context.Context, dto AddItemDTO) error
	HasItem(ctx context.Context, dto HasItemDTO) bool
	GetInventory(ctx context.Context, profileId string) (GetInventoryResponse, error)
}

func NewService(repo Repository, exchange exchange.Service) Service {
	return &svc{
		repo:     repo,
		exchange: exchange,
	}
}

type svc struct {
	repo     Repository
	exchange exchange.Service
}

func (s *svc) CreateProfile(ctx context.Context, p CreateProfileInput) (CreateProfileResponse, error) {
	genId := uuid.Generate().String()
	err := s.repo.CreateProfile(ctx, CreateProfileDTO{
		Id:   genId,
		Name: p.Name,
	})
	if err != nil {
		return CreateProfileResponse{}, err
	}
	fmt.Println("inside create profile service here aaa")
	profileSysExchangeId := s.exchange.GetProfileSystemExchangeId(genId)
	s.exchange.CreateExchange(ctx, &profileSysExchangeId)
	return CreateProfileResponse{Id: genId}, nil
}

func (s *svc) GetProfile(ctx context.Context, id string) (Profile, error) {
	return s.repo.GetProfile(ctx, id)
}
func (s *svc) DeleteProfile(ctx context.Context, id string) error {
	return s.repo.DeleteProfile(ctx, id)
}

func (s *svc) UpdateProfile(ctx context.Context, pdto UpdateProfileDTO) error {
	return s.repo.UpdateProfile(ctx, pdto)
}

func (s *svc) AddItem(ctx context.Context, dto AddItemDTO) error {
	return s.repo.AddItem(ctx, dto)
}

func (s *svc) HasItem(ctx context.Context, dto HasItemDTO) bool {
	return s.repo.HasItem(ctx, dto)
}

func (s *svc) GetInventory(ctx context.Context, profileId string) (GetInventoryResponse, error) {
	res, err := s.repo.GetInventory(ctx, profileId)
	if err != nil {
		return GetInventoryResponse{}, err
	}
	return GetInventoryResponse{
		Inventory: res,
	}, nil
}
