package cmdold

import "github.com/spf13/cobra"

func setSubCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use: "set",
	}
	return cmd
}
