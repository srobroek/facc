package config

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var (
	managerFlag  string
	usernameFlag string
	insecureFlag bool
	globalFlag   bool
	fileType     string = "json"
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

func (l *Backends) addBackend(b Backend) (err error) {
	jsonData, err := json.Marshal(l)
	if err != nil {
		return
	}
	jsonStr := string(jsonData)
	jsonQuery := fmt.Sprintf("Backends.#(Name==%s)#", b.Name)
	value := gjson.Get(jsonStr, jsonQuery)

	if len(value.Array()) == 1 {
		err = fmt.Errorf("object with name %s already exists", b.Name)
		return
	} else if len(value.Array()) > 1 {
		err = fmt.Errorf("multiple objects with name %s exist, check your configuration", b.Name)
		return
	}
	l.Backends = append(l.Backends, b)
	return
}

func addNsxtBackendCmd() (cmd *cobra.Command, err error) {

	cmd = &cobra.Command{
		Use: "nsxt",
		Run: func(cmd *cobra.Command, args []string) {
			//var configPath = filepath.Join(globalbackendConfig.GetString("configdir"), backendDir)
			var configPath = globalConfig.GetString("configdir")
			var err error
			fmt.Println(globalConfig.GetString("configdir"))

			viper.AddConfigPath(configPath)
			viper.SetConfigName("backends")
			viper.SetConfigType("json")
			l := Backends{}
			viper.ReadInConfig()
			viper.Unmarshal(&l)
			fmt.Println(viper.ConfigFileUsed())

			// c := NsxtBackendConfiguration{
			// 	Manager:  "managername",
			// 	Username: "admin",
			// 	Insecure: false,
			// 	Global:   true,
			// 	Version:  "3.2.0.1",
			// }
			// b := Backend{
			// 	Name:          "nsx backend",
			// 	BackendType:   "nsx",
			// 	Configuration: c,
			// }
			c := AwsBackendConfiguration{
				Region: "EU",
				Key:    "ABC",
			}
			b := Backend{
				Name:          "aws backend",
				BackendType:   "aws",
				Configuration: c,
			}
			err = l.addBackend(b)
			if err != nil {
				fmt.Println(err)
			}
			jsonData, _ := json.Marshal(l)
			viper.ReadConfig(bytes.NewBuffer(jsonData))
			err = viper.SafeWriteConfig()
			if err == nil {
				//Try to safe write first, then do unsafe write due to buggy writeconfig
				err = viper.WriteConfig()
				if err != nil {
					fmt.Println(err)
				}
			}
		},
	}

	return
}

// s1 := Student{
// 	Class: "bla",
// 	Age:   21,
// 	Name:  "Matt",
// }

// s2 := Student{
// 	Class: "blu",
// 	Age:   22,
// 	Name:  "Joe",
// }

// 			b3 := NsxtBackend{
// 				Name: "backend1",
// 				Configuration: NsxtBackendConfiguration{
// 					Manager:  "1.2.3.4",
// 					Username: "adminsjors",
// 					Insecure: false,
// 					Global:   false,
// 					Version:  "3.2.0.1",
// 				},
// 			}
// 			err := Add

// 			// students := Students{
// 			// 	Students: []Student{s1, s2},
// 			// }

// 			// jsonData, _ := json.Marshal(students)
// 			// fmt.Println(string(jsonData))
// 			// backendConfig.ReadConfig(bytes.NewBuffer(jsonData))
// 			l := NsxtBackendList{}
// 			backendConfig.ReadInConfig()
// 			backendConfig.Unmarshal(&l)
// 			AddBackend(l)
// 			//err := l.AddBackend(b3)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			jsonData, _ := json.Marshal(l)
// 			//fmt.Println(r)
// 			//fmt.Println(backendConfig.Get("students"))
// 			//fmt.Println(string(jsonData))
// 			backendConfig.ReadConfig(bytes.NewBuffer(jsonData))
// 			err = backendConfig.WriteConfig()
// 			if err != nil {
// 				fmt.Println(err)
// 			}

// 		},
// 	}

// 	cmd.Flags().StringVarP(&managerFlag, "manager", "m", "", "FQDN or IP of NSX-T manager")
// 	cmd.Flags().StringVarP(&usernameFlag, "username", "u", "admin", "NSX-T Username")
// 	cmd.Flags().BoolVarP(&insecureFlag, "insecure", "i", false, "Allow insecure SSL")
// 	cmd.Flags().BoolVarP(&globalFlag, "global", "g", false, "Use Global manager")
// 	cmd.Flags().BoolVar(&overwriteFlag, "overwrite", false, "Overwrite backend if it exists")
// 	cmd.MarkFlagRequired("manager")

// 	//BackendbackendConfig.BindPFlags(cmd.Flags())
// 	// BackendbackendConfig.Set("test.bla.yoho", "value")
// 	// s := BackendbackendConfig.Sub("test")
// 	// fmt.Println(s.GetString("bla.yoho"))

// 	return cmd
// }

// func getNsxtBackendCmd() *cobra.Command {

// 	var cmd = &cobra.Command{
// 		Use: "nsxt",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			var configPath = filepath.Join(globalbackendConfig.GetString("configdir"), backendDir)

// 			err := os.MkdirAll(configPath, 0700)
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 			fmt.Println(configPath)
// 			BackendbackendConfig.AddConfigPath(configPath)
// 			BackendbackendConfig.SetConfigName(BackendbackendConfig.GetString("name"))
// 			BackendbackendConfig.SetConfigType("json")
// 			if BackendbackendConfig.GetBool("overwrite") {
// 				err = BackendbackendConfig.WriteConfig()
// 			} else {
// 				err = BackendbackendConfig.SafeWriteConfig()
// 			}
// 			if err != nil {
// 				fmt.Println(err)
// 			}
// 		},
// 	}

// 	return cmd
// }
