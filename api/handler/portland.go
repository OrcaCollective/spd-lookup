package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/OrcaCollective/spd-lookup/api/data"
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
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	employeeId := r.URL.Query().Get("employee_id")
	helmetId := r.URL.Query().Get("helmet_id")
	helmetIdThreeDigit := r.URL.Query().Get("helmet_id_three_digit")

	if badge != "" {
		h.portlandGetOfficersByBadge(badge, w)
		return
	} else if employeeId != "" {
		h.portlandGetOfficersByEmployeeId(employeeId, w)
		return
	} else if helmetId != "" {
		h.portlandGetOfficersByHelmetId(helmetId, w)
		return
	} else if helmetIdThreeDigit != "" {
		h.portlandGetOfficersByHelmetIdThreeDigit(helmetIdThreeDigit, w)
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
			return officers[a].FirstName.String < officers[b].FirstName.String
		}
		return officers[a].LastName.String < officers[b].LastName.String
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&officers)
	if err != nil {
		return
	}
}

func (h *Handler) portlandGetOfficersByEmployeeId(employeeId string, w http.ResponseWriter) {
	officers, err := h.db.PortlandSearchOfficersByEmployeeId(employeeId)

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
			return officers[a].FirstName.String < officers[b].FirstName.String
		}
		return officers[a].LastName.String < officers[b].LastName.String
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&officers)
	if err != nil {
		return
	}
}

func (h *Handler) portlandGetOfficersByHelmetId(helmetId string, w http.ResponseWriter) {
	officers, err := h.db.PortlandSearchOfficersByHelmetId(helmetId)

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
			return officers[a].FirstName.String < officers[b].FirstName.String
		}
		return officers[a].LastName.String < officers[b].LastName.String
	})

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&officers)
	if err != nil {
		return
	}
}

func (h *Handler) portlandGetOfficersByHelmetIdThreeDigit(helmetIdThreeDigit string, w http.ResponseWriter) {
	officers, err := h.db.PortlandSearchOfficersByHelmetIdThreeDigit(helmetIdThreeDigit)

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
			return officers[a].FirstName.String < officers[b].FirstName.String
		}
		return officers[a].LastName.String < officers[b].LastName.String
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
			return officers[a].FirstName.String < officers[b].FirstName.String
		}
		return officers[a].LastName.String < officers[b].LastName.String
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
