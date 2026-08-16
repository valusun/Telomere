package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/valusun/Telomere/internal/db"
)

func makeDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

func makeDatabase(path string) error {
	conn, err := sql.Open("sqlite3", filepath.Join(path, "telomere.db"))
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer conn.Close()
	err = db.Initialize(conn)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	return nil
}

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	telomerePath := filepath.Join(home, ".telomere")
	workspacesPath := filepath.Join(telomerePath, "workspaces")

	err = makeDirectory(workspacesPath)
	if err != nil {
		log.Fatal(err)
	}

	err = makeDatabase(telomerePath)
	if err != nil {
		log.Fatal(err)
	}
}
