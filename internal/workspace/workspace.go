package workspace

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/valusun/Telomere/internal/db"
)

func Create(name string, ttlDays int) (string, error) {
	id := uuid.NewString()
	expiresAt := time.Now().AddDate(0, 0, ttlDays).Unix()

	// create target workspace directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home dir: %w", err)
	}
	path := filepath.Join(home, ".telomere", "workspaces", id)
	err = os.MkdirAll(path, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create workspace dir: %w", err)
	}

	dbPath := filepath.Join(home, ".telomere", "telomere.db")
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return "", fmt.Errorf("failed to open database: %w", err)
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
