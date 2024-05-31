package commands

import (
	"fmt"

	"github.com/starter-go/cli"
)

type subcmdGitAcmeDomainList struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeDomainList) name() string {
	return "git-acme-domain-list"
}

func (inst *subcmdGitAcmeDomainList) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeDomainList) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "向配置文件添加新的域名",
		Usage:   "git acme domain list",
		Content: "",
	}
}

func (inst *subcmdGitAcmeDomainList) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeDomainList) handle(t *cli.Task) error {

	return fmt.Errorf("no impl")

}
