package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// Fixtures
var (
	testGetOriginGroupResponse = `{
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
	testCreateOriginGroupResponse = `{
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
	testListOriginGroupsResponse = `[{
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
	testUpdateOriginGroupResponse = `{
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

func TestOriginGroupsService_List(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(originGroupsURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testListOriginGroupsResponse))
			if err != nil {
				t.Fatal(err)
			}
		})
	client := getAuthenticatedCommonClient()
	expected := testListOriginGroupsExpected
	got, _, err := client.OriginGroups.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Get(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testGetOriginGroupExpected
	mux.HandleFunc(fmt.Sprintf(originGroupURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testGetOriginGroupResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := getAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Get(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Create(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(originGroupsURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testCreateOriginGroupResponse))
			if err != nil {
				t.Fatal(err)
			}
		})
	expected := testCreateOriginGroupExpected
	body := CreateOriginGroupBody{
		Name:    "whatever.ru_wiggly.gcdn.co",
		UseNext: false,
		Origins: expected.Origins}

	client := getAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Create(context.Background(), &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Update(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testUpdateOriginGroupExpected

	mux.HandleFunc(fmt.Sprintf(originGroupURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testUpdateOriginGroupResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	body := UpdateOriginGroupBody{
		Name:    expected.Name,
		UseNext: expected.UseNext,
		Origins: expected.Origins,
	}

	client := getAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Update(context.Background(), expected.ID, &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestOriginGroupsService_Delete(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	mux.HandleFunc(fmt.Sprintf(originGroupURL, testGetOriginGroupExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	client := getAuthenticatedCommonClient()
	_, err := client.OriginGroups.Delete(context.Background(), testGetOriginGroupExpected.ID)
	if err != nil {
		t.Fatal(err)
	}
}
