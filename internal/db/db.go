package db

import (
	"database/sql"
	"fmt"
	"path/filepath"

	"github.com/valusun/Telomere/internal/config"
)

func Open() (*sql.DB, error) {
	dir, err := config.Dir()
	if err != nil {
		return nil, err
	}
	conn, err := sql.Open("sqlite3", filepath.Join(dir, "telomere.db"))
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return conn, nil
}
