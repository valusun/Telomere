package main

import (
	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/layout"
	"github.com/valusun/Telomere/internal/workspace"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all workspaces",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		asJSON, err := cmd.Flags().GetBool("json")
		if err != nil {
			return err
		}

		workspaces, err := workspace.List()
		if err != nil {
			return err
		}

		rows := make([]layout.WorkspaceView, 0, len(workspaces))
		for _, w := range workspaces {
			rows = append(rows, layout.WorkspaceView{
				Name:      w.Name,
				CreatedAt: w.CreatedAt,
				ExpiresAt: w.ExpiresAt,
			})
		}

		if asJSON {
			return layout.ViewListJSON(rows)
		}
		return layout.ViewList(rows)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("json", false, "output as JSON")
}
