package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"selfhosted/config"
	"selfhosted/database"
	"selfhosted/database/store"
	"selfhosted/handler"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type TestCase struct {
	Server *httptest.Server
	Client *http.Client
	T      *testing.T
}

func NewTestCase(t *testing.T) *TestCase {
	config.DatabaseDsn = ":memory:"

	database.Connect()

	server := httptest.NewServer(handler.NewRouter())

	return &TestCase{
		Server: server,
		Client: server.Client(),
		T:      t,
	}
}

func (tc *TestCase) Close() {
	tc.Server.Close()
}

func (tc *TestCase) CreateUser(name, email, password string) *store.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	user, err := database.New().CreateUser(context.Background(), store.CreateUserParams{
		Name:         name,
		Email:        name,
		PasswordHash: string(hash),
	})
	if err != nil {
		tc.T.Fatalf("unable to create new user: %v", err)
	}

	return &user
}

func (tc *TestCase) CreateSesionCookie(user *store.User) *http.Cookie {
	uuid := uuid.NewString()
	database.New().CreateSession(context.Background(), store.CreateSessionParams{
		Uuid:      uuid,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	})

	return &http.Cookie{
		Value:    uuid,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Name:     "session",
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false,
		HttpOnly: true,
	}
}
