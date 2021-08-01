package handler

import (
	"encoding/json"
	"net/http"

	"spd-lookup/api/data"
)

// Interface describes handler methods
type Interface interface {
	Ping(w http.ResponseWriter, r *http.Request)
	DescribeDepartments(w http.ResponseWriter, r *http.Request)
	SeattleOfficerMetadata(w http.ResponseWriter, r *http.Request)
	SeattleStrictMatch(w http.ResponseWriter, r *http.Request)
	SeattleStrictMatchHistorical(w http.ResponseWriter, r *http.Request)
	SeattleFuzzySearch(w http.ResponseWriter, r *http.Request)
	TacomaOfficerMetadata(w http.ResponseWriter, r *http.Request)
	TacomaStrictMatch(w http.ResponseWriter, r *http.Request)
	TacomaFuzzySearch(w http.ResponseWriter, r *http.Request)
	PortlandOfficerMetadata(w http.ResponseWriter, r *http.Request)
	PortlandStrictMatch(w http.ResponseWriter, r *http.Request)
	PortlandFuzzySearch(w http.ResponseWriter, r *http.Request)
	AuburnOfficerMetadata(w http.ResponseWriter, r *http.Request)
	AuburnStrictMatch(w http.ResponseWriter, r *http.Request)
	AuburnFuzzySearch(w http.ResponseWriter, r *http.Request)
	LakewoodOfficerMetadata(w http.ResponseWriter, r *http.Request)
	LakewoodStrictMatch(w http.ResponseWriter, r *http.Request)
	LakewoodFuzzySearch(w http.ResponseWriter, r *http.Request)
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
	_, err := w.Write([]byte("🏓 P O N G 🏓"))
	if err != nil {
		return
	}
}

// DescribeDepartments returns a list of departments and the fields supported for that department
func (h *Handler) DescribeDepartments(w http.ResponseWriter, r *http.Request) {
	departments := []*data.DepartmentMetadata{
		h.db.SeattleOfficerMetadata(),
		h.db.TacomaOfficerMetadata(),
		h.db.PortlandOfficerMetadata(),
		h.db.AuburnOfficerMetadata(),
		h.db.LakewoodOfficerMetadata(),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(departments)
	if err != nil {
		return
	}
}
