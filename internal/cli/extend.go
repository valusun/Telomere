package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

func NewExtendCmd(service *workspace.Service) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend <name>",
		Short: "Extend telomere expired",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]
			ttl, err := cmd.Flags().GetString("ttl")
			if err != nil {
				return err
			}
			if err := service.ExtendExpiry(cmd.Context(), name, ttl); err != nil {
				return err
			}
			fmt.Println("Successfully extended ttl")
			return nil
		},
	}
	cmd.Flags().StringP("ttl", "t", "", "workspace current expiredAt add <ttl> (e.g. 3d)")
	_ = cmd.MarkFlagRequired("ttl")
	return cmd
}
