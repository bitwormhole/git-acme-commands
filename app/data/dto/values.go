package dto

import "github.com/starter-go/base/lang"

// PublicKeyFingerprint 表示公钥指纹
type PublicKeyFingerprint lang.Hex

// DomainName ...
type DomainName string

////////////////////////////////////////////////////////////////////////////////

func (dn DomainName) String() string {
	return string(dn)
}
