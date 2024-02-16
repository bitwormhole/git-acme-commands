package test4gac
import (
    pa83404c2c "github.com/bitwormhole/git-acme-commands/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pa83404c2c.DemoUnit in package:github.com/bitwormhole/git-acme-commands/src/test/golang/unit
//
// id:com-a83404c2cb656b45-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type pa83404c2cb_unit_DemoUnit struct {
}

func (inst* pa83404c2cb_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-a83404c2cb656b45-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pa83404c2cb_unit_DemoUnit) new() any {
    return &pa83404c2c.DemoUnit{}
}

func (inst* pa83404c2cb_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pa83404c2c.DemoUnit)
	nop(ie, com)

	


    return nil
}


