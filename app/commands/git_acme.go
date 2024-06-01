package commands

import (
	"context"
	"os"

	"github.com/bitwormhole/git-acme-commands/app/core"

	"github.com/bitwormhole/gitlib"
	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/cli"
	"github.com/starter-go/keys"
)

type subcommand interface {
	Registration() *cli.HandlerRegistration
}

// GitACME 提供所有 git-acme-* 命令的注册信息
type GitACME struct {

	//starter:component

	_as func(cli.HandlerRegistry) //starter:as(".")

	CLI        cli.CLI             //starter:inject("#")
	FS         afs.FS              //starter:inject("#")
	Git        gitlib.Agent        //starter:inject("#")
	Contexts   core.Service        //starter:inject("#")
	KeyManager core.KeyManager     //starter:inject("#")
	KeyDrivers keys.DriverManager  //starter:inject("#")
	AppContext application.Context //starter:inject("context")

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

	sublist = append(sublist, &subcmdGitAcmeCerts{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeCurrent{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeDomainAdd{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeDomainList{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeFetch{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeHelp{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeInfo{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeInit{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeLatest{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeNewAccount{parent: inst})
	sublist = append(sublist, &subcmdGitAcmePrepare{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeRequest{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeUpdate{parent: inst})
	sublist = append(sublist, &subcmdGitAcmeVersion{parent: inst})

	sublist = append(sublist, inst.makeSubCommandProxy("git-acme"))
	sublist = append(sublist, inst.makeSubCommandProxy("git-acme-domain"))
	sublist = append(sublist, inst.makeSubCommandProxy("git-acme-cert"))

	all := make([]*cli.HandlerRegistration, 0)
	for _, subitem := range sublist {
		reg := subitem.Registration()
		all = append(all, reg)
	}
	return all
}

func (inst *GitACME) makeSubCommandProxy(name string) subcommand {
	sub := &subcmdGitAcmeDo{
		parent: inst,
		name:   name,
	}
	return sub
}

func (inst *GitACME) run() error {
	args := os.Args[1:]
	ctx := context.Background()
	ctx = inst.CLI.Bind(ctx)
	client := inst.CLI.GetClient()
	return client.RunCCA(ctx, "git-acme", args)
}
