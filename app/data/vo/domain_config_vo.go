package vo

import "github.com/bitwormhole/git-acme-commands/app/data/dto"

// DomainConfig ...
type DomainConfig struct {
	Name         dto.DomainName
	Key          dto.PublicKeyFingerprint
	Debug        bool
	FetchFromURL string
}
