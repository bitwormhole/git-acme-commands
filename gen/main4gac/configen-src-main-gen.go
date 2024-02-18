package main4gac
import (
    pdb462c5c7 "github.com/bitwormhole/git-acme-commands/app/commands"
    pb02666f39 "github.com/bitwormhole/git-acme-commands/app/contexts"
    p9daa3c6fa "github.com/bitwormhole/git-acme-commands/app/implements/icontexts"
    p45e104bf1 "github.com/bitwormhole/git-acme-commands/app/implements/ikeys"
    pfe1800bf5 "github.com/bitwormhole/git-acme-commands/app/implements/ikeys/pecdsa"
    p169246a6f "github.com/bitwormhole/git-acme-commands/app/implements/ikeys/prsa"
    p2e998826d "github.com/bitwormhole/git-acme-commands/app/keys"
    paeb460c7d "github.com/bitwormhole/gitlib"
    p0d2a11d16 "github.com/starter-go/afs"
    p1336d65ed "github.com/starter-go/cli"
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
    com.Keys = inst.getKeys(ie)


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


func (inst*pdb462c5c7a_commands_GitACME) getContexts(ie application.InjectionExt)pb02666f39.Service{
    return ie.GetComponent("#alias-b02666f395daf777103666ac99274a3d-Service").(pb02666f39.Service)
}


func (inst*pdb462c5c7a_commands_GitACME) getKeys(ie application.InjectionExt)p2e998826d.KeyPairProviderManager{
    return ie.GetComponent("#alias-2e998826d6036014cdd504594af76975-KeyPairProviderManager").(p2e998826d.KeyPairProviderManager)
}



// type p9daa3c6fa.ContextServiceImpl in package:github.com/bitwormhole/git-acme-commands/app/implements/icontexts
//
// id:com-9daa3c6fa26070b9-icontexts-ContextServiceImpl
// class:
// alias:alias-b02666f395daf777103666ac99274a3d-Service
// scope:singleton
//
type p9daa3c6fa2_icontexts_ContextServiceImpl struct {
}

func (inst* p9daa3c6fa2_icontexts_ContextServiceImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-9daa3c6fa26070b9-icontexts-ContextServiceImpl"
	r.Classes = ""
	r.Aliases = "alias-b02666f395daf777103666ac99274a3d-Service"
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



// type p45e104bf1.KeyPairProviderManagerImpl in package:github.com/bitwormhole/git-acme-commands/app/implements/ikeys
//
// id:com-45e104bf1ff44630-ikeys-KeyPairProviderManagerImpl
// class:
// alias:alias-2e998826d6036014cdd504594af76975-KeyPairProviderManager
// scope:singleton
//
type p45e104bf1f_ikeys_KeyPairProviderManagerImpl struct {
}

func (inst* p45e104bf1f_ikeys_KeyPairProviderManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-45e104bf1ff44630-ikeys-KeyPairProviderManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-2e998826d6036014cdd504594af76975-KeyPairProviderManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p45e104bf1f_ikeys_KeyPairProviderManagerImpl) new() any {
    return &p45e104bf1.KeyPairProviderManagerImpl{}
}

func (inst* p45e104bf1f_ikeys_KeyPairProviderManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p45e104bf1.KeyPairProviderManagerImpl)
	nop(ie, com)

	
    com.Providers = inst.getProviders(ie)


    return nil
}


func (inst*p45e104bf1f_ikeys_KeyPairProviderManagerImpl) getProviders(ie application.InjectionExt)[]p2e998826d.KeyPairProvider{
    dst := make([]p2e998826d.KeyPairProvider, 0)
    src := ie.ListComponents(".class-2e998826d6036014cdd504594af76975-KeyPairProvider")
    for _, item1 := range src {
        item2 := item1.(p2e998826d.KeyPairProvider)
        dst = append(dst, item2)
    }
    return dst
}



// type pfe1800bf5.Provider in package:github.com/bitwormhole/git-acme-commands/app/implements/ikeys/pecdsa
//
// id:com-fe1800bf56aac5b5-pecdsa-Provider
// class:class-2e998826d6036014cdd504594af76975-KeyPairProvider
// alias:
// scope:singleton
//
type pfe1800bf56_pecdsa_Provider struct {
}

func (inst* pfe1800bf56_pecdsa_Provider) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fe1800bf56aac5b5-pecdsa-Provider"
	r.Classes = "class-2e998826d6036014cdd504594af76975-KeyPairProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfe1800bf56_pecdsa_Provider) new() any {
    return &pfe1800bf5.Provider{}
}

func (inst* pfe1800bf56_pecdsa_Provider) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfe1800bf5.Provider)
	nop(ie, com)

	


    return nil
}



// type p169246a6f.Provider in package:github.com/bitwormhole/git-acme-commands/app/implements/ikeys/prsa
//
// id:com-169246a6f60e9b4b-prsa-Provider
// class:class-2e998826d6036014cdd504594af76975-KeyPairProvider
// alias:
// scope:singleton
//
type p169246a6f6_prsa_Provider struct {
}

func (inst* p169246a6f6_prsa_Provider) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-169246a6f60e9b4b-prsa-Provider"
	r.Classes = "class-2e998826d6036014cdd504594af76975-KeyPairProvider"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p169246a6f6_prsa_Provider) new() any {
    return &p169246a6f.Provider{}
}

func (inst* p169246a6f6_prsa_Provider) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p169246a6f.Provider)
	nop(ie, com)

	


    return nil
}


