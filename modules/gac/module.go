package gac

import (
	gitacmecommands "github.com/bitwormhole/git-acme-commands"
	"github.com/bitwormhole/git-acme-commands/gen/main4gac"
	"github.com/bitwormhole/git-acme-commands/gen/test4gac"
	"github.com/bitwormhole/gitlib/modules/gitlib"
	"github.com/starter-go/application"
	"github.com/starter-go/cli/modules/cli"
	"github.com/starter-go/keys/modules/keys"
	"github.com/starter-go/starter"
)

// Module  ...
func Module() application.Module {
	mb := gitacmecommands.NewMainModule()
	mb.Components(main4gac.ExportComponents)

	mb.Depend(starter.Module())
	mb.Depend(cli.Module())
	mb.Depend(gitlib.Module())
	mb.Depend(keys.ModuleForLib())

	mb.Depend(cli.ModuleExtention())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := gitacmecommands.NewTestModule()
	mb.Components(test4gac.ExportComponents)
	mb.Depend(Module())
	return mb.Create()
}
