package helpers

import "github.com/spf13/cobra"

func GetSubCommandsAsSlice(cmd *cobra.Command) []string {
	validArgs := cmd.ValidArgs
	for _, c := range cmd.Commands() {
		validArgs = append(validArgs, c.Name())
	}
	return validArgs
}
