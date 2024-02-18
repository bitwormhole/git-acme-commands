package commands

import (
	"context"
	"os"

	"github.com/bitwormhole/git-acme-commands/app/contexts"
	"github.com/bitwormhole/git-acme-commands/app/keys"
	"github.com/bitwormhole/gitlib"
	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/cli"
)

type subcommand interface {
	Registration() *cli.HandlerRegistration
}

// GitACME 提供所有 git-acme-* 命令的注册信息
type GitACME struct {

	//starter:component

	_as func(cli.HandlerRegistry) //starter:as(".")

	CLI      cli.CLI                     //starter:inject("#")
	FS       afs.FS                      //starter:inject("#")
	Git      gitlib.Agent                //starter:inject("#")
	Contexts contexts.Service            //starter:inject("#")
	Keys     keys.KeyPairProviderManager //starter:inject("#")

}

func (inst *GitACME) _impl() cli.HandlerRegistry {
	return inst
}

// Life 注册生命周期处理函数
func (inst *GitACME) Life() *application.Life {
	return &application.Life{
		OnLoop: inst.run,
	}
}

// GetHandlers 获取命令处理器的注册信息
func (inst *GitACME) GetHandlers() []*cli.HandlerRegistration {

	sublist := make([]subcommand, 0)

	sublist = append(sublist, &subcmdGitAcmeAdd{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeDo{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeHelp{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeInfo{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeInit{parent: inst})
	sublist = append(sublist, &subcmdGitAcmePrepare{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeRequest{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeUpdate{parent: inst})

	all := make([]*cli.HandlerRegistration, 0)
	for _, subitem := range sublist {
		reg := subitem.Registration()
		all = append(all, reg)
	}
	return all
}

func (inst *GitACME) run() error {
	args := os.Args[1:]
	ctx := context.Background()
	ctx = inst.CLI.Bind(ctx)
	client := inst.CLI.GetClient()
	return client.RunCCA(ctx, "git-acme", args)
}
