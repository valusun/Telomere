package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

type Workspace struct {
	ID        string
	Name      string
	Path      string
	CreatedAt int64
	ExpiresAt int64
}

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

func ListWorkspaces(conn *sql.DB) ([]Workspace, error) {
	rows, err := conn.Query("SELECT id, name, path, created_at, expires_at FROM workspaces ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("list workspaces: %w", err)
	}
	defer rows.Close()

	var workspaces []Workspace
	for rows.Next() {
		var workspace Workspace
		if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Path, &workspace.CreatedAt, &workspace.ExpiresAt); err != nil {
			return nil, fmt.Errorf("list workspaces: %w", err)
		}
		workspaces = append(workspaces, workspace)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("list workspaces: %w", err)
	}
	return workspaces, nil
}

func DeleteWorkspace(conn *sql.DB, name string) error {
	_, err := conn.Exec("DELETE FROM workspaces WHERE name = ?", name)
	if err != nil {
		return fmt.Errorf("delete workspace: %w", err)
	}
	return nil
}
