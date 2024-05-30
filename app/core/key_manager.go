package core

import (
	"crypto"

	"github.com/bitwormhole/git-acme-commands/app/data/dto"
)

// KeyHolder ...
type KeyHolder interface {
	Fingerprint() dto.PublicKeyFingerprint

	Signer() crypto.Signer
}

// KeyManager ...
type KeyManager interface {
	Exists(ctx *ContainerContext, fingerprint dto.PublicKeyFingerprint) bool

	Find(ctx *ContainerContext, fingerprint dto.PublicKeyFingerprint) (KeyHolder, error)

	CreateNew(ctx *ContainerContext) (KeyHolder, error)
}
