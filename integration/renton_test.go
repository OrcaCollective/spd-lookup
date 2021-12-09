package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Test Renton strict match endpoint
func testRentonStrict(ctx context.Context, t *testing.T, profile string) {
	t.Parallel()
	for _, tt := range [...]genericTestOptions{
		{
			name:              "NoParams",
			expectedStatus:    http.StatusBadRequest,
			expectedBody:      []byte("at least one of the following parameters must be provided: first_name, last_name"),
			expectedBodyCheck: EqualsBytes,
		},
		{
			name:               "FirstNameStrictSearch",
			firstName:          "Augustine",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameStrictSearch",
			lastName:           "Abram",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameStrictSearch",
			firstName:          "Augustine",
			lastName:           "Abram",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/renton/officer?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}

// Test Renton fuzzy endpoint
func testRentonFuzzy(ctx context.Context, t *testing.T, profile string) {
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
			firstName:          "Augustine",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameFuzzySearch",
			lastName:           "Abram",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameFuzzySearch",
			firstName:          "Augustine",
			lastName:           "Abram",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/renton/officer/search?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}
