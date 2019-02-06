package gcore

import (
	"context"
	"testing"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

var TestFakeAuthOptions = AuthOptions{
	Username: "whatever",
	Password: "whatever",
}

func TestNewCommonClient(t *testing.T) {
	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	common := NewCommonClient()
	common.BaseURL = testEnv.GetServerURL()

	err := common.Authenticate(context.Background(), TestFakeAuthOptions)
	if err != nil {
		t.Fatal(err)
	}

	if common.Token.Value != th.TestFakeToken {
		t.Errorf("Expected: %s, got %s", th.TestFakeToken, common.Token.Value)
	}
}

func TestNewResellerClient(t *testing.T) {
	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	reseller := NewResellerClient()
	reseller.BaseURL = testEnv.GetServerURL()

	err := reseller.Authenticate(context.Background(), TestFakeAuthOptions)
	if err != nil {
		t.Fatal(err)
	}

	if reseller.Token.Value != th.TestFakeToken {
		t.Errorf("Expected: %s, got %s", th.TestFakeToken, reseller.Token.Value)
	}
}
