package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// Fixtures
var (
	testGetClientResponse = `{
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
	testCreateClientResponse = `{
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
	testListClientsResponse = `[{
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
	testUpdateClientResponse = `{
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
  "name": "Another Name",
  "status": "trial",
  "created": "2018-04-09T11:31:40.000000Z",
  "updated": "2018-04-09T11:32:31.000000Z",
  "companyName": "Client 2 Company Name",
  "utilization_level": 0,
  "reseller": 1
}`

	testUserTokenResponse = `{"token": "123123123ololo"}`
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
		Name:             "Client 2 Name",
	}

	testUserTokenExpected = &Token{
		Value:  "123123123ololo",
		Expire: nil,
	}
)

func TestClientsService_Create(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(ResellUsersURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testCreateClientResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	resell := getAuthenticatedResellerClient()

	body := CreateClientBody{
		UserType: "common",
		Name:     "Client 2 Name",
		Company:  "Client 2 Company Name",
		Phone:    "Client 2 Company Phone",
		Email:    "common2@gcore.lu",
		Password: "123123123qwe",
	}

	got, _, err := resell.Clients.Create(context.Background(), &body)
	expected := testCreateClientExpected
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientService_Get(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(fmt.Sprintf(ResellClientURL, testGetClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testGetClientResponse))
			if err != nil {
				t.Fatal(err)
			}
		})
	resell := getAuthenticatedResellerClient()

	got, _, err := resell.Clients.Get(context.Background(), testGetClientExpected.ID)
	expected := testGetClientExpected
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_List(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(ResellClientsURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testListClientsResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	resell := getAuthenticatedResellerClient()

	got, _, err := resell.Clients.List(context.Background(), ListOpts{})
	expected := testListClientsExpected
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_Update(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(fmt.Sprintf(ResellClientURL, testUpdateClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testUpdateClientResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	resell := getAuthenticatedResellerClient()

	body := UpdateClientBody{Name: "Another Name"}
	testUpdateClientExpected.Name = "Another Name"

	expected := testUpdateClientExpected
	got, _, err := resell.Clients.Update(context.Background(), expected.ID, &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_GetCommonClient(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	resell := getAuthenticatedResellerClient()

	mux.HandleFunc(fmt.Sprintf(ResellUserTokenURL, testGetClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testUserTokenResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	common, _, err := resell.Clients.GetCommonClient(context.Background(), testGetClientExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(common.Token, testUserTokenExpected) {
		t.Errorf("Expected: %+v, got %+v\n", testUserTokenExpected, common.Token)
	}

}
