package commands

import (
	"context"
	"crypto/x509"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/git-acme-commands/app/core"
	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/bitwormhole/git-acme-commands/app/data/ls"
	"github.com/starter-go/cli"
	"github.com/starter-go/vlog"
)

type subcmdGitAcmeDomainList struct {
	parent *GitACME
}

func (inst *subcmdGitAcmeDomainList) name() string {
	return "git-acme-domains"
}

func (inst *subcmdGitAcmeDomainList) Registration() *cli.HandlerRegistration {
	return &cli.HandlerRegistration{
		Name:    inst.name(),
		OnInit:  inst.init,
		Handler: inst.handle,
		Help:    inst,
	}
}

func (inst *subcmdGitAcmeDomainList) GetHelp() *cli.HelpInfo {
	return &cli.HelpInfo{
		Name:    inst.name(),
		Title:   "向配置文件添加新的域名",
		Usage:   "git acme domains",
		Content: "",
	}
}

func (inst *subcmdGitAcmeDomainList) init(c *cli.Context) error {
	return nil
}

func (inst *subcmdGitAcmeDomainList) handle(t *cli.Task) error {
	t2 := new(subcmdGitAcmeDomainListTask)
	t2.parent = inst.parent
	return t2.run(t.Context)
}

////////////////////////////////////////////////////////////////////////////////

type subcmdGitAcmeDomainListItem struct {
	dn    dto.DomainName
	cert  *x509.Certificate
	err   error
	time1 time.Time
	time2 time.Time
}

////////////////////////////////////////////////////////////////////////////////

type subcmdGitAcmeDomainListTask struct {
	parent *GitACME
	cc     *core.ContainerContext
	ctx    context.Context

	items []*subcmdGitAcmeDomainListItem
}

func (inst *subcmdGitAcmeDomainListTask) run(ctx context.Context) error {
	steps := make([]func() error, 0)

	steps = append(steps, inst.prepareContext)
	steps = append(steps, inst.loadItems)
	steps = append(steps, inst.sortItems)
	steps = append(steps, inst.printItems)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *subcmdGitAcmeDomainListTask) prepareContext() error {
	ctx1 := inst.ctx
	ctx2, err := inst.parent.Contexts.LoadContainerContext(ctx1)
	if err != nil {
		return err
	}
	inst.cc = ctx2
	return nil
}

func (inst *subcmdGitAcmeDomainListTask) loadItems() error {
	src := inst.cc.Config.DomainList
	dst := inst.items
	for _, item1 := range src {
		item2, err := inst.loadItem(item1)
		if err != nil {
			// return err
			vlog.Error(err.Error())
			continue
		}
		dst = append(dst, item2)
	}
	inst.items = dst
	return nil
}

func (inst *subcmdGitAcmeDomainListTask) formatTime(t time.Time) string {
	return t.Format(time.DateOnly)
}

func (inst *subcmdGitAcmeDomainListTask) printItems() error {
	b := &strings.Builder{}
	list := inst.items
	for index, item := range list {
		row := inst.formatItem(index, item)
		b.WriteString(row)
		b.WriteString("\n")
	}
	fmt.Println(b.String())
	return nil
}

func (inst *subcmdGitAcmeDomainListTask) formatItem(index int, item *subcmdGitAcmeDomainListItem) string {
	const (
		tab = "\t"
	)
	cells := make([]string, 0)
	cells = append(cells, strconv.Itoa(index))
	cells = append(cells, item.dn.String())
	cells = append(cells, inst.formatTime(item.time1))
	cells = append(cells, inst.formatTime(item.time2))

	err := item.err
	if err != nil {
		cells = append(cells, "error:"+err.Error())
	}

	b := &strings.Builder{}
	for _, cell := range cells {
		b.WriteString(tab)
		b.WriteString(cell)
	}
	return b.String()
}

func (inst *subcmdGitAcmeDomainListTask) loadItem(src *dto.Domain) (*subcmdGitAcmeDomainListItem, error) {
	ctx := inst.ctx
	cc := inst.cc
	cfgFile1 := cc.MainConfigFile
	ref := src.Ref
	cfgFile2 := cfgFile1.GetParent().GetChild(ref)
	dc, err := inst.parent.Contexts.LoadDomainContextWithConfigFile(ctx, cfgFile2)
	if err != nil {
		return nil, err
	}

	item2 := new(subcmdGitAcmeDomainListItem)
	item2.dn = dc.DomainName

	cert, err := inst.loadCurrentCert(dc)
	if err != nil {
		item2.err = err
	} else {
		item2.cert = cert
		item2.time1 = cert.NotBefore
		item2.time2 = cert.NotAfter
	}

	return item2, nil
}

func (inst *subcmdGitAcmeDomainListTask) loadCurrentCert(dc *core.DomainContext) (*x509.Certificate, error) {

	currentFile := dc.CurrentFile
	name, err := currentFile.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	if name == "" {
		return nil, fmt.Errorf("current cert is nil")
	}

	certFile := currentFile.GetParent().GetChild(name)
	chain, err := ls.LoadCertificateChain(certFile)
	if err != nil {
		return nil, err
	}

	cer := chain.Leaf()
	return cer, nil
}

func (inst *subcmdGitAcmeDomainListTask) sortItems() error {
	sort.Sort(inst)
	return nil
}

func (inst *subcmdGitAcmeDomainListTask) Len() int {
	return len(inst.items)
}
func (inst *subcmdGitAcmeDomainListTask) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	s1 := o1.dn.String()
	s2 := o2.dn.String()
	return strings.Compare(s1, s2) > 0
}
func (inst *subcmdGitAcmeDomainListTask) Swap(i1, i2 int) {
	list := inst.items
	list[i1], list[i2] = list[i2], list[i1]
}

////////////////////////////////////////////////////////////////////////////////
