package gcore

type RulesService service

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

type CacheHTTPHeaders struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type CacheExpire struct {
	Enabled bool `json:"enabled"`
	Value   int  `json:"value"`
}

type AllowedHTTPMethods struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type CORS struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type CountryACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}

type DisableCache struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type FetchCompressed struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type ForceReturn struct {
	Enabled bool   `json:"enabled"`
	Code    int    `json:"code"`
	Body    string `json:"body"`
}

type GZIPOn struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type HostHeader struct {
	Enabled bool   `json:"enabled"`
	Value   string `json:"value"`
}

type IgnoreQueryString struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type IgnoreCookie struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type IPAddressACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}

type OverrideBrowserTTL struct {
	Enabled bool `json:"enabled"`
	Value   int  `json:"value"`
}

type ProxyCacheMethodsSet struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type ReferrerACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}

type Rewrite struct {
	Enabled bool   `json:"enabled"`
	Body    string `json:"body"`
	Flag    string `json:"flag"`
}

type SecureKey struct {
	Enabled bool   `json:"enabled"`
	Key     string `json:"body"`
	Type    int    `json:"type"`
}

type Slice struct {
	Enabled bool `json:"enabled"`
	Value   bool `json:"value"`
}

type Stale struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type StaticHeaders struct {
	Enabled bool     `json:"enabled"`
	Value   []string `json:"value"`
}

type UserAgentACL struct {
	Enabled        bool     `json:"enabled"`
	ExceptedValues []string `json:"excepted_values"`
	PolicyType     string   `json:"policy_type"`
}
