package main

import (
	"os"

	"github.com/bitwormhole/git-acme-commands/modules/gac"
	"github.com/starter-go/starter"
)

func main() {
	m := gac.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
