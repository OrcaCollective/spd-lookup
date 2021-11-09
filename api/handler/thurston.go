package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"spd-lookup/api/data"
)

// ThurstonCountyOfficerMetadata is the handler function for retrieving ThurstonCounty metadata
func (h *Handler) ThurstonCountyOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.ThurstonCountyOfficerMetadata())
	if err != nil {
		return
	}
}

// ThurstonCountyStrictMatch is the handler function for retrieving ThurstonCounty officers with a strict match
func (h *Handler) ThurstonCountyStrictMatch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := r.URL.Query().Get("first_name"), r.URL.Query().Get("last_name")

	if firstName != "" || lastName != "" {
		h.thurstonCountyGetOfficersByName(firstName, lastName, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("at least one of the following parameters must be provided: first_name, last_name"))
		if err != nil {
			return
		}
	}
}

func (h *Handler) thurstonCountyGetOfficersByName(firstName, lastName string, w http.ResponseWriter) {
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

	officers, err := h.db.ThurstonCountySearchOfficerByName(firstName, lastName)

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

// ThurstonCountyFuzzySearch is the handler function for retrieving APD officers through fuzzy search
func (h *Handler) ThurstonCountyFuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*data.ThurstonCountyOfficer{}
	var err error

	if firstName != "" && lastName != "" {
		officers, err = h.db.ThurstonCountyFuzzySearchByName(strings.Trim(firstName+" "+lastName, " "))
	} else if firstName != "" {
		officers, err = h.db.ThurstonCountyFuzzySearchByFirstName(firstName)
	} else if lastName != "" {
		officers, err = h.db.ThurstonCountyFuzzySearchByLastName(lastName)
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
