package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	testResourceID          = 42
	testGetResourceResponse = `{
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

	testCreateResourceResponse = `{
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

	testListResourcesResponse = `[{
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
)

func TestResourcesService_Get(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testGetResourceExpected
	mux.HandleFunc(fmt.Sprintf(resourceURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testGetResourceResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := getAuthenticatedCommonClient()
	got, _, err := client.Resources.Get(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestResourcesService_List(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(resourcesURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testListResourcesResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := getAuthenticatedCommonClient()
	expected := testListResourcesExpected
	got, _, err := client.Resources.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestResourcesService_Create(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(resourcesURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testCreateResourceResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	resourceBody := CreateResourceBody{
		Cname: "cdn.site.com",
		SecondaryHostnames: []string{
			"cdn1.yoursite.com",
			"cdn2.yoursite.com",
		},
	}

	client := getAuthenticatedCommonClient()
	expected := testCreateResourceExpected
	got, _, err := client.Resources.Create(context.Background(), &resourceBody)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}

}

func TestResourceService_Purge(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(fmt.Sprintf(resourcePurgeURL, testResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		})

	client := getAuthenticatedCommonClient()
	resp, err := client.Resources.Purge(context.Background(), testResourceID, []string{})
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d",
			http.StatusCreated,
			resp.StatusCode,
		)
	}
}

func TestResourceService_Prefetch(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(fmt.Sprintf(resourcePrefetchURL, testResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		})

	client := getAuthenticatedCommonClient()
	resp, err := client.Resources.Prefetch(context.Background(), testResourceID,
		[]string{"/file.jpg", "file2.jpg"})
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d",
			http.StatusCreated,
			resp.StatusCode,
		)
	}

}
