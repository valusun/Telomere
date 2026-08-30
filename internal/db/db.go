package db

import (
	"database/sql"
	"fmt"

	"github.com/valusun/Telomere/internal/config"
)

func Open() (*sql.DB, error) {
	paths, err := config.SetTelomerePaths()
	if err != nil {
		return nil, err
	}
	conn, err := sql.Open("sqlite3", paths.DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	return conn, nil
}
