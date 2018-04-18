package gcore

import (
	"context"
	"fmt"
	"net/http"
)

var (
	originGroupsURL = "/originGroups"
	originGroupURL  = "/originGroups/%d"
)

type OriginGroupsService service

type Origin struct {
	ID      int    `json:"id,omitempty"`
	Enabled bool   `json:"enabled,omitempty"`
	Backup  bool   `json:"backup"`
	Source  string `json:"source"`
}

type OriginGroup struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	UseNext   bool     `json:"useNext"`
	OriginIDs []Origin `json:"origin_ids,omitempty"`
	Origins   []Origin `json:"origins"`
}

type UpdateOriginGroupBody struct {
	Name    string   `json:"name"`
	UseNext bool     `json:"useNext"`
	Origins []Origin `json:"origins"`
}

type CreateOriginGroupBody struct {
	Name    string   `json:"name"`
	UseNext bool     `json:"useNext"`
	Origins []Origin `json:"origins"`
}

func (s *OriginGroupsService) List(ctx context.Context) ([]*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", originGroupsURL, nil)
	if err != nil {
		return nil, nil, err
	}

	originGroups := make([]*OriginGroup, 0)

	resp, err := s.client.Do(req, &originGroups)
	if err != nil {
		return nil, resp, err
	}

	return originGroups, resp, nil
}

func (s *OriginGroupsService) Get(ctx context.Context, originGroupID int) (*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "GET", fmt.Sprintf(originGroupURL, originGroupID), nil)
	if err != nil {
		return nil, nil, err
	}

	originGroup := &OriginGroup{}

	resp, err := s.client.Do(req, originGroup)
	if err != nil {
		return nil, resp, err
	}

	return originGroup, resp, nil
}

func (s *OriginGroupsService) Create(ctx context.Context, body CreateOriginGroupBody) (*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "POST", originGroupsURL, body)
	if err != nil {
		return nil, nil, err
	}

	originGroup := &OriginGroup{}

	resp, err := s.client.Do(req, originGroup)
	if err != nil {
		return nil, resp, err
	}

	return originGroup, resp, nil
}

func (s *OriginGroupsService) Update(ctx context.Context, originGroupID int, body UpdateOriginGroupBody) (*OriginGroup, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, "PUT", fmt.Sprintf(originGroupURL, originGroupID), body)
	if err != nil {
		return nil, nil, err
	}

	originGroup := &OriginGroup{}

	resp, err := s.client.Do(req, originGroup)
	if err != nil {
		return nil, resp, err
	}

	return originGroup, resp, nil
}

func (s *OriginGroupsService) Delete(ctx context.Context, originGroupID int) (*http.Response, error) {
	req, err := s.client.NewRequest(ctx, "DELETE", fmt.Sprintf(originGroupURL, originGroupID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
