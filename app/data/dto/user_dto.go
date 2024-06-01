package dto

import "github.com/bitwormhole/git-acme-commands/app/data/dxo"

// User ...
type User struct {
	Name  string
	Email string
	Key   dxo.Fingerprint
}
