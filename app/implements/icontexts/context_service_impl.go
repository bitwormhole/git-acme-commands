package icontexts

import (
	"context"
	"os"
	"time"

	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/ls"
	"github.com/bitwormhole/gitlib"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
)

// ContextServiceImpl ...
type ContextServiceImpl struct {

	//starter:component

	_as func(core.Service) //starter:as("#")

	FS  afs.FS       //starter:inject("#")
	Git gitlib.Agent //starter:inject("#")

}

func (inst *ContextServiceImpl) _impl() core.Service {
	return inst
}

// LoadGitContext ...
func (inst *ContextServiceImpl) LoadGitContext(c context.Context) (*core.GitContext, error) {

	if c == nil {
		c = context.Background()
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	wkdir := inst.FS.NewPath(wd)
	gl := inst.Git.GetLib()
	layout, err := gl.Locator().Locate(wkdir)
	if err != nil {
		return nil, err
	}

	wktree := layout.Workspace()

	res := &core.GitContext{
		Parent:   c,
		WD:       wkdir,
		Layout:   layout,
		Worktree: wktree,
	}

	return res, nil
}

// LoadContainerContext ...
func (inst *ContextServiceImpl) LoadContainerContext(c context.Context) (*core.ContainerContext, error) {

	parent, err := inst.LoadGitContext(c)
	if err != nil {
		return nil, err
	}

	// load config
	wktree := parent.Worktree
	configfile := wktree.GetChild("acme.config")
	cfg, err := ls.LoadContainerConfig(configfile)
	if err != nil {
		return nil, err
	}

	res := &core.ContainerContext{
		Parent:         parent,
		Config:         cfg,
		MainConfigFile: configfile,
	}

	// user
	res.UserEmail = cfg.User.Email
	res.UserName = cfg.User.Name
	res.UserSigner = nil

	//time
	now := lang.Now()
	interval := cfg.ACME.Interval
	res.Now = now.Time()
	res.SessionTime = inst.computeSessionTime(now, interval)
	res.SessionInterval = interval.Duration()

	return res, nil
}

func (inst *ContextServiceImpl) computeSessionTime(now lang.Time, interval lang.Milliseconds) time.Time {
	step := lang.Time(interval)
	if step < 1 {
		step = 1000
	}
	t2 := now - (now % step)
	return t2.Time()
}

// LoadDomainContext ...
func (inst *ContextServiceImpl) LoadDomainContext(c context.Context) (*core.DomainContext, error) {

	parent, err := inst.LoadContainerContext(c)
	if err != nil {
		return nil, err
	}

	wd := parent.Parent.WD

	domainConfigFile := wd.GetChild("domain.config")
	cfg, err := ls.LoadDomainConfig(domainConfigFile)
	if err != nil {
		return nil, err
	}

	dir := domainConfigFile.GetParent()
	currentFile := dir.GetChild("current")
	latestFile := dir.GetChild("latest")

	res := &core.DomainContext{
		Parent:           parent,
		Config:           cfg,
		DomainConfigFile: domainConfigFile,
		DomainName:       cfg.Name,
		DomainDirectory:  dir,
		LatestCertFile:   latestFile,
		CurrentCertFile:  currentFile,
	}
	return res, nil
}
