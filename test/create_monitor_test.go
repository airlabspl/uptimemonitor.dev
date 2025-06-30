package test

import (
	"net/http"
	"testing"
)

func TestCreateMonitor(t *testing.T) {
	t.Run("guests cannot create monitors", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		res, _ := tc.Client.Post(tc.Server.URL+"/v1/monitor", "application/json", nil)
		if res.StatusCode != http.StatusUnauthorized {
			t.Fatalf("expected 401 status code, got %v", res.StatusCode)
		}
	})
}
