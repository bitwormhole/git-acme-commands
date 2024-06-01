package commands

import (
	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/dxo"
	"github.com/starter-go/afs"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/cli"
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
	ctxser := inst.parent.Contexts
	ctx2, err := ctxser.LoadDomainContext(ctx1)
	if err != nil {
		return err
	}

	err = inst.prepareKeyForUser(ctx2.Parent)
	if err != nil {
		return err
	}

	err = inst.prepareKeyForDomain(ctx2)
	if err != nil {
		return err
	}

	return nil
}

func (inst *subcmdGitAcmePrepare) prepareKeyForUser(ctx *core.ContainerContext) error {

	userKey := ctx.Config.User.Key

	kh, err := inst.prepareKeyWithFingerprint(ctx, userKey)
	if err != nil {
		return err
	}

	if kh.Fingerprint() == userKey {
		return nil
	}

	// update key
	userKey = kh.Fingerprint()

	// save to config file ('acme.config')
	file := ctx.MainConfigFile
	props := make(map[string]string)
	props["user.key"] = userKey.String()
	return inst.updateConfigFile(file, props)
}

func (inst *subcmdGitAcmePrepare) prepareKeyForDomain(dc *core.DomainContext) error {

	cc := dc.Parent
	key := dc.Config.Key

	kh, err := inst.prepareKeyWithFingerprint(cc, key)
	if err != nil {
		return err
	}

	if kh.Fingerprint() == key {
		return nil
	}

	// update key
	key = kh.Fingerprint()

	// save to config file ('domain.config')
	file := dc.DomainConfigFile
	props := make(map[string]string)
	props["domain.key"] = key.String()
	return inst.updateConfigFile(file, props)
}

func (inst *subcmdGitAcmePrepare) prepareKeyWithFingerprint(ctx *core.ContainerContext, fp dxo.Fingerprint) (core.KeyHolder, error) {
	keyMan := inst.parent.KeyManager
	if keyMan.Exists(ctx, fp) {
		return keyMan.Find(ctx, fp)
	}
	return keyMan.CreateNew(ctx)
}

func (inst *subcmdGitAcmePrepare) updateConfigFile(file afs.Path, src map[string]string) error {

	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	dst, err := properties.Parse(text, nil)
	if err != nil {
		return err
	}

	for k, v := range src {
		dst.SetProperty(k, v)
	}

	text = properties.Format(dst, properties.FormatWithGroups)
	return file.GetIO().WriteText(text, nil)
}
