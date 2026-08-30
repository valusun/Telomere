package workspace

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/valusun/Telomere/internal/config"
	"github.com/valusun/Telomere/internal/db"
)

func Create(name string, ttlDays int) (string, error) {
	id := uuid.NewString()
	expiresAt := time.Now().AddDate(0, 0, ttlDays).Unix()

	// create target workspace directory
	paths, err := config.SetTelomerePaths()
	if err != nil {
		return "", fmt.Errorf("failed to get workspace dir: %w", err)
	}
	workspacePath := filepath.Join(paths.WorkspaceDir, id)
	err = os.MkdirAll(workspacePath, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create workspace dir: %w", err)
	}

	conn, err := db.Open()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	err = db.InsertWorkspace(conn, id, name, workspacePath, time.Now().Unix(), expiresAt)
	if err != nil {
		// 実害はないが邪魔なので消しておく
		os.RemoveAll(workspacePath)
		return "", fmt.Errorf("failed to insert workspace: %w", err)
	}

	return workspacePath, nil
}

func Find(name string) (db.Workspace, error) {
	conn, err := db.Open()
	if err != nil {
		return db.Workspace{}, err
	}
	defer conn.Close()
	return db.FindWorkspace(conn, name)
}

func List() ([]db.Workspace, error) {
	conn, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return db.ListWorkspaces(conn)
}

func Delete(name string) (string, error) {
	conn, err := db.Open()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	workspace, err := db.FindWorkspace(conn, name)
	if err != nil {
		return "", err
	}
	path := workspace.Path

	// DBを先に消しディレクトリの削除に失敗すると、ディレクトリの追跡が不可能になるため
	err = os.RemoveAll(path)
	if err != nil {
		return "", fmt.Errorf("failed to remove workspace dir: %w", err)
	}

	err = db.DeleteWorkspace(conn, name)
	if err != nil {
		return "", err
	}
	return path, nil
}

func FindExpiredWorkspaces() ([]db.Workspace, error) {
	conn, err := db.Open()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return db.FindExpiredWorkspaces(conn)
}

func ExtendExpiry(name string, ttlDays int) error {
	conn, err := db.Open()
	if err != nil {
		return err
	}
	defer conn.Close()
	data, err := db.FindWorkspace(conn, name)
	if err != nil {
		return err
	}

	// expiredAtに加算して更新
	updatedExpiresAt := time.Unix(data.ExpiresAt, 0).AddDate(0, 0, ttlDays)
	err = db.UpdateWorkspaceExpiry(conn, name, updatedExpiresAt.Unix())
	if err != nil {
		return err
	}
	return nil
}
