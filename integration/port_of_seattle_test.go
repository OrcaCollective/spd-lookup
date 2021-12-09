package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Test PortOfSeattle strict match endpoint
func testPortOfSeattleStrict(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]genericTestOptions{
		{
			name:              "StrictNoParams",
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      []byte("name must be provided"),
			expectedBodyCheck: EqualsBytes,
		},
		{
			name:               "BadgeStrictSearch",
			badge:              "10197",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstNameStrictSearch",
			searchName:         "Patrick",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  GreaterThanLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameStrictSearch",
			searchName:         "Addison",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameStrictSearch",
			searchName:         "Addison%2C%20Patrick",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
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

			checkBody(resp, tt, t)
		})
	}
}

// Test PortOfSeattle fuzzy endpoint
func testPortOfSeattleFuzzy(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]genericTestOptions{
		{
			name:              "FuzzyNoParams",
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      []byte("provided name must not be empty"),
			expectedBodyCheck: EqualsBytes,
		},
		{
			name:               "FirstNameFuzzySearch",
			searchName:         "Patrick",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  GreaterThanLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameFuzzySearch",
			searchName:         "Addison",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameFuzzySearch",
			searchName:         "Patrick%20Addison",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
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

			checkBody(resp, tt, t)
		})
	}
}
