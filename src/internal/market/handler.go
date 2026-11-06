package market

import (
	"net/http"

	"virtual.economy/src/internal/json"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) BuyItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input BuyItemInput
	err := json.Read(r, &input)
	if err != nil {
		json.Write(w, http.StatusBadRequest, err)
		return
	}

	itemId := r.PathValue("id")

	if itemId == "" {
		json.Write(w, http.StatusBadRequest, `{"error": "invalid item id"}`)
	}

	res, err := h.service.BuyItem(ctx, BuyItemDTO{
		ProfileId: input.ProfileId,
		ItemId: itemId,
	})

	if err != nil {
		json.Write(w, http.StatusBadRequest, err)
		return
	}

	json.Write(w, http.StatusOK, res)
}
