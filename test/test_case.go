package test

import (
	"bytes"
	"context"
	"encoding/json"
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
	Server       *httptest.Server
	Client       *http.Client
	T            *testing.T
	User         *store.User
	LastResponse *http.Response
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

func (tc *TestCase) ActingAs(user *store.User) {
	tc.User = user
}

func (tc *TestCase) Authenticated() {
	tc.User = tc.CreateUser("Test User", "test@example.com", "password")
}

func (tc *TestCase) Post(url string, data any) {
	body, _ := json.Marshal(data)

	req, _ := http.NewRequest(http.MethodPost, tc.Server.URL+url, bytes.NewBuffer(body))

	if tc.User != nil {
		cookie := tc.CreateSesionCookie(tc.User)
		req.AddCookie(cookie)
	}

	res, err := tc.Client.Do(req)
	if err != nil {
		tc.T.Fatalf("post request error: %v", err)
	}

	tc.LastResponse = res
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

func (tc *TestCase) AssertStatus(statusCode int) {
	if tc.LastResponse == nil {
		tc.T.Fatalf("no response for assertion availabe")
	}

	if tc.LastResponse.StatusCode != statusCode {
		tc.T.Fatalf("expected %v status code, but got: %v", statusCode, tc.LastResponse.StatusCode)
	}
}
