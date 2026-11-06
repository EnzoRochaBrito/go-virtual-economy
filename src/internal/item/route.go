package item

import "net/http"

func Router(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("POST /item", h.CreateItem)
	mux.HandleFunc("GET /item/{id}", h.GetItem)
	mux.HandleFunc("DELETE /item/{id}", h.DeleteItem)
}