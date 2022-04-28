package cli

import "github.com/spf13/cobra"

const showDesc = `
This command shows configuration elements
`

func getCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "show",
		Short: "show configurations",
		Long:  showDesc,
	}
	return cmd
}
