package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

var ErrWorkspaceNameExists = errors.New("workspace name already exists")
var ErrWorkspaceNotFound = errors.New("workspace not found")

func InsertWorkspace(conn *sql.DB, id, name, path string, createdAt, expiresAt int64) error {
	_, err := conn.Exec("INSERT INTO workspaces (id, name, path, created_at, expires_at) VALUES (?, ?, ?, ?, ?)", id, name, path, createdAt, expiresAt)
	if err == nil {
		return nil
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
		return ErrWorkspaceNameExists
	}
	return fmt.Errorf("insert workspace: %w", err)
}

func FindWorkspace(conn *sql.DB, name string) (string, error) {
	var path string
	err := conn.QueryRow("SELECT path FROM workspaces WHERE name = ?", name).Scan(&path)
	if errors.Is(err, sql.ErrNoRows) {
		return "", ErrWorkspaceNotFound
	}
	if err != nil {
		return "", fmt.Errorf("find workspace: %w", err)
	}
	return path, nil
}
