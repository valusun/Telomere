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

func Dir() (string, error) {
	root, err := config.Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, "workspaces"), nil
}

func Create(name string, ttlDays int) (string, error) {
	id := uuid.NewString()
	expiresAt := time.Now().AddDate(0, 0, ttlDays).Unix()

	// create target workspace directory
	workspaceDir, err := Dir()
	if err != nil {
		return "", fmt.Errorf("failed to get workspace dir: %w", err)
	}
	path := filepath.Join(workspaceDir, id)
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create workspace dir: %w", err)
	}

	conn, err := db.Open()
	if err != nil {
		return "", err
	}
	defer conn.Close()

	err = db.InsertWorkspace(conn, id, name, path, time.Now().Unix(), expiresAt)
	if err != nil {
		// 実害はないが邪魔なので消しておく
		os.RemoveAll(path)
		return "", fmt.Errorf("failed to insert workspace: %w", err)
	}

	return path, nil
}

func Find(name string) (string, error) {
	conn, err := db.Open()
	if err != nil {
		return "", err
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

	path, err := db.FindWorkspace(conn, name)
	if err != nil {
		return "", err
	}

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
