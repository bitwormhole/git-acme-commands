package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/bitwormhole/git-acme-commands/app/acme"
	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/starter-go/afs"
	"github.com/starter-go/cli"
	"github.com/starter-go/vlog"
)

type subcmdGitAcmeRequest struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeRequest) name() string {
	return "git-acme-request"
}

func (inst *subcmdGitAcmeRequest) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeRequest) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "请求创建证书",
		Usage:   "git acme request",
		Content: "",
	}
}

func (inst *subcmdGitAcmeRequest) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeRequest) handle(t *cli.Task) error {
	t2 := new(subcmdGitAcmeRequestTask)
	t2.root = inst.parent
	t2.debug = true
	return t2.run(t.Context)
}

////////////////////////////////////////////////////////////////////////////////

type subcmdGitAcmeRequestTask struct {
	root  *GitACME
	ctx   context.Context
	cc    *core.ContainerContext
	dc    *core.DomainContext
	debug bool

	req  *acme.RequestV2
	resp *acme.ResponseV2

	latest afs.Path
	target afs.Path
}

func (inst *subcmdGitAcmeRequestTask) run(ctx context.Context) error {

	inst.ctx = ctx
	steps := make([]func() error, 0)

	steps = append(steps, inst.prepareContext)
	steps = append(steps, inst.prepareRequest)
	steps = append(steps, inst.loadSigners)

	steps = append(steps, inst.sendRequest)

	steps = append(steps, inst.locateTargetCertFile)
	steps = append(steps, inst.saveCertToFile)
	steps = append(steps, inst.updateLatestFile)

	for i, step := range steps {
		fmt.Println("step: ", i)
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *subcmdGitAcmeRequestTask) prepareContext() error {

	ctx := inst.ctx
	dc, err := inst.root.Contexts.LoadDomainContext(ctx)
	if err != nil {
		return err
	}
	cc := dc.Parent

	inst.cc = cc
	inst.dc = dc
	return nil
}

func (inst *subcmdGitAcmeRequestTask) prepareRequest() error {

	req := &acme.RequestV2{}
	if inst.debug {
		req = req.PrepareForTest()
	} else {
		req = req.PrepareForProduction()
	}
	req.UserEmail = inst.cc.UserEmail
	req.DomainName = inst.dc.DomainName.String()

	inst.req = req
	return nil
}

func (inst *subcmdGitAcmeRequestTask) loadSigners() error {

	cc := inst.cc
	domainKeyFP := inst.dc.Config.Key
	userKeyFP := inst.cc.Config.User.Key
	keyman := inst.root.KeyManager

	domainKeyHolder, err := keyman.Find(cc, domainKeyFP)
	if err != nil {
		return err
	}

	userKeyHolder, err := keyman.Find(cc, userKeyFP)
	if err != nil {
		return err
	}

	req := inst.req
	req.DomainSigner = domainKeyHolder.Signer()
	req.UserSigner = userKeyHolder.Signer()
	return nil
}

func (inst *subcmdGitAcmeRequestTask) locateTargetCertFile() error {

	suffix := ".cer"
	dn := inst.dc.DomainName.String()
	dir := inst.dc.DomainDirectory
	now := inst.cc.SessionTime

	strTime := now.Format(time.DateOnly)
	name := dn + "-" + strTime + suffix

	file := dir.GetChild(name)
	inst.target = file
	inst.latest = inst.dc.LatestCertFile
	return nil
}

func (inst *subcmdGitAcmeRequestTask) sendRequest() error {

	req := inst.req
	req.DoMakeCert = true
	req.DoNewAccount = true

	resp, err := req.Send()
	if err != nil {
		return err
	}
	inst.resp = resp
	return nil
}

func (inst *subcmdGitAcmeRequestTask) saveCertToFile() error {

	file := inst.target

	if file.Exists() {
		path := file.GetPath()
		vlog.Warn("the target file is exists, skip to save. [path=%s]", path)
		return nil
	}

	resp := inst.resp
	data := resp.ChainPEM
	opt := afs.ToCreateFile()
	return file.GetIO().WriteBinary(data, opt)
}

func (inst *subcmdGitAcmeRequestTask) updateLatestFile() error {

	file := inst.latest
	filename := inst.target.GetName()

	opt := afs.ToWriteFile()
	if !file.Exists() {
		opt = afs.ToCreateFile()
	}

	return file.GetIO().WriteText(filename, opt)
}
