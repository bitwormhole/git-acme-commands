package contexts

import "context"

// Service 提供一组创建上下文的接口
type Service interface {
	NewGitRepoContext(c context.Context) (*GitRepoContext, error)

	NewCertRepoContext(c context.Context) (*CertRepoContext, error)

	NewDomainContext(parent *CertRepoContext, domain string) (*DomainContext, error)
}
