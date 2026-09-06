package main

import (
	"fmt"
	"os"

	"github.com/valusun/Telomere/internal/cli"
	"github.com/valusun/Telomere/internal/db"
	"github.com/valusun/Telomere/internal/workspace"
)

func run() error {
	conn, err := db.Open()
	if err != nil {
		return err
	}
	defer conn.Close()

	repository := workspace.NewRepository(conn)
	service := workspace.NewService(repository)
	return cli.NewRootCommand(service).Execute()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
