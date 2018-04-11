package gcore

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	// Mux is a multiplexer that can be used to register handlers.
	Mux *http.ServeMux

	// Server is an in-memory HTTP server for testing.
	Server *httptest.Server

	FakeToken = "eyJ0eXAiOiJKV1Q2lkIjo1MTEsImVtCJ1c2V2ns-QgWxERsywROmzASDniT3VHebO1DqgJ5ZIg"

	FakeTokenExpireDate = "2017-04-17T01:28:15.000Z"

	FakeAuthOpts = AuthOptions{Username: "whatever", Password: "whatever"}
)

// SetupHTTP prepares the Mux and Server.
func SetupHTTP() {
	Mux = http.NewServeMux()
	Server = httptest.NewServer(Mux)
}

// TeardownHTTP releases HTTP-related resources.
func TeardownHTTP() {
	Server.Close()
}

// Endpoint returns a fake endpoint that will actually target the Mux.
func Endpoint() string {
	return Server.URL
}

// Setup endpoints for getting GCore token
func SetupGCoreAuthServer() {
	Mux.HandleFunc(loginURL, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{
			"token": "%s",
			"expire": "%s"
		}`, FakeToken, FakeTokenExpireDate)
	})
}

func MockClientURL() *url.URL {
	return &url.URL{
		Scheme:     "http",
		User:       nil,
		Host:       Server.URL[7:],
		ForceQuery: false,
	}
}

func TestNewCommonClient(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	common := NewCommonClient(nil)
	common.BaseURL = MockClientURL()

	err := common.Authenticate(context.Background(), FakeAuthOpts)
	if err != nil {
		t.Fatal(err)
	}

	if common.Token.Value != FakeToken {
		t.Errorf("Expected: %s, got %s", FakeToken, common.Token.Value)
	}
}

func TestNewResellerClient(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	reseller := NewResellerClient(nil)
	reseller.BaseURL = MockClientURL()

	err := reseller.Authenticate(context.Background(), FakeAuthOpts)
	if err != nil {
		t.Fatal(err)
	}

	if reseller.Token.Value != FakeToken {
		t.Errorf("Expected: %s, got %s", FakeToken, reseller.Token.Value)
	}

}

func GetAuthenticatedCommonClient() *CommonClient {
	common := NewCommonClient(nil)
	common.BaseURL = MockClientURL()
	common.Authenticate(context.Background(), FakeAuthOpts)
	return common
}
