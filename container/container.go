package container

import (
	"errors"
	"io"
	"time"

	k8s_meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/topherbullock/garden-k8s/process"

	"code.cloudfoundry.org/garden"
)

// NOTE: k8s' smallest deployable unit is a pod, which can consist of multiple
// containers; usure what to do about that ... so for now let's just ignore it.
// let's pretend container == pod.
type container struct {
	handle    string
	namespace string
	client    *kubernetes.Clientset
}

//go:generate counterfeiter . Client

type Client interface {
	Pods(string) Pods
}

//go:generate counterfeiter . Pods

type Pods interface {
	Get(string)
	Delete(string, *k8s_meta.DeleteOptions) error
}

func New(handle, namespace string, client *kubernetes.Clientset) garden.Container {
	return &container{
		handle:    handle,
		namespace: namespace,
		client:    client,
	}
}

func (c *container) Handle() string {
	return c.handle
}

func (c *container) Stop(kill bool) error {
	podsClient := c.client.Pods(c.namespace)
	return podsClient.Delete(c.handle, &k8s_meta.DeleteOptions{})
}

func (c *container) Info() (garden.ContainerInfo, error) {
	podsClient := c.client.Pods(c.namespace)
	// IDEA:

	// pod, err := podsClient.Get(c.handle, &k8s_meta.GetOptions{})

	return garden.ContainerInfo{
		State:         "",
		Events:        []string{},
		HostIP:        "",
		ContainerIP:   "",
		ExternalIP:    "",
		ContainerPath: "",
		ProcessIDs:    []string{},
		Properties:    garden.Properties{},
		MappedPorts:   []garden.PortMapping{},
	}, errors.New("Not Implemented")
}

func (c *container) StreamIn(spec garden.StreamInSpec) error {
	return errors.New("Not Implemented")
}

func (c *container) StreamOut(spec garden.StreamOutSpec) (io.ReadCloser, error) {
	return nil, errors.New("Not Implemented")
}

func (c *container) CurrentBandwidthLimits() (garden.BandwidthLimits, error) {
	// IDEA: Add this eventually https://github.com/kubernetes/kubernetes/issues/2856
	return garden.BandwidthLimits{}, nil
}

func (c *container) CurrentCPULimits() (garden.CPULimits, error) {
	return garden.CPULimits{}, errors.New("Not Implemented")
}

func (c *container) CurrentDiskLimits() (garden.DiskLimits, error) {
	return garden.DiskLimits{}, errors.New("Not Implemented")
}

func (c *container) CurrentMemoryLimits() (garden.MemoryLimits, error) {
	return garden.MemoryLimits{}, errors.New("Not Implemented")
}

func (c *container) NetIn(hostPort, containerPort uint32) (uint32, uint32, error) {
	return 0, 0, errors.New("Not Implemented")
}

func (c *container) NetOut(netOutRule garden.NetOutRule) error {
	return errors.New("Not Implemented")
}

func (c *container) BulkNetOut(netOutRules []garden.NetOutRule) error {
	return errors.New("Not Implemented")
}

func (c *container) Run(garden.ProcessSpec, garden.ProcessIO) (garden.Process, error) {
	return &process.Process{}, errors.New("Not Implemented")
}

func (c *container) Attach(processID string, io garden.ProcessIO) (garden.Process, error) {
	return &process.Process{}, errors.New("Not Implemented")
}

func (c *container) Metrics() (garden.Metrics, error) {
	return garden.Metrics{}, errors.New("Not Implemented")
}

func (c *container) SetGraceTime(graceTime time.Duration) error {
	return errors.New("Not Implemented")
}

func (c *container) Properties() (garden.Properties, error) {
	return garden.Properties{}, errors.New("Not Implemented")
}

func (c *container) Property(name string) (string, error) {
	return "", errors.New("Not Implemented")
}

func (c *container) SetProperty(name string, value string) error {
	return errors.New("Not Implemented")
}

func (c *container) RemoveProperty(name string) error {
	return errors.New("Not Implemented")
}
