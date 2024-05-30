package ls

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/starter-go/afs"
)

// 定义 pem.Block 类型
const (
	BlockTypeRSAPrivateKey = "RSA PRIVATE KEY"
)

// LoadPrivateKeyRSA ...
func LoadPrivateKeyRSA(file afs.Path) (*rsa.PrivateKey, error) {
	loader := new(rsaPrivateKeyLoader)
	return loader.load(file)
}

// SavePrivateKeyRSA ...
func SavePrivateKeyRSA(pk *rsa.PrivateKey, dst afs.Path) error {
	s := new(rsaPrivateKeySaver)
	return s.save(pk, dst)
}

////////////////////////////////////////////////////////////////////////////////

type rsaPrivateKeyLoader struct{}

func (inst *rsaPrivateKeyLoader) load(file afs.Path) (*rsa.PrivateKey, error) {
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return nil, err
	}
	for {
		b, rest := pem.Decode(data)
		if b == nil {
			break
		}
		if inst.isPrivateKeyBlock(b) {
			return inst.loadPrivateKeyFromBlock(b)
		}
		data = rest
	}
	return nil, fmt.Errorf("EOF: no " + BlockTypeRSAPrivateKey)
}

func (inst *rsaPrivateKeyLoader) loadPrivateKeyFromBlock(block *pem.Block) (*rsa.PrivateKey, error) {
	der := block.Bytes
	return x509.ParsePKCS1PrivateKey(der)
}

func (inst *rsaPrivateKeyLoader) isPrivateKeyBlock(block *pem.Block) bool {
	t := strings.ToLower(block.Type)
	b1 := strings.Contains(t, "rsa")
	b2 := strings.Contains(t, "private")
	b3 := strings.Contains(t, "key")
	return b1 && b2 && b3
}

////////////////////////////////////////////////////////////////////////////////

type rsaPrivateKeySaver struct{}

func (inst *rsaPrivateKeySaver) save(pk *rsa.PrivateKey, dst afs.Path) error {
	buffer := &bytes.Buffer{}
	der := x509.MarshalPKCS1PrivateKey(pk)
	block := &pem.Block{
		Type:  BlockTypeRSAPrivateKey,
		Bytes: der,
	}
	err := pem.Encode(buffer, block)
	if err != nil {
		return err
	}
	opt := afs.ToCreateFile()
	data := buffer.Bytes()
	return dst.GetIO().WriteBinary(data, opt)
}

////////////////////////////////////////////////////////////////////////////////
