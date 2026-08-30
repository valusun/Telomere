package main

import (
	"fmt"
	"os"

	"github.com/valusun/Telomere/internal/cli"
	"github.com/valusun/Telomere/internal/db"
	"github.com/valusun/Telomere/internal/workspace"
)

func main() {
	conn, err := db.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "telomere:", err)
		os.Exit(1)
	}

	repository := workspace.NewRepository(conn)
	service := workspace.NewService(repository)
	root := cli.NewRootCommand(service)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "telomere:", err)
		os.Exit(1)
	}
}
