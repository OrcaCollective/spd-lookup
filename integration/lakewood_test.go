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

// Test Lakewood strict match endpoint
func testLakewoodStrict(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]struct {
		name           string
		firstName      string
		lastName       string
		badge          string
		expectedStatus int
		expectedBody   []byte
		expectedBodyCheck string
		expectedBodyLength int
	}{
		{
			name:           "NoParams",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte("at least one of the following parameters must be provided: first_name, last_name"),
			expectedBodyCheck: "EqualsBytes",
		},
		{
			name: "FirstNameStrictSearch",
			firstName: "Mike",
			expectedStatus: http.StatusOK,
			expectedBodyCheck: "GreaterThanLength",
			expectedBodyLength: 1,
		},
		{
			name:           "LastNameStrictSearch",
			lastName:       "Zaro",
			expectedStatus: http.StatusOK,
			expectedBodyCheck: "EqualsLength",
			expectedBodyLength: 1,
		},
		{
			name: "FirstAndLastNameStrictSearch",
			firstName: "Mike",
			lastName: "Zaro",
			expectedStatus: http.StatusOK,
			expectedBodyCheck: "EqualsLength",
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/lakewood/officer?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			if tt.expectedBodyCheck == "EqualsBytes" {
				if !bytes.Equal(tt.expectedBody, resp) {
					t.Errorf("\nTest: %s\nExpected resp %s; got %s", tt.name, tt.expectedBody, resp)
				}
			} else if tt.expectedBodyCheck == "EqualsLength" {
				var respJson []data.LakewoodOfficer
				err := json.Unmarshal(resp, &respJson)
				if err != nil {
					t.Errorf("\nTest: %s\nUnexpected error unmarsheling JSON response: %v", tt.name, err)
				}
				if len(respJson) != tt.expectedBodyLength {
					t.Errorf("\nTest: %s\nExpected body length %d; go %d", tt.name, tt.expectedBodyLength, len(respJson))
				}
			} else if tt.expectedBodyCheck == "GreaterThanLength" {
				var respJson []data.LakewoodOfficer
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

// Test Lakewood fuzzy endpoint
func testLakewoodFuzzy(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]struct {
		name           string
		firstName      string
		lastName       string
		badge          string
		expectedStatus int
		expectedBody   []byte
		expectedBodyCheck string
		expectedBodyLength int
	}{
		{
			name:           "NoParams",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   []byte("at least one of the following parameters must be provided: first_name, last_name"),
			expectedBodyCheck: "EqualsBytes",
		},
		{
			name: "FirstNameFuzzySearch",
			firstName: "Mike",
			expectedStatus: http.StatusOK,
			expectedBodyCheck: "GreaterThanLength",
			expectedBodyLength: 1,
		},
		{
			name:           "LastNameFuzzySearch",
			lastName:       "Zaro",
			expectedStatus: http.StatusOK,
			expectedBodyCheck: "EqualsLength",
			expectedBodyLength: 1,
		},
		{
			name: "FirstAndLastNameFuzzySearch",
			firstName: "Mike",
			lastName: "Zaro",
			expectedStatus: http.StatusOK,
			expectedBodyCheck: "GreaterThanLength",
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/lakewood/officer/search?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			if tt.expectedBodyCheck == "EqualsBytes" {
				if !bytes.Equal(tt.expectedBody, resp) {
					t.Errorf("\nTest: %s\nExpected resp %s; got %s", tt.name, tt.expectedBody, resp)
				}
			} else if tt.expectedBodyCheck == "EqualsLength" {
				var respJson []data.LakewoodOfficer
				err := json.Unmarshal(resp, &respJson)
				if err != nil {
					t.Errorf("\nTest: %s\nUnexpected error unmarsheling JSON response: %v", tt.name, err)
				}
				if len(respJson) != tt.expectedBodyLength {
					t.Errorf("\nTest: %s\nExpected body length %d; go %d", tt.name, tt.expectedBodyLength, len(respJson))
				}
			} else if tt.expectedBodyCheck == "GreaterThanLength" {
				var respJson []data.LakewoodOfficer
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