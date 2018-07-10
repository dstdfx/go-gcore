package resources

import (
	"time"

	"github.com/dstdfx/go-gcore/gcore"
)

var (
	TestGetResourceResponse = `{
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

	TestCreateResourceResponse = `{
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

	TestListResourcesResponse = `[{
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
	TestGetResourceExpected = &gcore.Resource{
		ID:                 4478,
		Name:               nil,
		CName:              "gcdn.example.me",
		Client:             2096,
		CompanyName:        "Whatever inc",
		Deleted:            false,
		Enabled:            true,
		OriginGroup:        7260,
		OriginProtocol:     "HTTPS",
		Active:             true,
		SecondaryHostnames: []string{},
		Options: &gcore.Options{
			Slice:             nil,
			GZIPOn:            nil,
			IgnoreQueryString: nil,
			HostHeader: &gcore.HostHeader{
				Enabled: true,
				Value:   "131231.example.ru",
			},
			StaticHeaders:      nil,
			AllowedHTTPMethods: nil,
			Stale: &gcore.Stale{
				Enabled: true,
				Value:   []string{"error", "updating"},
			},
			CORS:                 nil,
			ProxyCacheMethodsSet: nil,
			Rewrite:              nil,
			ForceReturn:          nil,
			SecureKey:            nil,
			CacheExpire: &gcore.CacheExpire{
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
		Rules:      []gcore.Rule{},
		SSLEnabled: true,
		SSLData:    gcore.IntPtr(1189),
		CreatedAt:  gcore.NewGCoreTime(time.Date(2018, 4, 9, 11, 31, 40, 0, time.UTC)),
		UpdatedAt:  gcore.NewGCoreTime(time.Date(2018, 4, 9, 11, 32, 31, 0, time.UTC)),
	}

	TestListResourcesExpected = []*gcore.Resource{{
		ID:                 4478,
		Name:               nil,
		CName:              "gcdn.example.me",
		Client:             2096,
		CompanyName:        "Whatever inc",
		Deleted:            false,
		Enabled:            true,
		Active:             true,
		OriginGroup:        7260,
		OriginProtocol:     "HTTPS",
		SecondaryHostnames: []string{},
		Options: &gcore.Options{
			Slice:             nil,
			GZIPOn:            nil,
			IgnoreQueryString: nil,
			HostHeader: &gcore.HostHeader{
				Enabled: true,
				Value:   "131231.example.ru",
			},
			StaticHeaders:      nil,
			AllowedHTTPMethods: nil,
			Stale: &gcore.Stale{
				Enabled: true,
				Value:   []string{"error", "updating"},
			},
			CORS:                 nil,
			ProxyCacheMethodsSet: nil,
			Rewrite:              nil,
			ForceReturn:          nil,
			SecureKey:            nil,
			CacheExpire: &gcore.CacheExpire{
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
		Rules:      []gcore.Rule{},
		SSLEnabled: true,
		SSLData:    gcore.IntPtr(1189),
		CreatedAt:  gcore.NewGCoreTime(time.Date(2018, 4, 9, 11, 31, 40, 0, time.UTC)),
		UpdatedAt:  gcore.NewGCoreTime(time.Date(2018, 4, 9, 11, 32, 31, 0, time.UTC)),
	}}

	TestCreateResourceExpected = &gcore.Resource{
		ID:                 4478,
		Name:               nil,
		CName:              "gcdn.example.me",
		Client:             2096,
		CompanyName:        "Whatever inc",
		Deleted:            false,
		Enabled:            true,
		Active:             true,
		OriginGroup:        7260,
		OriginProtocol:     "HTTPS",
		SecondaryHostnames: []string{},
		Options: &gcore.Options{
			Slice:             nil,
			GZIPOn:            nil,
			IgnoreQueryString: nil,
			HostHeader: &gcore.HostHeader{
				Enabled: true,
				Value:   "131231.example.ru",
			},
			StaticHeaders:      nil,
			AllowedHTTPMethods: nil,
			Stale: &gcore.Stale{
				Enabled: true,
				Value:   []string{"error", "updating"},
			},
			CORS:                 nil,
			ProxyCacheMethodsSet: nil,
			Rewrite:              nil,
			ForceReturn:          nil,
			SecureKey:            nil,
			CacheExpire: &gcore.CacheExpire{
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
		Rules:      []gcore.Rule{},
		SSLEnabled: true,
		SSLData:    gcore.IntPtr(1189),
		CreatedAt:  gcore.NewGCoreTime(time.Date(2018, 4, 9, 11, 31, 40, 0, time.UTC)),
		UpdatedAt:  gcore.NewGCoreTime(time.Date(2018, 4, 9, 11, 32, 31, 0, time.UTC)),
	}
)
