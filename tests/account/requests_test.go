package account

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-gcore/gcore"
	th "github.com/dstdfx/go-gcore/tests/testutils"
)

func TestAccountService_Details(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	th.Mux.HandleFunc(gcore.AccountDetailsURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(TestAccountDetailResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := th.GetAuthenticatedCommonClient()

	expected := TestAccountDetailExpected
	got, _, err := client.Account.Details(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
