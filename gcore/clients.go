package gcore

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

const (
	resellUsersURL     = "/users"
	resellClientsURL   = "/clients"
	resellClientURL    = "/clients/%d"
	resellUserTokenURL = "/users/%d/token"
)

// ClientsService handles communication with the client related methods
// of the G-Core CDN API.
type ClientsService service

// ClientAccount represents G-Core's client account.
type ClientAccount struct {
	ID               int     `json:"id"`
	Client           int     `json:"client"`
	Users            []*User `json:"users"`
	CurrentUser      int     `json:"currentUser"`
	Email            string  `json:"email"`
	Phone            string  `json:"phone"`
	Name             string  `json:"name"`
	Status           string  `json:"status"`
	Created          *Time   `json:"created"`
	Updated          *Time   `json:"updated"`
	CompanyName      string  `json:"companyName"`
	UtilizationLevel int     `json:"utilization_level"`
	Reseller         int     `json:"reseller"`
	Cname            string  `json:"cname,omitempty"`
}

// CreateClientBody represents request body for create client.
type CreateClientBody struct {
	UserType string `json:"user_type"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UpdateClientBody represents request body for update client.
type UpdateClientBody struct {
	Name        string `json:"name"`
	CompanyName string `json:"companyName"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Seller      int    `json:"seller,omitempty"`
}

// ListOpts represents list of additional options to filter client's list by.
type ListOpts struct {
	Email       string `param:"email,omitempty"`
	Name        string `param:"name,omitempty"`
	CompanyName string `param:"companyName,omitempty"`
	Deleted     bool   `param:"deleted,omitempty"`
	CDN         string `param:"cdn,omitempty"`
	Activated   bool   `param:"activated,omitempty"`
}

// Create method creates a new client, the client will be activated automatically.
func (s *ClientsService) Create(ctx context.Context, body *CreateClientBody) (*ClientAccount, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, resellUsersURL, body)
	if err != nil {
		return nil, nil, err
	}

	clientAccount := &ClientAccount{}

	resp, err := s.client.Do(req, clientAccount)
	if err != nil {
		return nil, resp, err
	}

	return clientAccount, resp, nil
}

// Get method returns data of a client by ID.
func (s *ClientsService) Get(ctx context.Context, clientID int) (*ClientAccount, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(resellClientURL, clientID), nil)
	if err != nil {
		return nil, nil, err
	}

	clientAccount := &ClientAccount{}

	resp, err := s.client.Do(req, clientAccount)
	if err != nil {
		return nil, resp, err
	}

	return clientAccount, resp, nil
}

// List method gets a list of all Clients assigned to a Reseller.
func (s *ClientsService) List(ctx context.Context, opts ListOpts) ([]*ClientAccount, *http.Response, error) {

	url := resellClientsURL
	queryParams, err := BuildQueryParameters(opts)
	if err != nil {
		return nil, nil, err
	}

	if queryParams != "" {
		url = strings.Join([]string{url, queryParams}, "?")
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	clients := make([]*ClientAccount, 0)

	resp, err := s.client.Do(req, &clients)
	if err != nil {
		return nil, resp, err
	}

	return clients, resp, nil
}

// Update method edits data of the client.
func (s *ClientsService) Update(ctx context.Context, clientID int, body *UpdateClientBody) (*ClientAccount, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodPut,
		fmt.Sprintf(resellClientURL, clientID), body)
	if err != nil {
		return nil, nil, err
	}

	client := &ClientAccount{}

	resp, err := s.client.Do(req, client)
	if err != nil {
		return nil, resp, err
	}

	return client, resp, nil
}

// GetCommonClient method returns CommonClient for the given userID.
// This feature has been taken from the admin web-panel, is not documented at all
// It allows to authenticate as a user (common client), common client can manage
// his own CDN resources, origins and etc.
func (s *ClientsService) GetCommonClient(ctx context.Context, userID int) (*CommonClient, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(resellUserTokenURL, userID), nil)
	if err != nil {
		return nil, nil, err
	}

	token := &Token{}

	resp, err := s.client.Do(req, token)
	if err != nil {
		return nil, resp, err
	}

	commonClient := NewCommonClient(s.client.log)
	commonClient.Token = token

	return commonClient, resp, nil
}
