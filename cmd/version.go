package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	bla    string
	commit = "n/a"
	date   = "n/a"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = fmt.Sprintf("%s-%s (%s)", bla, commit, date)
	},
}
