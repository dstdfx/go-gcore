package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

// Fixtures
const (
	testCreateRuleRawResponse = `{
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
	testGetRuleRawResponse = `{
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
	testListRuleRawResponse = `[
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

const (
	testCreateRuleRawRequest = `{  
   "rule":"",
   "name":"whatever",
   "ruleType":0,
   "options":{  
      "cache_http_headers":{  
         "enabled":true,
         "value":[  
            "x-token"
         ]
      },
      "cache_expire":null,
      "allowedHttpMethods":null,
      "cors":null,
      "country_acl":null,
      "disable_cache":null,
      "fetch_compressed":null,
      "force_return":null,
      "gzipOn":null,
      "hostHeader":null,
      "ignoreQueryString":null,
      "ignore_cookie":null,
      "ip_address_acl":null,
      "override_browser_ttl":null,
      "proxy_cache_methods_set":null,
      "referrer_acl":null,
      "rewrite":null,
      "secure_key":null,
      "slice":null,
      "stale":null,
      "staticHeaders":null,
      "user_agent_acl":null
   }
}`
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
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(rulesURL, fakeResourceID),
		RawResponse: testCreateRuleRawResponse,
		RawRequest:  testCreateRuleRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testCreateRuleExpected

	body := &CreateRuleBody{
		RuleType: 0,
		Name:     "whatever",
		Options: Options{
			CacheHTTPHeaders: &CacheHTTPHeaders{
				Enabled: true,
				Value:   []string{"x-token"},
			},
		},
	}

	got, _, err := client.Rules.Create(context.Background(), fakeResourceID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't create rules")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_Get(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(ruleURL, fakeResourceID, testGetRuleExpected.ID),
		RawResponse: testGetRuleRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testGetRuleExpected

	got, _, err := client.Rules.Get(context.Background(), fakeResourceID, testGetRuleExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get rules")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_Delete(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(ruleURL, fakeResourceID, testGetRuleExpected.ID),
		RawResponse: "",
		Method:      http.MethodDelete,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	_, err := client.Rules.Delete(context.Background(), fakeResourceID, testGetRuleExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't delete rules")
	}
}

func TestRulesService_List(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(rulesURL, fakeResourceID),
		RawResponse: testListRuleRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testListRuleExpected

	got, _, err := client.Rules.List(context.Background(), fakeResourceID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get rules")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
