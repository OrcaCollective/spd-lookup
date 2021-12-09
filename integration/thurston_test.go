package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Test Thurston strict match endpoint
func testThurstonStrict(ctx context.Context, t *testing.T, profile string) {
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
			firstName:          "Jennifer",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameStrictSearch",
			lastName:           "Mcaneney",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameStrictSearch",
			firstName:          "Jennifer",
			lastName:           "Mcaneney",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/thurston_county/officer?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}

// Test Thurston fuzzy endpoint
func testThurstonFuzzy(ctx context.Context, t *testing.T, profile string) {
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
			firstName:          "Jennifer",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  GreaterThanLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameFuzzySearch",
			lastName:           "Mcaneney",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameFuzzySearch",
			firstName:          "Jennifer",
			lastName:           "Mcaneney",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/thurston_county/officer/search?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}
