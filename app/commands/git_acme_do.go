package commands

import "github.com/starter-go/cli"

// subcmdGitAcmeDo 是 'git-acme' 的根命令处理器，它把接收到的任务委托给具体的子命令
type subcmdGitAcmeDo struct {
	parent *GitAcme
}

func (inst *subcmdGitAcmeDo) name() string {
	return "git-acme"
}

func (inst *subcmdGitAcmeDo) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeDo) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "把命令请求委托给 sub-command 处理",
		Usage:   "git acme [sub-command]",
		Content: "",
	}
}

func (inst *subcmdGitAcmeDo) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeDo) handle(t *cli.Task) error {

	args1 := t.Arguments
	args2 := make([]string, 0)
	subcmd := ""

	for i, ar := range args1 {
		if i == 0 {
			subcmd = ar
		} else {
			args2 = append(args2, ar)
		}
	}

	if subcmd == "" {
		subcmd = "help"
	}

	ctx := t.Context
	client := t.Client
	cmd := "git-acme-" + subcmd

	return client.RunCCA(ctx, cmd, args2)
}
