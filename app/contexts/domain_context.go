package contexts

import "github.com/starter-go/afs"

// DomainContext 包含关于一个域名的上下文信息
type DomainContext struct {
	Parent *CertRepoContext

	DomainName string

	DomainDirectory afs.Path

	CurrentCertFile afs.Path
	LatestCertFile  afs.Path
}
