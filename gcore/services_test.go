package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

const (
	fakeClientID  = 1234
	fakeServiceID = 21797
)

const (
	testListServicesRawResponse = `
[
 {
  "id": 21799,
  "name": "STREAMING",
  "client": 3989,
  "status": "new",
  "enabled": false,
  "start": "2018-04-09T11:31:40.000000Z"
 },
 {
  "id": 21798,
  "name": "STORAGE",
  "client": 3989,
  "status": "new",
  "enabled": false,
  "start": "2018-04-09T11:31:40.000000Z"
 },
 {
  "id": 21797,
  "name": "CDN",
  "client": 3989,
  "status": "active",
  "enabled": true,
  "start": "2018-04-09T11:31:40.000000Z"
 }
]`
	testUpdateServicesRawResponse = `
{
  "id": 21797,
  "name": "CDN",
  "client": 3989,
  "status": "paused",
  "enabled": false,
  "start": "2018-04-09T11:31:40.000000Z"
 }`
)

const testUpdateServiceRawRequest = `
{
    "enabled": false,
	"status": "paused",
}`

var (
	testListServicesExpected = []*Service{
		{
			ID:      21799,
			Name:    "STREAMING",
			Client:  3989,
			Status:  "new",
			Enabled: false,
			Start:   NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		},
		{
			ID:      21798,
			Name:    "STORAGE",
			Client:  3989,
			Status:  "new",
			Enabled: false,
			Start:   NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		},
		{
			ID:      21797,
			Name:    "CDN",
			Client:  3989,
			Status:  "active",
			Enabled: true,
			Start:   NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		},
	}
	testUpdateServiceExpected = &Service{
		ID:      21797,
		Name:    "CDN",
		Client:  3989,
		Status:  "paused",
		Enabled: false,
		Start:   NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
	}
)

func TestServicesService_ListServices(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(serviceListURL, fakeClientID),
		RawResponse: testListServicesRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	resell := NewResellerClient()
	resell.BaseURL = testEnv.GetServerURL()
	_ = resell.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testListServicesExpected

	got, _, err := resell.Services.List(context.Background(), fakeClientID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get client's services")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestServicesService_UpdateService(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(serviceUpdateURL, fakeClientID, fakeServiceID),
		RawResponse: testUpdateServicesRawResponse,
		RawRequest:  testUpdateServiceRawRequest,
		Method:      http.MethodPatch,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	resell := NewResellerClient()
	resell.BaseURL = testEnv.GetServerURL()
	_ = resell.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testUpdateServiceExpected

	got, _, err := resell.Services.Update(context.Background(),
		fakeClientID,
		fakeServiceID,
		&UpdateServiceBody{
			Enabled: false,
			Status:  "paused",
		})
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get client's services")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
