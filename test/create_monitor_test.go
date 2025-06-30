package test

import (
	"bytes"
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

	t.Run("url is required", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		user := tc.CreateUser("Test User", "test@example.com", "password")
		cookie := tc.CreateSesionCookie(user)

		req, _ := http.NewRequest(http.MethodPost, tc.Server.URL+"/v1/monitor", bytes.NewReader(nil))
		req.AddCookie(cookie)

		res, err := tc.Client.Do(req)
		if err != nil {
			t.Fatalf("error: %v", err)
		}

		if res.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected 400 status code, got %v", res.StatusCode)
		}
	})
}
