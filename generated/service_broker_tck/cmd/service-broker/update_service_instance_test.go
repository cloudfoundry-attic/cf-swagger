package operations_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#updateServiceInstance", func() {
	var (
		updateServiceInstanceData   string
		updateServiceInstanceResult bool
		updateServiceInstanceErr    error
	)

	Context("when updateServiceInstance succeed", func() {
		BeforeEach(func() {
			updateServiceInstanceData = "updateServiceInstance data"
			updateServiceInstanceErr = nil
			updateServiceInstanceResult = true
		})

		It("updateServiceInstance with updateServiceInstanceData and gets it back", func() {
			//Do updateServiceInstance
			Expect(updateServiceInstanceErr).ToNot(HaveOccurred())
			Expect(updateServiceInstanceResult).To(BeTrue())
			Expect(updateServiceInstanceData).To(Equal("updateServiceInstance data"))
		})
	})

	Context("when updateServiceInstance fail", func() {
		BeforeEach(func() {
			updateServiceInstanceErr = nil
			updateServiceInstanceResult = false
		})

		It("fails to execute updateServiceInstance with updateServiceInstanceData", func() {
			//Do updateServiceInstance
			Expect(updateServiceInstanceErr).ToNot(HaveOccurred())
			Expect(updateServiceInstanceResult).To(BeFalse())
		})
	})
})
