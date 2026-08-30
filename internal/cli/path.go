package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

func NewPathCmd(service *workspace.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "path <name>",
		Short: "View the path of a specific workspace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			w, err := service.Find(cmd.Context(), name)
			if err != nil {
				if errors.Is(err, workspace.ErrWorkspaceNotFound) {
					return fmt.Errorf("workspace %q not found", name)
				}
				return fmt.Errorf("failed to find workspace: %w", err)
			}
			fmt.Println(w.Path)
			return nil
		},
	}
}
