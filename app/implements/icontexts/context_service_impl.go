package icontexts

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/bitwormhole/git-acme-commands/app/contexts"
	"github.com/bitwormhole/gitlib"
	"github.com/starter-go/afs"
)

// ContextServiceImpl ...
type ContextServiceImpl struct {

	//starter:component

	_as func(contexts.Service) //starter:as("#")

	FS  afs.FS       //starter:inject("#")
	Git gitlib.Agent //starter:inject("#")

}

func (inst *ContextServiceImpl) _impl() contexts.Service {
	return inst
}

// NewGitRepoContext ...
func (inst *ContextServiceImpl) NewGitRepoContext(c context.Context) (*contexts.GitRepoContext, error) {

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	wkdir := inst.FS.NewPath(wd)

	layout, err := inst.Git.GetLib().Locator().Locate(wkdir)
	if err != nil {
		return nil, err
	}

	ctx := new(contexts.GitRepoContext)
	ctx.Parent = c
	ctx.WD = wkdir
	ctx.Layout = layout
	ctx.Worktree = layout.Workspace()

	return ctx, nil
}

// NewCertRepoContext ...
func (inst *ContextServiceImpl) NewCertRepoContext(c context.Context) (*contexts.CertRepoContext, error) {

	const (
		fileMainConfig  = "acme.json"
		fileLocalConfig = "acme-local.json"
		fileDomainList  = "acme-domains.list"
	)

	parent, err := inst.NewGitRepoContext(c)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	tree := parent.Worktree

	ctx := new(contexts.CertRepoContext)
	ctx.Parent = parent
	ctx.Now = now

	ctx.MainConfigFile = tree.GetChild(fileMainConfig)
	ctx.LocalConfigFile = tree.GetChild(fileLocalConfig)
	ctx.DomainListFile = tree.GetChild(fileDomainList)

	return ctx, nil
}

// NewDomainContext ...
func (inst *ContextServiceImpl) NewDomainContext(parent *contexts.CertRepoContext, domain string) (*contexts.DomainContext, error) {

	// parent, err := inst.NewCertRepoContext(c)
	// if err != nil {
	// 	return nil, err
	// }

	domain = inst.normalizeDomainName(domain)
	tree := parent.Parent.Worktree

	ctx := new(contexts.DomainContext)
	ctx.Parent = parent
	ctx.DomainName = domain
	ctx.DomainDirectory = tree.GetChild("domains").GetChild(domain)

	cfpb := &certFilePathBuilder{
		Domain: domain,
		Dir:    ctx.DomainDirectory,
		Time:   ctx.Parent.SessionTime,
	}
	ctx.CurrentCertFile = cfpb.createForCurrent()
	ctx.LatestCertFile = cfpb.createForLatest()

	return ctx, nil
}

func (inst *ContextServiceImpl) normalizeDomainName(domain string) string {
	domain = strings.TrimSpace(domain)
	domain = strings.ToLower(domain)
	return domain
}
