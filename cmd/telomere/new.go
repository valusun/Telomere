package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/db"
	"github.com/valusun/Telomere/internal/ttl"
	"github.com/valusun/Telomere/internal/workspace"
)

var newCmd = &cobra.Command{
	Use:   "new <name>",
	Short: "Create a new empty workspace",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		ttlStr, _ := cmd.Flags().GetString("ttl")
		ttlDays, err := ttl.Parse(ttlStr)
		if err != nil {
			return fmt.Errorf("invalid ttl: %w", err)
		}
		path, err := workspace.Create(name, ttlDays)
		if err != nil {
			if errors.Is(err, db.ErrWorkspaceNameExists) {
				return fmt.Errorf("workspace %q already exists", name)
			}
			return fmt.Errorf("failed to create workspace: %w", err)
		}
		fmt.Printf("✓ workspace %q created\n", name)
		fmt.Printf("  %s\n\n", path)
		fmt.Printf("→ cd \"$(telomere path %s)\"\n", name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("ttl", "t", "", "workspace date-to-live (e.g. 3d)")
	err := newCmd.MarkFlagRequired("ttl")
	if err != nil {
		log.Fatal("failed to mark ttl flag as required: ", err)
	}
}
