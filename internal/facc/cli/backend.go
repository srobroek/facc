package cli

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/srobroek/facc/internal/facc/require"
	"github.com/tidwall/gjson"
)

// import (
// 	"fmt"
// 	"io"
// 	"path/filepath"

// 	"github.com/spf13/cobra"
// 	"github.com/srobroek/facc/internal/cli/require"
// )

//Flags
var (
	nameFlag         string
	nsxtManagerFlag  string
	nsxtUsernameFlag string
	nsxtInsecureFlag bool
	nsxtGlobalFlag   bool
	nsxtVersionFlag  string
)

type Backends struct {
	Backends []Backend
}

type Backend struct {
	Name          string
	BackendType   string
	Configuration interface{}
}

type NsxtBackendConfiguration struct {
	Manager  string
	Username string
	Insecure bool
	Global   bool
	Version  string
}

type AwsBackendConfiguration struct {
	Region string
	Key    string
}

func (l *Backends) addBackend(b Backend) error {
	jsonData, err := json.Marshal(l)
	if err != nil {
		return fmt.Errorf("failed unmarshaling data: %w", err)
	}
	jsonStr := string(jsonData)
	jsonQuery := fmt.Sprintf("Backends.#(Name==%s)#", b.Name)
	value := gjson.Get(jsonStr, jsonQuery)

	if len(value.Array()) == 1 {

		return fmt.Errorf("backend with name %s already exists", b.Name)

	} else if len(value.Array()) > 1 {
		return fmt.Errorf("multiple objects with name %s exist, check your configuration", b.Name)

		//return err

	}
	l.Backends = append(l.Backends, b)
	return nil
}

func backendCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "backend",
		Short: "manage backends",
		Long:  backendDesc,
		// Args:  require.ExactArgs(1),
		// ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 	if len(args) == 0 {
		// 		// Allow file completion when completing the argument for the name
		// 		// which could be a path
		// 		return nil, cobra.ShellCompDirectiveDefault
		// 	}
		// 	// No more completions, so disable file completion
		// 	return nil, cobra.ShellCompDirectiveNoFileComp
		// },
		// RunE: func(cmd *cobra.Command, args []string) error {
		// 	o.name = args[0]
		// 	return o.run()
		// },
	}
	//cmd.PersistentFlags().VarP()
	//cmd.PersistentFlags().StringVarP(&nameFlag, "name", "n", "", "Name of the backend")

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add a backend",
		Long:  addBackendDesc,
	}
	addNsxtCmd := &cobra.Command{
		Use:   "nsxt NAME",
		Short: "add a nsxt backend",
		Long:  addNsxtBackendDesc,
		Args:  require.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				cmd.Help()

			} else {
				var configPath = cliConfig.GetString("configdir")
				backendConfig := viper.New()
				backendConfig.AddConfigPath(configPath)
				backendConfig.SetConfigName("backends")
				backendConfig.SetConfigType("json")
				backendList := Backends{}
				if err := backendConfig.ReadInConfig(); err != nil {
					return fmt.Errorf("failed reading config: %w", err)
				}

				if err := backendConfig.Unmarshal(&backendList); err != nil {
					return fmt.Errorf("failed unmarshaling data: %w", err)
				}

				config := NsxtBackendConfiguration{
					Manager:  cliConfig.GetString("manager"),
					Username: cliConfig.GetString("username"),
					Insecure: cliConfig.GetBool("insecure"),
					Global:   cliConfig.GetBool("global"),
				}

				backend := Backend{
					Name:          args[0],
					BackendType:   "nsx",
					Configuration: config,
				}

				// if err := backendList.addBackend(backend); err != nil {
				// 	//fmt.Println(err)
				// 	return fmt.Errorf("error: %w", err)

				// }
				if err := backendList.addBackend(backend); err != nil {
					return fmt.Errorf("failed adding backend to list: %w", err)

				}

				jsonData, err := json.Marshal(backendList)

				if err != nil {
					return fmt.Errorf("failed marshaling data: %w", err)
				}
				backendConfig.ReadConfig(bytes.NewBuffer(jsonData))
				pterm.Info.Println("success!")
				err = backendConfig.SafeWriteConfig()
				if err != nil {
					//Try to safe write first, then do unsafe write due to buggy writeconfig
					if err = backendConfig.WriteConfig(); err != nil {
						return fmt.Errorf("failed writing config: %w", err)
					}
					return nil
				}

			}
			pterm.Info.Println("success!")
			return nil
		},
	}
	addCmd.AddCommand(addNsxtCmd)
	addNsxtCmd.Flags().StringVarP(&nsxtManagerFlag, "manager", "m", "", "hostname or ip of the manager")
	addNsxtCmd.Flags().StringVarP(&nsxtUsernameFlag, "user", "u", "admin", "username")
	addNsxtCmd.Flags().BoolVarP(&nsxtInsecureFlag, "insecure", "i", false, "ignore SSL certificate validity")
	addNsxtCmd.Flags().BoolVarP(&nsxtGlobalFlag, "global", "g", false, "use the global manager API")
	addNsxtCmd.MarkFlagRequired("manager")
	cliConfig.BindPFlags(addNsxtCmd.Flags())

	showCmd := &cobra.Command{
		Use:   "show",
		Short: "show backends",
		Long:  showBackendDesc,
		Args:  require.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var configPath = cliConfig.GetString("configdir")
			backendConfig := viper.New()
			backendConfig.AddConfigPath(configPath)
			backendConfig.SetConfigName("backends")
			backendConfig.SetConfigType("json")
			backendConfig.ReadInConfig()
			backendList := Backends{}
			backendConfig.Unmarshal(&backendList)

			if len(args) == 0 {

				table := pterm.DefaultTable.WithHasHeader().WithData(pterm.TableData{
					{"Name", "Type"},
				})
				for _, b := range backendList.Backends {
					//fmt.Println(b.Name)
					data := pterm.TableData{
						{b.Name, b.BackendType},
					}
					table.Data = append(table.Data, data...)

				}
				table.Render()
				return nil
			} else {
				///TODO WRITE AN INTERFACE TO GET DATA FROM THE BACKEND CONFIG
				for _, b := range backendList.Backends {
					if b.Name == args[0] {
						fmt.Println(b.Name)
					}
				}
			}
			return nil
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete a backend",
		Long:  deleteBackendDesc,
		Args:  require.ExactArgs(1),
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update a backend",
		Long:  updateBackendDesc,
		Args:  require.ExactArgs(1),
	}

	//	cmd.PersistentFlags().StringVarP(&o.name, "name", "n", "", "name of the backend")
	cmd.AddCommand(
		addCmd,
		showCmd,
		deleteCmd,
		updateCmd,
	)

	return cmd
}
