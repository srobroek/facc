package cli

import (
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// var subCmds = []cobra.Command{
// 	addSubCmd(),
// 	getSubCmd(),
// 	setSubCmd(),
// 	deleteSubCmd(),
// }

//filled by linker
var (
	version string = "none"
	commit         = "none"
	date           = "none"
	builtby        = "none"
)

//Flags
var (
	configDirFlag string
	logDirFlag    string
	logLevelFlag  string
)

var cliConfig = viper.New()

var globalUsage = `The Firewall As Code Configurator
Common actions for Facc:
- facc backend: Manage backends
- facc show:   Show configuration elements
- facc update: Update configuration elements
- facc delete: Delete configuration elements
`

// Environment variables:
// | Name                               | Description                                                                       |
// |------------------------------------|-----------------------------------------------------------------------------------|
// | $HELM_CACHE_HOME                   | set an alternative location for storing cached files.                             |
// | $HELM_CONFIG_HOME                  | set an alternative location for storing Helm configuration.                       |
// | $HELM_DATA_HOME                    | set an alternative location for storing Helm data.                                |
// | $HELM_DEBUG                        | indicate whether or not Helm is running in Debug mode                             |
// | $HELM_DRIVER                       | set the backend storage driver. Values are: configmap, secret, memory, sql.       |
// | $HELM_DRIVER_SQL_CONNECTION_STRING | set the connection string the SQL storage driver should use.                      |
// | $HELM_MAX_HISTORY                  | set the maximum number of helm release history.                                   |
// | $HELM_NAMESPACE                    | set the namespace used for the helm operations.                                   |
// | $HELM_NO_PLUGINS                   | disable plugins. Set HELM_NO_PLUGINS=1 to disable plugins.                        |
// | $HELM_PLUGINS                      | set the path to the plugins directory                                             |
// | $HELM_REGISTRY_CONFIG              | set the path to the registry config file.                                         |
// | $HELM_REPOSITORY_CACHE             | set the path to the repository cache directory                                    |
// | $HELM_REPOSITORY_CONFIG            | set the path to the repositories file.                                            |
// | $KUBECONFIG                        | set an alternative Kubernetes configuration file (default "~/.kube/config")       |
// | $HELM_KUBEAPISERVER                | set the Kubernetes API Server Endpoint for authentication                         |
// | $HELM_KUBECAFILE                   | set the Kubernetes certificate authority file.                                    |
// | $HELM_KUBEASGROUPS                 | set the Groups to use for impersonation using a comma-separated list.             |
// | $HELM_KUBEASUSER                   | set the Username to impersonate for the operation.                                |
// | $HELM_KUBECONTEXT                  | set the name of the kubeconfig context.                                           |
// | $HELM_KUBETOKEN                    | set the Bearer KubeToken used for authentication.                                 |
// Helm stores cache, configuration, and data based on the following configuration order:
// - If a HELM_*_HOME environment variable is set, it will be used
// - Otherwise, on systems supporting the XDG base directory specification, the XDG variables will be used
// - When no other location is set a default location will be used based on the operating system
// By default, the default directories depend on the Operating System. The defaults are listed below:
// | Operating System | Cache Path                | Configuration Path             | Data Path               |
// |------------------|---------------------------|--------------------------------|-------------------------|
// | Linux            | $HOME/.cache/helm         | $HOME/.config/helm             | $HOME/.local/share/helm |
// | macOS            | $HOME/Library/Caches/helm | $HOME/Library/Preferences/helm | $HOME/Library/helm      |
// | Windows          | %TEMP%\helm               | %APPDATA%\helm                 | %APPDATA%\helm          |
// `

// func newRootCmd(actionConfig *action.Configuration, out io.Writer, args []string) (*cobra.Command, error) {
func newRootCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
		Use:          "facc",
		Short:        "The Firewall as Code Configurator",
		Long:         globalUsage,
		SilenceUsage: true,
		Version:      "0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			u, err := cmd.Flags().GetBool("update")
			if u == true {
				pcli.CheckForUpdates()

			} else {
				cmd.Help()
			}
			if err != nil {
				return err
			}
			return nil
		},
	}
	//flags := cmd.PersistentFlags()
	cmd.PersistentFlags().StringVarP(&configDirFlag, "configdir", "c", "./config", "configuration directory")
	cmd.PersistentFlags().StringVarP(&logDirFlag, "log", "l", "", "Log directory")
	cmd.MarkPersistentFlagDirname(("log"))
	cmd.MarkPersistentFlagDirname("configdir")
	cmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
	cmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
	cmd.Flags().Bool("update", false, "perform an update check")
	cliConfig.BindPFlags(cmd.PersistentFlags())
	//pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgBlack)
	pterm.ThemeDefault.DescriptionPrefixStyle = *pterm.NewStyle(pterm.BgLightGreen)
	pterm.ThemeDefault.BarLabelStyle = *pterm.NewStyle(pterm.BgBlue)

	cmd.AddCommand(
		//crud commands
		backendCmd(),
		//applicationCmd(),
		//stackCmd(),

		//createBackendCmd(),
		//showBackendCmd(),
		//updateBackendCmd(),
		//deleteBackendCmd(),
	)

	// settings.AddFlags(flags)
	// addKlogFlags(flags)
	// cmd.AddCommand(
	// //backend commands
	// //createBackendCmd(out),
	// //showBackendCmd(out),
	// //updateBackendCmd(out),
	// //deleteBackendCmd(out),
	// )

	// Change global PTerm theme

	// Execute cobra
	// if err := cmd.Execute(); err != nil {
	// 	return nil, err
	// 	//os.Exit(1)
	// }

	return cmd, nil
}
