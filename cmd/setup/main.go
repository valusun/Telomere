package main

import (
	"fmt"
	"log"
	"os"

	"github.com/valusun/Telomere/internal/config"
	"github.com/valusun/Telomere/internal/db"
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
	// なければ内部で作成する
	if _, err := config.SetTelomerePaths(); err != nil {
		log.Fatal(err)
	}

	if err := makeDatabase(); err != nil {
		log.Fatal(err)
	}
}
