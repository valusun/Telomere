package workspace

import "errors"

type Workspace struct {
	ID        string
	Name      string
	Path      string
	CreatedAt int64
	ExpiresAt int64
}

var (
	ErrWorkspaceNotFound   = errors.New("workspace not found")
	ErrWorkspaceNameExists = errors.New("workspace name already exists")
)
