package clients

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-gcore/gcore"
	th "github.com/dstdfx/go-gcore/tests/testutils"
)

func TestClientsService_Create(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(gcore.ResellUsersURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(TestCreateClientResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	resell := th.GetAuthenticatedResellerClient()

	body := gcore.CreateClientBody{
		UserType: "common",
		Name:     "Client 2 Name",
		Company:  "Client 2 Company Name",
		Phone:    "Client 2 Company Phone",
		Email:    "common2@gcore.lu",
		Password: "123123123qwe",
	}

	got, _, err := resell.Clients.Create(context.Background(), &body)
	expected := TestCreateClientExpected
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientService_Get(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(fmt.Sprintf(gcore.ResellClientURL, TestGetClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(TestGetClientResponse))
			if err != nil {
				t.Fatal(err)
			}
		})
	resell := th.GetAuthenticatedResellerClient()

	got, _, err := resell.Clients.Get(context.Background(), TestGetClientExpected.ID)
	expected := TestGetClientExpected
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_List(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(gcore.ResellClientsURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(TestListClientsResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	resell := th.GetAuthenticatedResellerClient()

	got, _, err := resell.Clients.List(context.Background(), gcore.ListOpts{})
	expected := TestListClientsExpected
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_Update(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(fmt.Sprintf(gcore.ResellClientURL, TestUpdateClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(TestUpdateClientResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	resell := th.GetAuthenticatedResellerClient()

	body := gcore.UpdateClientBody{Name: "Another Name"}
	TestUpdateClientExpected.Name = "Another Name"

	expected := TestUpdateClientExpected
	got, _, err := resell.Clients.Update(context.Background(), expected.ID, &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestClientsService_GetCommonClient(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	resell := th.GetAuthenticatedResellerClient()

	th.Mux.HandleFunc(fmt.Sprintf(gcore.ResellUserTokenURL, TestGetClientExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(TestUserTokenResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	common, _, err := resell.Clients.GetCommonClient(context.Background(), TestGetClientExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(common.Token, TestUserTokenExpected) {
		t.Errorf("Expected: %+v, got %+v\n", TestUserTokenExpected, common.Token)
	}

}
