package workspace

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

type Repository struct{ db *sql.DB }

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Insert(ctx context.Context, w Workspace) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO workspaces (id, name, path, created_at, expires_at) VALUES (?, ?, ?, ?, ?)", w.ID, w.Name, w.Path, w.CreatedAt, w.ExpiresAt)
	if err == nil {
		return nil
	}

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
		return ErrWorkspaceNameExists
	}
	return fmt.Errorf("failed to save workspace: %w", err)
}

func (r *Repository) GetWorkspaces(ctx context.Context) ([]Workspace, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, path, created_at, expires_at FROM workspaces ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("failed to get workspaces: %w", err)
	}
	defer rows.Close()

	var workspaces []Workspace
	for rows.Next() {
		var workspace Workspace
		if err := rows.Scan(&workspace.ID, &workspace.Name, &workspace.Path, &workspace.CreatedAt, &workspace.ExpiresAt); err != nil {
			return nil, fmt.Errorf("failed to get workspaces: %w", err)
		}
		workspaces = append(workspaces, workspace)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get workspaces: %w", err)
	}
	return workspaces, nil
}

func (r *Repository) FindByName(ctx context.Context, name string) (Workspace, error) {
	var ws Workspace
	err := r.db.QueryRowContext(ctx, "SELECT id, name, path, created_at, expires_at FROM workspaces WHERE name = ?", name).Scan(&ws.ID, &ws.Name, &ws.Path, &ws.CreatedAt, &ws.ExpiresAt)
	if errors.Is(err, sql.ErrNoRows) {
		return ws, ErrWorkspaceNotFound
	}
	if err != nil {
		return ws, fmt.Errorf("failed to find workspace: %w", err)
	}
	return ws, nil
}

func (r *Repository) Delete(ctx context.Context, name string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM workspaces WHERE name = ?", name)
	if err != nil {
		return fmt.Errorf("failed to delete workspace: %w", err)
	}
	return nil
}

func (r *Repository) UpdateExpiresAt(ctx context.Context, name string, expiresAt int64) error {
	_, err := r.db.ExecContext(ctx, "UPDATE workspaces SET expires_at = ? WHERE name = ?", expiresAt, name)
	if err != nil {
		return fmt.Errorf("failed to update workspace expiry: %w", err)
	}
	return nil
}
