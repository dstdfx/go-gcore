package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	geoRestrictionsBaseURL               = "/admin"
	geoRestrictionsListRegionsURL        = geoRestrictionsBaseURL + "/billing_regions"
	geoRestrictionsGetSetRestrictionsURL = geoRestrictionsBaseURL + "/clients/%d"
)

// GeoRestrictionsService handles communication with geo-restrictions related methods
// of the G-Core CDN API.
type GeoRestrictionsService service

// Region represents the region which is utilized in content delivery.
type Region struct {
	// ID represents the region identifier.
	ID int `json:"id"`

	// Name represents region abbreviation.
	Name string `json:"name"`

	// Description represents the region full name.
	Description string `json:"description"`

	// Required flag indicates if the region is might be excluded from the content delivery.
	Required bool `json:"required"`
}

// GeoRestrictions represents the list of regions that are utilized
// in content delivery for a client.
type GeoRestrictions struct {
	// IsIn flag shows if regions are utilized in the content delivery for a client or not:
	// true — only regions from the list are utilized in the content delivery for the client.
	// false — regions from the list are not participating in the content delivery for the client
	IsIn bool `json:"is_in"`

	// RegionList represents the list of regions IDs that are utilized
	// in the content delivery for a client.
	// Numbers 1-7 represent a region.
	RegionList []int `json:"region_list"`
}

// ListRegions method returns the regions that are available for content delivery.
func (s *GeoRestrictionsService) ListRegions(ctx context.Context) ([]*Region, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, geoRestrictionsListRegionsURL, nil)
	if err != nil {
		return nil, nil, err
	}

	regions := make([]*Region, 0)

	resp, err := s.client.Do(req, &regions)
	if err != nil {
		return nil, resp, err
	}
	return regions, resp, nil
}

// GetRestrictions method returns current client's geo restrictions.
// Returns 404 response code if restrictions haven't been set before.
func (s *GeoRestrictionsService) GetRestrictions(ctx context.Context,
	clientID int) (*GeoRestrictions, *http.Response, error) {

	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(geoRestrictionsGetSetRestrictionsURL, clientID), nil)
	if err != nil {
		return nil, nil, err
	}

	restrict := &GeoRestrictions{}

	resp, err := s.client.Do(req, restrict)
	if err != nil {
		return nil, resp, err
	}
	return restrict, resp, nil
}

// SetRestrictions method limits the regions for the content delivery for a client.
// In case of success it returns 204 response body.
func (s *GeoRestrictionsService) SetRestrictions(ctx context.Context,
	clientID int, restrictions *GeoRestrictions) (*http.Response, error) {

	req, err := s.client.NewRequest(ctx,
		http.MethodPost,
		fmt.Sprintf(geoRestrictionsGetSetRestrictionsURL, clientID), restrictions)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
