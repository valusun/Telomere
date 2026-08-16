package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/db"
	"github.com/valusun/Telomere/internal/workspace"
)

var killCmd = &cobra.Command{
	Use:   "kill <name>",
	Short: "Delete a workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		path, err := workspace.Delete(name)
		if err != nil {
			if errors.Is(err, db.ErrWorkspaceNotFound) {
				return fmt.Errorf("workspace %q not found", name)
			}
			return err
		}
		fmt.Println("workspace deleted: " + path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
}
