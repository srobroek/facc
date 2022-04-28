package require

import "github.com/spf13/cobra"

func createSubCmd(parent *cobra.Command, child *cobra.Command) {
	parent.AddCommand(child)
}
