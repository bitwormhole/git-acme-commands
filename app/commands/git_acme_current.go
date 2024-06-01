package commands

import (
	"github.com/bitwormhole/git-acme-commands/app/classes/certs"
	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/ls"
	"github.com/starter-go/cli"
)

type subcmdGitAcmeCurrent struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeCurrent) name() string {
	return "git-acme-current"
}

func (inst *subcmdGitAcmeCurrent) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeCurrent) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "显示当前线上使用的域名证书(fetch 命令的结果)",
		Usage:   "git acme current",
		Content: "",
	}
}

func (inst *subcmdGitAcmeCurrent) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeCurrent) handle(t *cli.Task) error {

	dc, err := inst.parent.Contexts.LoadDomainContext(t.Context)
	if err != nil {
		return err
	}

	certFile, err := dc.GetCurrentCertificateFile()
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
