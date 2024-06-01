package icontexts

import (
	"context"
	"fmt"
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

	FS     afs.FS          //starter:inject("#")
	Git    gitlib.Agent    //starter:inject("#")
	KeyMan core.KeyManager //starter:inject("#")

}

func (inst *ContextServiceImpl) _impl() core.Service {
	return inst
}

// LoadGitContext ...
func (inst *ContextServiceImpl) LoadGitContext(c context.Context) (*core.GitContext, error) {

	if c == nil {
		c = context.Background()
	}

	wd, err := inst.innerGetWD()
	if err != nil {
		return nil, err
	}

	gl := inst.Git.GetLib()
	layout, err := gl.Locator().Locate(wd)
	if err != nil {
		return nil, err
	}

	wktree := layout.Workspace()

	res := &core.GitContext{
		Parent:   c,
		WD:       wd,
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
		KeyManager:     inst.KeyMan,
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
	wd, err := inst.innerGetWD()
	if err != nil {
		return nil, err
	}
	cfg := wd.GetChild("domain.config")
	return inst.innerLoadDomainContextWithConfigFile(c, cfg)
}

// LoadDomainContextWithConfigFile ...
func (inst *ContextServiceImpl) LoadDomainContextWithConfigFile(c context.Context, cfg afs.Path) (*core.DomainContext, error) {
	return inst.innerLoadDomainContextWithConfigFile(c, cfg)
}

func (inst *ContextServiceImpl) innerGetWD() (afs.Path, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	dir := inst.FS.NewPath(wd)
	if !dir.IsDirectory() {
		return nil, fmt.Errorf("path of wd is not a dir [%s]", wd)
	}
	return dir, nil
}

func (inst *ContextServiceImpl) innerLoadDomainContextWithConfigFile(c context.Context, domainConfigFile afs.Path) (*core.DomainContext, error) {

	parent, err := inst.LoadContainerContext(c)
	if err != nil {
		return nil, err
	}

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
		LatestFile:       latestFile,
		CurrentFile:      currentFile,
	}
	return res, nil
}
