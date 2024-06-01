package ikeys

import (
	"crypto"
	"crypto/rsa"

	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/dxo"
	"github.com/bitwormhole/git-acme-commands/app/data/ls"
	"github.com/starter-go/afs"
)

////////////////////////////////////////////////////////////////////////////////

type keyHolder struct {
	file        afs.Path
	fingerprint dxo.Fingerprint
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

func (inst *keyHolder) Fingerprint() dxo.Fingerprint {
	return inst.fingerprint
}

func (inst *keyHolder) Signer() crypto.Signer {
	return inst.pk
}

func (inst *keyHolder) Algorithm() string {
	return "RSA"
}
