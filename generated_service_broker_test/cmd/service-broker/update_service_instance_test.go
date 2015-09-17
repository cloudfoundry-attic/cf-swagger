package operations_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	utils "github.com/maximilien/cf-swagger/utils"
)

var _ = Describe("#updateServiceInstance", func() {
	var (
		updateServiceInstanceResult bool
		updateServiceInstanceErr    error
	)

	Context("when updateServiceInstance succeed", func() {
		BeforeEach(func() {
			updateServiceInstanceErr = nil
			updateServiceInstanceResult = true
		})

		It("updateServiceInstance v2/createServiceInstance/aws-service-guid returns models.Empty", func() {

			parameters, err := utils.ReadTestFixtures("updateServiceInstance.json")
			Expect(err).ToNot(HaveOccurred())

			httpClient := utils.NewHttpClient("username", "apiKey")
			response, updateServiceInstanceErr := httpClient.DoRawHttpRequest("v2/createServiceInstance/aws-service-guid", "PATCH", parameters)
			if strings.Contains(string(response), "404") {
				updateServiceInstanceResult = false
			}
			Expect(updateServiceInstanceErr).ToNot(HaveOccurred())
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(updateServiceInstanceResult).To(BeTrue())
		})

	})

	Context("when updateServiceInstance fail", func() {
		BeforeEach(func() {
			updateServiceInstanceErr = nil
			updateServiceInstanceResult = false
		})

		Context("when parameters are empty", func() {
			It("PATCH v2/createServiceInstance/aws-service-guid with empty parameters", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, updateServiceInstanceErr := httpClient.DoRawHttpRequest("v2/createServiceInstance/aws-service-guid", "PATCH", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					updateServiceInstanceResult = false
				}

				Expect(updateServiceInstanceErr).ToNot(HaveOccurred())
				Expect(updateServiceInstanceResult).ToNot(BeTrue())
			})
		})

		Context("when HTTP method is incorrect", func() {

			It("Post  v2/createServiceInstance/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, updateServiceInstanceErr := httpClient.DoRawHttpRequest("v2/createServiceInstance/aws-service-guid", "Post", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					updateServiceInstanceResult = false
				}

				Expect(updateServiceInstanceErr).ToNot(HaveOccurred())
				Expect(updateServiceInstanceResult).ToNot(BeTrue())
			})

			It("Options  v2/createServiceInstance/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, updateServiceInstanceErr := httpClient.DoRawHttpRequest("v2/createServiceInstance/aws-service-guid", "Options", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					updateServiceInstanceResult = false
				}

				Expect(updateServiceInstanceErr).ToNot(HaveOccurred())
				Expect(updateServiceInstanceResult).ToNot(BeTrue())
			})

			It("Head  v2/createServiceInstance/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, updateServiceInstanceErr := httpClient.DoRawHttpRequest("v2/createServiceInstance/aws-service-guid", "Head", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					updateServiceInstanceResult = false
				}

				Expect(updateServiceInstanceErr).ToNot(HaveOccurred())
				Expect(updateServiceInstanceResult).ToNot(BeTrue())
			})

		})
	})
})
