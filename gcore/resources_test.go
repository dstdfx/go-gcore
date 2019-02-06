package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

const (
	testGetResourceRawResponse = `{
  "id": 4478,
  "deleted": false,
  "secondaryHostnames": [],
  "rules": [],
  "companyName": "Whatever inc",
  "client": 2096,
  "status": "active",
  "active": true,
  "enabled": true,
  "preset_applied": false,
  "options": {
    "slice": null,
    "gzipOn": null,
    "ignoreQueryString": null,
    "hostHeader": {
      "enabled": true,
      "value": "131231.example.ru"
    },
    "staticHeaders": null,
    "allowedHttpMethods": null,
    "stale": {
      "enabled": true,
      "value": [
        "error",
        "updating"
      ]
    },
    "cors": null,
    "proxy_cache_methods_set": null,
    "rewrite": null,
    "force_return": null,
    "secure_key": null,
    "cache_expire": {
      "enabled": true,
      "value": 345600
    },
    "disable_cache": null,
    "ignore_cookie": null,
    "cache_http_headers": null,
    "override_browser_ttl": null,
    "fetch_compressed": null,
    "country_acl": null,
    "referrer_acl": null,
    "user_agent_acl": null,
    "ip_address_acl": null
  },
  "name": null,
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "originProtocol": "HTTPS",
  "cname": "gcdn.example.me",
  "logTarget": "",
  "sslEnabled": true,
  "shielded": false,
  "shieldDatacenter": "",
  "originGroup": 7260,
  "sslData": 1189
}`
	testUpdateResourceRawResponse = `{
  "id": 4478,
  "deleted": false,
  "secondaryHostnames": [],
  "rules": [],
  "companyName": "Whatever inc",
  "client": 2096,
  "status": "active",
  "active": true,
  "enabled": true,
  "preset_applied": false,
  "options": {
    "slice": null,
    "gzipOn": null,
    "ignoreQueryString": null,
    "hostHeader": {
      "enabled": true,
      "value": "131231.example.ru"
    },
    "staticHeaders": null,
    "allowedHttpMethods": null,
    "stale": {
      "enabled": true,
      "value": [
        "error",
        "updating"
      ]
    },
    "cors": null,
    "proxy_cache_methods_set": null,
    "rewrite": null,
    "force_return": null,
    "secure_key": null,
    "cache_expire": {
      "enabled": true,
      "value": 345600
    },
    "disable_cache": null,
    "ignore_cookie": null,
    "cache_http_headers": null,
    "override_browser_ttl": null,
    "fetch_compressed": null,
    "country_acl": null,
    "referrer_acl": null,
    "user_agent_acl": null,
    "ip_address_acl": null
  },
  "name": null,
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "originProtocol": "HTTPS",
  "cname": "gcdn.example.me",
  "logTarget": "",
  "sslEnabled": true,
  "shielded": false,
  "shieldDatacenter": "",
  "originGroup": 7260,
  "sslData": 1189
}`

	testCreateResourceRawResponse = `{
  "id": 4478,
  "deleted": false,
  "secondaryHostnames": [],
  "rules": [],
  "companyName": "Whatever inc",
  "client": 2096,
  "status": "active",
  "active": true,
  "enabled": true,
  "preset_applied": false,
  "options": {
    "slice": null,
    "gzipOn": null,
    "ignoreQueryString": null,
    "hostHeader": {
      "enabled": true,
      "value": "131231.example.ru"
    },
    "staticHeaders": null,
    "allowedHttpMethods": null,
    "stale": {
      "enabled": true,
      "value": [
        "error",
        "updating"
      ]
    },
    "cors": null,
    "proxy_cache_methods_set": null,
    "rewrite": null,
    "force_return": null,
    "secure_key": null,
    "cache_expire": {
      "enabled": true,
      "value": 345600
    },
    "disable_cache": null,
    "ignore_cookie": null,
    "cache_http_headers": null,
    "override_browser_ttl": null,
    "fetch_compressed": null,
    "country_acl": null,
    "referrer_acl": null,
    "user_agent_acl": null,
    "ip_address_acl": null
  },
  "name": null,
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "originProtocol": "HTTPS",
  "cname": "gcdn.example.me",
  "logTarget": "",
  "sslEnabled": true,
  "shielded": false,
  "shieldDatacenter": "",
  "originGroup": 7260,
  "sslData": 1189
}`

	testListResourcesRawResponse = `[{
  "id": 4478,
  "deleted": false,
  "secondaryHostnames": [],
  "rules": [],
  "companyName": "Whatever inc",
  "client": 2096,
  "status": "active",
  "active": true,
  "enabled": true,
  "preset_applied": false,
  "options": {
    "slice": null,
    "gzipOn": null,
    "ignoreQueryString": null,
    "hostHeader": {
      "enabled": true,
      "value": "131231.example.ru"
    },
    "staticHeaders": null,
    "allowedHttpMethods": null,
    "stale": {
      "enabled": true,
      "value": [
        "error",
        "updating"
      ]
    },
    "cors": null,
    "proxy_cache_methods_set": null,
    "rewrite": null,
    "force_return": null,
    "secure_key": null,
    "cache_expire": {
      "enabled": true,
      "value": 345600
    },
    "disable_cache": null,
    "ignore_cookie": null,
    "cache_http_headers": null,
    "override_browser_ttl": null,
    "fetch_compressed": null,
    "country_acl": null,
    "referrer_acl": null,
    "user_agent_acl": null,
    "ip_address_acl": null
  },
  "name": null,
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "originProtocol": "HTTPS",
  "cname": "gcdn.example.me",
  "logTarget": "",
  "sslEnabled": true,
  "shielded": false,
  "shieldDatacenter": "",
  "originGroup": 7260,
  "sslData": 1189
}]`
)

const (
	testCreateResourceRawRequest = `{
	"cname":"cdn.site.com",
	"secondaryHostnames":["cdn1.yoursite.com","cdn2.yoursite.com"]
}
`
	testUpdateResourceRawRequest = `{
	"originGroup":7260,
	"secondaryHostnames":["cdn1.yoursite.com","cdn2.yoursite.com"]
}
`
	testPurgeResourceCacheRawRequest    = `{"paths":["/path/*.css", "/path/*.js"]}`
	testPrefetchResourceCacheRawRequest = `{"paths":["/path/*.css", "/path/*.js"]}`
)

var (
	testGetResourceExpected = &Resource{
		ID:                 4478,
		Name:               nil,
		Cname:              "gcdn.example.me",
		Client:             2096,
		CompanyName:        "Whatever inc",
		Deleted:            false,
		Enabled:            true,
		OriginGroup:        7260,
		OriginProtocol:     "HTTPS",
		Active:             true,
		SecondaryHostnames: []string{},
		Options: &Options{
			Slice:             nil,
			GZIPOn:            nil,
			IgnoreQueryString: nil,
			HostHeader: &HostHeader{
				Enabled: true,
				Value:   "131231.example.ru",
			},
			StaticHeaders:      nil,
			AllowedHTTPMethods: nil,
			Stale: &Stale{
				Enabled: true,
				Value:   []string{"error", "updating"},
			},
			CORS:                 nil,
			ProxyCacheMethodsSet: nil,
			Rewrite:              nil,
			ForceReturn:          nil,
			SecureKey:            nil,
			CacheExpire: &CacheExpire{
				Enabled: true,
				Value:   345600,
			},
			DisableCache:       nil,
			IgnoreCookie:       nil,
			CacheHTTPHeaders:   nil,
			OverrideBrowserTTL: nil,
			FetchCompressed:    nil,
			CountryACL:         nil,
			ReferrerACL:        nil,
			UserAgentACL:       nil,
			IPAddressACL:       nil,
		},
		Status:     "active",
		Rules:      []Rule{},
		SslEnabled: true,
		SslData:    IntPtr(1189),
		CreatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		UpdatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
	}

	testListResourcesExpected = []*Resource{{
		ID:                 4478,
		Name:               nil,
		Cname:              "gcdn.example.me",
		Client:             2096,
		CompanyName:        "Whatever inc",
		Deleted:            false,
		Enabled:            true,
		Active:             true,
		OriginGroup:        7260,
		OriginProtocol:     "HTTPS",
		SecondaryHostnames: []string{},
		Options: &Options{
			Slice:             nil,
			GZIPOn:            nil,
			IgnoreQueryString: nil,
			HostHeader: &HostHeader{
				Enabled: true,
				Value:   "131231.example.ru",
			},
			StaticHeaders:      nil,
			AllowedHTTPMethods: nil,
			Stale: &Stale{
				Enabled: true,
				Value:   []string{"error", "updating"},
			},
			CORS:                 nil,
			ProxyCacheMethodsSet: nil,
			Rewrite:              nil,
			ForceReturn:          nil,
			SecureKey:            nil,
			CacheExpire: &CacheExpire{
				Enabled: true,
				Value:   345600,
			},
			DisableCache:       nil,
			IgnoreCookie:       nil,
			CacheHTTPHeaders:   nil,
			OverrideBrowserTTL: nil,
			FetchCompressed:    nil,
			CountryACL:         nil,
			ReferrerACL:        nil,
			UserAgentACL:       nil,
			IPAddressACL:       nil,
		},
		Status:     "active",
		Rules:      []Rule{},
		SslEnabled: true,
		SslData:    IntPtr(1189),
		CreatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		UpdatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
	}}

	testCreateResourceExpected = &Resource{
		ID:                 4478,
		Name:               nil,
		Cname:              "gcdn.example.me",
		Client:             2096,
		CompanyName:        "Whatever inc",
		Deleted:            false,
		Enabled:            true,
		Active:             true,
		OriginGroup:        7260,
		OriginProtocol:     "HTTPS",
		SecondaryHostnames: []string{},
		Options: &Options{
			Slice:             nil,
			GZIPOn:            nil,
			IgnoreQueryString: nil,
			HostHeader: &HostHeader{
				Enabled: true,
				Value:   "131231.example.ru",
			},
			StaticHeaders:      nil,
			AllowedHTTPMethods: nil,
			Stale: &Stale{
				Enabled: true,
				Value:   []string{"error", "updating"},
			},
			CORS:                 nil,
			ProxyCacheMethodsSet: nil,
			Rewrite:              nil,
			ForceReturn:          nil,
			SecureKey:            nil,
			CacheExpire: &CacheExpire{
				Enabled: true,
				Value:   345600,
			},
			DisableCache:       nil,
			IgnoreCookie:       nil,
			CacheHTTPHeaders:   nil,
			OverrideBrowserTTL: nil,
			FetchCompressed:    nil,
			CountryACL:         nil,
			ReferrerACL:        nil,
			UserAgentACL:       nil,
			IPAddressACL:       nil,
		},
		Status:     "active",
		Rules:      []Rule{},
		SslEnabled: true,
		SslData:    IntPtr(1189),
		CreatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		UpdatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
	}

	testUpdateResourceExpected = &Resource{
		ID:                 4478,
		Name:               nil,
		Cname:              "gcdn.example.me",
		Client:             2096,
		CompanyName:        "Whatever inc",
		Deleted:            false,
		Enabled:            true,
		OriginGroup:        7260,
		OriginProtocol:     "HTTPS",
		Active:             true,
		SecondaryHostnames: []string{},
		Options: &Options{
			Slice:             nil,
			GZIPOn:            nil,
			IgnoreQueryString: nil,
			HostHeader: &HostHeader{
				Enabled: true,
				Value:   "131231.example.ru",
			},
			StaticHeaders:      nil,
			AllowedHTTPMethods: nil,
			Stale: &Stale{
				Enabled: true,
				Value:   []string{"error", "updating"},
			},
			CORS:                 nil,
			ProxyCacheMethodsSet: nil,
			Rewrite:              nil,
			ForceReturn:          nil,
			SecureKey:            nil,
			CacheExpire: &CacheExpire{
				Enabled: true,
				Value:   345600,
			},
			DisableCache:       nil,
			IgnoreCookie:       nil,
			CacheHTTPHeaders:   nil,
			OverrideBrowserTTL: nil,
			FetchCompressed:    nil,
			CountryACL:         nil,
			ReferrerACL:        nil,
			UserAgentACL:       nil,
			IPAddressACL:       nil,
		},
		Status:     "active",
		Rules:      []Rule{},
		SslEnabled: true,
		SslData:    IntPtr(1189),
		CreatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		UpdatedAt:  NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
	}
)

func TestResourcesService_Create(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         resourcesURL,
		RawResponse: testCreateResourceRawResponse,
		RawRequest:  testCreateResourceRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testCreateResourceExpected

	body := &CreateResourceBody{
		Cname: "cdn.site.com",
		SecondaryHostnames: []string{
			"cdn1.yoursite.com",
			"cdn2.yoursite.com",
		},
	}

	got, _, err := client.Resources.Create(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't create a resource")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestResourcesService_Get(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(resourceURL, testGetResourceExpected.ID),
		RawResponse: testGetResourceRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testGetResourceExpected

	got, _, err := client.Resources.Get(context.Background(), testGetResourceExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a resource")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestResourcesService_List(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         resourcesURL,
		RawResponse: testListResourcesRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testListResourcesExpected

	got, _, err := client.Resources.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a list of resources")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestResourcesService_Prefetch(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(resourcePrefetchURL, testGetResourceExpected.ID),
		RawResponse: "",
		RawRequest:  testPrefetchResourceCacheRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	body := []string{"/path/*.css", "/path/*.js"}

	_, err := client.Resources.Prefetch(context.Background(), testGetResourceExpected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't prefetch cache")
	}
}

func TestResourcesService_Purge(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(resourcePurgeURL, testGetResourceExpected.ID),
		RawResponse: "",
		RawRequest:  testPurgeResourceCacheRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	body := []string{"/path/*.css", "/path/*.js"}

	_, err := client.Resources.Purge(context.Background(), testGetResourceExpected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't purge cache")
	}
}

func TestResourcesService_Update(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(resourceURL, testUpdateResourceExpected.ID),
		RawResponse: testUpdateResourceRawResponse,
		RawRequest:  testUpdateResourceRawRequest,
		Method:      http.MethodPut,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testUpdateResourceExpected

	body := &UpdateResourceBody{
		OriginGroup: testUpdateResourceExpected.OriginGroup,
		SecondaryHostnames: []string{
			"cdn1.yoursite.com",
			"cdn2.yoursite.com",
		},
	}

	got, _, err := client.Resources.Update(context.Background(), testUpdateResourceExpected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't update a resource")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
