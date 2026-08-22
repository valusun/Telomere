package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/db"
	"github.com/valusun/Telomere/internal/workspace"
)

var pathCmd = &cobra.Command{
	Use:   "path <name>",
	Short: "View the path of a specific workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		workspace, err := workspace.Find(name)
		if err != nil {
			if errors.Is(err, db.ErrWorkspaceNotFound) {
				return fmt.Errorf("workspace %q not found", name)
			}
			return fmt.Errorf("failed to find workspace: %w", err)
		}
		fmt.Println(workspace.Path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(pathCmd)
}
