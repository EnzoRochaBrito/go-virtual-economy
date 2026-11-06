package profile

import (
	"time"

	"virtual.economy/src/pkg/uuid"
	veUtils "virtual.economy/src/shared/utils"
)

type Profile struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewProfile(name string) Profile {
	return Profile{
		Id: uuid.Generate().String(),
		Name: name,
		CreatedAt: time.Now(),
	}
}

type InventoryItem struct {
	ItemId string `json:"itemId"`
	ItemName string `json:"itemName"`
	PurchasedAt time.Time `json:"purchasedAt"`
	Value veUtils.Currency `json:"value"`
}