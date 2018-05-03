package rules

import "github.com/dstdfx/go-gcore/gcore"

var (
	TestCreateRuleResponse = `{
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
	TestGetRuleResponse = `{
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
	TestListRuleResponse = `[
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
	FakeResourceID         = 4538
	TestCreateRuleExpected = &gcore.Rule{
		ID:             1861,
		OriginProtocol: "HTTP",
		RuleType:       0,
		Weight:         2,
		Name:           "whatever",
		Rule:           "/images",
		PresetApplied:  false,
		Options: gcore.Options{
			CacheHTTPHeaders: &gcore.CacheHTTPHeaders{
				Enabled: true,
				Value:   []string{"content-length", "x-token", "connection", "date", "server", "content-type"},
			},
		},
	}
	TestGetRuleExpected = &gcore.Rule{
		ID:             1861,
		OriginProtocol: "HTTP",
		RuleType:       0,
		Weight:         2,
		Name:           "whatever",
		Rule:           "/images",
		PresetApplied:  false,
		Options: gcore.Options{
			CacheHTTPHeaders: &gcore.CacheHTTPHeaders{
				Enabled: true,
				Value:   []string{"content-length", "x-token", "connection", "date", "server", "content-type"},
			},
		},
	}
	TestListRuleExpected = []*gcore.Rule{
		{
			ID:             1861,
			OriginProtocol: "HTTP",
			RuleType:       0,
			Weight:         2,
			Name:           "whatever",
			Rule:           "/images",
			PresetApplied:  false,
			Options: gcore.Options{
				CacheHTTPHeaders: &gcore.CacheHTTPHeaders{
					Enabled: true,
					Value:   []string{"content-length", "x-token", "connection", "date", "server", "content-type"},
				},
			},
		},
	}
)
