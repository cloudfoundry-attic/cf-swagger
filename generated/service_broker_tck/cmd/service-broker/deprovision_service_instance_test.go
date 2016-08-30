package operations_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#deprovisionServiceInstance", func() {
	var (
		deprovisionServiceInstanceData   string
		deprovisionServiceInstanceResult bool
		deprovisionServiceInstanceErr    error
	)

	Context("when deprovisionServiceInstance succeed", func() {
		BeforeEach(func() {
			deprovisionServiceInstanceData = "deprovisionServiceInstance data"
			deprovisionServiceInstanceErr = nil
			deprovisionServiceInstanceResult = true
		})

		It("deprovisionServiceInstance with deprovisionServiceInstanceData and gets it back", func() {
			//Do deprovisionServiceInstance
			Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
			Expect(deprovisionServiceInstanceResult).To(BeTrue())
			Expect(deprovisionServiceInstanceData).To(Equal("deprovisionServiceInstance data"))
		})
	})

	Context("when deprovisionServiceInstance fail", func() {
		BeforeEach(func() {
			deprovisionServiceInstanceErr = nil
			deprovisionServiceInstanceResult = false
		})

		It("fails to execute deprovisionServiceInstance with deprovisionServiceInstanceData", func() {
			//Do deprovisionServiceInstance
			Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
			Expect(deprovisionServiceInstanceResult).To(BeFalse())
		})
	})
})
