package certs

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/dxo"
)

// DisplayCertificateDetail ... 显示证书的详细信息
func DisplayCertificateDetail(ctx *core.CertificateContext) error {
	task := &displayCertificateDetailTask{}
	return task.run(ctx)
}

type displayCertificateDetailTask struct {
	cc *core.CertificateContext

	certKey core.KeyHolder

	sha256CertificateFingerprint dxo.Fingerprint
	sha256PublicKeyFingerprint   dxo.Fingerprint
}

func (inst *displayCertificateDetailTask) run(cc *core.CertificateContext) error {

	inst.cc = cc
	steplist := make([]func() error, 0)

	steplist = append(steplist, inst.computeCertificateFingerprint)
	steplist = append(steplist, inst.computePublicKeyFingerprint)
	steplist = append(steplist, inst.loadPrivateKey)

	steplist = append(steplist, inst.displayChain)
	steplist = append(steplist, inst.displayPrivateKey)
	steplist = append(steplist, inst.displayCert)
	steplist = append(steplist, inst.displayOverview)
	steplist = append(steplist, inst.displayMeta)

	for _, step := range steplist {
		err := step()
		if err != nil {
			return err
		}
	}

	inst.displayBar("# END ")
	return nil
}

func (inst *displayCertificateDetailTask) loadPrivateKey() error {

	cc := inst.cc
	fingerprint := inst.sha256PublicKeyFingerprint
	containerCtx := cc.Parent.Parent
	km := containerCtx.KeyManager

	h, err := km.Find(containerCtx, fingerprint)
	if err != nil {
		return err
	}

	cc.CertSigner = h.Signer()
	inst.certKey = h
	return nil
}

func (inst *displayCertificateDetailTask) computeCertificateFingerprint() error {
	cert := inst.cc.Cert
	fp := core.ComputeCertificateFingerprint(cert)
	inst.sha256CertificateFingerprint = fp
	return nil
}

func (inst *displayCertificateDetailTask) computePublicKeyFingerprint() error {
	cert := inst.cc.Cert
	pub := cert.PublicKey
	fp := core.ComputePublicKeyFingerprint(pub)
	inst.sha256PublicKeyFingerprint = fp
	return nil
}

func (inst *displayCertificateDetailTask) displayOverview() error {
	inst.displayBar("# overview ")

	cert := inst.cc.Cert

	t1 := cert.NotBefore
	t2 := cert.NotAfter
	dnlist := cert.DNSNames
	issuer := cert.Issuer
	subject := cert.Subject

	fmt.Println("  From-Time: ", t1.Format(time.DateOnly))
	fmt.Println("    To-Time: ", t2.Format(time.DateOnly))
	fmt.Println("     Issuer: ", issuer.String())
	fmt.Println("    Subject: ", subject.String())
	fmt.Println("  DNS-Names: ", dnlist)

	fmt.Println("")

	fmt.Println("  SHA256-Certificate-Fingerprint: ", inst.sha256CertificateFingerprint)
	fmt.Println("  SHA256-Public-Key-Fingerprint:  ", inst.sha256PublicKeyFingerprint)

	return nil
}

func (inst *displayCertificateDetailTask) displayPrivateKey() error {
	inst.displayBar("# private key ")

	holder := inst.certKey
	key := holder.Signer()
	algorithm := holder.Algorithm()

	pri := key.(*rsa.PrivateKey)
	der := x509.MarshalPKCS1PrivateKey(pri)
	// if err != nil {
	// 	return err
	// }

	block := &pem.Block{
		Type:  algorithm + " PRIVATE KEY",
		Bytes: der,
	}
	bin := pem.EncodeToMemory(block)
	str := string(bin)
	fmt.Println(str)
	return nil
}

func (inst *displayCertificateDetailTask) displayCert() error {
	inst.displayBar("# leaf cert ")
	cer := inst.cc.Cert
	block := &pem.Block{}
	block.Type = "CERTIFICATE"
	block.Bytes = cer.Raw
	data := pem.EncodeToMemory(block)
	str := string(data)
	fmt.Println(str)
	return nil
}

func (inst *displayCertificateDetailTask) displayMeta() error {
	inst.displayBar("# meta ")

	const nl = "\n"
	file := inst.cc.CertFile
	info := file.GetInfo()

	fmt.Println("  Path: ", file.GetPath())
	fmt.Println("  Content-Length: ", info.Length())

	return nil
}

func (inst *displayCertificateDetailTask) displayChain() error {
	inst.displayBar("# chain ")
	file := inst.cc.CertFile
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}
	fmt.Println(text)
	return nil
}

func (inst *displayCertificateDetailTask) displayBar(tag string) {
	b := &strings.Builder{}
	for i := 100; i > 0; i-- {
		b.WriteRune('-')
	}
	fmt.Println(tag, b.String())
}
