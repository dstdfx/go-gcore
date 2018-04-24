package origingroups

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-gcore/gcore"
	th "github.com/dstdfx/go-gcore/tests/testutils"
)

func TestOriginGroupsService_List(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(gcore.OriginGroupsURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestListOriginGroupsResponse))
		})
	client := th.GetAuthenticatedCommonClient()
	expected := TestListOriginGroupsExpected
	got, _, err := client.OriginGroups.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Get(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestGetOriginGroupExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.OriginGroupURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestGetOriginGroupResponse))
		})

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Get(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Create(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(gcore.OriginGroupsURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestCreateOriginGroupResponse))
		})
	expected := TestCreateOriginGroupExpected
	body := gcore.CreateOriginGroupBody{
		Name:    "whatever.ru_wiggly.gcdn.co",
		UseNext: false,
		Origins: expected.Origins}

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Create(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Update(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestUpdateOriginGroupExpected

	th.Mux.HandleFunc(fmt.Sprintf(gcore.OriginGroupURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestUpdateOriginGroupResponse))
		})

	body := gcore.UpdateOriginGroupBody{
		Name:    expected.Name,
		UseNext: expected.UseNext,
		Origins: expected.Origins,
	}

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Update(context.Background(), expected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Delete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(fmt.Sprintf(gcore.OriginGroupURL, TestGetOriginGroupExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	client := th.GetAuthenticatedCommonClient()
	_, err := client.OriginGroups.Delete(context.Background(), TestGetOriginGroupExpected.ID)
	if err != nil {
		t.Fatal(err)
	}
}
