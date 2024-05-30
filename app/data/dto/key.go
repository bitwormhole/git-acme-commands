package dto

import "github.com/starter-go/base/lang"

func (k PublicKeyFingerprint) String() string {
	return string(k)
}

// Bytes ...
func (k PublicKeyFingerprint) Bytes() []byte {
	h := lang.Hex(k)
	return h.Bytes()
}
