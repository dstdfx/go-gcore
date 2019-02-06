package testhelper

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
)

const (
	TestFakeToken           = "eyJ0eXAiOiJKV1Q2lkIjo1MTEsImVtCJ1c2V2ns-QgWxERsywROmzASDniT3VHebO1DqgJ5ZIg"
	TestFakeTokenExpireDate = "2017-04-17T01:28:15.000Z"
)

// TestEnv represents a testing environment for all resources.
type TestEnv struct {
	Mux       *http.ServeMux
	Server    *httptest.Server
	ServerURL string
}

// SetupTestEnv prepares the new testing environment.
func SetupTestEnv() *TestEnv {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	testEnv := &TestEnv{
		Mux:    mux,
		Server: server,
	}

	testEnv.setupGCoreAuthServer()
	return testEnv
}

// GetServerURL returns server *url.URL.
func (tenv *TestEnv) GetServerURL() *url.URL {
	serverURL, _ := url.Parse(tenv.Server.URL)
	return serverURL
}

// TearDownTestEnv releases the testing environment.
func (tenv *TestEnv) TearDownTestEnv() {
	tenv.Server.Close()
	tenv.Server = nil
	tenv.Mux = nil
}

// setupGCoreAuthServer setups endpoints for getting GCore token.
func (tenv *TestEnv) setupGCoreAuthServer() {
	tenv.Mux.HandleFunc("/auth/signin", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintf(w, `{
			"token": "%s",
			"expire": "%s"
		}`, TestFakeToken, TestFakeTokenExpireDate)
	})
}
