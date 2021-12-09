package integration

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

// Test Tacoma strict match endpoint
func testTacomaStrict(ctx context.Context, t *testing.T, profile string) {
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
			firstName:          "Mary",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  GreaterThanLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameStrictSearch",
			lastName:           "Aguirre",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameStrictSearch",
			firstName:          "Mary",
			lastName:           "Aguirre",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/tacoma/officer?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}

// Test Tacoma fuzzy endpoint
func testTacomaFuzzy(ctx context.Context, t *testing.T, profile string) {
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
			firstName:          "Mary",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  GreaterThanLength,
			expectedBodyLength: 1,
		},
		{
			name:               "LastNameFuzzySearch",
			lastName:           "Aguirre",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
		{
			name:               "FirstAndLastNameFuzzySearch",
			firstName:          "Mary",
			lastName:           "Aguirre",
			expectedStatus:     http.StatusOK,
			expectedBodyCheck:  EqualsLength,
			expectedBodyLength: 1,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			res, _ := http.Get(fmt.Sprintf("%s/tacoma/officer/search?first_name=%s&last_name=%s", testServer, tt.firstName, tt.lastName))

			if res.StatusCode != tt.expectedStatus {
				t.Fatalf("Expected status %d, got %d", tt.expectedStatus, res.StatusCode)
			}

			defer res.Body.Close()
			resp, _ := ioutil.ReadAll(res.Body)

			checkBody(resp, tt, t)
		})
	}
}
