package cli

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

func NewKillCmd(service *workspace.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "kill <name>",
		Short: "Delete a workspace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			path, err := service.Delete(cmd.Context(), name)
			if err != nil {
				if errors.Is(err, workspace.ErrWorkspaceNotFound) {
					return fmt.Errorf("workspace %q not found", name)
				}
				return err
			}
			fmt.Println("workspace deleted: " + path)
			return nil
		},
	}
}
