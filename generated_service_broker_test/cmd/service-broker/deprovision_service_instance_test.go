package operations_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	utils "github.com/maximilien/cf-swagger/utils"
)

var _ = Describe("#deprovisionServiceInstance", func() {
	var (
		deprovisionServiceInstanceResult bool
		deprovisionServiceInstanceErr    error
	)

	Context("when deprovisionServiceInstance succeed", func() {
		BeforeEach(func() {
			deprovisionServiceInstanceErr = nil
			deprovisionServiceInstanceResult = true
		})

		It("deprovisionServiceInstance v2/service_instances/aws-service-guid returns models.Empty", func() {

			parameters, err := utils.ReadTestFixtures("deprovisionServiceInstance.json")
			Expect(err).ToNot(HaveOccurred())

			httpClient := utils.NewHttpClient("username", "apiKey")
			response, deprovisionServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "DELETE", parameters)
			if strings.Contains(string(response), "404") {
				deprovisionServiceInstanceResult = false
			}
			Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(deprovisionServiceInstanceResult).To(BeTrue())
		})

	})

	Context("when deprovisionServiceInstance fail", func() {
		BeforeEach(func() {
			deprovisionServiceInstanceErr = nil
			deprovisionServiceInstanceResult = false
		})

		Context("when parameters are empty", func() {
			It("DELETE v2/service_instances/aws-service-guid with empty parameters", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, deprovisionServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "DELETE", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					deprovisionServiceInstanceResult = false
				}

				Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
				Expect(deprovisionServiceInstanceResult).ToNot(BeTrue())
			})
		})

		Context("when HTTP method is incorrect", func() {

			It("Post  v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, deprovisionServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "Post", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					deprovisionServiceInstanceResult = false
				}

				Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
				Expect(deprovisionServiceInstanceResult).ToNot(BeTrue())
			})

			It("Options  v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, deprovisionServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "Options", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					deprovisionServiceInstanceResult = false
				}

				Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
				Expect(deprovisionServiceInstanceResult).ToNot(BeTrue())
			})

			It("Head  v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, deprovisionServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "Head", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					deprovisionServiceInstanceResult = false
				}

				Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
				Expect(deprovisionServiceInstanceResult).ToNot(BeTrue())
			})

		})
	})
})
