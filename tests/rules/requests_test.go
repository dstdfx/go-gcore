package rules

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-gcore/gcore"
	th "github.com/dstdfx/go-gcore/tests/testutils"
)

func TestRulesService_Create(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestCreateRuleExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.RulesURL, FakeResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestCreateRuleResponse))
		})

	body := gcore.CreateRuleBody{
		RuleType: 0,
		Name:     "whatever",
		Options: gcore.Options{
			CacheHTTPHeaders: &gcore.CacheHTTPHeaders{
				Enabled: true,
				Value:   []string{"x-token"},
			},
		},
	}

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.Rules.Create(context.Background(), FakeResourceID, &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_Get(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestGetRuleExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.RuleURL, FakeResourceID, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestGetRuleResponse))
		})

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.Rules.Get(context.Background(), FakeResourceID, expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_List(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestListRuleExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.RulesURL, FakeResourceID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestListRuleResponse))
		})

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.Rules.List(context.Background(), FakeResourceID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestRulesService_Delete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestGetRuleExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.RuleURL, FakeResourceID, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		})

	client := th.GetAuthenticatedCommonClient()
	_, err := client.Rules.Delete(context.Background(), FakeResourceID, expected.ID)
	if err != nil {
		t.Fatal(err)
	}
}
