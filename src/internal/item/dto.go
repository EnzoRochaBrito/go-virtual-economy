package item

import veUtils "virtual.economy/src/shared/utils"

type CreateItemDTO struct {
	CreatorId string           `json:"creatorId"`
	Name      string           `json:"name"`
	Value     veUtils.Currency `json:"value"`
}

type CreateItemInput struct {
	CreateItemDTO
}

type UpdateItemDTO struct {
	Name  string           `json:"name"`
	Value veUtils.Currency `json:"value"`
}

type UpdateItemInput struct {
	CreatorId string
	UpdateItemDTO
}

type DeleteItemInput struct {
	ItemId, ProfileId string
}
