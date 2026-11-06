package profile

type CreateProfileDTO struct {
	Id   string
	Name string
}

type CreateProfileInput struct {
	Name string `json:"name"`
}

type CreateProfileResponse struct {
	Id string `json:"id"`
}

type UpdateProfileDTO struct {
	Id   string
	Name string
}

type UpdateProfileInput struct {
	Name string `json:"name"`
}

type AddItemDTO struct {
	ProfileId string
	ItemId    string
}

type HasItemDTO struct {
	ProfileId string
	ItemId    string
}


type GetInventoryResponse struct {
	Inventory []InventoryItem `json:"inventory"`
}