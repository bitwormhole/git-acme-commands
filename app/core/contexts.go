package core

import (
	"context"
	"crypto"
	"crypto/x509"
	"time"

	"github.com/bitwormhole/git-acme-commands/app/data/dto"
	"github.com/bitwormhole/git-acme-commands/app/data/vo"
	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// GitContext 包含关于一个 git 仓库的上下文信息
type GitContext struct {
	Parent   context.Context
	Layout   repositories.Layout
	WD       afs.Path
	Worktree afs.Path
}

// ContainerContext 包含关于一个证书仓库的上下文信息
type ContainerContext struct {
	Parent *GitContext

	// files
	MainConfigFile afs.Path // at '.git/../acme.config'

	Config *vo.ContainerConfig

	// config
	// ConfigProperties properties.Table
	// MainConfig  *config.VO
	// LocalConfig *config.VO
	// MixedConfig *config.VO

	// domains
	DomainList []*DomainListItem
	Domains    map[string]*DomainListItem

	// session
	UserName   string
	UserEmail  string
	UserSigner crypto.Signer

	// time
	Now             time.Time
	SessionTime     time.Time
	SessionInterval time.Duration
}

// DomainContext 包含关于一个域名的上下文信息
type DomainContext struct {
	Parent *ContainerContext

	DomainDirectory  afs.Path
	DomainConfigFile afs.Path // 'domain.config'
	CurrentCertFile  afs.Path
	LatestCertFile   afs.Path

	Config *vo.DomainConfig

	DomainName   dto.DomainName
	DomainSigner crypto.Signer

	// DomainConfig     properties.Table
	// Current *CertificateContext
	// Latest  *CertificateContext
}

// CertificateContext ...
type CertificateContext struct {
	Parent *DomainContext

	CertFile   afs.Path
	Cert       x509.Certificate
	CertSigner crypto.Signer
}
