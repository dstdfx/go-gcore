package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var (
	getOriginGroupResp = `{
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
	createOriginGroupResp = getOriginGroupResp
	listOriginGroupReps   = `[{
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
	updateOriginGroupResp = `{
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
	getOriginGroupExpected = &OriginGroup{UseNext: false, ID: 7272,
		Name:      "whatever.ru_wiggly.gcdn.co",
		Origins:   []Origin{{Backup: false, Source: "whatever.ru", Enabled: true}},
		OriginIDs: []Origin{{ID: 9257, Source: "whatever.ru", Enabled: true, Backup: false}},
	}

	listOriginGroupExpected = &[]OriginGroup{*getOriginGroupExpected}

	createOriginGroupExpected = getOriginGroupExpected

	updateOriginGroupExpected = &OriginGroup{UseNext: false, ID: 7272,
		Name:      "whatever2.ru_wiggly.gcdn.co",
		Origins:   []Origin{{Backup: false, Source: "whatever.ru", Enabled: true}},
		OriginIDs: []Origin{{ID: 9257, Source: "whatever.ru", Enabled: true, Backup: false}},
	}
)

func TestOriginGroupsService_Create(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(originGroupsURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(createOriginGroupResp))
		})

	client := GetAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, listOriginGroupExpected) {
		t.Errorf("Expected: %+v, got %+v\n", listOriginGroupExpected, got)
	}
}

func TestOriginGroupsService_Get(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(fmt.Sprintf(originGroupURL, getOriginGroupExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(getOriginGroupResp))
		})

	client := GetAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Get(context.Background(), getOriginGroupExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, getOriginGroupExpected) {
		t.Errorf("Expected: %+v, got %+v\n", getOriginGroupExpected, got)
	}
}

func TestOriginGroupsService_List(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(originGroupsURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(createOriginGroupResp))
		})

	body := CreateOriginGroupBody{Name: "whatever.ru_wiggly.gcdn.co", UseNext: false, Origins: getOriginGroupExpected.Origins}

	client := GetAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Create(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, createOriginGroupExpected) {
		t.Errorf("Expected: %+v, got %+v\n", createOriginGroupExpected, got)
	}
}

func TestOriginGroupsService_Update(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(fmt.Sprintf(originGroupURL, getOriginGroupExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(updateOriginGroupResp))
		})

	body := UpdateOriginGroupBody{Name: updateOriginGroupExpected.Name,
		UseNext: updateOriginGroupExpected.UseNext, Origins: updateOriginGroupExpected.Origins}

	client := GetAuthenticatedCommonClient()
	got, _, err := client.OriginGroups.Update(context.Background(), getOriginGroupExpected.ID, body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, updateOriginGroupExpected) {
		t.Errorf("Expected: %+v, got %+v\n", updateOriginGroupExpected, got)
	}
}

func TestOriginGroupsService_Delete(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(fmt.Sprintf(originGroupURL, getOriginGroupExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	client := GetAuthenticatedCommonClient()
	_, err := client.OriginGroups.Delete(context.Background(), getOriginGroupExpected.ID)
	if err != nil {
		t.Fatal(err)
	}
}
