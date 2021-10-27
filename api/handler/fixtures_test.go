package handler

import (
	"fmt"
	"net/http"
	"spd-lookup/api/data"

	"github.com/gorilla/mux"
	"github.com/gobuffalo/nulls"
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

	router.HandleFunc("/portland/metadata", h.PortlandOfficerMetadata).Methods("GET")
	router.HandleFunc("/portland/officer", h.PortlandStrictMatch).Methods("GET")
	router.HandleFunc("/portland/officer/search", h.PortlandFuzzySearch).Methods("GET")

	router.HandleFunc("/auburn/metadata", h.AuburnOfficerMetadata).Methods("GET")
	router.HandleFunc("/auburn/officer", h.AuburnStrictMatch).Methods("GET")
	router.HandleFunc("/auburn/officer/search", h.AuburnFuzzySearch).Methods("GET")

	router.HandleFunc("/lakewood/metadata", h.LakewoodOfficerMetadata).Methods("GET")
	router.HandleFunc("/lakewood/officer", h.LakewoodStrictMatch).Methods("GET")
	router.HandleFunc("/lakewood/officer/search", h.LakewoodFuzzySearch).Methods("GET")

	router.HandleFunc("/olympia/metadata", h.OlympiaOfficerMetadata).Methods("GET")
	router.HandleFunc("/olympia/officer", h.OlympiaStrictMatch).Methods("GET")
	router.HandleFunc("/olympia/officer/search", h.OlympiaFuzzySearch).Methods("GET")
	return router
}

type MockDatabase struct {
	data.DatabaseInterface
}

var mayday = "1889-05-01"

func (m *MockDatabase) LakewoodOfficerMetadata() *data.DepartmentMetadata {
    return &data.DepartmentMetadata{
		Fields:                  []map[string]string{{"FieldName": "test", "Label": "Test"}},
		LastAvailableRosterDate: "yesterday",
		Name:                    "Lakewood PD",
		ID:                      "lpd",
		SearchRoutes: map[string]*data.SearchRouteMetadata{
			"exact": {
				Path:        "/lakewood/officer",
				QueryParams: []string{"first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/lakewood/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

var testLakewoodOfficer1 = &data.LakewoodOfficer{Date: mayday, FirstName: "first", LastName: "lak"}
var testLakewoodOfficer2 = &data.LakewoodOfficer{Date: mayday, FirstName: "first", LastName: "poo"}
var testLakewoodOfficer3 = &data.LakewoodOfficer{Date: mayday, FirstName: "test", LastName: "poo"}

func (m *MockDatabase) LakewoodSearchOfficerByName(firstName, lastName string) ([]*data.LakewoodOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("get officer by name db error")
	}
	return []*data.LakewoodOfficer{testLakewoodOfficer1, testLakewoodOfficer2, testLakewoodOfficer3}, nil
}

func (m *MockDatabase) LakewoodFuzzySearchByName(name string) ([]*data.LakewoodOfficer, error) {
	if name == "db_error db_error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.LakewoodOfficer{testLakewoodOfficer1}, nil
}

func (m *MockDatabase) LakewoodFuzzySearchByFirstName(name string) ([]*data.LakewoodOfficer, error) {
	if name == "db_error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.LakewoodOfficer{testLakewoodOfficer1}, nil
}

func (m *MockDatabase) LakewoodFuzzySearchByLastName(name string) ([]*data.LakewoodOfficer, error) {
	if name == "db_error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.LakewoodOfficer{testLakewoodOfficer1}, nil
}

func (m *MockDatabase) OlympiaOfficerMetadata() *data.DepartmentMetadata {
    return &data.DepartmentMetadata{
		Fields:                  []map[string]string{{"FieldName": "test", "Label": "Test"}},
		LastAvailableRosterDate: "yesterday",
		Name:                    "Olympia PD",
		ID:                      "opd",
		SearchRoutes: map[string]*data.SearchRouteMetadata{
			"exact": {
				Path:        "/olympia/officer",
				QueryParams: []string{"badge", "first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/olympia/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

var testOlympiaOfficer1 = &data.OlympiaOfficer{Date: mayday, Badge: "1", FirstName: "first", LastName: "oly"}
var testOlympiaOfficer2 = &data.OlympiaOfficer{Date: mayday, Badge: "2", FirstName: "first", LastName: "poo"}
var testOlympiaOfficer3 = &data.OlympiaOfficer{Date: mayday, Badge: "3", FirstName: "test", LastName: "poo"}

func (m *MockDatabase) OlympiaGetOfficerByBadge(badge string) ([]*data.OlympiaOfficer, error) {
	if badge == "db_error" {
		return nil, fmt.Errorf("get officer by badge db error")
	} else if badge == "badge_not_found" {
		return []*data.OlympiaOfficer{}, nil
	} else if badge == "1" {
	    return []*data.OlympiaOfficer{testOlympiaOfficer1}, nil
    }
	return []*data.OlympiaOfficer{testOlympiaOfficer1, testOlympiaOfficer2, testOlympiaOfficer3}, nil
}

func (m *MockDatabase) OlympiaSearchOfficerByName(firstName, lastName string) ([]*data.OlympiaOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("get officer by name db error")
	}
	return []*data.OlympiaOfficer{testOlympiaOfficer1, testOlympiaOfficer2, testOlympiaOfficer3}, nil
}

func (m *MockDatabase) OlympiaFuzzySearchByName(name string) ([]*data.OlympiaOfficer, error) {
	if name == "db error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.OlympiaOfficer{testOlympiaOfficer1}, nil
}

func (m *MockDatabase) OlympiaFuzzySearchByFirstName(firstName string) ([]*data.OlympiaOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by first name db error")
	}
	return []*data.OlympiaOfficer{testOlympiaOfficer1}, nil
}

func (m *MockDatabase) OlympiaFuzzySearchByLastName(lastName string) ([]*data.OlympiaOfficer, error) {
	if lastName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by last name db error")
	}
	return []*data.OlympiaOfficer{testOlympiaOfficer1}, nil
}

func (m *MockDatabase) AuburnOfficerMetadata() *data.DepartmentMetadata {
    return &data.DepartmentMetadata{
		Fields:                  []map[string]string{{"FieldName": "test", "Label": "Test"}},
		LastAvailableRosterDate: "yesterday",
		Name:                    "Auburn PD",
		ID:                      "apd",
		SearchRoutes: map[string]*data.SearchRouteMetadata{
			"exact": {
				Path:        "/auburn/officer",
				QueryParams: []string{"badge", "first_name", "last_name"},
			},
			"fuzzy": {
				Path:        "/auburn/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

var testAuburnOfficer1 = &data.AuburnOfficer{Date: mayday, Badge: "1", FirstName: "first", LastName: "aub"}
var testAuburnOfficer2 = &data.AuburnOfficer{Date: mayday, Badge: "2", FirstName: "first", LastName: "poo"}
var testAuburnOfficer3 = &data.AuburnOfficer{Date: mayday, Badge: "3", FirstName: "test", LastName: "poo"}

func (m *MockDatabase) AuburnGetOfficerByBadge(badge string) ([]*data.AuburnOfficer, error) {
	if badge == "db_error" {
		return nil, fmt.Errorf("get officer by badge db error")
	} else if badge == "badge_not_found" {
		return []*data.AuburnOfficer{}, nil
	} else if badge == "1" {
	    return []*data.AuburnOfficer{testAuburnOfficer1}, nil
    }
	return []*data.AuburnOfficer{testAuburnOfficer1, testAuburnOfficer2, testAuburnOfficer3}, nil
}

func (m *MockDatabase) AuburnSearchOfficerByName(firstName, lastName string) ([]*data.AuburnOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("get officer by name db error")
	}
	return []*data.AuburnOfficer{testAuburnOfficer1, testAuburnOfficer2, testAuburnOfficer3}, nil
}

func (m *MockDatabase) AuburnFuzzySearchByName(name string) ([]*data.AuburnOfficer, error) {
	if name == "db error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.AuburnOfficer{testAuburnOfficer1}, nil
}

func (m *MockDatabase) AuburnFuzzySearchByFirstName(firstName string) ([]*data.AuburnOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by first name db error")
	}
	return []*data.AuburnOfficer{testAuburnOfficer1}, nil
}

func (m *MockDatabase) AuburnFuzzySearchByLastName(lastName string) ([]*data.AuburnOfficer, error) {
	if lastName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by last name db error")
	}
	return []*data.AuburnOfficer{testAuburnOfficer1}, nil
}

func (m *MockDatabase) PortlandOfficerMetadata() *data.DepartmentMetadata {
    return &data.DepartmentMetadata{
		Fields:                  []map[string]string{{"FieldName": "test", "Label": "Test"}},
		LastAvailableRosterDate: "today",
		Name:                    "Portland PB",
		ID:                      "ppb",
		SearchRoutes: map[string]*data.SearchRouteMetadata{
			"exact": {
				Path:        "/portland/officer",
				QueryParams: []string{"badge", "first_name", "last_name", "employee_id", "helmet_id", "helmet_id_three_digit"},
			},
			"fuzzy": {
				Path:        "/portland/officer/search",
				QueryParams: []string{"first_name", "last_name"},
			},
		},
	}
}

var testPortlandOfficer1 = &data.PortlandOfficer{Badge: nulls.NewString("1"), EmployeeID: nulls.NewString("1"), HelmetID: nulls.NewString("1"), HelmetIDThreeDigit: nulls.NewString("111"), FirstName: nulls.NewString("first"), LastName: nulls.NewString("ppb")}
var testPortlandOfficer2 = &data.PortlandOfficer{Badge: nulls.NewString("2"), EmployeeID: nulls.NewString("2"), HelmetID: nulls.NewString("1"), HelmetIDThreeDigit: nulls.NewString("222"), FirstName: nulls.NewString("first"), LastName: nulls.NewString("poo")}
var testPortlandOfficer3 = &data.PortlandOfficer{Badge: nulls.NewString("3"), EmployeeID: nulls.NewString("3"), HelmetID: nulls.NewString("1"), HelmetIDThreeDigit: nulls.NewString("333"), FirstName: nulls.NewString("test"), LastName: nulls.NewString("poo")}

func (m *MockDatabase) PortlandSearchOfficersByBadge(badge string) ([]*data.PortlandOfficer, error) {
	if badge == "db_error" {
		return nil, fmt.Errorf("get officer by badge db error")
	} else if badge == "badge_not_found" {
		return []*data.PortlandOfficer{}, nil
	} else if badge == "1" {
	    return []*data.PortlandOfficer{testPortlandOfficer1}, nil
    }
	return []*data.PortlandOfficer{testPortlandOfficer1, testPortlandOfficer2, testPortlandOfficer3}, nil
}

func (m *MockDatabase) PortlandSearchOfficersByEmployeeId(employee_id string) ([]*data.PortlandOfficer, error) {
	if employee_id == "db_error" {
		return nil, fmt.Errorf("get officer by employee id db error")
	} else if employee_id == "employee_id_not_found" {
		return []*data.PortlandOfficer{}, nil
	} else if employee_id == "1" {
	    return []*data.PortlandOfficer{testPortlandOfficer1}, nil
    }
	return []*data.PortlandOfficer{testPortlandOfficer1, testPortlandOfficer2, testPortlandOfficer3}, nil
}

func (m *MockDatabase) PortlandSearchOfficersByHelmetId(helmet_id string) ([]*data.PortlandOfficer, error) {
	if helmet_id == "db_error" {
		return nil, fmt.Errorf("get officer by helmet id db error")
	} else if helmet_id == "helmet_id_not_found" {
		return []*data.PortlandOfficer{}, nil
	} else if helmet_id == "1" {
	    return []*data.PortlandOfficer{testPortlandOfficer1}, nil
    }
	return []*data.PortlandOfficer{testPortlandOfficer1, testPortlandOfficer2, testPortlandOfficer3}, nil
}

func (m *MockDatabase) PortlandSearchOfficersByHelmetIdThreeDigit(helmet_id_three_digit string) ([]*data.PortlandOfficer, error) {
	if helmet_id_three_digit == "db_error" {
		return nil, fmt.Errorf("get officer by helmet id three digits db error")
	} else if helmet_id_three_digit == "helmet_id_three_digit_not_found" {
		return []*data.PortlandOfficer{}, nil
	} else if helmet_id_three_digit == "111" {
	    return []*data.PortlandOfficer{testPortlandOfficer1}, nil
    }
	return []*data.PortlandOfficer{testPortlandOfficer1, testPortlandOfficer2, testPortlandOfficer3}, nil
}

func (m *MockDatabase) PortlandSearchOfficersByName(firstName, lastName string) ([]*data.PortlandOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("get officer by name db error")
	}
	return []*data.PortlandOfficer{testPortlandOfficer1, testPortlandOfficer2, testPortlandOfficer3}, nil
}

func (m *MockDatabase) PortlandFuzzySearchByName(name string) ([]*data.PortlandOfficer, error) {
	if name == "db error" {
		return nil, fmt.Errorf("fuzzy search by name db error")
	}
	return []*data.PortlandOfficer{testPortlandOfficer1}, nil
}

func (m *MockDatabase) PortlandFuzzySearchByFirstName(firstName string) ([]*data.PortlandOfficer, error) {
	if firstName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by first name db error")
	}
	return []*data.PortlandOfficer{testPortlandOfficer1}, nil
}

func (m *MockDatabase) PortlandFuzzySearchByLastName(lastName string) ([]*data.PortlandOfficer, error) {
	if lastName == "db_error" {
		return nil, fmt.Errorf("fuzzy search by last name db error")
	}
	return []*data.PortlandOfficer{testPortlandOfficer1}, nil
}

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
		return []*data.SeattleOfficer{}, nil
	} else if badge == "1" {
	    return []*data.SeattleOfficer{testSeattleOfficer1}, nil
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
