package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	ResourcesURL     = "/resources"
	ResourceURL      = "/resources/%d"
	ResourcePurgeURL = "/resources/%d/purge"
)

type ResourcesService service

// Resource represents G-Core's CDN Resource.
type Resource struct {
	ID                 int        `json:"id"`
	Name               *string    `json:"name"`
	Deleted            bool       `json:"deleted"`
	Active             bool       `json:"active"`
	Enabled            bool       `json:"enabled"`
	CompanyName        string     `json:"companyName"`
	Status             string     `json:"status"`
	Client             int        `json:"client"`
	OriginGroup        int        `json:"originGroup"`
	Cname              string     `json:"cname"`
	SecondaryHostnames []string   `json:"secondaryHostnames"`
	Options            *Options   `json:"options"`
	OriginProtocol     string     `json:"originProtocol"`
	Rules              []Rule     `json:"rules"`
	CreatedAt          *GCoreTime `json:"created"`
	UpdatedAt          *GCoreTime `json:"updated"`
	SslData            *int       `json:"sslData"`
	SslEnabled         bool       `json:"sslEnabled"`
}

type CreateResourceBody struct {
	Cname              string   `json:"cname"`
	Origin             string   `json:"origin,omitempty"`
	OriginGroupId      *int     `json:"originGroup,omitempty"`
	SecondaryHostnames []string `json:"secondaryHostnames,omitempty"`
	OriginProtocol     string   `json:"originProtocol,omitempty"`
	SslData            *int     `json:"sslData,omitempty"`
	SslEnabled         bool     `json:"sslEnabled,omitempty"`
	Options            *Options `json:"options,omitempty"`
}

type UpdateResourceBody struct {
	Active             *bool    `json:"active,omitempty"`
	Enabled            *bool    `json:"enabled,omitempty"`
	OriginGroup        int      `json:"originGroup,omitempty"`
	SecondaryHostnames []string `json:"secondaryHostnames,omitempty"`
	OriginProtocol     string   `json:"originProtocol,omitempty"`
	SslData            *int     `json:"sslData,omitempty"`
	SslEnabled         *bool    `json:"sslEnabled,omitempty"`
	Options            *Options `json:"options,omitempty"`
}

func (s *ResourcesService) Update(ctx context.Context, resourceID int, body UpdateResourceBody) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodPut,
		fmt.Sprintf(ResourceURL, resourceID), body)
	if err != nil {
		return nil, nil, err
	}

	resource := &Resource{}

	resp, err := s.client.Do(req, resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, nil
}

// Get information about all CDN Resources for this account.
func (s *ResourcesService) List(ctx context.Context) ([]*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, ResourcesURL, nil)
	if err != nil {
		return nil, nil, err
	}

	resources := make([]*Resource, 0)

	resp, err := s.client.Do(req, &resources)
	if err != nil {
		return nil, resp, err
	}

	return resources, resp, nil
}

// Get information about specific CDN Resource.
func (s *ResourcesService) Get(ctx context.Context, resourceID int) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(ResourceURL, resourceID), nil)
	if err != nil {
		return nil, nil, err
	}

	resource := &Resource{}

	resp, err := s.client.Do(req, resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, nil
}

// Create CDN Resource.
func (s *ResourcesService) Create(ctx context.Context, body CreateResourceBody) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, ResourcesURL, body)
	if err != nil {
		return nil, nil, err
	}

	resource := &Resource{}

	resp, err := s.client.Do(req, resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, nil
}

// Purge deletes cache from CDN servers. It is necessary for updating CDN content.
func (s *ResourcesService) Purge(ctx context.Context, resourceID int, paths []string) (*http.Response, error) {
	var pathsBody struct {
		Paths []string `json:"paths"`
	}
	pathsBody.Paths = paths

	req, err := s.client.NewRequest(ctx,
		http.MethodPost,
		fmt.Sprintf(ResourcePurgeURL, resourceID), pathsBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
