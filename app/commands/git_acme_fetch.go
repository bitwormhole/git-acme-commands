package commands

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/bitwormhole/git-acme-commands/app/data/ls"
	"github.com/starter-go/afs"
	"github.com/starter-go/cli"
	"github.com/starter-go/vlog"
)

type subcmdGitAcmeFetch struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeFetch) name() string {
	return "git-acme-fetch"
}

func (inst *subcmdGitAcmeFetch) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeFetch) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "从 web 服务器下载当前使用的证书",
		Usage:   "git acme fetch",
		Content: "",
	}
}

func (inst *subcmdGitAcmeFetch) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeFetch) handle(t *cli.Task) error {
	t2 := new(subcmdGitAcmeFetchTask)
	t2.parent = inst.parent
	return t2.run(t.Context)
}

////////////////////////////////////////////////////////////////////////////////

type subcmdGitAcmeFetchTask struct {
	parent *GitACME
	dc     *core.DomainContext
	ctx    context.Context

	url     string
	chain   dto.CertChain
	target  afs.Path
	current afs.Path
}

func (inst *subcmdGitAcmeFetchTask) run(ctx context.Context) error {
	inst.ctx = ctx
	steps := make([]func() error, 0)

	steps = append(steps, inst.prepareContext)
	steps = append(steps, inst.makeURL)
	steps = append(steps, inst.fetch)
	steps = append(steps, inst.locateFiles)
	steps = append(steps, inst.saveCert)
	steps = append(steps, inst.updateCurrentFile)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *subcmdGitAcmeFetchTask) prepareContext() error {
	dc, err := inst.parent.Contexts.LoadDomainContext(inst.ctx)
	if err != nil {
		return err
	}
	inst.dc = dc
	return nil
}

func (inst *subcmdGitAcmeFetchTask) makeURL() error {

	dc := inst.dc
	dn := dc.DomainName.String()
	src := dc.Config.FetchFromURL

	if src == "" {
		src = "https://" + dn + "/"
		inst.url = src
		return nil
	}

	x, err := url.Parse(src)
	if err != nil {
		return err
	}

	port := x.Port()
	if port == "" {
		x.Host = dn
	} else {
		x.Host = dn + ":" + port
	}

	src = x.String()
	inst.url = src
	return nil
}

func (inst *subcmdGitAcmeFetchTask) fetch() error {

	url := inst.url
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	cs := resp.TLS
	chain := cs.PeerCertificates
	inst.chain = chain
	return nil
}

func (inst *subcmdGitAcmeFetchTask) locateFiles() error {

	dn := inst.dc.DomainName
	cert := inst.chain.Leaf()
	dir := inst.dc.DomainDirectory

	date := cert.NotBefore
	dateStr := date.Format(time.DateOnly)

	builder := &strings.Builder{}
	builder.WriteString(dn.String())
	builder.WriteString("-")
	builder.WriteString(dateStr)
	builder.WriteString(".fetch.cer")

	name := builder.String()
	inst.current = inst.dc.CurrentCertFile
	inst.target = dir.GetChild(name)

	s1 := inst.url
	s2 := inst.target.GetPath()
	vlog.Info("fetch cert-chain:")
	vlog.Info("    from web [%s]", s1)
	vlog.Info("    to file  [%s]", s2)
	return nil
}

func (inst *subcmdGitAcmeFetchTask) saveCert() error {
	file := inst.target
	chain := inst.chain
	return ls.SaveCertificateChain(chain, file)
}

func (inst *subcmdGitAcmeFetchTask) updateCurrentFile() error {
	file := inst.current
	text := inst.target.GetName()
	opt := afs.ToCreateFile()
	if file.Exists() {
		opt = afs.ToWriteFile()
	}
	return file.GetIO().WriteText(text, opt)
}
