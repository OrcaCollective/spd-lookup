package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"spd-lookup/api/data"
)

// PortOfSeattleOfficerMetadata is the handler function for retrieving PortOfSeattle metadata
func (h *Handler) PortOfSeattleOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.PortOfSeattleOfficerMetadata())
	if err != nil {
		return
	}
}

// PortOfSeattleStrictMatch is the handler function for retrieving PortOfSeattle officers with a strict match
func (h *Handler) PortOfSeattleStrictMatch(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		h.portOfSeattleGetOfficersByName(name, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("name must be provided"))
		if err != nil {
			return
		}
	}
}

func (h *Handler) portOfSeattleGetOfficersByName(name string, w http.ResponseWriter) {
	if name == "" {
		name = "%"
	} else {
		name = strings.ReplaceAll(name, "*", "%")
	}

	officers, err := h.db.PortOfSeattleSearchOfficerByName(name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if errWrite != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&officers)
	if err != nil {
		return
	}
}

// PortOfSeattleFuzzySearch is the handler function for retrieving APD officers through fuzzy search
func (h *Handler) PortOfSeattleFuzzySearch(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	officers := []*data.PortOfSeattleOfficer{}
	var err error

	if name != "" {
		officers, err = h.db.PortOfSeattleFuzzySearchByName(name)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, writerErr := w.Write([]byte("provided name must not be empty"))
		if writerErr != nil {
			return
		}
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if writeErr != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&officers)
	if err != nil {
		return
	}
}
