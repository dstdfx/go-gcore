package gcore

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	th "github.com/dstdfx/go-gcore/gcore/internal/testhelper"
)

// Fixtures
const (
	testCreateSSLRawResponse = `{
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
	testGetSSLRawResponse = `{
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
	testListSSLRawResponse = `[
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

const (
	testCreateSSLRawRequest = `{  
   "name":"gcdn.example.me",
   "sslCertificate":"-----BEGIN CERTIFICATE----\nMIIDXTCCAkWgAwIBAgIJAIQqKEM2sJZYMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTcwNzAzMDU1NjU1WhcNMTgwNzAzMDU1NjU1WjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50\n-----END CERTIFICATE-----",
   "sslPrivateKey":"-----BEGIN PRIVATE KEY-----\\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZcNCZiNNHfX2O\\ndZpf12mv2rAZwqGZBAdpox0wntEPK3JciQ7ZRloLJeHuCNIJs9MidnH7Xk8zveju\\nmab6HmfIzvMJAAm88OYWMFQRiYe1ggJEHMe7yYPQbtXwTqWDYdWmjPPma3Ujqqmb\\nhmVX2rsYILD7cUjS+e0Ucfqx3QODQj/aujTt1rS0gFhJ0soY5m+C6VimPCx4Bjyw\\n5rhtskJDRrfXxrIhVXOvSPFRyxDSfjt3win8vjhhZ3oFPWgrl9lVhn0zaB5hjDsd\\n-----END PRIVATE KEY-----\\n"
}`
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

func TestCertService_Add(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         certificatesURL,
		RawResponse: testCreateSSLRawResponse,
		RawRequest:  testCreateSSLRawRequest,
		Method:      http.MethodPost,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testCreateSSLExpected

	body := &AddCertBody{
		Name:        "gcdn.example.me",
		Certificate: "-----BEGIN CERTIFICATE----\nMIIDXTCCAkWgAwIBAgIJAIQqKEM2sJZYMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV\nBAYTAkFVMRMwEQYDVQQIDApTb21lLVN0YXRlMSEwHwYDVQQKDBhJbnRlcm5ldCBX\naWRnaXRzIFB0eSBMdGQwHhcNMTcwNzAzMDU1NjU1WhcNMTgwNzAzMDU1NjU1WjBF\nMQswCQYDVQQGEwJBVTETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50\n-----END CERTIFICATE-----",
		PrivateKey:  `-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDZcNCZiNNHfX2O\ndZpf12mv2rAZwqGZBAdpox0wntEPK3JciQ7ZRloLJeHuCNIJs9MidnH7Xk8zveju\nmab6HmfIzvMJAAm88OYWMFQRiYe1ggJEHMe7yYPQbtXwTqWDYdWmjPPma3Ujqqmb\nhmVX2rsYILD7cUjS+e0Ucfqx3QODQj/aujTt1rS0gFhJ0soY5m+C6VimPCx4Bjyw\n5rhtskJDRrfXxrIhVXOvSPFRyxDSfjt3win8vjhhZ3oFPWgrl9lVhn0zaB5hjDsd\n-----END PRIVATE KEY-----\n`,
	}

	got, _, err := client.Certificates.Add(context.Background(), body)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't add a certificate")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestCertService_Get(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(certificateURL, testGetSSLExpected.ID),
		RawResponse: testGetSSLRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testGetSSLExpected

	got, _, err := client.Certificates.Get(context.Background(), testGetSSLExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a certificate")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}

func TestCertService_Delete(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         fmt.Sprintf(certificateURL, testGetSSLExpected.ID),
		RawResponse: "",
		Method:      http.MethodDelete,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	_, err := client.Certificates.Delete(context.Background(), testGetSSLExpected.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't delete a certificate")
	}
}

func TestCertService_List(t *testing.T) {
	endpointCalled := false

	testEnv := th.SetupTestEnv()
	defer testEnv.TearDownTestEnv()

	handleOpts := &th.HandleReqOpts{
		Mux:         testEnv.Mux,
		URL:         certificatesURL,
		RawResponse: testListSSLRawResponse,
		Method:      http.MethodGet,
		Status:      http.StatusCreated,
		CallFlag:    &endpointCalled,
	}

	th.HandleReqWithoutBody(t, handleOpts)

	client := NewCommonClient()
	client.BaseURL = testEnv.GetServerURL()
	_ = client.Authenticate(context.Background(), TestFakeAuthOptions)

	expected := testListSSLExpected

	got, _, err := client.Certificates.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if !endpointCalled {
		t.Fatal("didn't get a list of certificates")
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected: %+v, got %+v\n", expected, got)
	}
}
