package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	resourcesURL     = "/resources"
	resourceURL      = "/resources/%d"
	resourcePurgeURL = "/resources/%d/purge"
)

type ResourcesService service

type Resource struct {
	ID                 int        `json:"id"`
	Deleted            bool       `json:"deleted"`
	Enabled            bool       `json:"enabled"`
	CompanyName        string     `json:"companyName"`
	Status             string     `json:"status"`
	Client             int        `json:"client"`
	OriginGroup        int        `json:"originGroup"`
	CName              string     `json:"cname"`
	SecondaryHostnames []string   `json:"secondaryHostnames"`
	CreatedAt          *GCoreTime `json:"created"`
	UpdatedAt          *GCoreTime `json:"updated"`
	// TODO: Options, Rules, sslData, sslEnabled
}

type CreateResourceBody struct {
	CName              string   `json:"cname"`
	Origin             string   `json:"origin"`
	SecondaryHostnames []string `json:"secondaryHostnames"`
}

func (s *ResourcesService) List(ctx context.Context) ([]*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", resourcesURL, nil)
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

func (s *ResourcesService) Get(ctx context.Context, resourceID int) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", fmt.Sprintf(resourceURL, resourceID), nil)
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

func (s *ResourcesService) Create(ctx context.Context, body CreateResourceBody) (*Resource, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "POST", resourcesURL, body)
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

func (s *ResourcesService) Purge(ctx context.Context, resourceID int, paths []string) (*http.Response, error) {
	var pathsBody struct {
		Paths []string `json:"paths"`
	}
	pathsBody.Paths = paths

	req, err := s.client.NewRequest(ctx, "POST", fmt.Sprintf(resourcePurgeURL, resourceID), pathsBody)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
