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
	"github.com/libdns/libdns"
	"github.com/mholt/acmez"
	"github.com/mholt/acmez/acme"
	"go.uber.org/zap"
	"golang.org/x/net/idna"
)

// Request 。。。
type Request struct {
	Domains            []string
	ACMEAddress        string
	InsecureSkipVerify bool
	Email              string

	AccountSinger crypto.Signer
	CertSinger    crypto.Signer

	DoNewAccount bool
	DoMakeCert   bool

	// reseults
	ResultContentType string
	ResultContentData []byte
}

// ForTest ...
func (inst *Request) ForTest() *Request {

	inst.ACMEAddress = "https://acme-staging-v02.api.letsencrypt.org/directory"
	inst.Email = "you@example.com"
	inst.InsecureSkipVerify = true
	inst.Domains = []string{"test.bitwormhole.com"}

	return inst
}

// ForProduction ...
func (inst *Request) ForProduction() *Request {

	inst.ACMEAddress = "https://acme-v02.api.letsencrypt.org/directory"
	inst.Email = "you@example.com"
	inst.InsecureSkipVerify = false
	inst.Domains = []string{"test.bitwormhole.com"}

	return inst
}

// Run ...
func (inst *Request) Run() error {
	return highLevelExample(inst)
}

////////////////////////////////////////////////////////////////////////////////

// Run pebble (the ACME server) before running this example:
//
// PEBBLE_VA_ALWAYS_VALID=1 pebble -config ./test/config/pebble-config.json -strict

// func main() {
// 	err := highLevelExample()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func highLevelExample(wl *Request) error {
	// Put your domains here
	// domains := []string{"example.com"}
	domains := wl.Domains

	// A context allows us to cancel long-running ops
	ctx := context.Background()

	// Logging is important - replace with your own zap logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}

	// minimal example using Cloudflare
	solver := &certmagic.DNS01Solver{
		// DNSProvider: &cloudflare.Provider{APIToken: "topsecret"},
		DNSProvider: &myDNSProvider{},
	}

	// A high-level client embeds a low-level client and makes
	// the ACME flow much easier, but with less flexibility
	// than using the low-level API directly (see other example).
	//
	// One thing you will have to do is provide challenge solvers
	// for all the challenge types you wish to support. I recommend
	// supporting as many as possible in case there are errors. The
	// library will try all enabled challenge types, and certain
	// external factors can cause certain challenge types to fail,
	// where others might still succeed.
	//
	// Implementing challenge solvers is outside the scope of this
	// example, but you can find a high-quality, general-purpose
	// solver for the dns-01 challenge in CertMagic:
	// https://pkg.go.dev/github.com/caddyserver/certmagic#DNS01Solver
	client := acmez.Client{
		Client: &acme.Client{
			Directory: wl.ACMEAddress, // "https://127.0.0.1:14000/dir", // default pebble endpoint
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{
						InsecureSkipVerify: wl.InsecureSkipVerify, // true, // REMOVE THIS FOR PRODUCTION USE!
					},
				},
			},
			Logger: logger,
		},
		ChallengeSolvers: map[string]acmez.Solver{
			// acme.ChallengeTypeHTTP01:    mySolver{}, // provide these!
			// acme.ChallengeTypeTLSALPN01: mySolver{}, // provide these!
			acme.ChallengeTypeDNS01: solver, // provide these!
		},
	}

	// Before you can get a cert, you'll need an account registered with
	// the ACME CA; it needs a private key which should obviously be
	// different from any key used for certificates!

	// accountPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// if err != nil {
	// 	return fmt.Errorf("generating account key: %v", err)
	// }
	accountPrivateKey := wl.AccountSinger

	account := acme.Account{
		Contact:              []string{"mailto:" + wl.Email},
		TermsOfServiceAgreed: true,
		PrivateKey:           accountPrivateKey,
	}

	// If the account is new, we need to create it; only do this once!
	// then be sure to securely store the account key and metadata so
	// you can reuse it later!

	if wl.DoNewAccount {
		account, err = client.NewAccount(ctx, account)
		if err != nil {
			return fmt.Errorf("new account: %v", err)
		}
	}

	if !wl.DoMakeCert {
		return nil
	}

	// Every certificate needs a key.
	// certPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// if err != nil {
	// 	return fmt.Errorf("generating certificate key: %v", err)
	// }
	certPrivateKey := wl.CertSinger

	// Once your client, account, and certificate key are all ready,
	// it's time to request a certificate! The easiest way to do this
	// is to use ObtainCertificate() and pass in your list of domains
	// that you want on the cert. But if you need more flexibility, you
	// should create a CSR yourself and use ObtainCertificateUsingCSR().

	csr, err := makeCSR(domains, certPrivateKey)
	if err != nil {
		return err
	}

	// certs, err := client.ObtainCertificate(ctx, account, certPrivateKey, domains)
	certs, err := client.ObtainCertificateUsingCSR(ctx, account, csr)
	if err != nil {
		return fmt.Errorf("obtaining certificate: %v", err)
	}

	// ACME servers should usually give you the entire certificate chain
	// in PEM format, and sometimes even alternate chains! It's up to you
	// which one(s) to store and use, but whatever you do, be sure to
	// store the certificate and key somewhere safe and secure, i.e. don't
	// lose them!
	for _, cert := range certs {
		// fmt.Printf("Certificate %q:\n%s\n\n", cert.URL, cert.ChainPEM)

		wl.ResultContentType = "application/x-pem-cer"
		wl.ResultContentData = cert.ChainPEM
	}

	return nil
}

func makeCSR(domains []string, key crypto.Signer) (*x509.CertificateRequest, error) {

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

// // mySolver is a no-op acmez.Solver for example purposes only.
// type mySolver struct{}

// func (s mySolver) Present(ctx context.Context, chal acme.Challenge) error {
// 	log.Printf("[DEBUG] present: %#v", chal)
// 	return nil
// }

// func (s mySolver) CleanUp(ctx context.Context, chal acme.Challenge) error {
// 	log.Printf("[DEBUG] cleanup: %#v", chal)
// 	return nil
// }

////////////////////////////////////////////////////////////////////////////////

type myDNSProvider struct{}

func (inst *myDNSProvider) _impl() certmagic.ACMEDNSProvider {
	return inst
}

func (inst *myDNSProvider) AppendRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {

	for _, item := range recs {
		str := inst.formatRecord(zone, &item)
		fmt.Println("请 [添加] 域名解析记录：", str)
	}

	err := inst.waitForInput(ctx, "ok", "  请输入 [ok] 以便继续 ...")
	if err != nil {
		return nil, err
	}

	return recs, nil
}

func (inst *myDNSProvider) DeleteRecords(ctx context.Context, zone string, recs []libdns.Record) ([]libdns.Record, error) {

	for _, item := range recs {
		str := inst.formatRecord(zone, &item)
		fmt.Println("请 [删除] 域名解析记录：", str)
	}

	err := inst.waitForInput(ctx, "ok", "  请输入 [ok] 以便继续 ...")
	if err != nil {
		return nil, err
	}

	return recs, nil
}

func (inst *myDNSProvider) formatRecord(zone string, item *libdns.Record) string {
	name := item.Name
	typ := item.Type
	value := item.Value
	const f = "[dns.record zone:'%s' name:'%s' type:'%s' value:'%s']"
	return fmt.Sprintf(f, zone, name, typ, value)
}

func (inst *myDNSProvider) waitForInput(ctx context.Context, token string, tip string) error {

	bypass := true

	for {
		var text string
		fmt.Println(tip)

		if bypass {
			break
		}

		_, err := fmt.Scan(&text)
		if err != nil {
			return err
		}
		if text == token {
			return nil
		}

	}
	return nil
}
