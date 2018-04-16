package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	resellUsersURL          = "/users"
	resellClientsURL        = "/clients"
	resellClientURL         = "/clients/%d"
	resellUserTokenURL      = "/users/%d/token"
	resellClientServicesURL = "/clients/%d/services"
	resellClientServiceURL  = "/clients/%d/services/%d"
)

type ClientsService service

type ClientAccount struct {
	ID               int        `json:"id"`
	Users            []User     `json:"users"`
	CurrentUser      int        `json:"currentUser"`
	Email            string     `json:"email"`
	Phone            string     `json:"phone"`
	Name             string     `json:"name"`
	Status           string     `json:"status"`
	Created          *GCoreTime `json:"created"`
	Updated          *GCoreTime `json:"updated"`
	CompanyName      string     `json:"companyName"`
	UtilizationLevel int        `json:"utilization_level"`
	Reseller         int        `json:"reseller"`
}

type CreateClientBody struct {
	UserType string `json:"user_type"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateClientBody struct {
	Name        string `json:"name"`
	CompanyName string `json:"companyName"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Seller      int    `json:"seller,omitempty"`
}

type ListOpts struct {
	Email       string `url:"email,omitempty"`
	Name        string `url:"name,omitempty"`
	CompanyName string `url:"companyName,omitempty"`
	Deleted     bool   `url:"deleted,omitempty"`
	CDN         string `url:"cdn,omitempty"`
	Activated   bool   `url:"activated,omitempty"`
}

func (s *ClientsService) Create(ctx context.Context, body CreateClientBody) (*ClientAccount, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "POST", resellUsersURL, body)
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

func (s *ClientsService) Get(ctx context.Context, clientID int) (*ClientAccount, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", fmt.Sprintf(resellClientURL, clientID), nil)
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

func (s *ClientsService) List(ctx context.Context, opts ListOpts) (*[]ClientAccount, *http.Response, error) {
	url, err := addOptions(resellClientsURL, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	clients := make([]ClientAccount, 0)

	resp, err := s.client.Do(req, &clients)
	if err != nil {
		return nil, resp, err
	}

	return &clients, resp, nil
}

func (s *ClientsService) Update(ctx context.Context, clientID int, body UpdateClientBody) (*ClientAccount, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", fmt.Sprintf(resellClientURL, clientID), body)
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

func (s *ClientsService) GetCommonClient(ctx context.Context, userID int) (*CommonClient, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", fmt.Sprintf(resellUserTokenURL, userID), nil)
	if err != nil {
		return nil, nil, err
	}

	token := &Token{}

	resp, err := s.client.Do(req, token)
	if err != nil {
		return nil, resp, err
	}

	commonClient := NewCommonClient(s.client.client)
	commonClient.Token = token

	return commonClient, resp, nil
}

type PaidService struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// This feature has been taken from the admin web-panel, is not documented at all
// It allows to pause CDN service for specific client
func (s *ClientsService) SuspendCDN(ctx context.Context, clientID int) (*http.Response, error) {
	url, _ := addOptions(fmt.Sprintf(resellClientServicesURL, clientID), struct {
		Name string `url:"name"`
	}{"CDN"})

	req, err := s.client.NewRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	paidServices := make([]PaidService, 0)

	resp, err := s.client.Do(req, &paidServices)
	if err != nil {
		return resp, err
	}

	// The only one CDN service is supposed to be
	req, err = s.client.NewRequest(ctx, "PUT",
		fmt.Sprintf(resellClientServiceURL, clientID, paidServices[0].ID),
		struct {
			Enabled bool   `json:"enabled"`
			Status  string `json:"status"`
		}{false, "paused"})
	if err != nil {
		return nil, err
	}

	resp, err = s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// This feature has been taken from the admin web-panel, is not documented at all
// It allows to resume CDN service for specific client
func (s *ClientsService) ResumeCDN(ctx context.Context, clientID int) (*http.Response, error) {
	url, _ := addOptions(fmt.Sprintf(resellClientServicesURL, clientID), struct {
		Name string `url:"name"`
	}{"CDN"})

	req, err := s.client.NewRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	paidServices := make([]PaidService, 0)

	resp, err := s.client.Do(req, &paidServices)
	if err != nil {
		return resp, err
	}

	// The only one CDN service is supposed to be
	req, err = s.client.NewRequest(ctx, "PUT",
		fmt.Sprintf(resellClientServiceURL, clientID, paidServices[0].ID),
		struct {
			Enabled bool   `json:"enabled"`
			Status  string `json:"status"`
		}{true, "active"})
	if err != nil {
		return nil, err
	}

	resp, err = s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
