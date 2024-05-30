package ikeys

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/starter-go/keys"
)

// KeyManagerImpl ...
type KeyManagerImpl struct {

	//starter:component

	_as func(core.KeyManager) //starter:as("#")

	Drivers keys.DriverManager //starter:inject("#")
}

func (inst *KeyManagerImpl) _impl() core.KeyManager {
	return inst
}

// Exists ...
func (inst *KeyManagerImpl) Exists(ctx *core.ContainerContext, fingerprint dto.PublicKeyFingerprint) bool {
	h := inst.getKeyHolder(ctx, fingerprint)
	file := h.file
	return file.Exists()
}

// Find ...
func (inst *KeyManagerImpl) Find(ctx *core.ContainerContext, fingerprint dto.PublicKeyFingerprint) (core.KeyHolder, error) {
	h := inst.getKeyHolder(ctx, fingerprint)
	err := h.load()
	if err != nil {
		return nil, err
	}
	return h, nil
}

// CreateNew ...
func (inst *KeyManagerImpl) CreateNew(ctx *core.ContainerContext) (core.KeyHolder, error) {

	pk, err := rsa.GenerateKey(rand.Reader, 1024*2)
	if err != nil {
		return nil, err
	}

	fingerprint := computeFingerprint(pk)
	h := inst.getKeyHolder(ctx, fingerprint)
	h.pk = pk
	h.fingerprint = fingerprint

	err = h.save()
	if err != nil {
		return nil, err
	}

	return h, nil
}

func (inst *KeyManagerImpl) getKeyHolder(ctx *core.ContainerContext, fingerprint dto.PublicKeyFingerprint) *keyHolder {

	if fingerprint == "" {
		fingerprint = "0000"
	}

	layout := ctx.Parent.Layout
	repo := layout.Repository()
	file := repo.GetChild("keys/" + fingerprint.String() + ".key")

	h := &keyHolder{
		file:        file,
		fingerprint: fingerprint,
	}

	return h
}
