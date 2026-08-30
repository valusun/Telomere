package config

import (
	"fmt"
	"os"
	"path/filepath"
)

type Paths struct {
	Root         string
	WorkspaceDir string
	DatabasePath string
}

func makeDirs(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to make dirs: %w", err)
	}
	return nil
}

func SetTelomerePaths() (*Paths, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home dir: %w", err)
	}
	root := filepath.Join(home, ".telomere")
	workspaceDir := filepath.Join(root, "workspaces")

	for _, dir := range []string{root, workspaceDir} {
		if err := makeDirs(dir); err != nil {
			return nil, err
		}
	}

	// DBは接続時に作成されるので、ここで明示的に作成しない
	return &Paths{
		Root:         root,
		WorkspaceDir: workspaceDir,
		DatabasePath: filepath.Join(root, "telomere.db"),
	}, nil
}
