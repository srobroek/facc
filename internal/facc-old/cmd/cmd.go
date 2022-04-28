package cmdold

// import (
// 	"fmt"
// 	"os"
// 	"os/signal"

// 	"github.com/pterm/pcli"
// 	"github.com/pterm/pterm"
// 	"github.com/spf13/cobra"
// )

// var rootCmd = &cobra.Command{
// 	Use:     "facc",
// 	Short:   "Facc your firewall",
// 	Long:    ``,
// 	Example: `facc `,
// 	Version: version,
// }

// // var subCmds = []cobra.Command{
// // 	addSubCmd(),
// // 	getSubCmd(),
// // 	setSubCmd(),
// // 	deleteSubCmd(),
// // }

// var globalUsage = `The Firewall As Code Configurator
// Common actions for Facc:
// - facc add:    search for charts
// - facc get:      download a chart to your local directory to view
// - facc set:   upload the chart to Kubernetes
// - facc delete:      list releases of charts
// `

// // Environment variables:
// // | Name                               | Description                                                                       |
// // |------------------------------------|-----------------------------------------------------------------------------------|
// // | $HELM_CACHE_HOME                   | set an alternative location for storing cached files.                             |
// // | $HELM_CONFIG_HOME                  | set an alternative location for storing Helm configuration.                       |
// // | $HELM_DATA_HOME                    | set an alternative location for storing Helm data.                                |
// // | $HELM_DEBUG                        | indicate whether or not Helm is running in Debug mode                             |
// // | $HELM_DRIVER                       | set the backend storage driver. Values are: configmap, secret, memory, sql.       |
// // | $HELM_DRIVER_SQL_CONNECTION_STRING | set the connection string the SQL storage driver should use.                      |
// // | $HELM_MAX_HISTORY                  | set the maximum number of helm release history.                                   |
// // | $HELM_NAMESPACE                    | set the namespace used for the helm operations.                                   |
// // | $HELM_NO_PLUGINS                   | disable plugins. Set HELM_NO_PLUGINS=1 to disable plugins.                        |
// // | $HELM_PLUGINS                      | set the path to the plugins directory                                             |
// // | $HELM_REGISTRY_CONFIG              | set the path to the registry config file.                                         |
// // | $HELM_REPOSITORY_CACHE             | set the path to the repository cache directory                                    |
// // | $HELM_REPOSITORY_CONFIG            | set the path to the repositories file.                                            |
// // | $KUBECONFIG                        | set an alternative Kubernetes configuration file (default "~/.kube/config")       |
// // | $HELM_KUBEAPISERVER                | set the Kubernetes API Server Endpoint for authentication                         |
// // | $HELM_KUBECAFILE                   | set the Kubernetes certificate authority file.                                    |
// // | $HELM_KUBEASGROUPS                 | set the Groups to use for impersonation using a comma-separated list.             |
// // | $HELM_KUBEASUSER                   | set the Username to impersonate for the operation.                                |
// // | $HELM_KUBECONTEXT                  | set the name of the kubeconfig context.                                           |
// // | $HELM_KUBETOKEN                    | set the Bearer KubeToken used for authentication.                                 |
// // Helm stores cache, configuration, and data based on the following configuration order:
// // - If a HELM_*_HOME environment variable is set, it will be used
// // - Otherwise, on systems supporting the XDG base directory specification, the XDG variables will be used
// // - When no other location is set a default location will be used based on the operating system
// // By default, the default directories depend on the Operating System. The defaults are listed below:
// // | Operating System | Cache Path                | Configuration Path             | Data Path               |
// // |------------------|---------------------------|--------------------------------|-------------------------|
// // | Linux            | $HOME/.cache/helm         | $HOME/.config/helm             | $HOME/.local/share/helm |
// // | macOS            | $HOME/Library/Caches/helm | $HOME/Library/Preferences/helm | $HOME/Library/helm      |
// // | Windows          | %TEMP%\helm               | %APPDATA%\helm                 | %APPDATA%\helm          |
// // `

// // var globalConfig = viper.New()

// // //Flag pointers
// // var (
// // 	configDirFlag string
// // )

// // Execute adds all child commands to the root command and sets flags appropriately.
// // This is called by main.main(). It only needs to happen once to the rootCmd.
// func Execute() (err error) {
// 	// Fetch user interrupt
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	go func() {
// 		<-c
// 		pterm.DefaultBasicText.Println("user interrupt")
// 		_ = pcli.CheckForUpdates()
// 	}()
// 	if err = initialiseRootCmd(); err != nil {
// 		fmt.Errorf("initialising Root command failed with: %w", err)

// 	}
// 	return
// }

// func initialiseRootCmd() (err error) {

// 	rootCmd.PersistentFlags().BoolVarP(&pterm.PrintDebugMessages, "debug", "", false, "enable debug messages")
// 	rootCmd.PersistentFlags().BoolVarP(&pterm.RawOutput, "raw", "", false, "print unstyled raw output (set it if output is written to a file)")
// 	rootCmd.PersistentFlags().BoolVarP(&pcli.DisableUpdateChecking, "disable-update-checks", "", false, "disables update checks")
// 	rootCmd.PersistentFlags().StringVarP(&configDirFlag, "configdir", "c", "", "configuration directory")
// 	rootCmd.MarkPersistentFlagDirname("configdir")

// 	for _, s := range subCmds {
// 		fmt.Println(s)
// 		rootCmd.AddCommand(&s)
// 	}
// 	// addSubCmd, err := addSubCmd()
// 	// rootCmd.AddCommand(addSubCmd)
// 	// rootCmd.AddCommand(deleteSubCmd())
// 	// rootCmd.AddCommand(getSubCmd())
// 	// rootCmd.AddCommand(setSubCmd())

// 	globalConfig.BindPFlags(rootCmd.PersistentFlags())

// 	//rootCmd.Flags().String("test", "", "test")
// 	//rootCmd.AddCommand(backend.Execute())
// 	//rootCmd.AddCommand(InitialiseVersionCmd())
// 	//rootCmd.AddCommand(InitialiseAddCmd())
// 	// rootCmd.AddCommand(get.Initialise())
// 	// rootCmd.AddCommand(set.Initialise())
// 	//rootCmd.AddCommand(InitialiseDeleteCmd())
// 	// rootCmd.AddCommand(apply.Initialise())
// 	// rootCmd.AddCommand(destroy.Initialise())

// 	// Use https://github.com/pterm/pcli to style the output of cobra.
// 	_ = pcli.SetRepo("srobroek/facc")
// 	pcli.SetRootCmd(rootCmd)
// 	pcli.Setup()

// 	// Change global PTerm theme
// 	pterm.ThemeDefault.SectionStyle = *pterm.NewStyle(pterm.FgGray)

// 	// Execute cobra
// 	if err := rootCmd.Execute(); err != nil {
// 		_ = pcli.CheckForUpdates()
// 		os.Exit(1)
// 	}

// 	_ = pcli.CheckForUpdates()
// 	return
// }
