package commands

import (
	"github.com/bitwormhole/git-acme-commands/app/config"
	"github.com/starter-go/afs"
	"github.com/starter-go/cli"
	"github.com/starter-go/vlog"
)

type subcmdGitAcmePrepare struct {
	parent *GitACME
}

func (inst *subcmdGitAcmePrepare) name() string {
	return "git-acme-prepare"
}

func (inst *subcmdGitAcmePrepare) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmePrepare) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "为申请证书准备必须的密钥对",
		Usage:   "git acme prepare",
		Content: "",
	}
}

func (inst *subcmdGitAcmePrepare) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmePrepare) handle(t *cli.Task) error {

	ctx1 := t.Context
	ctx2, err := inst.parent.Contexts.NewCertRepoContext(ctx1)
	if err != nil {
		return err
	}

	err = ctx2.LoadConfig()
	if err != nil {
		return err
	}

	cfg := ctx2.MixedConfig
	kplist := cfg.KeyPairs
	for _, kp := range kplist {
		err = inst.prepareKeyPair(kp, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *subcmdGitAcmePrepare) prepareKeyPair(pair *config.KeyPairDTO, cfg *config.VO) error {

	dir, err := cfg.FindDirectory(pair.Directory)
	if err != nil {
		return err
	}

	file := inst.parent.FS.NewPath(dir.Path).GetChild(pair.FileName)
	path := file.GetPath()
	if file.Exists() {
		// skip
		vlog.Info("skip key file [%s]", path)
		return nil
	}

	// gen key
	provider, err := inst.parent.Keys.FindProvider(pair.Algorithm)
	if err != nil {
		return err
	}

	keypair, err := provider.Generator().Generate()
	if err != nil {
		return err
	}

	vlog.Info("generate key-pair and save to file [%s]", path)

	opt := afs.Todo().Create(true).Dir(true).Mkdirs(true).Options()
	file.MakeParents(opt)
	return provider.Saver().Save(keypair, file)
}
