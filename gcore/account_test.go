package gcore

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

// Fixtures
var testAccountDetailResponse = `{
    "currentUser": 511,
    "id": 505,
	"cname": "example.gcdn.co",
    "users": [
        {
            "client": 5,
            "company": "Your company",
            "deleted": false,
            "email": "user@yourcompany.com",
            "id": 513,
            "lang": "en",
            "name": "user",
            "phone": "+79882233443",
           "groups": [
               {
                "id": 2,
                "name": "users"
               }
           ]
        }
    ]
}
     `

var testAccountDetailExpected = &Account{
	CurrentUser: 511,
	ID:          505,
	Cname:       "example.gcdn.co",
	Users: []User{
		{
			Client:  5,
			Company: "Your company",
			Deleted: false,
			Email:   "user@yourcompany.com",
			ID:      513,
			Lang:    "en",
			Name:    "user",
			Phone:   "+79882233443",
			Groups: []*Group{
				{
					ID:   2,
					Name: "users",
				},
			},
		},
	},
}

func TestAccountService_Details(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         accountDetailsURL,
		RawResponse: testAccountDetailResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testAccountDetailExpected
	got, _, err := client.Account.Details(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get service account details")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
