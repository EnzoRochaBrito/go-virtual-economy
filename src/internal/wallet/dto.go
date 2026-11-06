package wallet

import (
	"time"

	veUtils "virtual.economy/src/shared/utils"
)

type GetWalletInput struct {
	ProfileId string `json:"profileId"`
}

type GetBalanceResult struct {
	Balance veUtils.Currency `json:"balance"`
}

type WalletLedge struct {
	ExchangeId string               `json:"exchangeId"`
	Amount     veUtils.Currency     `json:"amount"`
	Movement   veUtils.MovementType `json:"movement"`
	CreatedAt  time.Time            `json:"createdAt"`
}

type GetProfileLedgesResult struct {
	Ledges []WalletLedge `json:"movements"`
}