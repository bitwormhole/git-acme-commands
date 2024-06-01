package vo

import (
	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/bitwormhole/git-acme-commands/app/data/dxo"
)

// DomainConfig ...
type DomainConfig struct {
	Name         dto.DomainName
	Key          dxo.Fingerprint
	Debug        bool
	FetchFromURL string
}
