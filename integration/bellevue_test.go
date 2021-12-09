package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Test Bellevue strict match endpoint
func testBellevueStrict(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]genericTestOptions{
		{
			name:              "NoParams",
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      []byte("at least one of the following parameters must be provided: badge, first_name, last_name"),
			expectedBodyCheck: EqualsBytes,
		},
		{
			name:               "BadgeStrictSearch",
			badge:              "P420",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstNameStrictSearch",
			firstName:          "Kealii",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameStrictSearch",
			lastName:           "Akahane",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameStrictSearch",
			firstName:          "Kealii",
			lastName:           "Akahane",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/bellevue/officer?badge=%s&first_name=%s&last_name=%s", testServer, tt.badge, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}

// Test Bellevue fuzzy endpoint
func testBellevueFuzzy(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]genericTestOptions{
		{
			name:              "NoParams",
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      []byte("at least one of the following parameters must be provided: first_name, last_name"),
			expectedBodyCheck: EqualsBytes,
		},
		{
			name:               "FirstNameFuzzySearch",
			firstName:          "Kealii",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameFuzzySearch",
			lastName:           "Akahane",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameFuzzySearch",
			firstName:          "Kealii",
			lastName:           "Akahane",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/bellevue/officer/search?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}
