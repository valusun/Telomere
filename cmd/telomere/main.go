package main

import (
	"fmt"
	"os"

	"github.com/valusun/Telomere/internal/cli"
)

func main() {
	err := cli.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "telomere:", err)
		os.Exit(1)
	}
}
