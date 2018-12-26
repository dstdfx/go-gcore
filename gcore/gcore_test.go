package gcore

import (
	"context"
	"testing"
)

func TestNewCommonClient(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	common := NewCommonClient()
	common.BaseURL = mockClientURL()

	err := common.Authenticate(context.Background(), fakeAuthOpts)
	if err != nil {
		t.Fatal(err)
	}

	if common.Token.Value != fakeToken {
		t.Errorf("Expected: %s, got %s", fakeToken, common.Token.Value)
	}
}

func TestNewResellerClient(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	reseller := NewResellerClient()
	reseller.BaseURL = mockClientURL()

	err := reseller.Authenticate(context.Background(), fakeAuthOpts)
	if err != nil {
		t.Fatal(err)
	}

	if reseller.Token.Value != fakeToken {
		t.Errorf("Expected: %s, got %s", fakeToken, reseller.Token.Value)
	}

}
