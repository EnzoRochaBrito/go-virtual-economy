package profile

import "net/http"

func Router(h *Handler, mux *http.ServeMux) {
	mux.HandleFunc("POST /profile", h.CreateProfile)
	mux.HandleFunc("GET /profile/{id}", h.GetProfile)
	mux.HandleFunc("GET /inventory", h.GetInventory)
}