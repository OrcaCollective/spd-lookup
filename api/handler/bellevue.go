package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"spd-lookup/api/data"
)

// BellevueOfficerMetadata is the handler function for retrieving Bellevue metadata
func (h *Handler) BellevueOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.BellevueOfficerMetadata())
	if err != nil {
		return
	}
}

// BellevueStrictMatch is the handler function for retrieving Bellevue officers with a strict match
func (h *Handler) BellevueStrictMatch(w http.ResponseWriter, r *http.Request) {
	badge, firstName, lastName := r.URL.Query().Get("badge"), r.URL.Query().Get("first_name"), r.URL.Query().Get("last_name")

	if badge != "" {
		h.bellevueGetOfficersByBadge(badge, w)
	} else if firstName != "" || lastName != "" {
		h.bellevueGetOfficersByName(firstName, lastName, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("at least one of the following parameters must be provided: badge, first_name, last_name"))
		if err != nil {
			return
		}
	}
}

func (h *Handler) bellevueGetOfficersByBadge(badge string, w http.ResponseWriter) {
	officers, err := h.db.BellevueSearchOfficerByBadge(badge)

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

func (h *Handler) bellevueGetOfficersByName(firstName, lastName string, w http.ResponseWriter) {
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

	officers, err := h.db.BellevueSearchOfficerByName(firstName, lastName)

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

// BellevueFuzzySearch is the handler function for retrieving APD officers through fuzzy search
func (h *Handler) BellevueFuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*data.BellevueOfficer{}
	var err error

	if firstName != "" && lastName != "" {
		officers, err = h.db.BellevueFuzzySearchByName(strings.Trim(firstName+" "+lastName, " "))
	} else if firstName != "" {
		officers, err = h.db.BellevueFuzzySearchByFirstName(firstName)
	} else if lastName != "" {
		officers, err = h.db.BellevueFuzzySearchByLastName(lastName)
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
