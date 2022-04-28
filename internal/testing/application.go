package main

/////////////////////////

type applicationBuilder interface {
	//setName(n string) error
	Name() (string, error)
	addRole(r applicationRole) error
	Roles() ([]applicationRole, error)
	Role(n string) (applicationRole, error)
	removeRole(r applicationRole) error
	addProvider(p applicationProvider) error
	Providers() ([]applicationProvider, error)
	Provider(n string) (applicationProvider, error)
	removeProvider(p applicationProvider) error
	newApplication() (*application, error)
}

type roleBuilder interface {
	Name() (string, error)
}

type application struct {
	name      string
	roles     []applicationRole
	providers []applicationProvider
}

type applicationRole struct {
	name string
}

type applicationProvider struct {
}

type applicationConsumer struct {
}

func (a *application) newApplication() (*application, error) {
	return &application{}, nil
}

func (a *application) setName(n string) error {
	a.name = n
	return nil
}

func (a *application) Name() (string, error) {
	return a.name, nil

}
func (a *application) addRole(r applicationRole) error {
	a.roles = append(a.roles, r)
	return nil
}

func (a *application) Roles() ([]applicationRole, error) {
	return a.roles, nil
}

func (a *application) Role(n string) (applicationRole, error) {
	return a.roles[0], nil
}

func (a *application) removeRole(r applicationRole) error {
	i := 0
	a.roles = append(a.roles[:i], a.roles[i+1])

	return nil
}

func (a *application) addProvider(p applicationProvider) error {
	a.providers = append(a.providers, p)
	return nil
}

func (a *application) removeProvider(p applicationProvider) error {
	i := 0
	a.providers = append(a.providers[:i], a.providers[i+1])
	return nil
}

func (a *application) Providers() ([]applicationProvider, error) {
	return a.providers, nil
}

func (a *application) Provider(n string) (applicationProvider, error) {
	return a.providers[0], nil
}

func newApplication() applicationBuilder {
	return &application{}
}

/////////////////////

type serviceFactory interface {
}

func testApplication() {
	a := newApplication()
	a.setName("henk")

}
