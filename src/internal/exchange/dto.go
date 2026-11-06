package exchange

import veUtils "virtual.economy/src/shared/utils"

type PayInput struct {
	From, To string
	Amount   veUtils.Currency
	Exchange *string
}

type GiveAmountInput struct {
	Amount int `json:"amount"`
}

type GiveAmountDTO struct {
	ProfileId string
	Amount veUtils.Currency
}

type DrawAmountInput struct {
	Amount int `json:"amount"`
}

type DrawAmountDTO struct {
	ProfileId string
	Amount veUtils.Currency
}