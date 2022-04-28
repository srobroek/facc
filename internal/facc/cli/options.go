package cli

import "fmt"

type BackendOptions struct {
	name string
}

type NsxtBackendOptions struct {
	name     string
	manager  string
	insecure bool
	global   bool
	Username bool
}

type Flags struct {
	flag string
}

type rootOptions struct {
	configDir string
}

func (o *BackendOptions) run() error {
	fmt.Printf("Creating %s\n", o.name)
	var err error = nil
	// chartname := filepath.Base(o.name)
	// cfile := &chart.Metadata{
	// 	Name:        chartname,
	// 	Description: "A Helm chart for Kubernetes",
	// 	Type:        "application",
	// 	Version:     "0.1.0",
	// 	AppVersion:  "0.1.0",
	// 	APIVersion:  chart.APIVersionV2,
	// }

	// if o.starter != "" {
	// 	// Create from the starter
	// 	lstarter := filepath.Join(o.starterDir, o.starter)
	// 	// If path is absolute, we don't want to prefix it with helm starters folder
	// 	if filepath.IsAbs(o.starter) {
	// 		lstarter = o.starter
	// 	}
	// 	return chartutil.CreateFrom(cfile, filepath.Dir(o.name), lstarter)
	// }

	// chartutil.Stderr = out
	// _, err := chartutil.Create(chartname, filepath.Dir(o.name))
	// return err
	return err
}

func (o *NsxtBackendOptions) run() error {
	fmt.Printf("Creating %s\n", o.name)
	var err error = nil
	// chartname := filepath.Base(o.name)
	// cfile := &chart.Metadata{
	// 	Name:        chartname,
	// 	Description: "A Helm chart for Kubernetes",
	// 	Type:        "application",
	// 	Version:     "0.1.0",
	// 	AppVersion:  "0.1.0",
	// 	APIVersion:  chart.APIVersionV2,
	// }

	// if o.starter != "" {
	// 	// Create from the starter
	// 	lstarter := filepath.Join(o.starterDir, o.starter)
	// 	// If path is absolute, we don't want to prefix it with helm starters folder
	// 	if filepath.IsAbs(o.starter) {
	// 		lstarter = o.starter
	// 	}
	// 	return chartutil.CreateFrom(cfile, filepath.Dir(o.name), lstarter)
	// }

	// chartutil.Stderr = out
	// _, err := chartutil.Create(chartname, filepath.Dir(o.name))
	// return err
	return err
}
