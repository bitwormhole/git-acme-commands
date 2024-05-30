package main4gac
import (
    pdb462c5c7 "github.com/bitwormhole/git-acme-commands/app/commands"
    p7d2748932 "github.com/bitwormhole/git-acme-commands/app/core"
    p9daa3c6fa "github.com/bitwormhole/git-acme-commands/app/implements/icontexts"
    p45e104bf1 "github.com/bitwormhole/git-acme-commands/app/implements/ikeys"
    paeb460c7d "github.com/bitwormhole/gitlib"
    p0d2a11d16 "github.com/starter-go/afs"
    p1336d65ed "github.com/starter-go/cli"
    pc38c9ad22 "github.com/starter-go/keys"
     "github.com/starter-go/application"
)

// type pdb462c5c7.GitACME in package:github.com/bitwormhole/git-acme-commands/app/commands
//
// id:com-db462c5c7a634b6f-commands-GitACME
// class:class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry
// alias:
// scope:singleton
//
type pdb462c5c7a_commands_GitACME struct {
}

func (inst* pdb462c5c7a_commands_GitACME) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-db462c5c7a634b6f-commands-GitACME"
	r.Classes = "class-1336d65edeed550b78a5d5b61e92d726-HandlerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pdb462c5c7a_commands_GitACME) new() any {
    return &pdb462c5c7.GitACME{}
}

func (inst* pdb462c5c7a_commands_GitACME) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pdb462c5c7.GitACME)
	nop(ie, com)

	
    com.CLI = inst.getCLI(ie)
    com.FS = inst.getFS(ie)
    com.Git = inst.getGit(ie)
    com.Contexts = inst.getContexts(ie)
    com.KeyManager = inst.getKeyManager(ie)
    com.KeyDrivers = inst.getKeyDrivers(ie)


    return nil
}


func (inst*pdb462c5c7a_commands_GitACME) getCLI(ie application.InjectionExt)p1336d65ed.CLI{
    return ie.GetComponent("#alias-1336d65edeed550b78a5d5b61e92d726-CLI").(p1336d65ed.CLI)
}


func (inst*pdb462c5c7a_commands_GitACME) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*pdb462c5c7a_commands_GitACME) getGit(ie application.InjectionExt)paeb460c7d.Agent{
    return ie.GetComponent("#alias-aeb460c7d339df24b0b38a0d65e30102-Agent").(paeb460c7d.Agent)
}


func (inst*pdb462c5c7a_commands_GitACME) getContexts(ie application.InjectionExt)p7d2748932.Service{
    return ie.GetComponent("#alias-7d27489328b03a090b67e7d081689fc8-Service").(p7d2748932.Service)
}


func (inst*pdb462c5c7a_commands_GitACME) getKeyManager(ie application.InjectionExt)p7d2748932.KeyManager{
    return ie.GetComponent("#alias-7d27489328b03a090b67e7d081689fc8-KeyManager").(p7d2748932.KeyManager)
}


func (inst*pdb462c5c7a_commands_GitACME) getKeyDrivers(ie application.InjectionExt)pc38c9ad22.DriverManager{
    return ie.GetComponent("#alias-c38c9ad22b7867d5ce346589e145db9f-DriverManager").(pc38c9ad22.DriverManager)
}



// type p9daa3c6fa.ContextServiceImpl in package:github.com/bitwormhole/git-acme-commands/app/implements/icontexts
//
// id:com-9daa3c6fa26070b9-icontexts-ContextServiceImpl
// class:
// alias:alias-7d27489328b03a090b67e7d081689fc8-Service
// scope:singleton
//
type p9daa3c6fa2_icontexts_ContextServiceImpl struct {
}

func (inst* p9daa3c6fa2_icontexts_ContextServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9daa3c6fa26070b9-icontexts-ContextServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-7d27489328b03a090b67e7d081689fc8-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p9daa3c6fa2_icontexts_ContextServiceImpl) new() any {
    return &p9daa3c6fa.ContextServiceImpl{}
}

func (inst* p9daa3c6fa2_icontexts_ContextServiceImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p9daa3c6fa.ContextServiceImpl)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)
    com.Git = inst.getGit(ie)


    return nil
}


func (inst*p9daa3c6fa2_icontexts_ContextServiceImpl) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*p9daa3c6fa2_icontexts_ContextServiceImpl) getGit(ie application.InjectionExt)paeb460c7d.Agent{
    return ie.GetComponent("#alias-aeb460c7d339df24b0b38a0d65e30102-Agent").(paeb460c7d.Agent)
}



// type p45e104bf1.KeyManagerImpl in package:github.com/bitwormhole/git-acme-commands/app/implements/ikeys
//
// id:com-45e104bf1ff44630-ikeys-KeyManagerImpl
// class:
// alias:alias-7d27489328b03a090b67e7d081689fc8-KeyManager
// scope:singleton
//
type p45e104bf1f_ikeys_KeyManagerImpl struct {
}

func (inst* p45e104bf1f_ikeys_KeyManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-45e104bf1ff44630-ikeys-KeyManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-7d27489328b03a090b67e7d081689fc8-KeyManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p45e104bf1f_ikeys_KeyManagerImpl) new() any {
    return &p45e104bf1.KeyManagerImpl{}
}

func (inst* p45e104bf1f_ikeys_KeyManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p45e104bf1.KeyManagerImpl)
	nop(ie, com)

	
    com.Drivers = inst.getDrivers(ie)


    return nil
}


func (inst*p45e104bf1f_ikeys_KeyManagerImpl) getDrivers(ie application.InjectionExt)pc38c9ad22.DriverManager{
    return ie.GetComponent("#alias-c38c9ad22b7867d5ce346589e145db9f-DriverManager").(pc38c9ad22.DriverManager)
}


