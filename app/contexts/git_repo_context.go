package contexts

import (
	"context"

	"github.com/bitwormhole/gitlib/git/repositories"
	"github.com/starter-go/afs"
)

// GitRepoContext 包含关于一个 git 仓库的上下文信息
type GitRepoContext struct {
	Parent   context.Context
	Layout   repositories.Layout
	WD       afs.Path
	Worktree afs.Path
}
