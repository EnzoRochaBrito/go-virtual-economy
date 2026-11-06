package profile

import (
	"fmt"
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

func (h *Handler) CreateProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input CreateProfileInput
	json.Read(r, &input)

	fmt.Println("ERROR HERERERER")
	res, err := h.service.CreateProfile(ctx, input)

	if err != nil {
		json.Write(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusCreated, res)
}

func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.PathValue("id")
	res, err := h.service.GetProfile(ctx, id)
	if err != nil {
		json.Write(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusCreated, res)
}

func (h *Handler) GetInventory(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	profileId := r.Header.Get("ProfileId")
	res, err := h.service.GetInventory(ctx, profileId)
	if err != nil {
		json.Write(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.Write(w, http.StatusOK, res)
}