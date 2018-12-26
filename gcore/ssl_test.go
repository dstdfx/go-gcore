package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

// Fixtures
var (
	testCreateSSLResponse = `{
   "deleted" : false,
   "cert_subject_alt" : null,
   "hasRelatedResources" : true,
   "validity_not_after" : "2018-07-12T16:01:59Z",
   "id" : 1189,
   "sslCertificateChain" : "",
   "validity_not_before" : "2018-04-13T16:01:59Z",
   "name" : "gcdn.example.me",
   "cert_issuer" : "Let's Encrypt Authority X3",
   "cert_subject_cn" : "gcdn.example.me"
}`
	testGetSSLResponse = `{
   "deleted" : false,
   "cert_subject_alt" : null,
   "hasRelatedResources" : true,
   "validity_not_after" : "2018-07-12T16:01:59Z",
   "id" : 1189,
   "sslCertificateChain" : "",
   "validity_not_before" : "2018-04-13T16:01:59Z",
   "name" : "gcdn.example.me",
   "cert_issuer" : "Let's Encrypt Authority X3",
   "cert_subject_cn" : "gcdn.example.me"
}`
	testListSSLResponse = `[
   {
      "validity_not_before" : "2018-04-13T16:01:59Z",
      "validity_not_after" : "2018-07-12T16:01:59Z",
      "cert_issuer" : "Let's Encrypt Authority X3",
      "hasRelatedResources" : true,
      "name" : "gcdn.example.me",
      "sslCertificateChain" : "",
      "cert_subject_alt" : null,
      "cert_subject_cn" : "gcdn.example.me",
      "id" : 1189,
      "deleted" : false
   }
]`
)

var (
	testCreateSSLExpected = &CertSSL{
		ID:                  1189,
		Name:                "gcdn.example.me",
		Deleted:             false,
		CertSubjectAlt:      nil,
		CertSubjectCn:       "gcdn.example.me",
		HasRelatedResources: true,
		CertificateChain:    "",
		CertIssuer:          "Let's Encrypt Authority X3",
		ValidityNotAfter:    NewTime(time.Date(2018, time.July, 12, 16, 1, 59, 0, time.UTC)),
		ValidityNotBefore:   NewTime(time.Date(2018, time.April, 13, 16, 1, 59, 0, time.UTC)),
	}
	testGetSSLExpected = &CertSSL{
		ID:                  1189,
		Name:                "gcdn.example.me",
		Deleted:             false,
		CertSubjectAlt:      nil,
		CertSubjectCn:       "gcdn.example.me",
		HasRelatedResources: true,
		CertificateChain:    "",
		CertIssuer:          "Let's Encrypt Authority X3",
		ValidityNotAfter:    NewTime(time.Date(2018, time.July, 12, 16, 1, 59, 0, time.UTC)),
		ValidityNotBefore:   NewTime(time.Date(2018, time.April, 13, 16, 1, 59, 0, time.UTC)),
	}
	testListSSLExpected = []*CertSSL{
		{
			ID:                  1189,
			Name:                "gcdn.example.me",
			Deleted:             false,
			CertSubjectAlt:      nil,
			CertSubjectCn:       "gcdn.example.me",
			HasRelatedResources: true,
			CertificateChain:    "",
			CertIssuer:          "Let's Encrypt Authority X3",
			ValidityNotAfter:    NewTime(time.Date(2018, time.July, 12, 16, 1, 59, 0, time.UTC)),
			ValidityNotBefore:   NewTime(time.Date(2018, time.April, 13, 16, 1, 59, 0, time.UTC)),
		},
	}
)

func TestSSLService_Add(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testCreateSSLExpected
	mux.HandleFunc(certificatesURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testCreateSSLResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	body := AddCertBody{
		Name:        "gcdn.example.me",
		Certificate: "-----BEGIN CERTIFICATE----\nMIIDXTCCAkWgAwIBAgIJAIQqKEM2sJZYMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTcwNzAzMDU1NjU1WhcNMTgwNzAzMDU1NjU1WjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50\n-----END CERTIFICATE-----",
		PrivateKey:  `-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZcNCZiNNHfX2O\ndZpf12mv2rAZwqGZBAdpox0wntEPK3JciQ7ZRloLJeHuCNIJs9MidnH7Xk8zveju\nmab6HmfIzvMJAAm88OYWMFQRiYe1ggJEHMe7yYPQbtXwTqWDYdWmjPPma3Ujqqmb\nhmVX2rsYILD7cUjS+e0Ucfqx3QODQj/aujTt1rS0gFhJ0soY5m+C6VimPCx4Bjyw\n5rhtskJDRrfXxrIhVXOvSPFRyxDSfjt3win8vjhhZ3oFPWgrl9lVhn0zaB5hjDsd\n-----END PRIVATE KEY-----\n`,
	}

	client := getAuthenticatedCommonClient()
	got, _, err := client.Certificates.Add(context.Background(), &body)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestSSLService_Get(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testGetSSLExpected
	mux.HandleFunc(fmt.Sprintf(certificateURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testGetSSLResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := getAuthenticatedCommonClient()
	got, _, err := client.Certificates.Get(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestSSLService_List(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testListSSLExpected
	mux.HandleFunc(certificatesURL,
		func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte(testListSSLResponse))
			if err != nil {
				t.Fatal(err)
			}
		})

	client := getAuthenticatedCommonClient()
	got, _, err := client.Certificates.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestSSLService_Delete(t *testing.T) {
	setupHTTP()
	defer teardownHTTP()

	setupGCoreAuthServer()

	expected := testGetSSLExpected
	mux.HandleFunc(fmt.Sprintf(certificateURL, expected.ID),
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
		})

	client := getAuthenticatedCommonClient()
	_, err := client.Certificates.Delete(context.Background(), expected.ID)
	if err != nil {
		t.Fatal(err)
	}
}
