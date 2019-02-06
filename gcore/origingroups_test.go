package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

// Fixtures
const (
	testGetOriginGroupRawResponse = `{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}`
	testCreateOriginGroupRawResponse = `{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}`
	testListOriginGroupsRawResponse = `[{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}]`
	testUpdateOriginGroupRawResponse = `{
   "useNext" : false,
   "id" : 7272,
   "path" : "",
   "name" : "whatever2.ru_wiggly.gcdn.co",
   "origins" : [
      {
         "enabled" : true,
         "backup" : false,
         "source" : "whatever.ru"
      }
   ],
   "origin_ids" : [
      {
         "id" : 9257,
         "source" : "whatever.ru",
         "enabled" : true,
         "backup" : false
      }
   ]
}`
)

const (
	testCreateOriginGroupRawRequest = `{
	"name":"whatever.ru_wiggly.gcdn.co",
	"useNext":false,
	"origins":[
	{
		"enabled":true,
		"backup":false,
		"source":"whatever.ru"
	}]
}
`
	testUpdateOriginGroupRawRequest = `{
	"originGroup":7260,
	"secondaryHostnames":["cdn1.yoursite.com","cdn2.yoursite.com"]
}
`
)

var (
	testGetOriginGroupExpected = &OriginGroup{
		UseNext: false,
		ID:      7272,
		Name:    "whatever.ru_wiggly.gcdn.co",
		Origins: []Origin{
			{
				Backup:  false,
				Source:  "whatever.ru",
				Enabled: true,
			},
		},
		OriginIDs: []Origin{
			{
				ID:      9257,
				Source:  "whatever.ru",
				Enabled: true,
				Backup:  false,
			},
		},
	}

	testListOriginGroupsExpected = []*OriginGroup{
		{
			UseNext: false,
			ID:      7272,
			Name:    "whatever.ru_wiggly.gcdn.co",
			Origins: []Origin{
				{
					Backup:  false,
					Source:  "whatever.ru",
					Enabled: true,
				},
			},
			OriginIDs: []Origin{
				{
					ID:      9257,
					Source:  "whatever.ru",
					Enabled: true,
					Backup:  false,
				},
			},
		},
	}

	testCreateOriginGroupExpected = &OriginGroup{
		UseNext: false,
		ID:      7272,
		Name:    "whatever.ru_wiggly.gcdn.co",
		Origins: []Origin{
			{
				Backup:  false,
				Source:  "whatever.ru",
				Enabled: true,
			},
		},
		OriginIDs: []Origin{
			{
				ID:      9257,
				Source:  "whatever.ru",
				Enabled: true,
				Backup:  false,
			},
		},
	}

	testUpdateOriginGroupExpected = &OriginGroup{
		UseNext: false,
		ID:      7272,
		Name:    "whatever2.ru_wiggly.gcdn.co",
		Origins: []Origin{
			{
				Backup:  false,
				Source:  "whatever.ru",
				Enabled: true,
			},
		},
		OriginIDs: []Origin{
			{
				ID:      9257,
				Source:  "whatever.ru",
				Enabled: true,
				Backup:  false,
			},
		},
	}
)

func TestOriginGroupsService_Create(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         originGroupsURL,
		RawResponse: testCreateOriginGroupRawResponse,
		RawRequest:  testCreateOriginGroupRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testCreateOriginGroupExpected

	body := &CreateOriginGroupBody{
		Name:    "whatever.ru_wiggly.gcdn.co",
		UseNext: false,
		Origins: expected.Origins,
	}

	got, _, err := client.OriginGroups.Create(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't create an origin group")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Get(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(originGroupURL, testGetOriginGroupExpected.ID),
		RawResponse: testGetOriginGroupRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testGetOriginGroupExpected

	got, _, err := client.OriginGroups.Get(context.Background(), testGetOriginGroupExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get an origin group")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Delete(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(originGroupURL, testGetOriginGroupExpected.ID),
		RawResponse: "",
		Method:      http.MethodDelete,
		Status:      http.StatusNoContent,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	_, err := client.OriginGroups.Delete(context.Background(), testGetOriginGroupExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't delete an origin group")
	}
}

func TestOriginGroupsService_List(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         originGroupsURL,
		RawResponse: testListOriginGroupsRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testListOriginGroupsExpected

	got, _, err := client.OriginGroups.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a list of origin groups")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Update(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(originGroupURL, testGetOriginGroupExpected.ID),
		RawResponse: testUpdateOriginGroupRawResponse,
		RawRequest:  testUpdateOriginGroupRawRequest,
		Method:      http.MethodPut,
		Status:      http.StatusOK,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testUpdateOriginGroupExpected

	body := &UpdateOriginGroupBody{
		Name:    expected.Name,
		UseNext: expected.UseNext,
		Origins: expected.Origins,
	}

	got, _, err := client.OriginGroups.Update(context.Background(), expected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't update an origin group")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
