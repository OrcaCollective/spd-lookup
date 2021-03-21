package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"spd-lookup/api/data"
)

// PortlandOfficerMetadata is the handler function for retrieving SPD metadata
func (h *Handler) PortlandOfficerMetadata(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(h.db.PortlandOfficerMetadata())
	if err != nil {
		return
	}
}

// PortlandStrictMatch is the handler function for retrieving SPD officers with a strict match
func (h *Handler) PortlandStrictMatch(w http.ResponseWriter, r *http.Request) {
    badge := r.URL.Query().Get("badge")
    first_name := r.URL.Query().Get("first_name")
    last_name := r.URL.Query().Get("last_name")
    employee_id := r.URL.Query().Get("employee_id")
    helmet_id := r.URL.Query().Get("helmet_id")
    helmet_id_three_digit := r.URL.Query().Get("helmet_id_three_digit")

	if badge != "" {
		h.portlandGetOfficersByBadge(badge, w)
		return
    } else if employee_id != "" {
		h.portlandGetOfficersByEmployeeId(employee_id, w)
		return
    } else if helmet_id != "" {
		h.portlandGetOfficersByHelmetId(helmet_id, w)
		return
    } else if helmet_id_three_digit != "" {
		h.portlandGetOfficersByHelmetIdThreeDigit(helmet_id_three_digit, w)
		return
	} else if firstName != "" || lastName != "" {
		h.portlandGetOfficersByName(firstName, lastName, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("at least one of the following parameters must be provided: badge, first_name, last_name"))
		if err != nil {
			return
		}
	}
}

func (h *Handler) portlandGetOfficersByBadge(badge string, w http.ResponseWriter) {
	officers, err := h.db.PortlandSearchOfficersByBadge(badge)

	if err != nil {
		if err.Error() == "no rows in result set" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode([]*data.PortlandOfficer{})
			if err != nil {
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if err != nil {
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

func (h *Handler) portlandGetOfficersByEmployeeId(employee_id string, w http.ResponseWriter) {
	officers, err := h.db.PortlandSearchOfficersByEmployeeId(employee_id)

	if err != nil {
		if err.Error() == "no rows in result set" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode([]*data.PortlandOfficer{})
			if err != nil {
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if err != nil {
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

func (h *Handler) portlandGetOfficersByHelmetId(helmet_id string, w http.ResponseWriter) {
	officers, err := h.db.PortlandSearchOfficersByHelmetId(helmet_id)

	if err != nil {
		if err.Error() == "no rows in result set" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode([]*data.PortlandOfficer{})
			if err != nil {
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if err != nil {
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

func (h *Handler) portlandGetOfficersByHelmetIdThreeDigit(helmet_id_three_digit string, w http.ResponseWriter) {
	officers, err := h.db.PortlandSearchOfficersByHelmetIdThreeDigit(helmet_id_three_digit)

	if err != nil {
		if err.Error() == "no rows in result set" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode([]*data.PortlandOfficer{})
			if err != nil {
				return
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		if err != nil {
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

func (h *Handler) portlandGetOfficersByName(firstName, lastName string, w http.ResponseWriter) {
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

	officers, err := h.db.PortlandSearchOfficersByName(firstName, lastName)

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

// PortlandFuzzySearch is the handler function for retrieving SPD officers through fuzzy search
func (h *Handler) PortlandFuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*data.PortlandOfficer{}
	var err error

	if firstName != "" && lastName != "" {
		officers, err = h.db.PortlandFuzzySearchByName(strings.Trim(firstName+" "+lastName, " "))
	} else if firstName != "" {
		officers, err = h.db.PortlandFuzzySearchByFirstName(firstName)
	} else if lastName != "" {
		officers, err = h.db.PortlandFuzzySearchByLastName(lastName)
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
