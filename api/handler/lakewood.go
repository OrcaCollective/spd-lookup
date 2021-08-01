package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"spd-lookup/api/data"
)

// LakewoodOfficerMetadata is the handler function for retrieving SPD metadata
func (h *Handler) LakewoodOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.LakewoodOfficerMetadata())
	if err != nil {
		return
	}
}

// LakewoodStrictMatch is the handler function for retrieving SPD officers with a strict match
func (h *Handler) LakewoodStrictMatch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := r.URL.Query().Get("first_name"), r.URL.Query().Get("last_name")

	if firstName != "" || lastName != "" {
		h.lakewoodGetOfficersByName(firstName, lastName, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("at least one of the following parameters must be provided: first_name, last_name"))
		if err != nil {
			return
		}
	}
}

func (h *Handler) lakewoodGetOfficersByName(firstName, lastName string, w http.ResponseWriter) {
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

	officers, err := h.db.LakewoodSearchOfficerByName(firstName, lastName)

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

// LakewoodFuzzySearch is the handler function for retrieving APD officers through fuzzy search
func (h *Handler) LakewoodFuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*data.LakewoodOfficer{}
	var err error

	if firstName != "" && lastName != "" {
		officers, err = h.db.LakewoodFuzzySearchByName(strings.Trim(firstName+" "+lastName, " "))
	} else if firstName != "" {
		officers, err = h.db.LakewoodFuzzySearchByFirstName(firstName)
	} else if lastName != "" {
		officers, err = h.db.LakewoodFuzzySearchByLastName(lastName)
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
