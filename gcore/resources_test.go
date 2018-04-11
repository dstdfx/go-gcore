package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	getResourceResp = `{
    "client": 170,
    "cname": "cdn.site.com",
    "companyName": "Your Company",
    "created": "2018-04-09T11:31:40.000000Z",
    "deleted": false,
    "enabled": true,
    "id": 220,
    "origin": "example.com",
    "originGroup": 80,
    "secondaryHostnames": [
        "cdn1.yoursite.com",
        "cdn2.yoursite.com"
    ],
    "status": "active",
    "updated": "2018-04-09T11:32:31.000000Z"
}
     `
	listResourcesResp  = ``
	createResourceResp = ``

	getResourceExpected = &Resource{
		ID:          220,
		CName:       "cdn.site.com",
		Client:      170,
		CompanyName: "Your Company",
		Deleted:     false,
		Enabled:     true,
		Origin:      "example.com",
		OriginGroup: 80,
		SecondaryHostnames: []string{
			"cdn1.yoursite.com",
			"cdn2.yoursite.com",
		},
		Status:    "active",
		CreatedAt: NewGCoreTime(time.Date(2018, 4, 9,11, 31,40, 0, time.UTC)),
		UpdatedAt: NewGCoreTime(time.Date(2018, 4, 9,11, 32,31, 0, time.UTC)),
	}
)

func TestResourcesService_GetResource(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	Mux.HandleFunc(fmt.Sprintf(resourceURL, getResourceExpected.ID),
		func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getResourceResp))
	})


	client := GetAuthenticatedCommonClient()
	got, _, err := client.Resources.GetResource(context.Background(), getResourceExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, getResourceExpected) {
		t.Errorf("Expected: %+v, got %+v\n", getResourceExpected, got)
	}
}
