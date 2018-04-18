package gcore

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

var (
	// Mocked response
	accountDetailsResp = `{
    "currentUser": 511,
    "id": 505,
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
	// Expected result
	accountDetaildExpected = &Account{CurrentUser: 511, ID: 505, Users: []User{{Client: 5, Company: "Your company", Deleted: false, Email: "user@yourcompany.com",
		ID: 513, Lang: "en", Name: "user", Phone: "+79882233443", Groups: []*Group{{ID: 2, Name: "users"}}}}}
)

func TestAccountService_Details(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(accountDetailsURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(accountDetailsResp))
		})

	client := GetAuthenticatedCommonClient()
	got, _, err := client.Account.Details(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, accountDetaildExpected) {
		t.Errorf("Expected: %+v, got %+v\n", accountDetaildExpected, got)
	}
}
