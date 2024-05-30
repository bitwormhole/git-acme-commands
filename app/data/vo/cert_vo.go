package vo

import "crypto/x509"

// CertChain ...
type CertChain struct {
	certs []*x509.Certificate
}
