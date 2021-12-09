package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Test Olympia strict match endpoint
func testOlympiaStrict(ctx context.Context, t *testing.T, profile string) {
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
			badge:              "OLY2",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstNameStrictSearch",
			firstName:          "RICH",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameStrictSearch",
			lastName:           "ALLEN",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameStrictSearch",
			firstName:          "RICH",
			lastName:           "ALLEN",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/olympia/officer?badge=%s&first_name=%s&last_name=%s", testServer, tt.badge, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}

// Test Olympia fuzzy endpoint
func testOlympiaFuzzy(ctx context.Context, t *testing.T, profile string) {
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
			firstName:          "RICH",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameFuzzySearch",
			lastName:           "ALLEN",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameFuzzySearch",
			firstName:          "RICH",
			lastName:           "ALLEN",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/olympia/officer/search?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}
