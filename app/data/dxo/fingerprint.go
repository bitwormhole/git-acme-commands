package dxo

import "github.com/starter-go/base/lang"

// Fingerprint 表示(公钥|证书)指纹
type Fingerprint lang.Hex

func (f Fingerprint) String() string {
	return string(f)
}

// Bytes ...
func (f Fingerprint) Bytes() []byte {
	h := lang.Hex(f)
	return h.Bytes()
}

////////////////////////////////////////////////////////////////////////////////
