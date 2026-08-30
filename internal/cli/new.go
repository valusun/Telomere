package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

// NewNewCommandになるのは気持ち悪いので仕方なく…
func NewCreatecmd(service *workspace.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new <name>",
		Short: "Create a new empty workspace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ttlText, err := cmd.Flags().GetString("ttl")
			if err != nil {
				return err
			}

			name := args[0]
			created, err := service.Create(cmd.Context(), name, ttlText)
			if err != nil {
				return fmt.Errorf("create workspace: %w", err)
			}

			fmt.Printf("✓ workspace %q created\n", created.Name)
			fmt.Printf("  %s\n", created.Path)
			return nil
		},
	}

	cmd.Flags().StringP("ttl", "t", "", "workspace TTL (e.g. 3d)")
	_ = cmd.MarkFlagRequired("ttl")

	return cmd
}
