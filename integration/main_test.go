package integration

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

// Default test server
var testServer = "http://localhost:1312"

// validateFunc are for subtests that share a single setup
type validateFunc func(context.Context, *testing.T, string)

// Timeout value for integration tests
const ctxTimeout = 30 * time.Second

// Main integration test runner
func TestIntegrations(t *testing.T) {
	profile := "integration"
	ctx, _ := context.WithTimeout(context.Background(), ctxTimeout)

	// Serial tests
	t.Run("serial", func(t *testing.T) {
		tests := []struct {
			name      string
			validator validateFunc
		}{
			{"TestPing", testPing},
		}
		for _, tc := range tests {
			tc := tc
			if ctx.Err() == context.DeadlineExceeded {
				t.Fatalf("Unable to run more tests (deadline exceeded)")
			}
			t.Run(tc.name, func(t *testing.T) {
				tc.validator(ctx, t, profile)
			})
		}
	})
}

// Test ping endpoint
func testPing(ctx context.Context, t *testing.T, profile string) {
	expectedResponse := []byte(`🏓 P O N G 🏓`)
	resp, err := http.Get(testServer + "/ping")
	if err != nil {
		t.Errorf("Unspecified error with request: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Unspecified error with reading response: %v", err)
	} else if !bytes.Equal(expectedResponse, body) {
		t.Errorf("Ping failed, got:%s, want:%s", body, expectedResponse)
	}
}