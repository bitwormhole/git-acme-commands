package config

import "github.com/starter-go/base/lang"

// AccountName ...
type AccountName string

// DomainName ...
type DomainName string

// KeyPairName ...
type KeyPairName string

// DirectoryName ...
type DirectoryName string

////////////////////////////////////////////////////////////////////////////////

// ACME ...
type ACME struct {
	Account  AccountName       `json:"account"`
	Interval lang.Milliseconds `json:"interval"`
}

// AccountDTO ...
type AccountDTO struct {
	Name    AccountName `json:"name"`
	KeyPair KeyPairName `json:"keypair"`
	Email   string      `json:"email"`
	URL     string      `json:"url"`
}

// DomainDTO ...
type DomainDTO struct {
	Name    DomainName  `json:"name"`
	KeyPair KeyPairName `json:"keypair"`
}

// KeyPairDTO ...
type KeyPairDTO struct {
	Name      KeyPairName   `json:"name"`
	Algorithm string        `json:"algorithm"`
	FileName  string        `json:"file"`
	Directory DirectoryName `json:"directory"`
}

// DirectoryDTO ...
type DirectoryDTO struct {
	Name DirectoryName `json:"name"`
	Path string        `json:"path"`
}
