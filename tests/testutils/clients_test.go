package testutils

import (
	"context"
	"testing"

	"github.com/dstdfx/go-gcore/gcore"
)

func TestNewCommonClient(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	common := gcore.NewCommonClient()
	common.BaseURL = MockClientURL()

	err := common.Authenticate(context.Background(), FakeAuthOpts)
	if err != nil {
		t.Fatal(err)
	}

	if common.Token.Value != FakeToken {
		t.Errorf("Expected: %s, got %s", FakeToken, common.Token.Value)
	}
}

func TestNewResellerClient(t *testing.T) {
	SetupHTTP()
	defer TeardownHTTP()

	SetupGCoreAuthServer()

	reseller := gcore.NewResellerClient()
	reseller.BaseURL = MockClientURL()

	err := reseller.Authenticate(context.Background(), FakeAuthOpts)
	if err != nil {
		t.Fatal(err)
	}

	if reseller.Token.Value != FakeToken {
		t.Errorf("Expected: %s, got %s", FakeToken, reseller.Token.Value)
	}

}
