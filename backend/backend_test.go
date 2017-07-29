package backend_test

import (
	"time"

	"code.cloudfoundry.org/garden"
	"github.com/topherbullock/garden-k8s/backend"
	"github.com/topherbullock/garden-k8s/backend/backendfakes"
	"github.com/topherbullock/garden-k8s/v1fakes"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/pkg/api/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Backend", func() {
	var (
		k8s        garden.Backend
		fakeClient *backendfakes.FakeClient
		fakePods   *v1fakes.FakePodInterface
	)

	BeforeEach(func() {
		fakeClient = new(backendfakes.FakeClient)
		fakePods = new(v1fakes.FakePodInterface)
		fakeClient.PodsReturns(fakePods)
		k8s = backend.New("test-namespace", fakeClient)
	})

	Describe("Ping", func() {
		It("returns an error", func() {
			err := k8s.Ping()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Capacity", func() {
		It("returns an error", func() {
			_, err := k8s.Capacity()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Create", func() {
		var spec garden.ContainerSpec

		It("returns an error", func() {
			_, err := k8s.Create(spec)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Destroy", func() {
		var handle string

		It("returns an error", func() {
			err := k8s.Destroy(handle)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Containers", func() {
		var props garden.Properties
		BeforeEach(func() {
			fakePods.ListReturns(&v1.PodList{
				Items: []v1.Pod{
					v1.Pod{
						meta_v1.TypeMeta{},
						meta_v1.ObjectMeta{
							Name: "free-willy",
							Annotations: map[string]string{
								"type": "killer",
							},
						},
						v1.PodSpec{},
						v1.PodStatus{},
					},
					v1.Pod{
						meta_v1.TypeMeta{},
						meta_v1.ObjectMeta{
							Name: "shamu",
							Annotations: map[string]string{
								"type": "killer",
							},
						},
						v1.PodSpec{},
						v1.PodStatus{},
					},
					v1.Pod{
						meta_v1.TypeMeta{},
						meta_v1.ObjectMeta{
							Name: "some-other-whale",
							Annotations: map[string]string{
								"type": "beluga",
							},
						},
						v1.PodSpec{},
						v1.PodStatus{},
					},
				},
			}, nil)
		})

		Context("without any filter properties", func() {
			It("returns a list of containers", func() {
				containers, err := k8s.Containers(props)
				Expect(err).ToNot(HaveOccurred())

				Expect(len(containers)).To(Equal(3))

				By("keeping the ordering from the k8s client")
				Expect(containers[0].Handle()).To(Equal("free-willy"))
				Expect(containers[1].Handle()).To(Equal("shamu"))
				Expect(containers[2].Handle()).To(Equal("some-other-whale"))
			})
		})

	})

	Describe("BulkInfo", func() {
		var handles []string

		It("returns an error", func() {
			_, err := k8s.BulkInfo(handles)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("BulkMetrics", func() {
		var handles []string

		It("returns an error", func() {
			_, err := k8s.BulkMetrics(handles)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Lookup", func() {
		var handle string
		It("returns an error", func() {
			_, err := k8s.Lookup(handle)
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GraceTime", func() {
		var container garden.Container

		It("returns a bogus grace time", func() {
			graceTime := k8s.GraceTime(container)
			Expect(graceTime).To(Equal(time.Duration(0)))
		})
	})

	Describe("Start", func() {
		It("returns an error", func() {
			err := k8s.Start()
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("Stop", func() {
		It("does nothing", func() {
			k8s.Stop()
		})
	})

})
