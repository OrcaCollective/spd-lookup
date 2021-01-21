package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	h := &handler{
		db: newDBClient(),
	}

	router := mux.NewRouter()
	router.HandleFunc("/ping", h.ping).Methods("GET")
	router.HandleFunc("/officer", h.getOfficers).Methods("GET")
	router.HandleFunc("/officer/search", h.fuzzySearch).Methods("GET")

	port := os.Getenv("PORT")
	fmt.Println("starting server on port", port)
	http.ListenAndServe(":"+port, router)
}

type handler struct {
	db databaseInterface
}

func (h *handler) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("🏓 P O N G 🏓"))
}

func (h *handler) getOfficers(w http.ResponseWriter, r *http.Request) {
	badge, firstName, lastName := r.URL.Query().Get("badge"), r.URL.Query().Get("first_name"), r.URL.Query().Get("last_name")

	if badge != "" {
		h.getOfficerByBadge(badge, w)
		return
	} else if firstName != "" || lastName != "" {
		h.getOfficersByName(firstName, lastName, w)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("at least one of the following parameters must be provided: badge, first_name, last_name")))
	}
}

func (h *handler) getOfficerByBadge(badge string, w http.ResponseWriter) {
	ofc, err := h.db.getOfficerByBadge(badge)

	if err != nil {
		if err.Error() == "no rows in result set" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode([]*officer{})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]*officer{ofc})
}

func (h *handler) getOfficersByName(firstName, lastName string, w http.ResponseWriter) {
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

	officers, err := h.db.searchOfficerByName(firstName, lastName)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		return
	}

	alphabetize(officers)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&officers)
}

func (h *handler) fuzzySearch(w http.ResponseWriter, r *http.Request) {
	firstName, lastName := strings.TrimSpace(r.URL.Query().Get("first_name")), strings.TrimSpace(r.URL.Query().Get("last_name"))

	officers := []*officer{}
	var err error

	if firstName == "" && lastName == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("at least one of the following parameters must be provided: first_name, last_name")))
		return
	} else if firstName == "" {
		officers, err = h.db.fuzzySearchByLastName(lastName)
	} else if lastName == "" {
		officers, err = h.db.fuzzySearchByFirstName(firstName)
	} else {
		officers, err = h.db.fuzzySearchByName(firstName + " " + lastName)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error getting officer: %s", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&officers)
}

func alphabetize(officers []*officer) {
	sort.Slice(officers, func(a, b int) bool {
		if officers[a].LastName == officers[b].LastName {
			return officers[a].FirstName < officers[b].FirstName
		}
		return officers[a].LastName < officers[b].LastName
	})
}
