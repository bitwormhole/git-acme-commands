package acme

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"

	"github.com/caddyserver/certmagic"
	"github.com/mholt/acmez"
	"github.com/mholt/acmez/acme"
	"go.uber.org/zap"
	"golang.org/x/net/idna"
)

// RequestV2 ...
type RequestV2 struct {
	ServiceURL   string
	DoNewAccount bool
	DoMakeCert   bool
	DoSkipVerify bool

	UserEmail  string
	UserSigner crypto.Signer

	DomainName   string
	DomainSigner crypto.Signer
}

// PrepareForTest ...
func (inst *RequestV2) PrepareForTest() *RequestV2 {

	inst.ServiceURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
	inst.UserEmail = "test@example.com"
	inst.DomainName = "test.bitwormhole.com"
	inst.DoSkipVerify = true

	return inst
}

// PrepareForProduction ...
func (inst *RequestV2) PrepareForProduction() *RequestV2 {

	inst.ServiceURL = "https://acme-v02.api.letsencrypt.org/directory"
	inst.UserEmail = "todo@example.com"
	inst.DomainName = "todo.bitwormhole.com"
	inst.DoSkipVerify = false

	return inst
}

// Send ...
func (inst *RequestV2) Send() (*ResponseV2, error) {

	domains := []string{inst.DomainName}
	ctx := context.Background()

	// prepare logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}

	// prepare solver
	solver := &certmagic.DNS01Solver{
		DNSProvider: &myDNSProvider{},
	}

	// prepare client
	client := acmez.Client{
		Client: &acme.Client{
			Directory: inst.ServiceURL,
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: inst.DoSkipVerify,
					},
				},
			},
			Logger: logger,
		},
		ChallengeSolvers: map[string]acmez.Solver{
			acme.ChallengeTypeDNS01: solver,
		},
	}

	// prepare account
	account := acme.Account{
		Contact:              []string{"mailto:" + inst.UserEmail},
		TermsOfServiceAgreed: true,
		PrivateKey:           inst.UserSigner,
	}

	if inst.DoNewAccount {
		account, err = client.NewAccount(ctx, account)
		if err != nil {
			return nil, fmt.Errorf("new account: %v", err)
		}
	}

	if !inst.DoMakeCert {
		return nil, nil
	}

	// prepare csr
	csr, err := inst.makeCSR(domains)
	if err != nil {
		return nil, err
	}

	// send request
	certs, err := client.ObtainCertificateUsingCSR(ctx, account, csr)
	if err != nil {
		return nil, fmt.Errorf("obtaining certificate: %v", err)
	}

	// make result
	for _, cert := range certs {
		return inst.makeResult(&cert)
	}
	return nil, fmt.Errorf("no result cert")
}

func (inst *RequestV2) makeResult(src *acme.Certificate) (*ResponseV2, error) {
	resp := new(ResponseV2)
	resp.Certificate = *src
	return resp, nil
}

func (inst *RequestV2) makeCSR(domains []string) (*x509.CertificateRequest, error) {

	key := inst.DomainSigner
	template := new(x509.CertificateRequest)

	template.Subject.Country = []string{"CN"}
	template.Subject.Province = []string{"guangxi"}
	template.Subject.Locality = []string{"guilin"}
	template.Subject.Organization = []string{"bitwormhole"}
	template.Subject.OrganizationalUnit = []string{"IT"}

	for _, name := range domains {
		normalizedName, err := idna.ToASCII(name)
		if err != nil {
			return nil, fmt.Errorf("converting identifier '%s' to ASCII: %v", name, err)
		}
		template.DNSNames = append(template.DNSNames, normalizedName)
	}

	der, err := x509.CreateCertificateRequest(rand.Reader, template, key)
	if err != nil {
		return nil, err
	}

	return x509.ParseCertificateRequest(der)
}

////////////////////////////////////////////////////////////////////////////////

// ResponseV2 ...
type ResponseV2 struct {
	acme.Certificate
}

// ToX509 ...
func (inst *ResponseV2) ToX509() ([]*x509.Certificate, error) {

	return nil, fmt.Errorf("no impl: ResponseV2.ToX509")
}
