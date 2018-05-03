package ssl

import (
	"time"

	"github.com/dstdfx/go-gcore/gcore"
)

var (
	TestCreateSSLResponse = `{
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
	TestGetSSLResponse = `{
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
	TestListSSLResponse = `[
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
	TestCreateSSLExpected = &gcore.CertSSL{
		ID:                  1189,
		Name:                "gcdn.example.me",
		Deleted:             false,
		CertSubjectAlt:      nil,
		CertSubjectCn:       "gcdn.example.me",
		HasRelatedResources: true,
		CertificateChain:    "",
		CertIssuer:          "Let's Encrypt Authority X3",
		ValidityNotAfter:    gcore.NewGCoreTime(time.Date(2018, 7, 12, 16, 1, 59, 0, time.UTC)),
		ValidityNotBefore:   gcore.NewGCoreTime(time.Date(2018, 4, 13, 16, 1, 59, 0, time.UTC)),
	}
	TestGetSSLExpected = &gcore.CertSSL{
		ID:                  1189,
		Name:                "gcdn.example.me",
		Deleted:             false,
		CertSubjectAlt:      nil,
		CertSubjectCn:       "gcdn.example.me",
		HasRelatedResources: true,
		CertificateChain:    "",
		CertIssuer:          "Let's Encrypt Authority X3",
		ValidityNotAfter:    gcore.NewGCoreTime(time.Date(2018, 7, 12, 16, 1, 59, 0, time.UTC)),
		ValidityNotBefore:   gcore.NewGCoreTime(time.Date(2018, 4, 13, 16, 1, 59, 0, time.UTC)),
	}
	TestListSSLExpected = []*gcore.CertSSL{
		{
			ID:                  1189,
			Name:                "gcdn.example.me",
			Deleted:             false,
			CertSubjectAlt:      nil,
			CertSubjectCn:       "gcdn.example.me",
			HasRelatedResources: true,
			CertificateChain:    "",
			CertIssuer:          "Let's Encrypt Authority X3",
			ValidityNotAfter:    gcore.NewGCoreTime(time.Date(2018, 7, 12, 16, 1, 59, 0, time.UTC)),
			ValidityNotBefore:   gcore.NewGCoreTime(time.Date(2018, 4, 13, 16, 1, 59, 0, time.UTC)),
		},
	}
)
