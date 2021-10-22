package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_TacomaOfficerMetadata(t *testing.T) {
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/tacoma/metadata")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d", res.StatusCode)
	}
	defer res.Body.Close()
	resp, _ := ioutil.ReadAll(res.Body)

	expected := []byte(`{"id":"tpd","name":"Tacoma PD","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/tacoma/officer","query_params":["last_name"]}}}` + "\n")

	if !bytes.Equal(expected, resp) {
		t.Errorf("Expected resp: %s, got %s", expected, resp)
	}
}

func Test_Handler_TacomaStrictMatch(t *testing.T) {
	t.Parallel()
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	for _, tt := range [...]struct {
		name           string
		firstName      string
		lastName       string
		badge          string
		expectedStatus int
		expectedBody   []byte
	}{
		{
			name:           "no params",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte("at least one of the following parameters must be provided: first_name, last_name"),
		},
		{
			name:           "badge provided",
			badge:          "test",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte("At this time we do not have the badge numbers available for Tacoma PD. Please attempt searches by first or last name only."),
		},
		{
			name:           "name search, db error",
			firstName:      "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: get officer by name db error"),
		},
		{
			name:           "name search, officers found",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"date":"1889-05-01","first_name":"first","last_name":"poo","salary":null},{"date":"1889-05-01","first_name":"test","last_name":"poo","salary":null},{"date":"1889-05-01","first_name":"first","last_name":"tac","salary":null}]` + "\n"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/tacoma/officer?badge=%s&first_name=%s&last_name=%s", ts.URL, tt.badge, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)
			if !bytes.Equal(tt.expectedBody, resp) {
				t.Errorf("Expected resp %s; got %s", tt.expectedBody, resp)
			}
		})
	}
}

func Test_Handler_TacomaFuzzySearch(t *testing.T) {
	t.Parallel()
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	for _, tt := range [...]struct {
		name           string
		firstName      string
		lastName       string
		expectedStatus int
		expectedBody   []byte
	}{
		{
			name:           "no params",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte("at least one of the following parameters must be provided: first_name, last_name"),
		},
		{
			name:           "name fuzzy search, db error",
			firstName:      "db",
			lastName:       "error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: fuzzy search by name db error"),
		},
		{
			name:           "name fuzzy search, officers found",
			firstName:      "test",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"date":"1889-05-01","first_name":"first","last_name":"tac","salary":null}]` + "\n"),
		},
		{
			name:           "first name fuzzy search, db error",
			firstName:      "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: fuzzy search by first name db error"),
		},
		{
			name:           "first name fuzzy search, officers found",
			firstName:      "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"date":"1889-05-01","first_name":"first","last_name":"tac","salary":null}]` + "\n"),
		},
		{
			name:           "last name fuzzy search, db error",
			lastName:       "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: fuzzy search by last name db error"),
		},
		{
			name:           "last name fuzzy search, officers found",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"date":"1889-05-01","first_name":"first","last_name":"tac","salary":null}]` + "\n"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/tacoma/officer/search?first_name=%s&last_name=%s", ts.URL, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)
			if !bytes.Equal(tt.expectedBody, resp) {
				t.Errorf("Expected resp %s; got %s", tt.expectedBody, resp)
			}
		})
	}
}
