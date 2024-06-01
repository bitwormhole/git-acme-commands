package core

import (
	"crypto"

	"github.com/bitwormhole/git-acme-commands/app/data/dxo"
)

// KeyHolder ...
type KeyHolder interface {
	Fingerprint() dxo.Fingerprint

	Signer() crypto.Signer

	Algorithm() string
}

// KeyManager ...
type KeyManager interface {
	Exists(ctx *ContainerContext, fingerprint dxo.Fingerprint) bool

	Find(ctx *ContainerContext, fingerprint dxo.Fingerprint) (KeyHolder, error)

	CreateNew(ctx *ContainerContext) (KeyHolder, error)
}
