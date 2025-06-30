package test

import (
	"net/http"
	"testing"
)

func TestCreateMonitor(t *testing.T) {
	t.Run("guests cannot create monitors", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		tc.Post("/v1/monitor", struct{}{})
		tc.AssertStatus(http.StatusUnauthorized)
	})

	t.Run("url is required", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		tc.Authenticated()
		tc.Post("/v1/monitor", struct{}{})
		tc.AssertStatus(http.StatusBadRequest)
	})
}
