package configuration

type backendType int

const (
	nsxt backendType = iota
	aws
	azure
)

type backendFactory interface {
	setName(n string)
	setConfiguration(c backendConfiguration)
	setType(t backendType)
	//	addApplication(application application)
	getName() string
	getType() backendType
	//	getConfiguration() string
	//	getApplications() []*applic
}

type backend struct {
	name string
	backendType
}
