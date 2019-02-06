package testhelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

// HandleReqOpts represents options for the testing utils package handlers.
type HandleReqOpts struct {
	// Mux represents HTTP Mux for a testing handler.
	Mux *http.ServeMux

	// URL represents handler's HTTP URL.
	URL string

	// RawResponse represents raw string HTTP response that needs to be returned
	// by the handler.
	RawResponse string

	// RawRequest represents raw string HTTP request that needs to be compared
	// with the actual request that will be provided by the caller.
	RawRequest string

	// QueryParams represents query params that has to be sent in URL.
	QueryParams map[string]string

	// Method contains HTTP method that needs to be compared against real method
	// provided by the caller.
	Method string

	// Status represents HTTP status that will be returned by the handler.
	Status int

	// CallFlag can be used to check if caller sent a request to a handler.
	CallFlag *bool
}

// HandleReqWithoutBody provides the HTTP endpoint to test requests without body.
func HandleReqWithoutBody(t *testing.T, opts *HandleReqOpts) {
	opts.Mux.HandleFunc(opts.URL, func(w http.ResponseWriter, r *http.Request) {
		checkQueryParams(t, r.URL, opts)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(opts.Status)
		fmt.Fprint(w, opts.RawResponse)

		if r.Method != opts.Method {
			t.Fatalf("expected %s method but got %s", opts.Method, r.Method)
		}

		*opts.CallFlag = true
	})
}

// HandleReqWithBody provides the HTTP endpoint to test requests with body.
func HandleReqWithBody(t *testing.T, opts *HandleReqOpts) {
	opts.Mux.HandleFunc(opts.URL, func(w http.ResponseWriter, r *http.Request) {
		checkQueryParams(t, r.URL, opts)

		if r.Method != opts.Method {
			t.Fatalf("expected %s method but got %s", opts.Method, r.Method)
		}

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read the request body: %v", err)
		}
		defer r.Body.Close()

		var actualRequest interface{}
		err = json.Unmarshal(b, &actualRequest)
		if err != nil {
			t.Errorf("unable to unmarshal the request body: %v", err)
		}

		var expectedRequest interface{}
		err = json.Unmarshal([]byte(opts.RawRequest), &expectedRequest)
		if err != nil {
			t.Errorf("unable to unmarshal expected raw request: %v", err)
		}

		if !reflect.DeepEqual(expectedRequest, actualRequest) {
			t.Fatalf("expected %#v request, but got %#v", expectedRequest, actualRequest)
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(opts.Status)
		fmt.Fprint(w, opts.RawResponse)

		*opts.CallFlag = true
	})
}

// checkQueryParams compares actual and expected query params.
func checkQueryParams(t *testing.T, requestURL *url.URL, opts *HandleReqOpts) {
	gotQueryParams := requestURL.Query()
	if opts.QueryParams != nil {
		if len(gotQueryParams) != len(opts.QueryParams) {
			t.Fatalf("expected %d query params, but got %d",
				len(opts.QueryParams),
				len(gotQueryParams))
		}

		for param, value := range opts.QueryParams {
			v, ok := gotQueryParams[param]
			if !ok {
				t.Fatalf("query param %s hasn't been sent", param)
			}

			if v[0] != value {
				t.Fatalf("query params %s values mistmatch, expected %s, but got %s",
					param,
					value,
					v)
			}
		}
	}
}
