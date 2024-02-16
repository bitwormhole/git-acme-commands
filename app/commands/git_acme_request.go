package commands

import "github.com/starter-go/cli"

type subcmdGitAcmeRequest struct {
	parent *GitAcme
}

func (inst *subcmdGitAcmeRequest) name() string {
	return "git-acme-request"
}

func (inst *subcmdGitAcmeRequest) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeRequest) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "请求创建证书",
		Usage:   "git acme request",
		Content: "",
	}
}

func (inst *subcmdGitAcmeRequest) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeRequest) handle(t *cli.Task) error {
	return nil
}
