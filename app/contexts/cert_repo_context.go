package contexts

import (
	"time"

	"github.com/bitwormhole/git-acme-commands/app/config"
	"github.com/starter-go/afs"
)

// CertRepoContext 包含关于一个证书仓库的上下文信息
type CertRepoContext struct {
	Parent *GitRepoContext

	// files
	DomainListFile  afs.Path // at '.git/../acme.config'
	LocalConfigFile afs.Path // at '.git/../acme-local.config'
	MainConfigFile  afs.Path // at '.git/../acme-domains.list'

	// config
	// ConfigProperties properties.Table
	MainConfig  *config.VO
	LocalConfig *config.VO
	MixedConfig *config.VO

	// domains
	DomainList []*DomainListItem
	Domains    map[string]*DomainListItem

	// time
	Now             time.Time
	SessionTime     time.Time
	SessionInterval time.Duration
}

// LoadConfig ...
func (inst *CertRepoContext) LoadConfig() error {

	fileM := inst.MainConfigFile
	fileL := inst.LocalConfigFile

	cfgM, err := config.Load(fileM)
	if err != nil {
		return err
	}

	cfgL, err := config.Load(fileL)
	if err != nil {
		return err
	}

	cfgX := config.Mix(cfgM, cfgL)

	inst.MixedConfig = cfgX
	inst.MainConfig = cfgM
	inst.LocalConfig = cfgL
	return nil
}

// LoadDomainList ...
func (inst *CertRepoContext) LoadDomainList() error {

	file := inst.DomainListFile
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return err
	}

	list, err := parseDomainList(text)
	if err != nil {
		return err
	}

	table := make(map[string]*DomainListItem)
	for _, item := range list {
		if item.Enabled {
			table[item.Domain] = item
		}
	}

	inst.DomainList = list
	inst.Domains = table
	return nil
}
