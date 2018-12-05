package resources

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-gcore/gcore"
	th "github.com/dstdfx/go-gcore/tests/testutils"
)

func TestResourcesService_Get(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestGetResourceExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.ResourceURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestGetResourceResponse))
		})

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.Resources.Get(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestResourcesService_List(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(gcore.ResourcesURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestListResourcesResponse))
		})

	client := th.GetAuthenticatedCommonClient()
	expected := TestListResourcesExpected
	got, _, err := client.Resources.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestResourcesService_Create(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(gcore.ResourcesURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestCreateResourceResponse))
		})

	resourceBody := gcore.CreateResourceBody{
		Cname: "cdn.site.com",
		SecondaryHostnames: []string{
			"cdn1.yoursite.com",
			"cdn2.yoursite.com",
		},
	}

	client := th.GetAuthenticatedCommonClient()
	expected := TestCreateResourceExpected
	got, _, err := client.Resources.Create(context.Background(), &resourceBody)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}

}

func TestResourceService_Purge(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(fmt.Sprintf(gcore.ResourcePurgeURL, TestResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		})

	client := th.GetAuthenticatedCommonClient()
	resp, err := client.Resources.Purge(context.Background(), TestResourceID, []string{})
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d",
			http.StatusCreated,
			resp.StatusCode,
		)
	}
}

func TestResourceService_Prefetch(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(fmt.Sprintf(gcore.ResourcePrefetchURL, TestResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
		})

	client := th.GetAuthenticatedCommonClient()
	resp, err := client.Resources.Prefetch(context.Background(), TestResourceID,
		[]string{"/file.jpg", "file2.jpg"})
	if err != nil {
		t.Errorf("Expected no error, but got: %s", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d",
			http.StatusCreated,
			resp.StatusCode,
		)
	}

}
