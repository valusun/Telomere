package cli

import (
	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/ttl"
	"github.com/valusun/Telomere/internal/workspace"
)

var extendCmd = &cobra.Command{
	Use:   "extend <name>",
	Short: "Extend telomere expired",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		ttlStr := cmd.Flag("ttl").Value.String()
		ttlDays, err := ttl.Parse(ttlStr)
		if err != nil {
			return err
		}
		err = workspace.ExtendExpiry(name, ttlDays)
		if err != nil {
			return err
		}
		cmd.Println("Successfully extended ttl")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(extendCmd)
	extendCmd.Flags().StringP("ttl", "t", "", "workspace current expiredAt add <ttl> (e.g. 3d)")
	extendCmd.MarkFlagRequired("ttl")
}
