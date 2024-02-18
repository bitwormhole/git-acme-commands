package pecdsa

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/bitwormhole/git-acme-commands/app/keys"
	"github.com/starter-go/afs"
)

const (
	pemTypeName = "ECDSA PRIVATE KEY"
)

// Provider ...
type Provider struct {

	//starter:component

	_as func(keys.KeyPairProvider) //starter:as(".")

}

func (inst *Provider) _impl() keys.KeyPairProvider {
	return inst
}

// Info ...
func (inst *Provider) Info() *keys.KeyPairProviderInfo {
	return &keys.KeyPairProviderInfo{
		Algorithm: "ECDSA",
		Provider:  inst,
	}
}

func (inst *Provider) getFactory() *factory {
	return &factory{provider: inst}
}

// Generator ...
func (inst *Provider) Generator() keys.KeyPairGenerator {
	return inst.getFactory()
}

// Loader ...
func (inst *Provider) Loader() keys.KeyPairLoader {
	return inst.getFactory()
}

// Saver ...
func (inst *Provider) Saver() keys.KeyPairSaver {
	return inst.getFactory()
}

////////////////////////////////////////////////////////////////////////////////

type factory struct {
	provider *Provider
}

func (inst *factory) Generate() (keys.KeyPair, error) {
	c := elliptic.P256()
	r := rand.Reader
	key, err := ecdsa.GenerateKey(c, r)
	if err != nil {
		return nil, err
	}

	kp := &keypair{
		provider: inst.provider,
		raw:      key,
	}
	return kp, nil
}

func (inst *factory) loadAsPEM(file afs.Path) (*pem.Block, error) {
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return nil, err
	}
	for {
		block, rest := pem.Decode(data)
		if block == nil {
			break
		}
		if block.Type == pemTypeName {
			return block, nil
		}
		data = rest
	}
	const ptn = pemTypeName
	path := file.GetPath()
	return nil, fmt.Errorf("cannot find PEM block [%s] in file [%s]", ptn, path)
}

func (inst *factory) Load(file afs.Path) (keys.KeyPair, error) {

	b, err := inst.loadAsPEM(file)
	if err != nil {
		return nil, err
	}

	der := b.Bytes
	key, err := x509.ParseECPrivateKey(der)
	if err != nil {
		return nil, err
	}

	kp := &keypair{
		provider: inst.provider,
		raw:      key,
	}
	return kp, nil
}

func (inst *factory) Save(kp keys.KeyPair, file afs.Path) error {

	kp2, ok := kp.(*keypair)
	if !ok {
		return fmt.Errorf("bad ECDSA key pair")
	}

	key := kp2.raw
	der, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		return err
	}

	block := &pem.Block{
		Type:  pemTypeName,
		Bytes: der,
	}

	opt := afs.Todo().Create(true).File(true).FromBegin(true).Options()
	data := pem.EncodeToMemory(block)
	return file.GetIO().WriteBinary(data, opt)
}

////////////////////////////////////////////////////////////////////////////////

type keypair struct {
	raw      *ecdsa.PrivateKey
	provider *Provider
}

func (inst *keypair) _impl() keys.KeyPair {
	return inst
}

func (inst *keypair) Provider() keys.KeyPairProvider {
	return inst.provider
}

func (inst *keypair) Signer() crypto.Signer {
	return inst.raw
}

////////////////////////////////////////////////////////////////////////////////
