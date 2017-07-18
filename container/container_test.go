package container_test

import (
	"code.cloudfoundry.org/garden"
	. "github.com/topherbullock/garden-k8s/container"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container", func() {

	var container garden.Container

	BeforeEach(func() {
		container := NewContainer(handle, namespace, client)
	})

})
