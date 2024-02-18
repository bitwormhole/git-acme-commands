package main4gac

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p169246a6f6_prsa_Provider{})
    inst.register(&p45e104bf1f_ikeys_KeyPairProviderManagerImpl{})
    inst.register(&p9daa3c6fa2_icontexts_ContextServiceImpl{})
    inst.register(&pdb462c5c7a_commands_GitACME{})
    inst.register(&pfe1800bf56_pecdsa_Provider{})


    return nil
}
