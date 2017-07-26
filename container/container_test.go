package container_test

import (
	"code.cloudfoundry.org/garden"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topherbullock/garden-k8s/container"
	"github.com/topherbullock/garden-k8s/v1fakes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/pkg/api/v1"
)

var _ = Describe("Container", func() {

	var c garden.Container
	var handle = "messiah"
	var namespace string
	var fakeClient *v1fakes.FakePodInterface

	BeforeEach(func() {
		fakeClient = new(v1fakes.FakePodInterface)
	})

	JustBeforeEach(func() {
		c = container.New(handle, namespace, fakeClient)
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
				fakeClient.DeleteReturns(nil)
			})

			It("deletes the pod by handle", func() {
				err := c.Stop(true)
				Expect(err).ToNot(HaveOccurred())
				Expect(fakeClient.DeleteCallCount()).To(Equal(1))
				handel, _ := fakeClient.DeleteArgsForCall(0)
				Expect(handel).To(Equal("messiah"))
			})
		})
	})

	Describe("Info", func() {
		Context("When the pod exists", func() {
			BeforeEach(func() {
				fakeClient.GetReturns(&v1.Pod{
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
				fakeClient.GetReturns(&v1.Pod{
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
				fakeClient.GetReturns(&v1.Pod{
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

	Describe("SetProperty", func() {

		var existingPod *v1.Pod

		Context("When the pod exists and has annotations", func() {
			existingPod = &v1.Pod{
				metav1.TypeMeta{},
				metav1.ObjectMeta{
					Annotations: map[string]string{
						"hello": "world",
						"foo":   "bar",
					},
				},
				v1.PodSpec{},
				v1.PodStatus{},
			}

			BeforeEach(func() {
				fakeClient.GetReturns(existingPod, nil)
			})

			It("sets an existing property to a new ", func() {
				err := c.SetProperty("foo", "fufu")
				Expect(err).ToNot(HaveOccurred())

				Expect(fakeClient.UpdateCallCount()).To(Equal(1))
				annotations := existingPod.GetAnnotations()
				Expect(annotations["foo"]).To(Equal("fufu"))
			})
		})
	})

	Describe("RemoveProperty", func() {
		var existingPod *v1.Pod

		Context("When the pod exists and has annotations", func() {
			existingPod = &v1.Pod{
				metav1.TypeMeta{},
				metav1.ObjectMeta{
					Annotations: map[string]string{
						"hello": "world",
						"foo":   "bar",
					},
				},
				v1.PodSpec{},
				v1.PodStatus{},
			}

			BeforeEach(func() {
				fakeClient.GetReturns(existingPod, nil)
			})

			It("deletes an existing property", func() {
				err := c.RemoveProperty("foo")
				Expect(err).ToNot(HaveOccurred())

				Expect(fakeClient.UpdateCallCount()).To(Equal(1))
				annotations := existingPod.GetAnnotations()
				val, found := annotations["foo"]
				Expect(val).To(BeEmpty())
				Expect(found).To(BeFalse())
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

})
