package commands

import "github.com/starter-go/cli"

type subcmdGitAcmeUpdate struct {
	parent *GitAcme
}

func (inst *subcmdGitAcmeUpdate) name() string {
	return "git-acme-update"
}

func (inst *subcmdGitAcmeUpdate) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeUpdate) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "更新配置文件",
		Usage:   "git acme update",
		Content: "",
	}
}

func (inst *subcmdGitAcmeUpdate) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeUpdate) handle(t *cli.Task) error {
	return nil
}
