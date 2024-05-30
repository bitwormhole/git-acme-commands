package commands

import (
	"fmt"

	"github.com/starter-go/cli"
)

type subcmdGitAcmeNewAccount struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeNewAccount) name() string {
	return "git-acme-new-account"
}

func (inst *subcmdGitAcmeNewAccount) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeNewAccount) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "向 ACME 服务商注册新账号",
		Usage:   "git acme new-account",
		Content: "",
	}
}

func (inst *subcmdGitAcmeNewAccount) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeNewAccount) handle(t *cli.Task) error {

	return fmt.Errorf("no impl")

}
