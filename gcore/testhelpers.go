package gcore

//import (
//	"context"
//	"fmt"
//	"net/http"
//	"net/http/httptest"
//	"net/url"
//)
//
//var (
//	// mux is a multiplexer that can be used to register handlers.
//	mux *http.ServeMux
//
//	// server is an in-memory HTTP server for testing.
//	server *httptest.Server
//
//	fakeToken = "eyJ0eXAiOiJKV1Q2lkIjo1MTEsImVtCJ1c2V2ns-QgWxERsywROmzASDniT3VHebO1DqgJ5ZIg"
//
//	fakeTokenExpireDate = "2017-04-17T01:28:15.000Z"
//
//	fakeAuthOpts = AuthOptions{
//		Username: "whatever",
//		Password: "whatever",
//	}
//)
//
//// setupHTTP prepares the mux and server.
//func setupHTTP() {
//	mux = http.NewServeMux()
//	server = httptest.NewServer(mux)
//}
//
//// teardownHTTP releases HTTP-related resources.
//func teardownHTTP() {
//	server.Close()
//}
//
//// Setup endpoints for getting GCore token
//func setupGCoreAuthServer() {
//	mux.HandleFunc(loginURL, func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, `{
//			"token": "%s",
//			"expire": "%s"
//		}`, fakeToken, fakeTokenExpireDate)
//	})
//}
//
//// TODO: fix me
//func mockClientURL() *url.URL {
//	return &url.URL{
//		Scheme:     "http",
//		User:       nil,
//		Host:       server.URL[7:],
//		ForceQuery: false,
//	}
//}
//
//func getAuthenticatedCommonClient() *CommonClient {
//	common := NewCommonClient()
//	common.BaseURL = mockClientURL()
//	_ = common.Authenticate(context.Background(), fakeAuthOpts)
//	return common
//}
//
//func getAuthenticatedResellerClient() *ResellerClient {
//	resell := NewResellerClient()
//	resell.BaseURL = mockClientURL()
//	_ = resell.Authenticate(context.Background(), fakeAuthOpts)
//	return resell
//}
