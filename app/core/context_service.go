package core

import (
	"context"

	"github.com/starter-go/afs"
)

// Service 提供一组创建上下文的接口
type Service interface {
	LoadGitContext(c context.Context) (*GitContext, error)

	LoadContainerContext(c context.Context) (*ContainerContext, error)

	LoadDomainContext(c context.Context) (*DomainContext, error)

	LoadDomainContextWithConfigFile(c context.Context, cfg afs.Path) (*DomainContext, error)
}
