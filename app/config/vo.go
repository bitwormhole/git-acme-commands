package config

import (
	"encoding/json"
	"fmt"

	"github.com/starter-go/afs"
)

// VO ...
type VO struct {
	ACME     *ACME           `json:"acme"`
	Accounts []*AccountDTO   `json:"accounts"`
	Domains  []*DomainDTO    `json:"domains"`
	KeyPairs []*KeyPairDTO   `json:"keypairs"`
	Dirs     []*DirectoryDTO `json:"directories"`
}

// GetACME ...
func (inst *VO) GetACME() (*ACME, error) {
	p := inst.ACME
	if p == nil {
		return nil, fmt.Errorf("ACME info in config file is not found")
	}
	return p, nil
}

// FindAccount ...
func (inst *VO) FindAccount(name AccountName) (*AccountDTO, error) {
	list := inst.Accounts
	for _, item := range list {
		if item.Name == name {
			return item, nil
		}
	}
	return nil, fmt.Errorf("account with name [%s] is not found", name)
}

// FindDomain ...
func (inst *VO) FindDomain(name DomainName) (*DomainDTO, error) {
	list := inst.Domains
	for _, item := range list {
		if item.Name == name {
			return item, nil
		}
	}
	return nil, fmt.Errorf("domain with name [%s] is not found", name)
}

// FindKeyPair ...
func (inst *VO) FindKeyPair(name KeyPairName) (*KeyPairDTO, error) {
	list := inst.KeyPairs
	for _, item := range list {
		if item.Name == name {
			return item, nil
		}
	}
	return nil, fmt.Errorf("key-pair with name [%s] is not found", name)
}

// FindDirectory ...
func (inst *VO) FindDirectory(name DirectoryName) (*DirectoryDTO, error) {
	list := inst.Dirs
	for _, item := range list {
		if item.Name == name {
			return item, nil
		}
	}
	return nil, fmt.Errorf("directory with name [%s] is not found", name)
}

////////////////////////////////////////////////////////////////////////////////

// Save 把 obj 保存到文件
func Save(obj *VO, file afs.Path) error {

	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	opt := afs.Todo().Create(true).File(true).ResetLength(true).Options()
	err = file.GetIO().WriteBinary(data, opt)
	return err
}

// Load 从文件加载 obj
func Load(file afs.Path) (*VO, error) {

	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return nil, err
	}

	obj := new(VO)
	err = json.Unmarshal(data, obj)
	if err != nil {
		return nil, err
	}

	return obj, nil
}
