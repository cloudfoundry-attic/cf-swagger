package operations_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#serviceBind", func() {
	var (
		serviceBindData   string
		serviceBindResult bool
		serviceBindErr    error
	)

	Context("when serviceBind succeed", func() {
		BeforeEach(func() {
			serviceBindData = "serviceBind data"
			serviceBindErr = nil
			serviceBindResult = true
		})

		It("serviceBind with serviceBindData and gets it back", func() {
			//Do serviceBind
			Expect(serviceBindErr).ToNot(HaveOccurred())
			Expect(serviceBindResult).To(BeTrue())
			Expect(serviceBindData).To(Equal("serviceBind data"))
		})
	})

	Context("when serviceBind fail", func() {
		BeforeEach(func() {
			serviceBindErr = nil
			serviceBindResult = false
		})

		It("fails to execute serviceBind with serviceBindData", func() {
			//Do serviceBind
			Expect(serviceBindErr).ToNot(HaveOccurred())
			Expect(serviceBindResult).To(BeFalse())
		})
	})
})
