package ls

import (
	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/bitwormhole/git-acme-commands/app/data/vo"
	"github.com/starter-go/afs"
	"github.com/starter-go/application/properties"
)

// LoadDomainConfig ...
func LoadDomainConfig(file afs.Path) (*vo.DomainConfig, error) {
	loader := new(DomainConfigLoader)
	return loader.Load(file)
}

////////////////////////////////////////////////////////////////////////////////

// DomainConfigLoader ...
type DomainConfigLoader struct {
}

// Load ...
func (inst *DomainConfigLoader) Load(file afs.Path) (*vo.DomainConfig, error) {

	props, err := inst.loadProperties(file)
	if err != nil {
		return nil, err
	}

	getter := props.Getter()
	dn := getter.GetString("domain.name")
	key := getter.GetString("domain.key")

	dst := &vo.DomainConfig{
		Name: dto.DomainName(dn),
		Key:  dto.PublicKeyFingerprint(key),
	}

	err = getter.Error()
	return dst, err
}

func (inst *DomainConfigLoader) loadProperties(file afs.Path) (properties.Table, error) {
	text, err := file.GetIO().ReadText(nil)
	if err != nil {
		return nil, err
	}
	return properties.Parse(text, nil)
}

////////////////////////////////////////////////////////////////////////////////

type DomainConfigSaver struct {
}

////////////////////////////////////////////////////////////////////////////////
