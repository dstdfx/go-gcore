package testutils

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/dstdfx/go-gcore/gcore"
)

var (
	// Mux is a multiplexer that can be used to register handlers.
	Mux *http.ServeMux

	// Server is an in-memory HTTP server for testing.
	Server *httptest.Server

	FakeToken = "eyJ0eXAiOiJKV1Q2lkIjo1MTEsImVtCJ1c2V2ns-QgWxERsywROmzASDniT3VHebO1DqgJ5ZIg"

	FakeTokenExpireDate = "2017-04-17T01:28:15.000Z"

	FakeAuthOpts = gcore.AuthOptions{
		Username: "whatever",
		Password: "whatever",
	}
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
	Mux.HandleFunc(gcore.LoginURL, func(w http.ResponseWriter, r *http.Request) {
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

func GetAuthenticatedCommonClient() *gcore.CommonClient {
	common := gcore.NewCommonClient()
	common.BaseURL = MockClientURL()
	_ = common.Authenticate(context.Background(), FakeAuthOpts)
	return common
}

func GetAuthenticatedResellerClient() *gcore.ResellerClient {
	resell := gcore.NewResellerClient()
	resell.BaseURL = MockClientURL()
	_ = resell.Authenticate(context.Background(), FakeAuthOpts)
	return resell
}
