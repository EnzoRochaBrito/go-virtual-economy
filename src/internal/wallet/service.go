package wallet

import (
	"context"
	"fmt"
	"time"

	veUtils "virtual.economy/src/shared/utils"
)

type Service interface {
	GetBalance(ctx context.Context, profileId string) (GetBalanceResult, error)
	SaveSnapshot(ctx context.Context, snapshot WalletSnapshot) error
	GetProfileLedges(ctx context.Context, profileId string) (GetProfileLedgesResult, error)
}

type svc struct {
	repo   Repository
	policy WalletPolicy
}

func NewService(repo Repository) Service {
	return &svc{
		repo: repo,
	}
}

// GetBalance implements [Service].
func (s *svc) GetBalance(ctx context.Context, profileId string) (GetBalanceResult, error) {
	walletSnapshot, err := s.repo.GetLastSnapshot(ctx, profileId)
	if err != nil {
		if err.Error() != "no rows in result set" {
			return GetBalanceResult{}, err
		}
		walletSnapshot = WalletSnapshot{
			ProfileId: profileId,
			Balance:   0,
			EffectiveAt: time.Time{},
		}
	}
	ledges, err := s.repo.GetProfileLastLedgesFromDate(ctx, profileId, walletSnapshot.EffectiveAt)
	if err != nil {
		if err.Error() != "no rows in result set" {
			return GetBalanceResult{}, err
		}
		ledges = []UserLedge{}
	}

	var balance veUtils.Currency = walletSnapshot.Balance

	if len(ledges) == 0 {
		return GetBalanceResult{
			Balance: balance,
		}, nil
	}

	for _, l := range ledges {
		fmt.Print("\n")
		fmt.Printf("balance: %v ", balance)
		fmt.Printf("l.Movement: %v ", l.Movement)
		fmt.Printf("l.Amount: %v -> ", l.Amount)
		balance = l.Movement.Result(balance, l.Amount)
		fmt.Printf("balance: %v\n", balance)
	}

	if s.policy.ShouldSnapshot(walletSnapshot, len(ledges)) {
		// this needs to be scheduled:
		s.SaveSnapshot(ctx, WalletSnapshot{
			ProfileId:   profileId,
			EffectiveAt: time.Now(),
			Balance:     balance,
		})
	}

	return GetBalanceResult{
		Balance: balance,
	}, nil
}

// SaveSnapshot implements [Service].
func (s *svc) SaveSnapshot(ctx context.Context, snapshot WalletSnapshot) error {
	return s.repo.SaveSnapshot(ctx, snapshot)
}

func (s *svc) GetProfileLedges(ctx context.Context, profileId string) (GetProfileLedgesResult, error) {
	ledges, err := s.repo.GetProfileLastLedgesFromDate(ctx, profileId, time.Time{})
	if err != nil {
		return GetProfileLedgesResult{}, err
	}
	var walletLedges []WalletLedge = []WalletLedge{}
	for _, l := range ledges {
		walletLedges = append(walletLedges, WalletLedge{
			ExchangeId: l.ExchangeId,
			Amount: l.Amount,
			Movement: l.Movement,
			CreatedAt: l.CreatedAt,
		})
	}
	return GetProfileLedgesResult{
		Ledges: walletLedges,
	}, nil
}