package backend_test

import (
	"time"

	"code.cloudfoundry.org/garden"
	"github.com/topherbullock/garden-k8s/backend"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	restclient "k8s.io/client-go/rest"
)

var _ = Describe("Backend", func() {
	var k8s garden.Backend
	BeforeEach(func() {
		k8s = backend.New(&restclient.Config{})
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

		It("returns an error", func() {
			_, err := k8s.Containers(props)
			Expect(err).To(HaveOccurred())
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
