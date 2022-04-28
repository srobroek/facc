package cli

import "github.com/spf13/cobra"

const updateDesc = `
This command updates configuration elements
`

func setCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "update",
		Short: "update configurations",
		Long:  updateDesc,
	}
	return cmd
}
