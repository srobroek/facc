package cmd

import (
	"os"
	"os/signal"

	_ "embed"

	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// var (
// 	//Variables are all set by goReleaser
// 	version string
// 	commit  string
// 	date    string
// )

var rootCmd = &cobra.Command{
	Use:   "facc",
	Short: "Facc your firewall",
	Long:  ``,
	Example: `facc 
facc
`,
	Version: bla, // <---VERSION---> Updating this version, will also create a new GitHub release.
	// Uncomment the following lines if your bare application has an action associated with it:
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	// Your code here
	//
	// 	return nil
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		_ = pcli.CheckForUpdates()
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		_ = pcli.CheckForUpdates()
		os.Exit(1)
	}
	print(bla)

	_ = pcli.CheckForUpdates()
}

func Initialise() {
	// Adds global flags for PTerm settings.
	// Fill the empty strings with the shorthand variant (if you like to have one).
	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "disables update checks")
	rootCmd.AddCommand(versionCmd)

	// Use https://github.com/pterm/pcli to style the output of cobra.
	_ = pcli.SetRepo("srobroek/facc")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()

	// Change global PTerm theme
	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgCyan)
}
