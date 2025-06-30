package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"selfhosted/config"
	"selfhosted/database"
	"selfhosted/database/store"
	"selfhosted/handler"
	"strings"
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

func (tc *TestCase) Get(url string) {
	req, _ := http.NewRequest(http.MethodGet, tc.Server.URL+url, nil)

	if tc.User != nil {
		cookie := tc.CreateSesionCookie(tc.User)
		req.AddCookie(cookie)
	}

	res, err := tc.Client.Do(req)
	if err != nil {
		tc.T.Fatalf("get request error: %v", err)
	}

	tc.LastResponse = res
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

func (tc *TestCase) AssertDatabaseCount(table string, count int) {
	query := fmt.Sprintf(`SELECT COUNT(*) AS count FROM %s`, table)
	row := database.DB().QueryRow(query)

	var actual int
	err := row.Scan(&actual)
	if err != nil {
		tc.T.Fatalf("unexpected scan error: %v", err)
	}

	if actual != count {
		tc.T.Fatalf("expected %v rows in %v table, got %v instead", count, table, actual)
	}
}

func (tc *TestCase) AssertDatabaseHas(table string, filters map[string]any) {
	if len(filters) == 0 {
		tc.T.Fatalf("no filters provided for AssertDatabaseHas")
	}

	where := ""
	args := make([]any, 0, len(filters))
	for k, v := range filters {
		if where != "" {
			where += " AND "
		}
		where += fmt.Sprintf("%s = ?", k)
		args = append(args, v)
	}

	query := fmt.Sprintf(`SELECT COUNT(*) AS count FROM %s WHERE %s`, table, where)
	row := database.DB().QueryRow(query, args...)

	var actual int
	err := row.Scan(&actual)
	if err != nil {
		tc.T.Fatalf("unexpected scan error: %v", err)
	}

	if actual == 0 {
		tc.T.Fatalf("expected at least 1 row in %v table matching filters, got 0", table)
	}
}

func (tc *TestCase) AssertJSONKeys(expected any) {
	if tc.LastResponse == nil {
		tc.T.Fatalf("no response for assertion available")
	}
	defer tc.LastResponse.Body.Close()
	body, err := io.ReadAll(tc.LastResponse.Body)
	if err != nil {
		tc.T.Fatalf("failed to read response body: %v", err)
	}

	var actualMap map[string]any
	if err := json.Unmarshal(body, &actualMap); err != nil {
		tc.T.Fatalf("response is not valid JSON object: %v", err)
	}

	var expectedKeys []string
	switch v := expected.(type) {
	case []string:
		expectedKeys = v
	default:
		expectedKeys = structFieldNames(expected)
	}

	for _, key := range expectedKeys {
		if _, ok := actualMap[key]; !ok {
			tc.T.Fatalf("expected key '%s' in response JSON, but it was missing", key)
		}
	}
	for key := range actualMap {
		found := false
		for _, expKey := range expectedKeys {
			if key == expKey {
				found = true
				break
			}
		}
		if !found {
			tc.T.Fatalf("unexpected key '%s' in response JSON", key)
		}
	}
}

func (tc *TestCase) AssertJSONEquals(expected any) {
	if tc.LastResponse == nil {
		tc.T.Fatalf("no response for assertion available")
	}
	defer tc.LastResponse.Body.Close()
	body, err := io.ReadAll(tc.LastResponse.Body)
	if err != nil {
		tc.T.Fatalf("failed to read response body: %v", err)
	}

	expectedJSON, err := json.Marshal(expected)
	if err != nil {
		tc.T.Fatalf("failed to marshal expected struct: %v", err)
	}
	var expectedMap map[string]any
	if err := json.Unmarshal(expectedJSON, &expectedMap); err != nil {
		tc.T.Fatalf("expected struct is not a valid JSON object: %v", err)
	}

	var actualMap map[string]any
	if err := json.Unmarshal(body, &actualMap); err != nil {
		tc.T.Fatalf("response is not valid JSON object: %v", err)
	}

	if !reflect.DeepEqual(expectedMap, actualMap) {
		tc.T.Fatalf("expected JSON response to equal: %v, got: %v", expectedMap, actualMap)
	}
}

func (tc *TestCase) AssertJSONDoesNotContain(s string) {
	if tc.LastResponse == nil {
		tc.T.Fatalf("no response for assertion available")
	}
	defer tc.LastResponse.Body.Close()
	body, err := io.ReadAll(tc.LastResponse.Body)
	if err != nil {
		tc.T.Fatalf("failed to read response body: %v", err)
	}
	if strings.Contains(string(body), s) {
		tc.T.Fatalf("did not expect to find '%s' in JSON response, but it was present", s)
	}
}

func structFieldNames(s any) []string {
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	var names []string
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		if tag == "-" {
			continue
		}
		name := strings.Split(tag, ",")[0]
		if name == "" {
			name = f.Name
		}
		names = append(names, name)
	}
	return names
}
