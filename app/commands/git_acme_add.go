package commands

import "github.com/starter-go/cli"

type subcmdGitAcmeAdd struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeAdd) name() string {
	return "git-acme-add"
}

func (inst *subcmdGitAcmeAdd) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeAdd) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "向配置文件添加新的域名",
		Usage:   "git acme add",
		Content: "",
	}
}

func (inst *subcmdGitAcmeAdd) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeAdd) handle(t *cli.Task) error {
	return nil
}
