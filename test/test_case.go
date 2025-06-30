package test

import (
	"net/http"
	"net/http/httptest"
	"selfhosted/config"
	"selfhosted/handler"
	"testing"
)

type TestCase struct {
	Server *httptest.Server
	Client *http.Client
	T      *testing.T
}

func NewTestCase(t *testing.T) *TestCase {
	config.DatabaseDsn = ":memory:"

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
