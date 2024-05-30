package vo

import "github.com/bitwormhole/git-acme-commands/app/data/dto"

// ContainerConfig ...
type ContainerConfig struct {
	ACME *dto.ACME

	User *dto.User

	DomainList []*dto.Domain
}
