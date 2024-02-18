package commands

import (
	"crypto"
	"sort"

	"github.com/bitwormhole/git-acme-commands/app/acme"
	"github.com/bitwormhole/git-acme-commands/app/config"
	"github.com/bitwormhole/git-acme-commands/app/contexts"
	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
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

	ctx1, err := inst.parent.Contexts.NewCertRepoContext(t.Context)
	if err != nil {
		return err
	}

	err = ctx1.LoadConfig()
	if err != nil {
		return err
	}

	err = ctx1.LoadDomainList()
	if err != nil {
		return err
	}

	inst.prepareTimeFieldsForContext(ctx1)

	domainlist := inst.getDomainList(ctx1)
	for i, domain := range domainlist {

		vlog.Info("certificate for domain [%s]", domain)

		ctx2, err := inst.parent.Contexts.NewDomainContext(ctx1, domain)
		if err != nil {
			return err
		}

		current := ctx2.CurrentCertFile
		if current.Exists() {
			vlog.Info("skip: certificate file exists, at path [%s]", current.GetPath())
			continue
		}

		makeNewAccount := (i == 0)
		err = inst.requestCert(ctx2, makeNewAccount)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *subcmdGitAcmeRequest) prepareTimeFieldsForContext(c *contexts.CertRepoContext) {

	now := lang.Now()
	cfg := c.MixedConfig

	interval := cfg.ACME.Interval
	t0 := now - (now % lang.Time(interval))

	c.Now = now.Time()
	c.SessionTime = t0.Time()
	c.SessionInterval = interval.Duration()
}

func (inst *subcmdGitAcmeRequest) getDomainList(c *contexts.CertRepoContext) []string {
	src := c.Domains
	dst := make([]string, 0)
	for _, item := range src {
		dst = append(dst, item.Domain)
	}
	sort.Strings(dst)
	return dst
}

func (inst *subcmdGitAcmeRequest) requestCert(dc *contexts.DomainContext, makeNewAccount bool) error {

	cfg := dc.Parent.MixedConfig
	acmeInfo, err := cfg.GetACME()
	if err != nil {
		return err
	}

	account, err := cfg.FindAccount(acmeInfo.Account)
	if err != nil {
		return err
	}

	domainName := dc.DomainName
	domain, err := cfg.FindDomain(config.DomainName(domainName))
	if err != nil {
		return err
	}

	accountKey, err := inst.loadSinger(account.KeyPair, dc.Parent)
	if err != nil {
		return err
	}

	certKey, err := inst.loadSinger(domain.KeyPair, dc.Parent)
	if err != nil {
		return err
	}

	req := new(acme.Request)
	req.Domains = []string{domainName}
	req.ACMEAddress = account.URL
	req.Email = account.Email
	req.AccountSinger = accountKey
	req.CertSinger = certKey
	req.DoNewAccount = makeNewAccount
	req.DoMakeCert = true
	err = req.Run()
	if err != nil {
		return err
	}

	err = inst.saveCert(req.ResultContentData, dc.CurrentCertFile)
	if err != nil {
		return err
	}

	err = inst.saveCert(req.ResultContentData, dc.LatestCertFile)
	if err != nil {
		return err
	}

	return nil
}

func (inst *subcmdGitAcmeRequest) saveCert(data []byte, file afs.Path) error {
	if file.Exists() {
		file.Delete()
	}
	opt := afs.Todo().File(true).Create(true).Write(true).Options()
	file.MakeParents(opt)
	return file.GetIO().WriteBinary(data, opt)
}

func (inst *subcmdGitAcmeRequest) loadSinger(name config.KeyPairName, c *contexts.CertRepoContext) (crypto.Signer, error) {

	cfg := c.MixedConfig
	key, err := cfg.FindKeyPair(name)
	if err != nil {
		return nil, err
	}

	dir, err := cfg.FindDirectory(key.Directory)
	if err != nil {
		return nil, err
	}

	file := inst.parent.FS.NewPath(dir.Path).GetChild(key.FileName)
	provider, err := inst.parent.Keys.FindProvider(key.Algorithm)
	if err != nil {
		return nil, err
	}

	keypair, err := provider.Loader().Load(file)
	if err != nil {
		return nil, err
	}

	return keypair.Signer(), nil
}
