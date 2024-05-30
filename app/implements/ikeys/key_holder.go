package ikeys

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"

	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/bitwormhole/git-acme-commands/app/data/ls"
	"github.com/starter-go/afs"
)

func computeFingerprint(pk *rsa.PrivateKey) dto.PublicKeyFingerprint {
	pub := pk.PublicKey
	der := x509.MarshalPKCS1PublicKey(&pub)
	sum := sha1.Sum(der)
	str := hex.EncodeToString(sum[:])
	return dto.PublicKeyFingerprint(str)
}

////////////////////////////////////////////////////////////////////////////////

type keyHolder struct {
	file        afs.Path
	fingerprint dto.PublicKeyFingerprint
	pk          *rsa.PrivateKey
}

func (inst *keyHolder) _impl() core.KeyHolder {
	return inst
}

func (inst *keyHolder) load() error {
	pk, err := ls.LoadPrivateKeyRSA(inst.file)
	if err != nil {
		return err
	}
	inst.pk = pk
	return nil
}

func (inst *keyHolder) save() error {
	return ls.SavePrivateKeyRSA(inst.pk, inst.file)
}

func (inst *keyHolder) Fingerprint() dto.PublicKeyFingerprint {
	return inst.fingerprint
}

func (inst *keyHolder) Signer() crypto.Signer {
	return inst.pk
}
