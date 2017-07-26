package backend

import (
	"errors"
	"time"

	"code.cloudfoundry.org/garden"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

var errNotImpl = errors.New("Not implemented")

type backend struct {
	client Client
}

//go:generate counterfeiter . Client

type Client interface {
	Pods(string) v1.PodInterface
}

func New(client Client) garden.Backend {

	return &backend{
		client: client,
	}
}

func (b *backend) Ping() error {
	return errNotImpl
}

func (b *backend) Capacity() (garden.Capacity, error) {
	return garden.Capacity{}, errNotImpl
}

func (b *backend) Create(garden.ContainerSpec) (garden.Container, error) {
	return nil, errNotImpl
}

func (b *backend) Destroy(handle string) error {
	return errNotImpl
}

func (b *backend) Containers(garden.Properties) ([]garden.Container, error) {
	return []garden.Container{}, errNotImpl
}

func (b *backend) BulkInfo(handles []string) (map[string]garden.ContainerInfoEntry, error) {
	return map[string]garden.ContainerInfoEntry{}, errNotImpl
}

func (b *backend) BulkMetrics(handles []string) (map[string]garden.ContainerMetricsEntry, error) {
	return map[string]garden.ContainerMetricsEntry{}, errNotImpl
}

func (b *backend) Lookup(handle string) (garden.Container, error) {
	return nil, errNotImpl
}

func (b *backend) GraceTime(container garden.Container) time.Duration {
	return 0
}

func (b *backend) Start() error {
	return errNotImpl
}

func (b *backend) Stop() {
}
