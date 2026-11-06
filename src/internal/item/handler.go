package item

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

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body CreateItemInput
	json.Read(r, &body)
	item, err := h.service.CreateItem(ctx, body)
	if err != nil {
		json.Write(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusCreated, item)
}

func (h *Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	item, err := h.service.GetItem(ctx, id)
	if err != nil {
		json.Write(w, http.StatusNotFound, err.Error())
		return
	}
	json.Write(w, http.StatusCreated, item)
}

func (h *Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")

	err := h.service.DeleteItem(ctx, DeleteItemInput{
		ItemId:    id,
		ProfileId: "MOCK-ID",
	})

	if err != nil {
		json.Write(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusNoContent, `{"success": true}`)
}