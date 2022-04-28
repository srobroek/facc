package cmdold

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var BackendConfig = viper.New()

var (
	nameFlag      string
	overwriteFlag bool
	backendDir    string = "backends"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a backend",
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get backend properties",
}

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set backend properties",
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a backend",
}

func addBackendSubCmd() (cmd *cobra.Command, err error) {
	cmd = &cobra.Command{
		Use:   "backend",
		Short: "Manage backends",
	}
	cmd.PersistentFlags().String("name", "", "name of the backend")
	cmd.MarkPersistentFlagRequired("name")
	//subCmd, err := addNsxtBackendCmd()
	//cmd.AddCommand(subCmd)

	BackendConfig.BindPFlags(cmd.PersistentFlags())

	return
}

func getBackendSubCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "backend",
		Short: "Manage backends",
	}
	cmd.PersistentFlags().String("name", "", "name of the backend")
	cmd.MarkPersistentFlagRequired("name")
	//cmd.AddCommand(getNsxtBackendCmd())
	return cmd
}

func setBackendSubCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "backend",
		Short: "Manage backends",
	}
	return cmd
}

func deletebackendSubCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "backend",
		Short: "Manage backends",
	}
	return cmd
}

func initialiseBackendConfig() {

	fmt.Println(filepath.Join(BackendConfig.GetString("configdir"), backendDir))
	//fmt.Println(filepath.Join(rootCmd.PersistentFlags().Lookup("configdir").Value.String(), backendDir))
}

// func initialiseSubCmd() {
// 	subCmd.PersistentFlags().StringVarP(&nameFlag, "name", "n", "", "Backend name")
// 	subCmd.PersistentFlags().BoolVar(&overwriteFlag, "overwrite", false, "overwrite existing configuration")
// 	//BackendConfig.BindPFlag("name", cmd.Parent().PersistentFlags().Lookup("name"))
// 	subCmd.PersistentFlags().StringVarP(&configDirFlag, "configdir", "c", "./configs", "configuration directory")
// 	// cmd.MarkPersistentFlagDirname("configdir")
// 	// cmd.MarkPersistentFlagRequired("name")

// 	BackendConfig.BindPFlags(cmd.Flags())
// 	BackendConfig.BindPFlags(cmd.PersistentFlags())
// 	subCmd.AddCommand(addCmd)
// 	addCmd.AddCommand(nsxtAddCmd())

// 	//fmt.Println(cmd.PersistentFlags().Lookup("configdir"))
// 	BackendConfig.SetConfigName(BackendConfig.GetString("name"))
// 	BackendConfig.AddConfigPath(BackendConfig.GetString("configpath") + "backends")
// 	BackendConfig.SetConfigType("json")

// 	// err := BackendConfig.SafeWriteConfig()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// }

// func initialiseConfig() {

// }
