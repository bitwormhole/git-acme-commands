package commands

import (
	"fmt"

	"github.com/starter-go/cli"
)

type subcmdGitAcmeDomainAdd struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeDomainAdd) name() string {
	return "git-acme-domain-add"
}

func (inst *subcmdGitAcmeDomainAdd) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeDomainAdd) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "向配置文件添加新的域名",
		Usage:   "git acme domain add",
		Content: "",
	}
}

func (inst *subcmdGitAcmeDomainAdd) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeDomainAdd) handle(t *cli.Task) error {

	return fmt.Errorf("no impl")

}
