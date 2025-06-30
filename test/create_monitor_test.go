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

	t.Run("url validation", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		tc.Authenticated()
		tc.Post("/v1/monitor", map[string]string{
			"url": "invalid",
		})
		tc.AssertStatus(http.StatusBadRequest)
	})

	t.Run("monitor created", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		tc.Authenticated()
		tc.Post("/v1/monitor", map[string]string{
			"url": "https://google.com",
		})
		tc.AssertStatus(http.StatusCreated)
		tc.AssertDatabaseCount("monitors", 1)
		tc.AssertDatabaseHas("monitors", map[string]any{
			"user_id": 1,
			"url":     "https://google.com",
		})
	})
}
