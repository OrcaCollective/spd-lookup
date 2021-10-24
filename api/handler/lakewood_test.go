package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Handler_LakewoodOfficerMetadata(t *testing.T) {
	router := NewRouter(NewHandler(&MockDatabase{}))
	ts := httptest.NewServer(router)
	defer ts.Close()

	res, _ := http.Get(ts.URL + "/lakewood/metadata")
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status %d", res.StatusCode)
	}
	defer res.Body.Close()
	resp, _ := ioutil.ReadAll(res.Body)

	expected := []byte(`{"id":"lpd","name":"Lakewood PD","last_available_roster_date":"yesterday","fields":[{"FieldName":"test","Label":"Test"}],"search_routes":{"exact":{"path":"/lakewood/officer","query_params":["first_name","last_name"]},"fuzzy":{"path":"/lakewood/officer/search","query_params":["first_name","last_name"]}}}` + "\n")

	if !bytes.Equal(expected, resp) {
		t.Errorf("Expected resp: %s, got %s", expected, resp)
	}
}

func Test_Handler_LakewoodStrictMatch(t *testing.T) {
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
			name:           "name search, db error",
			firstName:      "db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: get officer by name db error"),
		},
		{
			name:           "name search, officers found",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"date":"1889-05-01","last_name":"lak","first_name":"first"},{"date":"1889-05-01","last_name":"poo","first_name":"first"},{"date":"1889-05-01","last_name":"poo","first_name":"test"}]` + "\n"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/lakewood/officer?first_name=%s&last_name=%s", ts.URL, tt.firstName, tt.lastName))

			if err != nil{
				t.Fatalf("Got unexpected error while testing:\ntest name: %s\nerror: %s", tt.name, err)
			}
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

func Test_Handler_LakewoodFuzzySearch(t *testing.T) {
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
			name:           "name search, db error",
			firstName:      "db_error",
			lastName:		"db_error",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   []byte("error getting officer: fuzzy search by name db error"),
		},
		{
			name:           "name search, officers found",
			lastName:       "test",
			expectedStatus: http.StatusOK,
			expectedBody:   []byte(`[{"date":"1889-05-01","last_name":"lak","first_name":"first"}]` + "\n"),
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, err := http.Get(fmt.Sprintf("%s/lakewood/officer/search?first_name=%s&last_name=%s", ts.URL, tt.firstName, tt.lastName))

			if err != nil{
				t.Fatalf("Got unexpected error while testing:\ntest name: %s\nerror: %s", tt.name, err)
			}
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