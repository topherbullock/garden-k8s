package container_test

import (
	"code.cloudfoundry.org/garden"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topherbullock/garden-k8s/container"
	"github.com/topherbullock/garden-k8s/container/containerfakes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

var _ = Describe("Container", func() {

	var c garden.Container
	var handle = "messiah"
	var namespace string
	var client *containerfakes.FakeClient
	var fakePods *containerfakes.FakePods

	BeforeEach(func() {
		client = new(containerfakes.FakeClient)
		fakePods = new(containerfakes.FakePods)
		client.PodsReturns(fakePods)
	})

	JustBeforeEach(func() {
		c = container.New(handle, namespace, client)
	})

	Describe("Handle", func() {
		It("returns the container's handle", func() {
			handel := c.Handle()
			Expect(handel).To(Equal("messiah"))
		})
	})

	Describe("Stop", func() {
		Context("When deletion succeeds", func() {
			BeforeEach(func() {
				fakePods.DeleteReturns(nil)
			})

			It("deletes the pod by handle", func() {
				err := c.Stop(true)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakePods.DeleteCallCount()).To(Equal(1))
				handel, _ := fakePods.DeleteArgsForCall(0)
				Expect(handel).To(Equal("messiah"))
			})
		})
	})

	Describe("Info", func() {
		Context("When the pod exists", func() {
			BeforeEach(func() {
				fakePods.GetReturns(&v1.Pod{
					metav1.TypeMeta{},
					metav1.ObjectMeta{},
					v1.PodSpec{},
					v1.PodStatus{
						Phase:  v1.PodRunning,
						HostIP: "127.0.0.1",
					},
				}, nil)
			})

			It("returns info about the pod", func() {
				info, err := c.Info()
				Expect(err).ToNot(HaveOccurred())
				Expect(info.State).To(Equal("Running"))
				Expect(info.HostIP).To(Equal("127.0.0.1"))
			})
		})
	})

	Describe("Properties", func() {
		Context("When the pod exists and has annotations", func() {
			BeforeEach(func() {
				fakePods.GetReturns(&v1.Pod{
					metav1.TypeMeta{},
					metav1.ObjectMeta{
						Annotations: map[string]string{
							"hello": "world",
							"foo":   "bar",
						},
					},
					v1.PodSpec{},
					v1.PodStatus{},
				}, nil)
			})

			It("returns the annotations as garden properties", func() {
				props, err := c.Properties()
				Expect(err).ToNot(HaveOccurred())
				Expect(props["hello"]).To(Equal("world"))
			})
		})
	})

	Describe("Property", func() {
		Context("When the pod exists and has annotations", func() {
			BeforeEach(func() {
				fakePods.GetReturns(&v1.Pod{
					metav1.TypeMeta{},
					metav1.ObjectMeta{
						Annotations: map[string]string{
							"hello": "world",
							"foo":   "bar",
						},
					},
					v1.PodSpec{},
					v1.PodStatus{},
				}, nil)
			})

			It("returns the annotation value", func() {
				val, err := c.Property("foo")
				Expect(err).ToNot(HaveOccurred())
				Expect(val).To(Equal("bar"))
			})
		})
	})

	Describe("StreamIn", func() {

	})

	Describe("StreamOut", func() {

	})

	Describe("CurrentBandwidthLimits", func() {

	})

	Describe("CurrentCPULimits", func() {

	})

	Describe("CurrentDiskLimits", func() {

	})

	Describe("CurrentMemoryLimits", func() {

	})

	Describe("NetIn", func() {

	})

	Describe("NetOut", func() {

	})

	Describe("BulkNetOut", func() {

	})

	Describe("Run", func() {

	})

	Describe("Attach", func() {

	})

	Describe("Metrics", func() {

	})

	Describe("SetGraceTime", func() {

	})

	Describe("SetProperty", func() {

	})

	Describe("RemoveProperty", func() {

	})

})
