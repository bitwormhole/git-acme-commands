package ls

import (
	"fmt"

	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/bitwormhole/git-acme-commands/app/data/vo"

	"github.com/starter-go/afs"
	"github.com/starter-go/application/properties"
	"github.com/starter-go/base/lang"
)

// LoadContainerConfig ...
func LoadContainerConfig(file afs.Path) (*vo.ContainerConfig, error) {
	loader := new(ContainerConfigLoader)
	return loader.Load(file)
}

////////////////////////////////////////////////////////////////////////////////

// ContainerConfigLoader ...
type ContainerConfigLoader struct {
}

// Load ...
func (inst *ContainerConfigLoader) Load(file afs.Path) (*vo.ContainerConfig, error) {

	props, err := inst.loadProperties(file)
	if err != nil {
		return nil, err
	}
	getter := props.Getter()

	acme, err := inst.loadACME(getter)
	if err != nil {
		return nil, err
	}

	user, err := inst.loadUser(getter)
	if err != nil {
		return nil, err
	}

	domains, err := inst.loadDomainList(props)
	if err != nil {
		return nil, err
	}

	dst := &vo.ContainerConfig{
		ACME:       acme,
		User:       user,
		DomainList: domains,
	}
	return dst, nil
}

func (inst *ContainerConfigLoader) loadProperties(file afs.Path) (properties.Table, error) {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return properties.Parse(text, nil)
}

func (inst *ContainerConfigLoader) loadACME(g1 properties.Getter) (*dto.ACME, error) {
	const (
		pInterval = "acme.interval"
	)
	n := g1.GetInt64(pInterval, 3600*1000)
	dst := &dto.ACME{}
	dst.Interval = lang.Milliseconds(n)
	err := g1.Error()
	return dst, err
}

func (inst *ContainerConfigLoader) loadUser(p properties.Getter) (*dto.User, error) {
	const (
		pName  = "user.name"
		pEmail = "user.email"
		pKey   = "user.key"
	)
	p = p.Required()
	name := p.GetString(pName, "")
	email := p.GetString(pEmail, "")
	keypair := p.GetString(pKey, "")
	dst := &dto.User{
		Name:  name,
		Email: email,
		Key:   dto.PublicKeyFingerprint(keypair),
	}
	err := p.Error()
	return dst, err
}

func (inst *ContainerConfigLoader) loadDomainList(p properties.Table) ([]*dto.Domain, error) {
	const (
		prefix = "domain."
		suffix = ".name"
	)
	getter := p.Getter()
	ids := getter.ListItems(prefix, suffix)
	items := make([]*dto.Domain, 0)
	for _, id := range ids {
		item, err := inst.loadDomainItem(id, getter)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func (inst *ContainerConfigLoader) loadDomainItem(id string, src properties.Getter) (*dto.Domain, error) {
	prefix := fmt.Sprintf("domain.%s.", id)
	name := src.GetString(prefix + "name")
	ref := src.GetString(prefix + "ref")
	item := &dto.Domain{
		ID:   id,
		Ref:  ref,
		Name: dto.DomainName(name),
	}
	err := src.Error()
	return item, err
}

////////////////////////////////////////////////////////////////////////////////

// ContainerConfigSaver ...
type ContainerConfigSaver struct {
}

////////////////////////////////////////////////////////////////////////////////
