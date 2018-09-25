package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	RulesURL = "/resources/%d/rules"
	RuleURL  = "/resources/%d/rules/%d"
)

type RulesService service

// Rule represent G-Core's rule for CDN Resource.
type Rule struct {
	ID             int     `json:"id"`
	Rule           string  `json:"rule"`
	Name           string  `json:"name"`
	OriginGroup    int     `json:"originGroup"`
	RuleType       int     `json:"ruleType"`
	Options        Options `json:"options"`
	Weight         int     `json:"weight"`
	PresetApplied  bool    `json:"preset_applied"`
	OriginProtocol string  `json:"originProtocol"`
}

// Options represent possible params for a Rule.
type Options struct {
	CacheHTTPHeaders     *CacheHTTPHeaders     `json:"cache_http_headers"`
	CacheExpire          *CacheExpire          `json:"cache_expire"`
	AllowedHTTPMethods   *AllowedHTTPMethods   `json:"allowedHttpMethods"`
	CORS                 *CORS                 `json:"cors"`
	CountryACL           *CountryACL           `json:"country_acl"`
	DisableCache         *DisableCache         `json:"disable_cache"`
	FetchCompressed      *FetchCompressed      `json:"fetch_compressed"`
	ForceReturn          *ForceReturn          `json:"force_return"`
	GZIPOn               *GZIPOn               `json:"gzipOn"`
	HostHeader           *HostHeader           `json:"hostHeader"`
	IgnoreQueryString    *IgnoreQueryString    `json:"ignoreQueryString"`
	IgnoreCookie         *IgnoreCookie         `json:"ignore_cookie"`
	IPAddressACL         *IPAddressACL         `json:"ip_address_acl"`
	OverrideBrowserTTL   *OverrideBrowserTTL   `json:"override_browser_ttl"`
	ProxyCacheMethodsSet *ProxyCacheMethodsSet `json:"proxy_cache_methods_set"`
	ReferrerACL          *ReferrerACL          `json:"referrer_acl"`
	Rewrite              *Rewrite              `json:"rewrite"`
	SecureKey            *SecureKey            `json:"secure_key"`
	Slice                *Slice                `json:"slice"`
	Stale                *Stale                `json:"stale"`
	StaticHeaders        *StaticHeaders        `json:"staticHeaders"`
	UserAgentACL         *UserAgentACL         `json:"user_agent_acl"`
}

type CreateRuleBody struct {
	Rule     string  `json:"rule"`
	Name     string  `json:"name"`
	RuleType int     `json:"ruleType"`
	Options  Options `json:"options"`
}

// List HTTP Headers that must be included in the response.
type CacheHTTPHeaders struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

// Specifies cache expiration time in seconds.
type CacheExpire struct {
	Enabled bool `json:"enabled"`
	Value   int  `json:"value"`
}

// The list of allowed HTTP methods.
type AllowedHTTPMethods struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

// Option allows you to add Access-Control-Allow-Origin for the specified domains or for all domains.
// It has two parameters:
// For all domains
// "value": ["*"]
// For the specified list of domains
// "value": ["domain.com", "second.dom.com"]
type CORS struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

// Control access to the content for specified countries.
type CountryACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}

// When enabled the content caching is completely disabled.
type DisableCache struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// A CDN request and cache already compressed content.
// Your server should support compression.
// CDN servers won't ungzip your content even if a user's browser doesn't accept compression
// (nowadays almost all browsers support it). By default, option is disabled (enabled: false).
// Not supported with gzipON option enabled. Only one of these options can be used at the same time.
// fetch_compressed overrides gzipOn.
type FetchCompressed struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// Allows to apply custom HTTP code to the CDN content.
// Specify HTTP-code you need and text or URL if you're going to set up redirection.
type ForceReturn struct {
	Enabled bool   `json:"enabled"`
	Code    int    `json:"code"`
	Body    string `json:"body"`
}

// The option allows to compress content with gzip on the CDN`s end.
// CDN servers will request only uncompressed content from the origin.
// Not supported with fetch_compressed.
type GZIPOn struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// Specify the Host header that CDN servers use when request content from an origin server.
// Your server must be able to process requests with the chosen header.
// If the option is in NULL state Host Header value is taken from the parent CDN resource's value.
type HostHeader struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

// This option determines how files with different query strings will be cached:
// either as one object (when this option is enabled) or
// as different objects (when this option is disabled).
type IgnoreQueryString struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// By default, files pulled from an origin source with cookies are not cached in a CDN.
// Enable this option to cache such objects.
type IgnoreCookie struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// Control access to the CDN Resource content for specified IP addresses.
// If you wish to use IPs from our CDN servers IP list for IP ACL configuration,
// you need to independently monitor its relevance.
// We recommend you use a script for automatically update IP ACL.
// Read more https://docs.gcorelabs.com/cdn/#operation--public-ip-list-get.
type IPAddressACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}

// Сache content according to origin Cache-Control header.
// When enabled Origin Source Cache-Control is inherited and respected.
// It overrides the cache_expire option value.
// Specify cache expiry time in seconds for the end user’s browser.
type OverrideBrowserTTL struct {
	Enabled bool `json:"enabled"`
	Value   int  `json:"value"`
}

// Allows caching for GET, HEAD and POST requests.
type ProxyCacheMethodsSet struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// Control access to the CDN Resource content for specified domain names.
type ReferrerACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}

// The pattern for Rewrite. At least one group should be specified.
// For Example: /rewrite_from/(.*) /rewrite_to/$1
// Read more about Rewrite option here http://nginx.org/en/docs/http/ngx_http_rewrite_module.html#rewrite.
type Rewrite struct {
	Enabled bool   `json:"enabled"`
	Body    string `json:"body"`
	Flag    string `json:"flag"`
}

// The option allows configuring an access with tokenized URLs.
// It makes impossible to access content without a valid (unexpired) hash key.
// When enabled you need to specify a key that you use to generate a token.
type SecureKey struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"body"`
	Type    int    `json:"type"`
}

// Files larger than 10 MB will be requested and cached in parts (no larger than 10 MB each part).
// It reduces time to first byte.The origin must support HTTP Range requests.
// By default the option is disabled.
type Slice struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

// The list of errors which the option is applied for.
type Stale struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

// Specify custom HTTP Headers that a CDN server adds to response.
type StaticHeaders struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

// Control access to the content for specified user-agent.
type UserAgentACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}

// Get list of the Rules for given CDN Resource
func (s *RulesService) List(ctx context.Context, resourceID int) ([]*Rule, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(RulesURL, resourceID), nil)
	if err != nil {
		return nil, nil, err
	}

	rules := make([]*Rule, 0)

	resp, err := s.client.Do(req, &rules)
	if err != nil {
		return nil, resp, err
	}

	return rules, resp, nil
}

// Create Rule for specific CDN Resource
func (s *RulesService) Create(ctx context.Context, resourceID int, body CreateRuleBody) (*Rule, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodPost,
		fmt.Sprintf(RulesURL, resourceID), body)
	if err != nil {
		return nil, nil, err
	}

	rule := &Rule{}

	resp, err := s.client.Do(req, rule)
	if err != nil {
		return nil, resp, err
	}

	return rule, resp, nil
}

// Get specific Rule for specific CDN Resource
func (s *RulesService) Get(ctx context.Context, resourceID int, ruleID int) (*Rule, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(RuleURL, resourceID, ruleID), nil)
	if err != nil {
		return nil, nil, err
	}

	rule := &Rule{}

	resp, err := s.client.Do(req, rule)
	if err != nil {
		return nil, resp, err
	}

	return rule, resp, nil
}

// Delete specific Rule for specific CDN Resource
func (s *RulesService) Delete(ctx context.Context, resourceID int, ruleID int) (*http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodDelete,
		fmt.Sprintf(RuleURL, resourceID, ruleID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
