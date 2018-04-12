package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	resellUsersURL     = "/users"
	resellClientURL    = "/clients/%d"
	resellUserTokenURL = "/users/%d/token"
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
	Manager          string     `json:"manager"`
	CompanyOwner     string     `json:"company_owner"`
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
