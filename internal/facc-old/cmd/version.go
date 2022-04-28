package cmdold

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// These are filled by the linker
var (
	version string
	commit  string
	date    string
)

// var (
// 	//Variables are all set by goReleaser
// 	version string
// 	commit  string
// 	date    string
// )

func InitialiseVersionCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "version",
		Short: "Shows more detailed version information",
		Run: func(cmd *cobra.Command, args []string) {
			pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("FACC")).Render()
			pterm.Printf("Version: %s\n", version)
			pterm.Printf("Commit: %s\n", commit)
			pterm.Printf("Build Date: %s\n", date)
		},
	}

	return cmd
}
