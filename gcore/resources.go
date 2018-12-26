package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	resourcesURL        = "/resources"
	resourceURL         = "/resources/%d"
	resourcePurgeURL    = "/resources/%d/purge"
	resourcePrefetchURL = "/resources/%d/prefetch"
)

// ResourcesService handles communication with the resource related methods
// of the G-Core CDN API.
type ResourcesService service

// Resource represents G-Core's CDN Resource.
type Resource struct {
	ID                 int      `json:"id"`
	Name               *string  `json:"name"`
	Deleted            bool     `json:"deleted"`
	Active             bool     `json:"active"`
	Enabled            bool     `json:"enabled"`
	CompanyName        string   `json:"companyName"`
	Status             string   `json:"status"`
	Client             int      `json:"client"`
	OriginGroup        int      `json:"originGroup"`
	Cname              string   `json:"cname"`
	SecondaryHostnames []string `json:"secondaryHostnames"`
	Options            *Options `json:"options"`
	OriginProtocol     string   `json:"originProtocol"`
	Rules              []Rule   `json:"rules"`
	CreatedAt          *Time    `json:"created"`
	UpdatedAt          *Time    `json:"updated"`
	SslData            *int     `json:"sslData"`
	SslEnabled         bool     `json:"sslEnabled"`
}

// CreateResourceBody represents request body for resource create.
type CreateResourceBody struct {
	Cname              string   `json:"cname"`
	Origin             string   `json:"origin,omitempty"`
	OriginGroupID      *int     `json:"originGroup,omitempty"`
	SecondaryHostnames []string `json:"secondaryHostnames,omitempty"`
	OriginProtocol     string   `json:"originProtocol,omitempty"`
	SslData            *int     `json:"sslData,omitempty"`
	SslEnabled         bool     `json:"sslEnabled,omitempty"`
	Options            *Options `json:"options,omitempty"`
}

// UpdateResourceBody represents request body for resource update.
type UpdateResourceBody struct {
	Active             *bool    `json:"active,omitempty"`
	Enabled            *bool    `json:"enabled,omitempty"`
	OriginGroup        int      `json:"originGroup,omitempty"`
	SecondaryHostnames []string `json:"secondaryHostnames"`
	OriginProtocol     string   `json:"originProtocol,omitempty"`
	SslData            *int     `json:"sslData,omitempty"`
	SslEnabled         *bool    `json:"sslEnabled,omitempty"`
	Options            *Options `json:"options,omitempty"`
}

// Update method updates resource by given body.
func (s *ResourcesService) Update(ctx context.Context, resourceID int, body *UpdateResourceBody) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodPut,
		fmt.Sprintf(resourceURL, resourceID), body)
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

// List method returns all resources for this account.
func (s *ResourcesService) List(ctx context.Context) ([]*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, resourcesURL, nil)
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

// Get method returns resource by given resourceID.
func (s *ResourcesService) Get(ctx context.Context, resourceID int) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(resourceURL, resourceID), nil)
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

// Create method creates resource.
func (s *ResourcesService) Create(ctx context.Context, body *CreateResourceBody) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, resourcesURL, body)
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

// Purge method deletes cache from CDN servers for given paths.
// If `paths` is empty - purges all cache.
func (s *ResourcesService) Purge(ctx context.Context, resourceID int, paths []string) (*http.Response, error) {
	var pathsBody struct {
		Paths []string `json:"paths"`
	}
	pathsBody.Paths = paths

	req, err := s.client.NewRequest(ctx,
		http.MethodPost,
		fmt.Sprintf(resourcePurgeURL, resourceID), pathsBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Prefetch method pre-loads objects from given paths to
// CDN-servers cache.
func (s *ResourcesService) Prefetch(ctx context.Context, resourceID int, paths []string) (*http.Response, error) {
	var pathsBody struct {
		Paths []string `json:"paths"`
	}
	pathsBody.Paths = paths

	req, err := s.client.NewRequest(ctx,
		http.MethodPost,
		fmt.Sprintf(resourcePrefetchURL, resourceID), pathsBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
