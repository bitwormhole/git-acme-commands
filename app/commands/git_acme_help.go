package commands

import "github.com/starter-go/cli"

type subcmdGitAcmeHelp struct {
	parent *GitAcme
}

func (inst *subcmdGitAcmeHelp) name() string {
	return "git-acme-help"
}

func (inst *subcmdGitAcmeHelp) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeHelp) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "显示帮助信息",
		Usage:   "git acme help",
		Content: "",
	}
}

func (inst *subcmdGitAcmeHelp) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeHelp) handle(t *cli.Task) error {

	ctx := t.Context
	client := t.Client
	args := []string{}

	return client.RunCCA(ctx, "help", args)
}
