package main

import (
	"fmt"
	"log"

	"github.com/valusun/Telomere/internal/config"
	"github.com/valusun/Telomere/internal/db"
)

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
