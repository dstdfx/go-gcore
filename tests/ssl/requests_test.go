package ssl

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/dstdfx/go-gcore/gcore"
	th "github.com/dstdfx/go-gcore/tests/testutils"
)

func TestSSLService_Add(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestCreateSSLExpected
	th.Mux.HandleFunc(gcore.CertificatesURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestCreateSSLResponse))
		})

	body := gcore.AddCertBody{
		Name:        "gcdn.example.me",
		Certificate: "-----BEGIN CERTIFICATE----\nMIIDXTCCAkWgAwIBAgIJAIQqKEM2sJZYMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTcwNzAzMDU1NjU1WhcNMTgwNzAzMDU1NjU1WjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50\n-----END CERTIFICATE-----",
		PrivateKey:  `-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZcNCZiNNHfX2O\ndZpf12mv2rAZwqGZBAdpox0wntEPK3JciQ7ZRloLJeHuCNIJs9MidnH7Xk8zveju\nmab6HmfIzvMJAAm88OYWMFQRiYe1ggJEHMe7yYPQbtXwTqWDYdWmjPPma3Ujqqmb\nhmVX2rsYILD7cUjS+e0Ucfqx3QODQj/aujTt1rS0gFhJ0soY5m+C6VimPCx4Bjyw\n5rhtskJDRrfXxrIhVXOvSPFRyxDSfjt3win8vjhhZ3oFPWgrl9lVhn0zaB5hjDsd\n-----END PRIVATE KEY-----\n`,
	}

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.Certificates.Add(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestSSLService_Get(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestGetSSLExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.CertificateURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestGetSSLResponse))
		})

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.Certificates.Get(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestSSLService_List(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestListSSLExpected
	th.Mux.HandleFunc(gcore.CertificatesURL,
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(TestListSSLResponse))
		})

	client := th.GetAuthenticatedCommonClient()
	got, _, err := client.Certificates.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestSSLService_Delete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.SetupGCoreAuthServer()

	expected := TestGetSSLExpected
	th.Mux.HandleFunc(fmt.Sprintf(gcore.CertificateURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		})

	client := th.GetAuthenticatedCommonClient()
	_, err := client.Certificates.Delete(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}
}
