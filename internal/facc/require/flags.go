package require

import "github.com/spf13/viper"

func initSettings() *viper.Viper {
	settings := viper.New()
	return settings
}

// func getSettings() *viper.Viper {
// 	return settings
// }

// type EnvSettings struct {
// 	//configdir string
// 	// namespace string
// 	// // KubeConfig is the path to the kubeconfig file
// 	// KubeConfig string
// 	// // KubeContext is the name of the kubeconfig context.
// 	// KubeContext string
// 	// // Bearer KubeToken used for authentication
// 	// KubeToken string
// 	// // Username to impersonate for the operation
// 	// KubeAsUser string
// 	// // Groups to impersonate for the operation, multiple groups parsed from a comma delimited list
// 	// KubeAsGroups []string
// 	// // Kubernetes API Server Endpoint for authentication
// 	// KubeAPIServer string
// 	// // Custom certificate authority file.
// 	// KubeCaFile string
// 	// // Debug indicates whether or not Helm is running in Debug mode.
// 	// Debug bool
// 	// // RegistryConfig is the path to the registry config file.
// 	// RegistryConfig string
// 	// // RepositoryConfig is the path to the repositories file.
// 	// RepositoryConfig string
// 	// // RepositoryCache is the path to the repository cache directory.
// 	// RepositoryCache string
// 	// // PluginsDirectory is the path to the plugins directory.
// 	// PluginsDirectory string
// 	// // MaxHistory is the max release history maintained.
// 	// MaxHistory int
// }

// func New() *EnvSettings {
// 	env := &EnvSettings{
// 		configdir: "",
// 	}
// 	viper.Env
// }
// func (s *EnvSettings) AddFlags(fs *pflag.FlagSet) {
// 	fs.StringVarP(&s.namespace, "namespace", "n", s.namespace, "namespace scope for this request")
// 	fs.StringVar(&s.KubeConfig, "kubeconfig", "", "path to the kubeconfig file")
// 	fs.StringVar(&s.KubeContext, "kube-context", s.KubeContext, "name of the kubeconfig context to use")
// 	fs.StringVar(&s.KubeToken, "kube-token", s.KubeToken, "bearer token used for authentication")
// 	fs.StringVar(&s.KubeAsUser, "kube-as-user", s.KubeAsUser, "username to impersonate for the operation")
// 	fs.StringArrayVar(&s.KubeAsGroups, "kube-as-group", s.KubeAsGroups, "group to impersonate for the operation, this flag can be repeated to specify multiple groups.")
// 	fs.StringVar(&s.KubeAPIServer, "kube-apiserver", s.KubeAPIServer, "the address and the port for the Kubernetes API server")
// 	fs.StringVar(&s.KubeCaFile, "kube-ca-file", s.KubeCaFile, "the certificate authority file for the Kubernetes API server connection")
// 	fs.BoolVar(&s.Debug, "debug", s.Debug, "enable verbose output")
// 	fs.StringVar(&s.RegistryConfig, "registry-config", s.RegistryConfig, "path to the registry config file")
// 	fs.StringVar(&s.RepositoryConfig, "repository-config", s.RepositoryConfig, "path to the file containing repository names and URLs")
// 	fs.StringVar(&s.RepositoryCache, "repository-cache", s.RepositoryCache, "path to the file containing cached repository indexes")
// }
