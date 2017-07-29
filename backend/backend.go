package backend

import (
	"errors"
	"time"

	"github.com/topherbullock/garden-k8s/container"

	"code.cloudfoundry.org/garden"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/typed/core/v1"
)

var errNotImpl = errors.New("Not implemented")

type backend struct {
	client    Client
	namespace string
}

//go:generate counterfeiter . Client

type Client interface {
	Pods(string) v1.PodInterface
}

func New(namespace string, client Client) garden.Backend {

	return &backend{
		client:    client,
		namespace: namespace,
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

func (b *backend) Containers(props garden.Properties) ([]garden.Container, error) {
	var containers []garden.Container
	podsClient := b.client.Pods(b.namespace)

	pods, err := podsClient.List(meta_v1.ListOptions{})
	if err != nil {
		return containers, err
	}

	for _, pod := range pods.Items {
		containers = append(containers, container.New(pod.Name, pod.Namespace, podsClient))
	}

	return containers, nil
}

func listOptsFromProperties(props garden.Properties) {

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
