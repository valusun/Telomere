package cli

import (
	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

func NewRootCmd(service *workspace.Service) *cobra.Command {
	root := &cobra.Command{
		Use:           "telomere",
		Short:         "Telomere CLI",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	root.AddCommand(NewCreateCmd(service))
	root.AddCommand(NewListCmd(service))
	root.AddCommand(NewPathCmd(service))
	root.AddCommand(NewKillCmd(service))
	root.AddCommand(NewGCCmd(service))
	root.AddCommand(NewExtendCmd(service))
	return root
}
