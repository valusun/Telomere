package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

func NewGCCmd(service *workspace.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:    "gc",
		Short:  "Delete expired workspaces",
		Args:   cobra.NoArgs,
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Deleting expired workspaces...")
			ws, err := service.FindExpiredWorkspaces(cmd.Context())
			if err != nil {
				return err
			}

			failed := 0
			for _, w := range ws {
				path, err := service.Delete(cmd.Context(), w.Name)
				if err != nil {
					fmt.Fprintf(os.Stderr, "failed to delete workspace %q: %v\n", w.Name, err)
					failed++
					continue
				}
				fmt.Println("workspace deleted: " + path)
			}

			if failed > 0 {
				return fmt.Errorf("failed to delete %d of %d expired workspaces", failed, len(ws))
			}
			fmt.Println("all expired workspaces deleted")
			return nil
		},
	}
	return cmd
}
