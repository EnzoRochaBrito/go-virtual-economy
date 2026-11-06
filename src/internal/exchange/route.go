package exchange

import "net/http"

func Router(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("POST /profile/{profileId}/currency/give", h.GiveAmount)
	mux.HandleFunc("POST /profile/{profileId}/currency/draw", h.DrawAmount)
}
