package core

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestPublicKeyFingerprint(t *testing.T) {

	bits := 1024 * 2
	pri, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		t.Error(err)
		return
	}

	pub := &pri.PublicKey

	f1 := computePublicKeyFingerprintAny(pub)
	f2 := computePublicKeyFingerprintRSA(pub)

	t.Logf("rsa.publickey.fingerprint.1=%s", f1)
	t.Logf("rsa.publickey.fingerprint.2=%s", f2)

}
