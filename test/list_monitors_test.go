package test

import (
	"context"
	"net/http"
	"selfhosted/database"
	"selfhosted/database/store"
	"testing"

	"github.com/google/uuid"
)

func TestListMonitors(t *testing.T) {
	t.Run("guests cannot list monitors", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		tc.Get("/v1/monitors")
		tc.AssertStatus(http.StatusUnauthorized)
	})

	t.Run("users cannot list monitors from other users", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		uuid := uuid.NewString()
		url := "http://google.com"
		user := tc.CreateUser("Other User", "other@example.com", "password")
		database.New().CreateMonitor(context.Background(), store.CreateMonitorParams{
			UserID: user.ID,
			Uuid:   uuid,
			Url:    url,
		})

		tc.Authenticated()
		tc.Get("/v1/monitors")

		tc.AssertJSONDoesNotContain(uuid)
		tc.AssertJSONDoesNotContain(url)
	})

	t.Run("users can list their monitors", func(t *testing.T) {
		tc := NewTestCase(t)
		defer tc.Close()

		tc.Authenticated()

		uuid := uuid.NewString()
		url := "http://google.com"

		database.New().CreateMonitor(context.Background(), store.CreateMonitorParams{
			UserID: 1,
			Uuid:   uuid,
			Url:    url,
		})

		tc.Get("/v1/monitors")
		tc.AssertStatus(http.StatusOK)
		tc.AssertJSONContains(uuid)
		tc.AssertJSONContains(url)
	})
}
