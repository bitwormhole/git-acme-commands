package ls

import (
	"bytes"
	"crypto/x509"
	"encoding/pem"

	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/starter-go/afs"
)

// LoadCertificateChain ...
func LoadCertificateChain(file afs.Path) (dto.CertChain, error) {
	loader := new(certChainLoader)
	return loader.Load(file)
}

// SaveCertificateChain ...
func SaveCertificateChain(chain dto.CertChain, file afs.Path) error {
	saver := new(certChainSaver)
	return saver.Save(chain, file)
}

////////////////////////////////////////////////////////////////////////////////

type certChainLoader struct {
}

func (inst *certChainLoader) Load(file afs.Path) (dto.CertChain, error) {
	list := make(dto.CertChain, 0)
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return nil, err
	}
	for {
		block, rest := pem.Decode(data)
		if block == nil {
			break
		}
		data = rest
		cer, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, err
		}
		list = append(list, cer)
	}
	return list, nil
}

////////////////////////////////////////////////////////////////////////////////

type certChainSaver struct {
}

func (inst *certChainSaver) Save(chain dto.CertChain, file afs.Path) error {
	buffer := bytes.NewBuffer(nil)
	for _, cer := range chain {
		der := cer.Raw
		block := new(pem.Block)
		block.Type = "CERTIFICATE"
		block.Bytes = der
		err := pem.Encode(buffer, block)
		if err != nil {
			return err
		}
	}
	opt := afs.ToCreateFile()
	if file.Exists() {
		opt = afs.ToWriteFile()
	}
	data := buffer.Bytes()
	return file.GetIO().WriteBinary(data, opt)
}

////////////////////////////////////////////////////////////////////////////////
