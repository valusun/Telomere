package main

import (
	"fmt"
	"log"
	"os"

	"github.com/valusun/Telomere/internal/db"
	"github.com/valusun/Telomere/internal/workspace"
)

func makeDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

func makeDatabase() error {
	conn, err := db.Open()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = db.Initialize(conn)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	return nil
}

func main() {
	// MkdirAll が ~/.telomere ごと作るので、workspaces までを一度に掘れば足りる
	workspacesPath, err := workspace.Dir()
	if err != nil {
		log.Fatal(err)
	}

	err = makeDirectory(workspacesPath)
	if err != nil {
		log.Fatal(err)
	}

	err = makeDatabase()
	if err != nil {
		log.Fatal(err)
	}
}
