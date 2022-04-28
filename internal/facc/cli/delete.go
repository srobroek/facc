package cli

import "github.com/spf13/cobra"

const deleteDesc = `
This command deletes configuration elements
`

func deleteCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete configurations",
		Long:  deleteDesc,
	}
	return cmd
}
