package core

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	"github.com/bitwormhole/git-acme-commands/app/data/dxo"
	"github.com/starter-go/base/lang"
)

// ComputePublicKeyFingerprint ... 计算公钥指纹
func ComputePublicKeyFingerprint(pk crypto.PublicKey) dxo.Fingerprint {
	hex := computePublicKeyFingerprintAny(pk)
	return dxo.Fingerprint(hex)
}

func computePublicKeyFingerprintAny(pk crypto.PublicKey) lang.Hex {
	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		// vlog.Error(err.Error())
		// return "0000"
		panic(err)
	}
	sum := sha256.Sum256(der)
	return lang.HexFromBytes(sum[:])
}

func computePublicKeyFingerprintRSA(pk *rsa.PublicKey) lang.Hex {
	der := x509.MarshalPKCS1PublicKey(pk)
	sum := sha256.Sum256(der)
	return lang.HexFromBytes(sum[:])
}
