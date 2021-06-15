package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"spd-lookup/api/data"
)

// SeattleOfficerMetadata is the handler function for retrieving SPD metadata
func (h *Handler) SeattleOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.SeattleOfficerMetadata())
	if err != nil {
		return
	}
}

// SeattleStrictMatch is the handler function for retrieving SPD officers with a strict match
func (h *Handler) SeattleStrictMatch(w http.ResponseWriter, r *http.Request) {
	badge, firstName, lastName := r.URL.Query().Get("badge"), r.URL.Query().Get("first_name"), r.URL.Query().Get("last_name")

	if badge != "" {
		h.seattleGetOfficerByBadge(badge, w)
		return
	} else if firstName != "" || lastName != "" {
		h.seattleGetOfficersByName(firstName, lastName, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("at least one of the following parameters must be provided: badge, first_name, last_name"))
		if err != nil {
			return
		}
	}
}

// SeattleStrictMatchHistorical is the handler function for retrieving SPD officers with a strict match
func (h *Handler) SeattleStrictMatchHistorical(w http.ResponseWriter, r *http.Request) {
	badge := r.URL.Query().Get("badge")

	if badge != "" {
		h.seattleGetOfficerByBadgeHistorical(badge, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("at least one of the following parameters must be provided: badge"))
		if err != nil {
			return
		}
	}
}

func (h *Handler) seattleGetOfficerByBadge(badge string, w http.ResponseWriter) {
	officers, err := h.db.SeattleGetOfficerByBadge(badge)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if errWrite != nil {
			return
		}
		return
	}

	sort.Slice(officers, func(a, b int) bool {
		if officers[a].LastName == officers[b].LastName {
			return officers[a].FirstName < officers[b].FirstName
		}
		return officers[a].LastName < officers[b].LastName
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&officers)
	if err != nil {
		return
	}
}

func (h *Handler) seattleGetOfficersByName(firstName, lastName string, w http.ResponseWriter) {
	if firstName == "" {
		firstName = "%"
	} else {
		firstName = strings.ReplaceAll(firstName, "*", "%")
	}

	if lastName == "" {
		lastName = "%"
	} else {
		lastName = strings.ReplaceAll(lastName, "*", "%")
	}

	officers, err := h.db.SeattleSearchOfficerByName(firstName, lastName)

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

func (h *Handler) seattleGetOfficerByBadgeHistorical(badge string, w http.ResponseWriter) {
	officers, err := h.db.SeattleGetOfficerByBadgeHistorical(badge)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if errWrite != nil {
			return
		}
		return
	}

	sort.Slice(officers, func(a, b int) bool {
		if officers[a].Date == officers[b].Date {
			return officers[a].Date > officers[b].Date
		}
		return officers[a].Date > officers[b].Date
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&officers)
	if err != nil {
		return
	}
}

// SeattleFuzzySearch is the handler function for retrieving SPD officers through fuzzy search
func (h *Handler) SeattleFuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*data.SeattleOfficer{}
	var err error

	if firstName != "" && lastName != "" {
		officers, err = h.db.SeattleFuzzySearchByName(strings.Trim(firstName+" "+lastName, " "))
	} else if firstName != "" {
		officers, err = h.db.SeattleFuzzySearchByFirstName(firstName)
	} else if lastName != "" {
		officers, err = h.db.SeattleFuzzySearchByLastName(lastName)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, writerErr := w.Write([]byte("at least one of the following parameters must be provided: first_name, last_name"))
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
