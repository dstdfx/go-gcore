package gcore

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	LibraryVersion = "2.0.0"
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
	Value  string `json:"token"`
	Expire *Time  `json:"expire"`
}

// Authenticate gets API Token, if client already took a token, check if it's valid.
// If it's not, get new one.
func (c *Client) Authenticate(ctx context.Context, authOpts AuthOptions) error {
	req, err := c.NewRequest(ctx, http.MethodPost, LoginURL, authOpts)
	if err != nil {
		return err
	}

	c.Lock()
	defer c.Unlock()
	if c.Token == nil || c.Token.Expire.Before(time.Now().UTC()) {
		// Renew token if expired
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

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: UserAgent,
		log:       SelectLogger(logger...),
	}
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

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: UserAgent,
		log:       SelectLogger(logger...),
	}
	c.common.client = c

	resellerServices := ResellerServices{}
	resellerServices.Clients = (*ClientsService)(&c.common)
	resellClient := &ResellerClient{Client: c, ResellerServices: resellerServices}

	return resellClient
}

// NewRequest method returns new request by given options.
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

// Do method executes request and checks response body.
func (c *Client) Do(req *http.Request, to interface{}) (*http.Response, error) {
	c.log.Debugf("REQ  %v %v", req.Method, req.URL)

	resp, err := c.client.Do(req)
	if err != nil {
		c.log.Errorf("Request failed with error: %s", err)
		return nil, err
	}

	c.log.Debugf("RESP   %v %v %v", req.Method, req.URL, resp.StatusCode)

	if resp.StatusCode >= http.StatusBadRequest &&
		resp.StatusCode <= http.StatusNetworkAuthenticationRequired {

		var respErr error

		if resp.Body != nil {

			body, _ := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			// To able to read response twice
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(body))
			resp.Body = rdr2

			c.log.Debugf("RESP BODY  %s", string(body))
			respErr = fmt.Errorf("gcore: got the %d error status code from the server with body: %s",
				resp.StatusCode, string(body))
		} else {
			respErr = fmt.Errorf("gcore: got the %d error status code from the server", resp.StatusCode)
		}

		return resp, respErr
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

// IntPtr returns pointer to int.
func IntPtr(v int) *int {
	return &v
}

// StringPtr returns pointer to string.
func StringPtr(v string) *string {
	return &v
}

// BuildQueryParameters converts provided options struct to the string of URL parameters.
func BuildQueryParameters(opts interface{}) (string, error) {
	optsValue := reflect.ValueOf(opts)
	if optsValue.Kind() != reflect.Struct {
		return "", errors.New("provided options is not a structure")
	}
	optsType := reflect.TypeOf(opts)

	params := url.Values{}

	for i := 0; i < optsValue.NumField(); i++ {
		fieldValue := optsValue.Field(i)
		fieldType := optsType.Field(i)

		queryTag := fieldType.Tag.Get("param")
		if queryTag != "" {
			if isZero(fieldValue) {
				continue
			}

			tags := strings.Split(queryTag, ",")
		loop:
			switch fieldValue.Kind() {
			case reflect.Ptr:
				fieldValue = fieldValue.Elem()
				goto loop
			case reflect.String:
				params.Add(tags[0], fieldValue.String())
			case reflect.Int:
				params.Add(tags[0], strconv.FormatInt(fieldValue.Int(), 10))
			case reflect.Bool:
				params.Add(tags[0], strconv.FormatBool(fieldValue.Bool()))
			}
		}
	}

	return params.Encode(), nil
}

// isZero checks if provided value is zero.
func isZero(v reflect.Value) bool {
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	z := reflect.Zero(v.Type())

	return v.Interface() == z.Interface()
}
