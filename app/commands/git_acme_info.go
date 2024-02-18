package commands

import "github.com/starter-go/cli"

type subcmdGitAcmeInfo struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeInfo) name() string {
	return "git-acme-info"
}

func (inst *subcmdGitAcmeInfo) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeInfo) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "显示当前仓库中的证书信息（大概）",
		Usage:   "git acme info",
		Content: "",
	}
}

func (inst *subcmdGitAcmeInfo) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeInfo) handle(t *cli.Task) error {

	ctx1 := t.Context

	ctx2, err := inst.parent.Contexts.NewCertRepoContext(ctx1)
	if err != nil {
		return err
	}

	err = ctx2.LoadConfig()
	if err != nil {
		return err
	}

	err = ctx2.LoadDomainList()
	if err != nil {
		return err
	}

	return nil
}
