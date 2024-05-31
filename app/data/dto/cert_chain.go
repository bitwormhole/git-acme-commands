package dto

import (
	"crypto/x509"
)

// CertChain 表示一个证书链
type CertChain []*x509.Certificate

// Leaf ...
func (cc CertChain) Leaf() *x509.Certificate {
	// 找出生效时间最晚的证书
	var cer *x509.Certificate
	list := cc
	for _, item := range list {
		if item == nil {
			continue
		}
		if cer == nil {
			cer = item
			continue
		}
		t1 := cer.NotBefore
		t2 := item.NotBefore
		if t2.After(t1) {
			cer = item
		}
	}
	return cer
}

// Chain ...
func (cc CertChain) Chain() []*x509.Certificate {
	return cc
}
