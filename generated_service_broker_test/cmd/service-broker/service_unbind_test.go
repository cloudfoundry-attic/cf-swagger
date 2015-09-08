package operations_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("#serviceUnbind", func() {
	var (
		serviceUnbindData   string
		serviceUnbindResult bool
		serviceUnbindErr    error
	)

	Context("when serviceUnbind succeed", func() {
		BeforeEach(func() {
			serviceUnbindData = "serviceUnbind data"
			serviceUnbindErr = nil
			serviceUnbindResult = true
		})

		It("serviceUnbind with serviceUnbindData and gets it back", func() {
			//Do serviceUnbind
			Expect(serviceUnbindErr).ToNot(HaveOccurred())
			Expect(serviceUnbindResult).To(BeTrue())
			Expect(serviceUnbindData).To(Equal("serviceUnbind data"))
		})
	})

	Context("when serviceUnbind fail", func() {
		BeforeEach(func() {
			serviceUnbindErr = nil
			serviceUnbindResult = false
		})

		It("fails to execute serviceUnbind with serviceUnbindData", func() {
			//Do serviceUnbind
			Expect(serviceUnbindErr).ToNot(HaveOccurred())
			Expect(serviceUnbindResult).To(BeFalse())
		})
	})
})
