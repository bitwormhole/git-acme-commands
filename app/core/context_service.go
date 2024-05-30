package core

import "context"

// Service 提供一组创建上下文的接口
type Service interface {

	// NewGitRepoContext(c context.Context) (*GitContext, error)

	// NewCertRepoContext(c context.Context) (*ContainerContext, error)

	// NewDomainContext(parent *ContainerContext, domain string) (*DomainContext, error)

	LoadGitContext(c context.Context) (*GitContext, error)

	LoadContainerContext(c context.Context) (*ContainerContext, error)

	LoadDomainContext(c context.Context) (*DomainContext, error)
}
