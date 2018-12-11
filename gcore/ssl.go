package gcore

import (
	"context"
	"fmt"
	"net/http"
)

const (
	CertificatesURL = "/sslData"
	CertificateURL  = "/sslData/%d"
)

// CertService handles communication with the SSL/TLS certificate related methods
// of the G-Core CDN API.
type CertService service

// CertSSL represents G-Core's CertSSL certificate
type CertSSL struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	CName               string  `json:"cname"`
	Deleted             bool    `json:"deleted"`
	CertSubjectAlt      *string `json:"cert_subject_alt"`
	HasRelatedResources bool    `json:"hasRelatedResources"`
	ValidityNotAfter    *Time   `json:"validity_not_after"`
	ValidityNotBefore   *Time   `json:"validity_not_before"`
	CertificateChain    string  `json:"sslCertificateChain"`
	CertIssuer          string  `json:"cert_issuer"`
	CertSubjectCn       string  `json:"cert_subject_cn"`
}

// AddCertBody represents SSL certificate body for add certificate.
type AddCertBody struct {
	Name        string `json:"name"`
	Certificate string `json:"sslCertificate"`
	PrivateKey  string `json:"sslPrivateKey"`
}

// List returns list of all SSL certificates.
func (s *CertService) List(ctx context.Context) ([]*CertSSL, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, CertificatesURL, nil)
	if err != nil {
		return nil, nil, err
	}

	certs := make([]*CertSSL, 0)

	resp, err := s.client.Do(req, &certs)
	if err != nil {
		return nil, resp, err
	}

	return certs, resp, nil
}

// Get method returns specific SSL certificate.
func (s *CertService) Get(ctx context.Context, certID int) (*CertSSL, *http.Response, error) {
	req, err := s.client.NewRequest(ctx,
		http.MethodGet,
		fmt.Sprintf(CertificateURL, certID), nil)
	if err != nil {
		return nil, nil, err
	}

	cert := &CertSSL{}

	resp, err := s.client.Do(req, cert)
	if err != nil {
		return nil, resp, err
	}

	return cert, resp, nil
}

// Delete method deletes specific SSL certificate.
func (s *CertService) Delete(ctx context.Context, certID int) (*http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodDelete, fmt.Sprintf(CertificateURL, certID), nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Add method adds SSL certificate to deliver content via HTTPS protocol.
// Paste all strings of the certificate(s) and the private key in one string parameter.
// Each certificate in chain and the private key should be divided with the "\n" symbol.
func (s *CertService) Add(ctx context.Context, body *AddCertBody) (*CertSSL, *http.Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, CertificatesURL, body)
	if err != nil {
		return nil, nil, err
	}

	cert := &CertSSL{}

	resp, err := s.client.Do(req, cert)
	if err != nil {
		return nil, resp, err
	}

	return cert, resp, nil
}
