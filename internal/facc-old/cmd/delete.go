package cmdold

import "github.com/spf13/cobra"

func deleteSubCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use: "delete",
	}
	return cmd
}
