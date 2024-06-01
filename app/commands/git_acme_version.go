package commands

import (
	"fmt"

	"github.com/starter-go/cli"
)

type subcmdGitAcmeVersion struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeVersion) name() string {
	return "git-acme-version"
}

func (inst *subcmdGitAcmeVersion) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeVersion) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "显示 git-acme 应用程序的版本",
		Usage:   "git acme version",
		Content: "",
	}
}

func (inst *subcmdGitAcmeVersion) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeVersion) handle(t *cli.Task) error {
	mod := inst.parent.AppContext.GetMainModule()
	ver := mod.Version()
	r := mod.Revision()
	fmt.Printf("Git ACME %s (r%d)", ver, r)
	return nil
}
