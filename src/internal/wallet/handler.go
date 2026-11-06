package wallet

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

func (h *Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profileId := r.Header.Get("ProfileId") // no jwt auth yet

	balance, err := h.service.GetBalance(ctx, profileId)
	if err != nil {
		json.Write(w, http.StatusInternalServerError, err)
		return
	}
	json.Write(w, http.StatusOK, balance)
}

func (h *Handler) GetLedges(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	profileId := r.Header.Get("ProfileId")

	res, err := h.service.GetProfileLedges(ctx, profileId)

	if err != nil {
		json.Write(w, http.StatusInternalServerError, err)
		return
	}
	json.Write(w, http.StatusOK, res)
}