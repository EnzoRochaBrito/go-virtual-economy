package wallet

import "net/http"

func Router(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("GET /wallet", h.GetBalance)
	mux.HandleFunc("GET /wallet/movement", h.GetLedges)
}