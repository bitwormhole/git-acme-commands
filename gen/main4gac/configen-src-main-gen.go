package main4gac
import (
    pdb462c5c7 "github.com/bitwormhole/git-acme-commands/app/commands"
    p1336d65ed "github.com/starter-go/cli"
     "github.com/starter-go/application"
)

// type pdb462c5c7.GitAcme in package:github.com/bitwormhole/git-acme-commands/app/commands
//
// id:com-db462c5c7a634b6f-commands-GitAcme
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type pdb462c5c7a_commands_GitAcme struct {
}

func (inst* pdb462c5c7a_commands_GitAcme) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-db462c5c7a634b6f-commands-GitAcme"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdb462c5c7a_commands_GitAcme) new() any {
    return &pdb462c5c7.GitAcme{}
}

func (inst* pdb462c5c7a_commands_GitAcme) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdb462c5c7.GitAcme)
	nop(ie, com)

	
    com.CLI = inst.getCLI(ie)


    return nil
}


func (inst*pdb462c5c7a_commands_GitAcme) getCLI(ie application.InjectionExt)p1336d65ed.CLI{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-CLI").(p1336d65ed.CLI)
}


