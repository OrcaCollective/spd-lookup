package handler

import (
	"fmt"
	"net/http"
	"spd-lookup/api/data"

	"github.com/gorilla/mux"
)

func NewRouter(h Interface) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/ping", h.Ping).Methods("GET")
	router.HandleFunc("/departments", h.DescribeDepartments).Methods("GET")

	router.HandleFunc("/seattle/metadata", h.SeattleOfficerMetadata).Methods("GET")
	router.HandleFunc("/seattle/officer", h.SeattleStrictMatch).Methods("GET")
	router.HandleFunc("/seattle/officer/search", h.SeattleFuzzySearch).Methods("GET")

	router.HandleFunc("/tacoma/metadata", h.TacomaOfficerMetadata).Methods("GET")
	router.HandleFunc("/tacoma/officer", h.TacomaStrictMatch).Methods("GET")
	router.HandleFunc("/tacoma/officer/search", h.TacomaFuzzySearch).Methods("GET")
	return router
}

type MockDatabase struct {
	data.DatabaseInterface
}

var mayday = "1889-05-01"

func (m *MockDatabase) SeattleOfficerMetadata() *data.DepartmentMetadata {
	return &data.DepartmentMetadata{
		Fields:                  []map[string]string{{"FieldName": "test", "Label": "Test"}},
		LastAvailableRosterDate: "today",
		Name:                    "Seattle PD",
		ID:                      "spd",
		SearchRoutes: map[string]*data.SearchRouteMetadata{
			"exact": {
				Path:        "/seattle/officer",
				QueryParams: []string{"badge", "first_name", "last_name"},
			},
		},
	}
}

var testSeattleOfficer1 = &data.SeattleOfficer{Date: mayday, Badge: "1", FirstName: "first", LastName: "sea"}
var testSeattleOfficer2 = &data.SeattleOfficer{Date: mayday, Badge: "2", FirstName: "first", LastName: "poo"}
var testSeattleOfficer3 = &data.SeattleOfficer{Date: mayday, Badge: "3", FirstName: "test", LastName: "poo"}

func (m *MockDatabase) SeattleGetOfficerByBadge(badge string) ([]*data.SeattleOfficer, error) {
	if badge == "db_error" {
		return nil, fmt.Errorf("get officer by badge db error")
	} else if badge == "badge_not_found" {
		return nil, fmt.Errorf("no rows in result set")
	}
	return []*data.SeattleOfficer{testSeattleOfficer1, testSeattleOfficer2, testSeattleOfficer3}, nil
}

func (m *MockDatabase) SeattleSearchOfficerByName(firstName, lastName string) ([]*data.SeattleOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("get officer by name db error")
	}
	return []*data.SeattleOfficer{testSeattleOfficer1, testSeattleOfficer2, testSeattleOfficer3}, nil
}

func (m *MockDatabase) SeattleFuzzySearchByName(name string) ([]*data.SeattleOfficer, error) {
	if name == "db error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.SeattleOfficer{testSeattleOfficer1}, nil
}

func (m *MockDatabase) SeattleFuzzySearchByFirstName(firstName string) ([]*data.SeattleOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by first name db error")
	}
	return []*data.SeattleOfficer{testSeattleOfficer1}, nil
}

func (m *MockDatabase) SeattleFuzzySearchByLastName(lastName string) ([]*data.SeattleOfficer, error) {
	if lastName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by last name db error")
	}
	return []*data.SeattleOfficer{testSeattleOfficer1}, nil
}

func (m *MockDatabase) TacomaOfficerMetadata() *data.DepartmentMetadata {
	return &data.DepartmentMetadata{
		Fields:                  []map[string]string{{"FieldName": "test", "Label": "Test"}},
		LastAvailableRosterDate: "yesterday",
		Name:                    "Tacoma PD",
		ID:                      "tpd",
		SearchRoutes: map[string]*data.SearchRouteMetadata{
			"exact": {
				Path:        "/tacoma/officer",
				QueryParams: []string{"last_name"},
			},
		},
	}
}

var testTacomaOfficer1 = &data.TacomaOfficer{Date: mayday, FirstName: "first", LastName: "tac"}
var testTacomaOfficer2 = &data.TacomaOfficer{Date: mayday, FirstName: "first", LastName: "poo"}
var testTacomaOfficer3 = &data.TacomaOfficer{Date: mayday, FirstName: "test", LastName: "poo"}

func (m *MockDatabase) TacomaSearchOfficerByName(firstName, lastName string) ([]*data.TacomaOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("get officer by name db error")
	}
	return []*data.TacomaOfficer{testTacomaOfficer1, testTacomaOfficer2, testTacomaOfficer3}, nil
}

func (m *MockDatabase) TacomaFuzzySearchByName(name string) ([]*data.TacomaOfficer, error) {
	if name == "db error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.TacomaOfficer{testTacomaOfficer1}, nil
}

func (m *MockDatabase) TacomaFuzzySearchByFirstName(firstName string) ([]*data.TacomaOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by first name db error")
	}
	return []*data.TacomaOfficer{testTacomaOfficer1}, nil
}

func (m *MockDatabase) TacomaFuzzySearchByLastName(lastName string) ([]*data.TacomaOfficer, error) {
	if lastName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by last name db error")
	}
	return []*data.TacomaOfficer{testTacomaOfficer1}, nil
}
