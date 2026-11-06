package item

import (
	"time"

	"virtual.economy/src/pkg/uuid"
	veUtils "virtual.economy/src/shared/utils"
)

type Item struct {
	Id        string           `json:"id"`
	CreatorId string           `json:"creatorId"`
	Name      string           `json:"name"`
	Value     veUtils.Currency `json:"value"`
	CreatedAt time.Time        `json:"createdAt"`
}

func NewItem(creatorId, name string, value veUtils.Currency) Item {
	return Item{
		Id:        uuid.Generate().String(),
		CreatorId: creatorId,
		Name:      name,
		Value:     value,
		CreatedAt: time.Now(),
	}
}

// CREATE TABLE IF NOT EXISTS item (
//     id TEXT NOT NULL,
//     creator_id TEXT NOT NULL,
//     name TEXT NOT NULL,
//     value INT,
//     created_at DATE DEFAULT CURRENT_DATE,
