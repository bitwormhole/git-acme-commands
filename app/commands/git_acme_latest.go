package commands

import (
	"github.com/bitwormhole/git-acme-commands/app/classes/certs"
	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/ls"
	"github.com/starter-go/cli"
)

type subcmdGitAcmeLatest struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeLatest) name() string {
	return "git-acme-latest"
}

func (inst *subcmdGitAcmeLatest) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeLatest) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "显示最新的域名证书",
		Usage:   "git acme latest",
		Content: "",
	}
}

func (inst *subcmdGitAcmeLatest) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeLatest) handle(t *cli.Task) error {

	dc, err := inst.parent.Contexts.LoadDomainContext(t.Context)
	if err != nil {
		return err
	}

	certFile, err := dc.GetLatestCertificateFile()
	if err != nil {
		return err
	}

	chain, err := ls.LoadCertificateChain(certFile)
	if err != nil {
		return err
	}

	cer := chain.Leaf()
	cc := &core.CertificateContext{
		Parent:   dc,
		CertFile: certFile,
		Chain:    chain,
		Cert:     cer,
	}
	return certs.DisplayCertificateDetail(cc)
}
