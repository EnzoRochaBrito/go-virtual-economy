package exchange

import (
	"net/http"

	"virtual.economy/src/internal/json"
	veUtils "virtual.economy/src/shared/utils"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GiveAmount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profileId := r.PathValue("profileId")

	var input GiveAmountInput

	err := json.Read(r, &input)

	if err != nil {
		json.Write(w, http.StatusBadRequest, err)
		return
	}

	err = h.service.GiveAmount(ctx, GiveAmountDTO{
		ProfileId: profileId,
		Amount:    veUtils.Currency(input.Amount),
	})
	if err != nil {
		json.Write(w, http.StatusInternalServerError, err)
		return
	}
	json.Write(w, http.StatusOK, `{"success": true}`)
}

func (h *Handler) DrawAmount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	profileId := r.PathValue("profileId")

	var input DrawAmountInput

	err := json.Read(r, &input)

	if err != nil {
		json.Write(w, http.StatusBadRequest, err)
		return
	}
	err = h.service.DrawAmount(ctx, DrawAmountDTO{
		ProfileId: profileId,
		Amount:    veUtils.Currency(input.Amount),
	})
	if err != nil {
		json.Write(w, http.StatusInternalServerError, err)
		return
	}
	json.Write(w, http.StatusOK, `{"success": true}`)
}
