package handler

import (
	"encoding/json"
	"net/http"

	"spd-lookup/api/data"
)

// HandlerInterface describes handler methods
type HandlerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
	DescribeDepartments(w http.ResponseWriter, r *http.Request)
	SeattleOfficerMetadata(w http.ResponseWriter, r *http.Request)
	SeattleStrictMatch(w http.ResponseWriter, r *http.Request)
	SeattleFuzzySearch(w http.ResponseWriter, r *http.Request)
	TacomaOfficerMetadata(w http.ResponseWriter, r *http.Request)
	TacomaStrictMatch(w http.ResponseWriter, r *http.Request)
	TacomaFuzzySearch(w http.ResponseWriter, r *http.Request)
}

// Handler is the struct for route handler functions
type Handler struct {
	db data.DatabaseInterface
}

// NewHandler is the constructor for the handler
func NewHandler(db data.DatabaseInterface) *Handler {
	return &Handler{
		db: db,
	}
}

// Ping pong :^)
func (h *Handler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🏓 P O N G 🏓"))
}

// DescribeDepartments returns a list of departments and the fields supported for that department
func (h *Handler) DescribeDepartments(w http.ResponseWriter, r *http.Request) {
	departments := []*data.DepartmentMetadata{
		h.db.SeattleOfficerMetadata(),
		h.db.TacomaOfficerMetadata(),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(departments)
}
