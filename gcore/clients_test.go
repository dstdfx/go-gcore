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

// Fixtures
const (
	testGetClientRawResponse = `{
  "id": 2,
  "users": [
    {
      "id": 6,
      "deleted": false,
      "email": "common2@gcore.lu",
      "name": "Client 2 Name",
      "client": 2,
      "company": "Client 2 Company Name",
      "lang": "en",
      "phone": "1232323",
      "groups": [
        {
          "id": 2,
          "name": "Administrators"
        }
      ]
    }
  ],
  "currentUser": 7,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "Client 2 Name",
  "status": "trial",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1,
  "cname": "example.gcdn.co"
}`
	testCreateClientRawResponse = `{
  "id": 2,
  "users": [],
  "currentUser": 7,
  "client": 5,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "Client 2 Name",
  "status": "active",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1
}`
	testListClientsRawResponse = `[{
  "id": 2,
  "users": [
    {
      "id": 6,
      "deleted": false,
      "email": "common2@gcore.lu",
      "name": "Client 2 Name",
      "client": 2,
      "company": "Client 2 Company Name",
      "lang": "en",
      "phone": "1232323",
      "groups": [
        {
          "id": 2,
          "name": "Administrators"
        }
      ]
    }
  ],
  "currentUser": 7,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "Client 2 Name",
  "status": "trial",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1,
  "cname": "example.gcdn.co"
}]`
	testUpdateClientRawResponse = `{
  "id": 2,
  "users": [
    {
      "id": 6,
      "deleted": false,
      "email": "common2@gcore.lu",
      "name": "Client 2 Name",
      "client": 2,
      "company": "Client 2 Company Name",
      "lang": "en",
      "phone": "1232323",
      "groups": [
        {
          "id": 2,
          "name": "Administrators"
        }
      ]
    }
  ],
  "currentUser": 7,
  "email": "common2@gcore.lu",
  "phone": "Client 2 Company Phone",
  "name": "New name",
  "status": "trial",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1
}`

	testUserTokenRawResponse = `{"token": "123123123ololo"}`
)

const (
	testCreateClientRawRequest = `{
  "user_type":"common",
  "name":"Client 2 Name",
  "company":"Client 2 Company Name",
  "phone":"Client 2 Company Phone",
  "email":"common2@gcore.lu",
  "password":"123123123qwe"
}`

	testUpdateClientRawRequest = `{
	"name": "New name",
	"companyName":"Client 2 Company Name",
  	"phone":"Client 2 Company Phone",
  	"email":"common2@gcore.lu"
	}`
)

// Expected results
var (
	testCreateClientExpected = &ClientAccount{
		ID:               2,
		Users:            []*User{},
		CurrentUser:      7,
		Status:           "active",
		Created:          NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		Updated:          NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName:      "Client 2 Company Name",
		UtilizationLevel: 0,
		Reseller:         1,
		Email:            "common2@gcore.lu",
		Phone:            "Client 2 Company Phone",
		Name:             "Client 2 Name",
		Client:           5,
	}

	testGetClientExpected = &ClientAccount{
		ID: 2,
		Users: []*User{
			{
				ID:      6,
				Name:    "Client 2 Name",
				Deleted: false,
				Email:   "common2@gcore.lu",
				Client:  2,
				Company: "Client 2 Company Name",
				Lang:    "en",
				Phone:   "1232323",
				Groups: []*Group{
					{
						ID:   2,
						Name: "Administrators",
					},
				},
			},
		},
		CurrentUser:      7,
		Status:           "trial",
		Created:          NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		Updated:          NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName:      "Client 2 Company Name",
		UtilizationLevel: 0,
		Reseller:         1,
		Email:            "common2@gcore.lu",
		Phone:            "Client 2 Company Phone",
		Name:             "Client 2 Name",
		Cname:            "example.gcdn.co",
	}

	testListClientsExpected = []*ClientAccount{
		{
			ID: 2,
			Users: []*User{
				{
					ID:      6,
					Name:    "Client 2 Name",
					Deleted: false,
					Email:   "common2@gcore.lu",
					Client:  2,
					Company: "Client 2 Company Name",
					Lang:    "en",
					Phone:   "1232323",
					Groups: []*Group{
						{
							ID:   2,
							Name: "Administrators",
						},
					},
				},
			},
			CurrentUser:      7,
			Status:           "trial",
			Created:          NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
			Updated:          NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
			CompanyName:      "Client 2 Company Name",
			UtilizationLevel: 0,
			Reseller:         1,
			Email:            "common2@gcore.lu",
			Phone:            "Client 2 Company Phone",
			Name:             "Client 2 Name",
			Cname:            "example.gcdn.co",
		},
	}

	testUpdateClientExpected = &ClientAccount{
		ID: 2,
		Users: []*User{
			{
				ID:      6,
				Name:    "Client 2 Name",
				Deleted: false,
				Email:   "common2@gcore.lu",
				Client:  2,
				Company: "Client 2 Company Name",
				Lang:    "en",
				Phone:   "1232323",
				Groups: []*Group{
					{
						ID:   2,
						Name: "Administrators",
					},
				},
			},
		},
		CurrentUser:      7,
		Status:           "trial",
		Created:          NewTime(time.Date(2018, time.April, 9, 11, 31, 40, 0, time.UTC)),
		Updated:          NewTime(time.Date(2018, time.April, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName:      "Client 2 Company Name",
		UtilizationLevel: 0,
		Reseller:         1,
		Email:            "common2@gcore.lu",
		Phone:            "Client 2 Company Phone",
		Name:             "New name",
	}

	testUserTokenExpected = &Token{
		Value:  "123123123ololo",
		Expire: nil,
	}
)

func TestClientsService_Create(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         resellUsersURL,
		RawResponse: testCreateClientRawResponse,
		RawRequest:  testCreateClientRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	resell := NewResellerClient()
	resell.BaseURL = testEnv.GetServerURL()
	_ = resell.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testCreateClientExpected

	body := &CreateClientBody{
		UserType: "common",
		Name:     "Client 2 Name",
		Company:  "Client 2 Company Name",
		Phone:    "Client 2 Company Phone",
		Email:    "common2@gcore.lu",
		Password: "123123123qwe",
	}

	got, _, err := resell.Clients.Create(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't create new client account")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_Get(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(resellClientURL, testGetClientExpected.ID),
		RawResponse: testGetClientRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	resell := NewResellerClient()
	resell.BaseURL = testEnv.GetServerURL()
	_ = resell.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testGetClientExpected

	got, _, err := resell.Clients.Get(context.Background(), testGetClientExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a client account")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_List(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         resellClientsURL,
		RawResponse: testListClientsRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	resell := NewResellerClient()
	resell.BaseURL = testEnv.GetServerURL()
	_ = resell.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testListClientsExpected

	got, _, err := resell.Clients.List(context.Background(), ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a list of accounts")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_Update(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(resellClientURL, testGetClientExpected.ID),
		RawResponse: testUpdateClientRawResponse,
		RawRequest:  testUpdateClientRawRequest,
		Method:      http.MethodPut,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	resell := NewResellerClient()
	resell.BaseURL = testEnv.GetServerURL()
	_ = resell.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testUpdateClientExpected

	body := &UpdateClientBody{
		Name:        "New name",
		CompanyName: expected.CompanyName,
		Email:       expected.Email,
		Phone:       expected.Phone,
	}

	got, _, err := resell.Clients.Update(context.Background(), testGetClientExpected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't update client account")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_GetCommonClient(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(resellUserTokenURL, testGetClientExpected.ID),
		RawResponse: testUserTokenRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	resell := NewResellerClient()
	resell.BaseURL = testEnv.GetServerURL()
	_ = resell.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testUserTokenExpected

	got, _, err := resell.Clients.GetCommonClient(context.Background(), testGetClientExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a client account")
	}

	if !reflect.DeepEqual(got.Token, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
