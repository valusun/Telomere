package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/valusun/Telomere/internal/workspace"
)

var gcCmd = &cobra.Command{
	Use:    "gc",
	Short:  "Delete expired workspaces",
	Args:   cobra.NoArgs,
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Deleting expired workspaces...")
		workspaces, err := workspace.FindExpiredWorkspaces()
		if err != nil {
			return err
		}

		failed := 0
		for _, w := range workspaces {
			path, err := workspace.Delete(w.Name)
			if err != nil {
				// 1件の失敗で打ち切ると残りが溜まり続けるため、最後まで走らせて件数だけ返す
				fmt.Fprintf(os.Stderr, "failed to delete workspace %q: %v\n", w.Name, err)
				failed++
				continue
			}
			fmt.Println("workspace deleted: " + path)
		}

		if failed > 0 {
			return fmt.Errorf("failed to delete %d of %d expired workspaces", failed, len(workspaces))
		}
		fmt.Println("all expired workspaces deleted")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gcCmd)
}
