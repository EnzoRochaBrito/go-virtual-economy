package market

import "net/http"

func Router(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("POST /market/item/{id}", h.BuyItem)
}