package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// Fixtures
var (
	testCreateRuleResponse = `{
   "rule" : "/images",
   "originGroup" : null,
   "preset_applied" : false,
   "options" : {
      "force_return" : null,
      "user_agent_acl" : null,
      "allowedHttpMethods" : null,
      "staticHeaders" : null,
      "ignoreQueryString" : null,
      "disable_cache" : null,
      "ignore_cookie" : null,
      "referrer_acl" : null,
      "secure_key" : null,
      "stale" : null,
      "gzipOn" : null,
      "slice" : null,
      "country_acl" : null,
      "ip_address_acl" : null,
      "cors" : null,
      "fetch_compressed" : null,
      "hostHeader" : null,
      "override_browser_ttl" : null,
      "proxy_cache_methods_set" : null,
      "cache_http_headers" : {
         "value" : [
            "content-length",
            "x-token",
            "connection",
            "date",
            "server",
            "content-type"
         ],
         "enabled" : true
      },
      "cache_expire" : null,
      "rewrite" : null
   },
   "ruleType" : 0,
   "name" : "whatever",
   "id" : 1861,
   "weight" : 2,
   "originProtocol" : "HTTP"
}`
	testGetRuleResponse = `{
   "rule" : "/images",
   "originGroup" : null,
   "preset_applied" : false,
   "options" : {
      "force_return" : null,
      "user_agent_acl" : null,
      "allowedHttpMethods" : null,
      "staticHeaders" : null,
      "ignoreQueryString" : null,
      "disable_cache" : null,
      "ignore_cookie" : null,
      "referrer_acl" : null,
      "secure_key" : null,
      "stale" : null,
      "gzipOn" : null,
      "slice" : null,
      "country_acl" : null,
      "ip_address_acl" : null,
      "cors" : null,
      "fetch_compressed" : null,
      "hostHeader" : null,
      "override_browser_ttl" : null,
      "proxy_cache_methods_set" : null,
      "cache_http_headers" : {
         "value" : [
            "content-length",
            "x-token",
            "connection",
            "date",
            "server",
            "content-type"
         ],
         "enabled" : true
      },
      "cache_expire" : null,
      "rewrite" : null
   },
   "ruleType" : 0,
   "name" : "whatever",
   "id" : 1861,
   "weight" : 2,
   "originProtocol" : "HTTP"
}`
	testListRuleResponse = `[
{
   "rule" : "/images",
   "originGroup" : null,
   "preset_applied" : false,
   "options" : {
      "force_return" : null,
      "user_agent_acl" : null,
      "allowedHttpMethods" : null,
      "staticHeaders" : null,
      "ignoreQueryString" : null,
      "disable_cache" : null,
      "ignore_cookie" : null,
      "referrer_acl" : null,
      "secure_key" : null,
      "stale" : null,
      "gzipOn" : null,
      "slice" : null,
      "country_acl" : null,
      "ip_address_acl" : null,
      "cors" : null,
      "fetch_compressed" : null,
      "hostHeader" : null,
      "override_browser_ttl" : null,
      "proxy_cache_methods_set" : null,
      "cache_http_headers" : {
         "value" : [
            "content-length",
            "x-token",
            "connection",
            "date",
            "server",
            "content-type"
         ],
         "enabled" : true
      },
      "cache_expire" : null,
      "rewrite" : null
   },
   "ruleType" : 0,
   "name" : "whatever",
   "id" : 1861,
   "weight" : 2,
   "originProtocol" : "HTTP"}]`
)

var (
	fakeResourceID         = 4538
	testCreateRuleExpected = &Rule{
		ID:             1861,
		OriginProtocol: "HTTP",
		RuleType:       0,
		Weight:         2,
		Name:           "whatever",
		Rule:           "/images",
		PresetApplied:  false,
		Options: Options{
			CacheHTTPHeaders: &CacheHTTPHeaders{
				Enabled: true,
				Value:   []string{"content-length", "x-token", "connection", "date", "server", "content-type"},
			},
		},
	}
	testGetRuleExpected = &Rule{
		ID:             1861,
		OriginProtocol: "HTTP",
		RuleType:       0,
		Weight:         2,
		Name:           "whatever",
		Rule:           "/images",
		PresetApplied:  false,
		Options: Options{
			CacheHTTPHeaders: &CacheHTTPHeaders{
				Enabled: true,
				Value:   []string{"content-length", "x-token", "connection", "date", "server", "content-type"},
			},
		},
	}
	testListRuleExpected = []*Rule{
		{
			ID:             1861,
			OriginProtocol: "HTTP",
			RuleType:       0,
			Weight:         2,
			Name:           "whatever",
			Rule:           "/images",
			PresetApplied:  false,
			Options: Options{
				CacheHTTPHeaders: &CacheHTTPHeaders{
					Enabled: true,
					Value:   []string{"content-length", "x-token", "connection", "date", "server", "content-type"},
				},
			},
		},
	}
)

func TestRulesService_Create(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testCreateRuleExpected
	mux.HandleFunc(fmt.Sprintf(RulesURL, fakeResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testCreateRuleResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	body := CreateRuleBody{
		RuleType: 0,
		Name:     "whatever",
		Options: Options{
			CacheHTTPHeaders: &CacheHTTPHeaders{
				Enabled: true,
				Value:   []string{"x-token"},
			},
		},
	}

	client := getAuthenticatedCommonClient()
	got, _, err := client.Rules.Create(context.Background(), fakeResourceID, &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_Get(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testGetRuleExpected
	mux.HandleFunc(fmt.Sprintf(RuleURL, fakeResourceID, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testGetRuleResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := getAuthenticatedCommonClient()
	got, _, err := client.Rules.Get(context.Background(), fakeResourceID, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_List(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testListRuleExpected
	mux.HandleFunc(fmt.Sprintf(RulesURL, fakeResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testListRuleResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := getAuthenticatedCommonClient()
	got, _, err := client.Rules.List(context.Background(), fakeResourceID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_Delete(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testGetRuleExpected
	mux.HandleFunc(fmt.Sprintf(RuleURL, fakeResourceID, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		})

	client := getAuthenticatedCommonClient()
	_, err := client.Rules.Delete(context.Background(), fakeResourceID, expected.ID)
	if err != nil {
		t.Fatal(err)
	}
}
