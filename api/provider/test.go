package provider

import (
	"io"

	"github.com/convox/rack/Godeps/_workspace/src/github.com/stretchr/testify/mock"
	"github.com/convox/rack/api/structs"
)

var TestProvider = &TestProviderRunner{}

type TestProviderRunner struct {
	mock.Mock
	App       structs.App
	Build     structs.Build
	Capacity  structs.Capacity
	Instances structs.Instances
	Release   structs.Release
	Releases  structs.Releases
}

func (p *TestProviderRunner) AppGet(name string) (*structs.App, error) {
	p.Called(name)
	return &p.App, nil
}

func (p *TestProviderRunner) BuildCreateTar(app string, src io.Reader, manifest, description string, cache bool) (*structs.Build, error) {
	p.Called(app, src, manifest, description, cache)
	return &p.Build, nil
}

func (p *TestProviderRunner) BuildDelete(app, id string) (*structs.Build, error) {
	p.Called(app, id)
	return &p.Build, nil
}

func (p *TestProviderRunner) BuildGet(app, id string) (*structs.Build, error) {
	p.Called(app, id)
	return &p.Build, nil
}

func (p *TestProviderRunner) BuildRelease(b *structs.Build) (*structs.Release, error) {
	p.Called(b)
	return &p.Release, nil
}

func (p *TestProviderRunner) BuildSave(b *structs.Build, logdir string) error {
	p.Called(b, logdir)
	return nil
}

func (p *TestProviderRunner) CapacityGet() (*structs.Capacity, error) {
	p.Called()
	return &p.Capacity, nil
}

func (p *TestProviderRunner) InstanceList() (structs.Instances, error) {
	p.Called()
	return p.Instances, nil
}

func (p *TestProviderRunner) ReleaseGet(app, id string) (*structs.Release, error) {
	p.Called(app, id)
	return &p.Release, nil
}

func (p *TestProviderRunner) ReleaseList(app string) (structs.Releases, error) {
	p.Called(app)
	return p.Releases, nil
}

func (p *TestProviderRunner) ReleaseSave(r *structs.Release, logdir, key string) error {
	p.Called(r, logdir, key)
	return nil
}

func (p *TestProviderRunner) SystemGet() (*structs.System, error) {
	p.Called()
	return nil, nil
}

func (p *TestProviderRunner) SystemSave(system structs.System) error {
	p.Called(system)
	return nil
}
