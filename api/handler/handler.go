package handler

import (
	"net/http"
	"spd-lookup/api/provider"
)

// Handler is the struct for route handler functions
type Handler struct {
	db provider.DatabaseInterface
}

// NewHandler is the constructor for the handler
func NewHandler() *Handler {
	return &Handler{
		db: provider.NewDBClient(),
	}
}

// Ping pong :^)
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🏓 P O N G 🏓"))
}
