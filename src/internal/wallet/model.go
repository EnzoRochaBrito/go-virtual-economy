package wallet

import (
	"time"

	veUtils "virtual.economy/src/shared/utils"
)

type WalletSnapshot struct {
	ProfileId   string           `json:"profileId"`
	EffectiveAt time.Time        `json:"effectiveAt"`
	Balance     veUtils.Currency `json:"balance"`
}

type UserLedge struct {
	ExchangeId string               `json:"exchangeId"`
	Amount     veUtils.Currency     `json:"amount"`
	Movement   veUtils.MovementType `json:"movement"`
	CreatedAt  time.Time            `json:"createdAt"`
}