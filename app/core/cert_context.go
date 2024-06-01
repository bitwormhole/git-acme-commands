package core

import (
	"crypto/sha256"
	"crypto/x509"

	"github.com/bitwormhole/git-acme-commands/app/data/dxo"
	"github.com/starter-go/base/lang"
)

// ComputeCertificateFingerprint ... 计算证书指纹
func ComputeCertificateFingerprint(cert *x509.Certificate) dxo.Fingerprint {
	der := cert.Raw
	sum := sha256.Sum256(der)
	hex := lang.HexFromBytes(sum[:])
	return dxo.Fingerprint(hex)
}
