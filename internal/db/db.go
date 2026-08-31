package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/valusun/Telomere/internal/config"
	"github.com/valusun/Telomere/internal/workspace"
)

func makeDatabaseFile(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to create database file: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("failed to close database file: %w", err)
	}
	// OpenFile は umask の影響を受けること、既存ファイルのモードは変わらないため揃える
	if err := os.Chmod(path, 0600); err != nil {
		return fmt.Errorf("failed to chmod database file: %w", err)
	}
	return nil
}

func Open() (*sql.DB, error) {
	paths, err := config.SetTelomerePaths()
	if err != nil {
		return nil, err
	}
	if err := makeDatabaseFile(paths.DatabasePath); err != nil {
		return nil, err
	}
	conn, err := sql.Open("sqlite3", paths.DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return conn, nil
}

func Initialize(conn *sql.DB) error {
	return workspace.MakeSchema(conn)
}
