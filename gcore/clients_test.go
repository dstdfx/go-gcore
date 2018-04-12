package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// Mocked responses
var (
	getClientResp = `{
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
  "reseller": 1
}`
	createClientResp = `{
  "id": 2,
  "users": [],
  "currentUser": 7,
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
	listClientResp = `[{
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
  "reseller": 1
}]`
	updateClientResp = `{
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

	userTokenResp = `{"token": "123123123ololo"}`
)

// Expected results
var (
	createClientExpected = &ClientAccount{ID: 2, Users: []User{}, CurrentUser: 7, Status: "active",
		Created:     NewGCoreTime(time.Date(2018, 4, 9, 11, 31, 40, 0, time.UTC)),
		Updated:     NewGCoreTime(time.Date(2018, 4, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName: "Client 2 Company Name", UtilizationLevel: 0, Reseller: 1,
		Email: "common2@gcore.lu", Phone: "Client 2 Company Phone", Name: "Client 2 Name"}

	getClientExpected = &ClientAccount{ID: 2, Users: []User{
		{ID: 6, Name: "Client 2 Name", Deleted: false, Email: "common2@gcore.lu", Client: 2,
			Company: "Client 2 Company Name", Lang: "en", Phone: "1232323",
			Groups: []Group{{ID: 2, Name: "Administrators"}}}}, CurrentUser: 7, Status: "trial",
		Created:     NewGCoreTime(time.Date(2018, 4, 9, 11, 31, 40, 0, time.UTC)),
		Updated:     NewGCoreTime(time.Date(2018, 4, 9, 11, 32, 31, 0, time.UTC)),
		CompanyName: "Client 2 Company Name", UtilizationLevel: 0, Reseller: 1,
		Email: "common2@gcore.lu", Phone: "Client 2 Company Phone", Name: "Client 2 Name"}

	listClientExpected = &[]ClientAccount{*getClientExpected}

	updateClientExpected = getClientExpected

	userTokenExpected = &Token{Value: "123123123ololo", Expire: nil}
)

func TestClientsService_Create(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(resellUsersURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(createClientResp))
		})

	resell := GetAuthenticatedResellerClient()

	body := CreateClientBody{UserType: "common", Name: "Client 2 Name", Company: "Client 2 Company Name"}

	got, _, err := resell.Clients.Create(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	createClientExpected.Users = []User{}

	if !reflect.DeepEqual(got, createClientExpected) {
		t.Errorf("Expected: %+v, got %+v\n", createClientExpected, got)
	}
}

func TestClientsService_Get(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(fmt.Sprintf(resellClientURL, getClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(getClientResp))
		})

	resell := GetAuthenticatedResellerClient()

	got, _, err := resell.Clients.Get(context.Background(), getClientExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, getClientExpected) {
		t.Errorf("Expected: %+v, got %+v\n", getClientExpected, got)
	}
}

func TestClientsService_List(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(resellClientsURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(listClientResp))
		})

	resell := GetAuthenticatedResellerClient()

	got, _, err := resell.Clients.List(context.Background(), ListOpts{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, listClientExpected) {
		t.Errorf("Expected: %+v, got %+v\n", listClientExpected, got)
	}
}

func TestClientsService_Update(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(fmt.Sprintf(resellClientURL, updateClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(updateClientResp))
		})

	resell := GetAuthenticatedResellerClient()

	body := UpdateClientBody{Name: "Another Name"}
	updateClientExpected.Name = "Another Name"

	got, _, err := resell.Clients.Update(context.Background(), updateClientExpected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, updateClientExpected) {
		t.Errorf("Expected: %+v, got %+v\n", updateClientExpected, got)
	}
}

func TestClientsService_GetCommonClient(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	resell := GetAuthenticatedResellerClient()

	Mux.HandleFunc(fmt.Sprintf(resellUserTokenURL, getClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(userTokenResp))
		})

	common, _, err := resell.Clients.GetCommonClient(context.Background(), getClientExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(common.Token, userTokenExpected) {
		t.Errorf("Expected: %+v, got %+v\n", userTokenExpected, common.Token)
	}

}
