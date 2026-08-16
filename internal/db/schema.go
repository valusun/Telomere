package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func makeWorkspaces(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS workspaces (
			id text Primary key NOT NULL,
			name text NOT NULL,
			path text NOT NULL,
			created_at integer NOT NULL,
			expires_at integer NOT NULL,
			CONSTRAINT expires_check CHECK (expires_at > created_at)
		);
		CREATE UNIQUE INDEX IF NOT EXISTS workspace_name_unique_index ON workspaces (name);
		CREATE UNIQUE INDEX IF NOT EXISTS workspace_path_unique_index ON workspaces (path);
	`)
	if err != nil {
		return fmt.Errorf("create workspaces table: %w", err)
	}
	return nil
}

func Initialize(conn *sql.DB) error {
	return makeWorkspaces(conn)
}
