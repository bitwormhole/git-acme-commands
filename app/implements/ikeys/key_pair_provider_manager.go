package ikeys

import (
	"fmt"
	"strings"

	"github.com/bitwormhole/git-acme-commands/app/keys"
)

// KeyPairProviderManagerImpl ...
type KeyPairProviderManagerImpl struct {

	//starter:component

	_as func(keys.KeyPairProviderManager) //starter:as("#")

	Providers []keys.KeyPairProvider //starter:inject(".")
}

func (inst *KeyPairProviderManagerImpl) _impl() keys.KeyPairProviderManager {
	return inst
}

// FindProvider ...
func (inst *KeyPairProviderManagerImpl) FindProvider(algorithm string) (keys.KeyPairProvider, error) {
	all := inst.Providers
	a1 := strings.ToLower(algorithm)
	for _, item := range all {
		info := item.Info()
		a2 := strings.ToLower(info.Algorithm)
		if a1 == a2 {
			return item, nil
		}
	}
	return nil, fmt.Errorf("cannot find provider with algorithm name: %s", a1)
}
