package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/OrcaCollective/spd-lookup/api/data"
)

// TacomaOfficerMetadata is the handler function for retrieving Tacoma metadata
func (h *Handler) TacomaOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.TacomaOfficerMetadata())
	if err != nil {
		return
	}
}

// TacomaStrictMatch is the handler function for retrieving Tacoma officers with a strict match
func (h *Handler) TacomaStrictMatch(w http.ResponseWriter, r *http.Request) {
	badge, firstName, lastName := r.URL.Query().Get("badge"), strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	if badge != "" && firstName == "" && lastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, writerErr := w.Write([]byte("At this time we do not have the badge numbers available for Tacoma PD. Please attempt searches by first or last name only."))
		if writerErr != nil {
			return
		}
		return
	}

	if firstName == "" && lastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, writerErr := w.Write([]byte("at least one of the following parameters must be provided: first_name, last_name"))
		if writerErr != nil {
			return
		}
		return
	}

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

	officers, err := h.db.TacomaSearchOfficerByName(firstName, lastName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writerErr := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if writerErr != nil {
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

// TacomaFuzzySearch is the handler function for retrieving Tacoma officers through fuzzy search
func (h *Handler) TacomaFuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*data.TacomaOfficer{}
	var err error

	if firstName != "" && lastName != "" {
		officers, err = h.db.TacomaFuzzySearchByName(strings.Trim(firstName+" "+lastName, " "))
	} else if firstName != "" {
		officers, err = h.db.TacomaFuzzySearchByFirstName(firstName)
	} else if lastName != "" {
		officers, err = h.db.TacomaFuzzySearchByLastName(lastName)
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
		_, writerErr := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if writerErr != nil {
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
