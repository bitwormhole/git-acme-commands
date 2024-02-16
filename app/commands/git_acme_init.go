package commands

import "github.com/starter-go/cli"

type subcmdGitAcmeInit struct {
	parent *GitAcme
}

func (inst *subcmdGitAcmeInit) name() string {
	return "git-acme-init"
}

func (inst *subcmdGitAcmeInit) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeInit) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "把当前目录下的 git 仓库初始化为 ACME 证书的存储库",
		Usage:   "git acme init",
		Content: "",
	}
}

func (inst *subcmdGitAcmeInit) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeInit) handle(t *cli.Task) error {
	return nil
}
