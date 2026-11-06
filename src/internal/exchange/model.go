package exchange

import (
	"time"

	veUtils "virtual.economy/src/shared/utils"
)

type Exchange struct {
	Id     string `json:"id"`
	Ledges []ViewLedge
}

type Ledge struct {
	ProfileId  string               `json:"profileId"`
	ExchangeId string               `json:"exchangeId"`
	Amount     veUtils.Currency     `json:"amount"`
	Movement   veUtils.MovementType `json:"movement"`
	CreatedAt  time.Time            `json:"createdAt"`
}

type ViewLedge struct {
	ProfileId string               `json:"profileId"`
	Amount    veUtils.Currency     `json:"amount"`
	Movement  veUtils.MovementType `json:"movement"`
	CreatedAt time.Time            `json:"createdAt"`
}

// profile_id TEXT NOT NULL,
//     exchange_id TEXT NOT NULL,
//     amount INT,
//     movement movement_type NOT NULL,
//     created_at TIMESTAMPTZ DEFAULT NOW(),
