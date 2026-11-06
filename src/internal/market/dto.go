package market

type BuyItemInput struct {
	ProfileId string `json:"profileId"`
}

type BuyItemDTO struct {
	ProfileId string
	ItemId    string
}

type BuyItemOutput struct {
	ExchangeId string
}
