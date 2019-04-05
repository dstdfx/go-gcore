package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

const (
	testListRegionsRawResponse = `
[
   {
      "required" : false,
      "description" : "North America",
      "id" : 1,
      "name" : "na"
   },
   {
      "name" : "eu",
      "id" : 2,
      "required" : true,
      "description" : "Europe"
   },
   {
      "name" : "cis",
      "id" : 3,
      "required" : false,
      "description" : "CIS"
   },
   {
      "required" : false,
      "description" : "Asia",
      "id" : 4,
      "name" : "asia"
   },
   {
      "name" : "au",
      "id" : 5,
      "description" : "Australia",
      "required" : false
   },
   {
      "required" : false,
      "description" : "Latin America",
      "name" : "latam",
      "id" : 6
   },
   {
      "description" : "Middle East",
      "required" : false,
      "name" : "me",
      "id" : 7
   },
   {
      "name" : "ru2",
      "id" : 8,
      "required" : false,
      "description" : "Russia 2"
   }
]`

	testGetRestrictionsRawResponse = `
{
   "is_in" : true,
   "region_list" : [
      2,
      3,
      8
   ]
}
`

	testGetRestrictionsNotFound = `
	{
		“error”: “No such client_id”
	}
`
	testSetRestrictionsExcludeRequired = `
	{
		"error": "Some required billing regions absent"
	}
`
)

const (
	testSetRestrictionsRawRequest = `
{
   "is_in" : true,
   "region_list" : [
      2,
      3,
      8
   ]
}
`
	testSetRestrictionsExcludeRequiredRawRequest = `
{
   "is_in" : true,
   "region_list" : [
      1
   ]
}
`
)

var (
	testListRegionsExpected = []*Region{
		{
			ID:          1,
			Name:        "na",
			Description: "North America",
			Required:    false,
		},
		{
			ID:          2,
			Name:        "eu",
			Description: "Europe",
			Required:    true,
		},
		{
			ID:          3,
			Name:        "cis",
			Description: "CIS",
			Required:    false,
		},
		{
			ID:          4,
			Name:        "asia",
			Description: "Asia",
			Required:    false,
		},
		{
			ID:          5,
			Name:        "au",
			Description: "Australia",
			Required:    false,
		},
		{
			ID:          6,
			Name:        "latam",
			Description: "Latin America",
			Required:    false,
		},
		{
			ID:          7,
			Name:        "me",
			Description: "Middle East",
			Required:    false,
		},
		{
			ID:          8,
			Name:        "ru2",
			Description: "Russia 2",
			Required:    false,
		},
	}

	testGetRestrictionsExpected = &GeoRestrictions{
		IsIn:       true,
		RegionList: []int{2, 3, 8},
	}
)

func TestGeoRestrictionsService_ListRegions(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         geoRestrictionsListRegionsURL,
		RawResponse: testListRegionsRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewResellerClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testListRegionsExpected
	got, _, err := client.GeoRestrictions.ListRegions(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a list of regions")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestGeoRestrictionsService_GetRestrictions(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(geoRestrictionsGetSetRestrictionsURL, 1),
		RawResponse: testGetRestrictionsRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewResellerClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testGetRestrictionsExpected
	got, _, err := client.GeoRestrictions.GetRestrictions(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get restrictions")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestGeoRestrictionsService_GetRestrictions_NoRestrictions(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(geoRestrictionsGetSetRestrictionsURL, 2),
		RawResponse: testGetRestrictionsNotFound,
		Method:      http.MethodGet,
		Status:      http.StatusNotFound,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewResellerClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	_, _, err := client.GeoRestrictions.GetRestrictions(context.Background(), 2)
	if err == nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get restrictions")
	}
}

func TestGeoRestrictionsService_SetRestrictions(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(geoRestrictionsGetSetRestrictionsURL, 1),
		RawResponse: testGetRestrictionsRawResponse,
		RawRequest:  testSetRestrictionsRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusNoContent,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewResellerClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	body := &GeoRestrictions{
		IsIn:       true,
		RegionList: []int{2, 3, 8},
	}

	_, err := client.GeoRestrictions.SetRestrictions(context.Background(), 1, body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't set restrictions")
	}
}

func TestGeoRestrictionsService_SetRestrictions_ExcludeRequired(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(geoRestrictionsGetSetRestrictionsURL, 1),
		RawResponse: testSetRestrictionsExcludeRequired,
		RawRequest:  testSetRestrictionsExcludeRequiredRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusConflict,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewResellerClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	body := &GeoRestrictions{
		IsIn:       true,
		RegionList: []int{1},
	}

	_, err := client.GeoRestrictions.SetRestrictions(context.Background(), 1, body)
	if err == nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't set restrictions")
	}
}
