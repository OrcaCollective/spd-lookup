package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/OrcaCollective/spd-lookup/api/data"
)

// OlympiaOfficerMetadata is the handler function for retrieving SPD metadata
func (h *Handler) OlympiaOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.OlympiaOfficerMetadata())
	if err != nil {
		return
	}
}

// OlympiaStrictMatch is the handler function for retrieving SPD officers with a strict match
func (h *Handler) OlympiaStrictMatch(w http.ResponseWriter, r *http.Request) {
	badge, firstName, lastName := r.URL.Query().Get("badge"), r.URL.Query().Get("first_name"), r.URL.Query().Get("last_name")

	if badge != "" {
		h.olympiaGetOfficerByBadge(badge, w)
		return
	} else if firstName != "" || lastName != "" {
		h.olympiaGetOfficersByName(firstName, lastName, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("at least one of the following parameters must be provided: badge, first_name, last_name"))
		if err != nil {
			return
		}
	}
}

func (h *Handler) olympiaGetOfficerByBadge(badge string, w http.ResponseWriter) {
	officers, err := h.db.OlympiaGetOfficerByBadge(badge)

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

func (h *Handler) olympiaGetOfficersByName(firstName, lastName string, w http.ResponseWriter) {
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

	officers, err := h.db.OlympiaSearchOfficerByName(firstName, lastName)

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

// OlympiaFuzzySearch is the handler function for retrieving APD officers through fuzzy search
func (h *Handler) OlympiaFuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*data.OlympiaOfficer{}
	var err error

	if firstName != "" && lastName != "" {
		officers, err = h.db.OlympiaFuzzySearchByName(strings.Trim(firstName+" "+lastName, " "))
	} else if firstName != "" {
		officers, err = h.db.OlympiaFuzzySearchByFirstName(firstName)
	} else if lastName != "" {
		officers, err = h.db.OlympiaFuzzySearchByLastName(lastName)
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
