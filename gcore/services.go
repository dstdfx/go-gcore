package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	serviceListURL   = "/clients/%d/services"
	serviceUpdateURL = serviceListURL + "/%d"
)

// ServicesService handles communication with service related methods
// of the G-Core CDN API.
type ServicesService service

// Service represents G-Core client's service.
type Service struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Client  int    `json:"client"`
	Status  string `json:"status"`
	Enabled bool   `json:"enabled"`
	Start   *Time  `json:"start"`
	// TODO: add trial options, options
}

// UpdateServiceBody represents request body for service updating.
type UpdateServiceBody struct {
	Enabled bool   `json:"enabled"`
	Status  string `json:"status"`
}

// List method returns list of the client's services.
func (s *ServicesService) List(ctx context.Context, clientID int) ([]*Service, *http.Response, error) {
	url := fmt.Sprintf(serviceListURL, clientID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}

	services := make([]*Service, 0)

	resp, err := s.client.Do(req, &services)
	if err != nil {
		return nil, resp, err
	}

	return services, resp, nil
}

// Update method updates service status.
func (s *ServicesService) Update(ctx context.Context, clientID, serviceID int, body *UpdateServiceBody) (*Service, *http.Response, error) {
	url := fmt.Sprintf(serviceUpdateURL, clientID, serviceID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, url, body)
	if err != nil {
		return nil, nil, err
	}

	service := &Service{}

	resp, err := s.client.Do(req, service)
	if err != nil {
		return nil, resp, err
	}

	return service, resp, nil
}
