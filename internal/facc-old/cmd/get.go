package cmdold

import "github.com/spf13/cobra"

func getSubCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use: "get",
	}
	cmd.AddCommand(getBackendSubCmd())
	return cmd
}
