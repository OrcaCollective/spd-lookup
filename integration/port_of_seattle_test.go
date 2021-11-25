package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/OrcaCollective/spd-lookup/api/data"
)

// Test PortOfSeattle strict match endpoint
func testPortOfSeattleStrict(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]struct {
		name               string
		badge              string
		searchName         string
		expectedStatus     int
		expectedBody       []byte
		expectedBodyCheck  string
		expectedBodyLength int
	}{
		{
			name:              "StrictNoParams",
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      []byte("name must be provided"),
			expectedBodyCheck: "EqualsBytes",
		},
		{
			name:               "BadgeStrictSearch",
			badge:              "10197",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  "EqualsLength",
			expectedBodyLength: 1,
		},
		{
			name:               "FirstNameStrictSearch",
			searchName:         "Patrick",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  "GreaterThanLength",
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameStrictSearch",
			searchName:         "Addison",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  "EqualsLength",
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameStrictSearch",
			searchName:         "Addison%2C%20Patrick",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  "EqualsLength",
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/port_of_seattle/officer?badge=%s&name=%s", testServer, tt.badge, tt.searchName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("\nTest: %s\nExpected status %d; got %d", tt.name, tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			if tt.expectedBodyCheck == "EqualsBytes" {
				if !bytes.Equal(tt.expectedBody, resp) {
					t.Errorf("\nTest: %s\nExpected resp %s; got %s", tt.name, tt.expectedBody, resp)
				}
			} else if tt.expectedBodyCheck == "EqualsLength" {
				var respJson []data.PortOfSeattleOfficer
				err := json.Unmarshal(resp, &respJson)
				if err != nil {
					t.Errorf("\nTest: %s\nUnexpected error unmarsheling JSON response: %v", tt.name, err)
				}
				if len(respJson) != tt.expectedBodyLength {
					t.Errorf("\nTest: %s\nExpected body length %d; go %d", tt.name, tt.expectedBodyLength, len(respJson))
				}
			} else if tt.expectedBodyCheck == "GreaterThanLength" {
				var respJson []data.PortOfSeattleOfficer
				err := json.Unmarshal(resp, &respJson)
				if err != nil {
					t.Errorf("\nTest: %s\nUnexpected error unmarsheling JSON response: %v", tt.name, err)
				}
				if len(respJson) <= tt.expectedBodyLength {
					t.Errorf("\nTest: %s\nExpected body length > %d; go %d", tt.name, tt.expectedBodyLength, len(respJson))
				}
			} else {
				t.Errorf("\nTest: %s\nInvalid body check passed: %s", tt.name, tt.expectedBodyCheck)
			}
		})
	}
}

// Test PortOfSeattle fuzzy endpoint
func testPortOfSeattleFuzzy(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]struct {
		name               string
		searchName         string
		badge              string
		expectedStatus     int
		expectedBody       []byte
		expectedBodyCheck  string
		expectedBodyLength int
	}{
		{
			name:              "FuzzyNoParams",
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      []byte("provided name must not be empty"),
			expectedBodyCheck: "EqualsBytes",
		},
		{
			name:               "FirstNameFuzzySearch",
			searchName:         "Patrick",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  "GreaterThanLength",
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameFuzzySearch",
			searchName:         "Addison",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  "EqualsLength",
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameFuzzySearch",
			searchName:         "Patrick%20Addison",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  "EqualsLength",
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/port_of_seattle/officer/search?name=%s", testServer, tt.searchName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("\nTest: %s\nExpected status %d, got %d", tt.name, tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			if tt.expectedBodyCheck == "EqualsBytes" {
				if !bytes.Equal(tt.expectedBody, resp) {
					t.Errorf("\nTest: %s\nExpected resp %s; got %s", tt.name, tt.expectedBody, resp)
				}
			} else if tt.expectedBodyCheck == "EqualsLength" {
				var respJson []data.PortOfSeattleOfficer
				err := json.Unmarshal(resp, &respJson)
				if err != nil {
					t.Errorf("\nTest: %s\nUnexpected error unmarsheling JSON response: %v", tt.name, err)
				}
				if len(respJson) != tt.expectedBodyLength {
					t.Errorf("\nTest: %s\nExpected body length %d; go %d", tt.name, tt.expectedBodyLength, len(respJson))
				}
			} else if tt.expectedBodyCheck == "GreaterThanLength" {
				var respJson []data.PortOfSeattleOfficer
				err := json.Unmarshal(resp, &respJson)
				if err != nil {
					t.Errorf("\nTest: %s\nUnexpected error unmarsheling JSON response: %v", tt.name, err)
				}
				if len(respJson) <= tt.expectedBodyLength {
					t.Errorf("\nTest: %s\nExpected body length > %d; go %d", tt.name, tt.expectedBodyLength, len(respJson))
				}
			} else {
				t.Errorf("\nTest: %s\nInvalid body check passed: %s", tt.name, tt.expectedBodyCheck)
			}
		})
	}
}
