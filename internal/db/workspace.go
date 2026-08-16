package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

var ErrWorkspaceNameExists = errors.New("workspace name already exists")

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
