package gcore

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	LibraryVersion = "0.0.1"
	DefaultBaseURL = "https://api.gcdn.co"
	UserAgent      = "go-gcore/" + LibraryVersion

	LoginURL = "/auth/signin"
)

// Client manages communication with G-Core CDN API.
type Client struct {
	sync.Mutex

	// HTTP client used to communicate with the GC API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client.
	UserAgent string

	log GenericLogger

	common service

	// Token to communicate with G-Core API.
	Token *Token
}

type service struct {
	client *Client
}

// CommonClient represents API of basic G-Core account.
type CommonClient struct {
	*Client
	CommonServices
}

// ResellerClient represents API of reseller G-Core account.
type ResellerClient struct {
	*Client
	ResellerServices
}

// CommonServices represent specific account type features.
type CommonServices struct {
	Account      *AccountService
	Resources    *ResourcesService
	OriginGroups *OriginGroupsService
	Rules        *RulesService
	Certificates *CertService
}

// ResellerServices represent specific account type features.
type ResellerServices struct {
	Clients *ClientsService
}

// G-Core account credentials.
type AuthOptions struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token to access G-Core API.
type Token struct {
	Value  string     `json:"token"`
	Expire *GCoreTime `json:"expire"`
}

// Authenticate gets API Token, if client already took a token, check if it's valid.
// If it's not, get new one.
func (c *Client) Authenticate(ctx context.Context, authOpts AuthOptions) error {
	req, err := c.NewRequest(ctx, "POST", LoginURL, authOpts)
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()
	if c.Token == nil || c.Token.Expire.Before(time.Now().UTC()) {
		token := &Token{}
		_, err = c.Do(req, token)
		if err != nil {
			return err
		}

		c.Token = token
	}

	return nil
}

// NewCommonClient creates basic G-Core client.
func NewCommonClient(httpClient *http.Client, logger ...GenericLogger) *CommonClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(DefaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: UserAgent, log: SelectLogger(logger...)}
	c.common.client = c

	commonServices := CommonServices{}
	commonServices.Account = (*AccountService)(&c.common)
	commonServices.Resources = (*ResourcesService)(&c.common)
	commonServices.OriginGroups = (*OriginGroupsService)(&c.common)
	commonServices.Rules = (*RulesService)(&c.common)
	commonServices.Certificates = (*CertService)(&c.common)

	commonClient := &CommonClient{Client: c, CommonServices: commonServices}

	return commonClient
}

// NewResellerClient creates reseller G-Core client.
func NewResellerClient(httpClient *http.Client, logger ...GenericLogger) *ResellerClient {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(DefaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, UserAgent: UserAgent, log: SelectLogger(logger...)}
	c.common.client = c

	resellerServices := ResellerServices{}
	resellerServices.Clients = (*ClientsService)(&c.common)
	resellClient := &ResellerClient{Client: c, ResellerServices: resellerServices}

	return resellClient
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
		b, _ := json.Marshal(body)
		c.log.Debugf("REQ BODY %s", string(b))
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	if c.Token != nil {
		req.Header.Add("Authorization", "Token "+c.Token.Value)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

func (c *Client) Do(req *http.Request, to interface{}) (*http.Response, error) {
	c.log.Debugf("REQ  %v %v", req.Method, req.URL)

	resp, err := c.client.Do(req)
	if err != nil {
		c.log.Errorf("Request failed with error: %s", err)
		return nil, err
	}

	c.log.Debugf("RESP   %v %v %v", req.Method, req.URL, resp.StatusCode)

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		gcoreErr := &GCoreError{Code: resp.StatusCode}

		if resp.Body != nil {
			body, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			err = json.Unmarshal(body, gcoreErr)
			if err != nil {
				err = fmt.Errorf("gcore: got the %d error status code from the server", resp.StatusCode)
				return resp, err
			}

			err = gcoreErr
		} else {
			err = fmt.Errorf("gcore: got the %d error status code from the server", resp.StatusCode)
		}

		return resp, err
	}

	if to != nil {
		if err = ExtractResult(resp, to); err != nil {
			return resp, err
		}
	}

	return resp, nil
}

// ExtractResult reads response body and unmarshal it to given interface
func ExtractResult(resp *http.Response, to interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, to)
	return err
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func IntPtr(v int) *int {
	return &v
}

func StringPtr(v string) *string {
	return &v
}
