package main

import "fmt"

type backendType int

///Define backend type
const (
	nsxt backendType = iota
	aws
)

func (t backendType) String() string {
	return [...]string{"nsxt", "aws"}[t]
}

func newBackend(n string, t backendType, c configurationFactory) (*backend, error) {
	return &backend{
		name:        n,
		backendType: t,
	}, nil

}

type backendFactory interface {
	Name() error

	Configuration() configurationFactory
	setConfiguration(c *configurationFactory) error
}

type backend struct {
	name                 string
	backendType          backendType
	backendConfiguration configurationFactory
}

func (b *backend) Name() (string, error) {
	return b.name, nil
}

func (b *backend) Configuration() (configurationFactory, error) {
	switch b.backendType {
	case nsxt:
		return b.backendConfiguration.(*nsxtBackendConfiguration), nil
	case aws:
		return b.backendConfiguration.(*awsBackendConfiguration), nil
	}
	return nil, nil
}

func (b *backend) setConfiguration(c configurationFactory) error {
	b.backendConfiguration = c
	return nil

}

type configurationFactory interface {
	Type() string
}

func newNSXTConfiguration() (*nsxtBackendConfiguration, error) {
	return &nsxtBackendConfiguration{
		backendType: "nsxt",
	}, nil
}

func newAWSConfiguration() (*awsBackendConfiguration, error) {
	return &awsBackendConfiguration{
		backendType: "aws",
	}, nil
}

// NSXT backend configuration type
type nsxtBackendConfiguration struct {
	manager     string
	username    string
	insecure    bool
	global      bool
	version     string
	backendType string
}

func (c *nsxtBackendConfiguration) newConfiguration() (configurationFactory, error) {
	return &nsxtBackendConfiguration{}, nil
}

func (c *nsxtBackendConfiguration) Manager() string {
	return c.manager
}
func (c *nsxtBackendConfiguration) Username() string {
	return c.username
}
func (c *nsxtBackendConfiguration) Insecure() bool {
	return c.insecure
}
func (c *nsxtBackendConfiguration) Global() bool {
	return c.global
}
func (c *nsxtBackendConfiguration) Version() string {
	return c.version
}
func (c *nsxtBackendConfiguration) Type() string {
	return c.backendType
}

func (c *nsxtBackendConfiguration) setManager(m string) {
	c.manager = m
}
func (c *nsxtBackendConfiguration) setUsername(u string) {
	c.username = u
}

func (c *nsxtBackendConfiguration) setInsecure(i bool) {
	c.insecure = i
}
func (c *nsxtBackendConfiguration) setGlobal(g bool) {
	c.global = g
}
func (c *nsxtBackendConfiguration) setVersion(v string) {
	c.version = v
}

// AWS backend configuration type. This doesn't do anything yet.
type awsBackendConfiguration struct {
	region      string
	key         string
	backendType string
}

func (c *awsBackendConfiguration) newConfiguration() (configurationFactory, error) {
	return &awsBackendConfiguration{}, nil
}
func (c *awsBackendConfiguration) Type() string {
	return c.backendType
}

func testBackend() {
	//cf1, _ := newConfigurationFactory(nsxt)
	//c1, _ := newConfiguration(nsxt)
	c1, _ := newNSXTConfiguration()
	b1, _ := newBackend("backend1", nsxt, c1)

	b1.setConfiguration(c1)

	//x := c1.(*nsxtBackendConfiguration)
	//x.setManager("hello")
	//x.setUsername("admin")
	//x.setVersion("3.2.0.1")
	// fmt.Println(x.getManager())
	// fmt.Println(b1.getConfiguration())
	//y, _ := b1.getConfiguration()
	//z := y.(*nsxtBackendConfiguration)
	//z.getManager()
	//z.setManager("henk")
	//	fmt.Println(x.version)
	//fmt.Println(z.getManager())
	//c1.setManager("asdf")
	//fmt.Println(z.getManager())
	fmt.Printf("%v", b1)

	//fmt.Println(y.Version)
	// fmt.Println(y.(*nsxtBackendConfiguration).username)
	// fmt.Printf("%v", b1)

}
