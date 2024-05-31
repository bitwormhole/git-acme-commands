package commands

import "github.com/starter-go/cli"

type subcmdGitAcmeCerts struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeCerts) name() string {
	return "git-acme-certs"
}

func (inst *subcmdGitAcmeCerts) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeCerts) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "显示证书列表",
		Usage:   "git acme certs",
		Content: "",
	}
}

func (inst *subcmdGitAcmeCerts) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeCerts) handle(t *cli.Task) error {

	ctx := t.Context
	client := t.Client
	args := []string{}

	return client.RunCCA(ctx, "help", args)
}
