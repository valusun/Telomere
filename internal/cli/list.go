package cli

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

func NewListCmd(service *workspace.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all workspaces",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			asJson, err := cmd.Flags().GetBool("json")
			if err != nil {
				return err
			}

			ws, err := service.List(cmd.Context())
			if err != nil {
				return err
			}

			rows := make([]WorkspaceView, 0, len(ws))
			for _, w := range ws {
				rows = append(rows, WorkspaceView{
					Name:      w.Name,
					Path:      w.Path,
					CreatedAt: w.CreatedAt,
					ExpiresAt: w.ExpiresAt,
				})
			}
			if asJson {
				return ViewListJSON(rows)
			}
			return ViewList(rows)
		},
	}
	cmd.Flags().Bool("json", false, "output as a JSON")
	err := cmd.Flags().MarkHidden("json")
	if err != nil {
		log.Fatal("failed to mark json flag as hidden: ", err)
	}
	return cmd
}
