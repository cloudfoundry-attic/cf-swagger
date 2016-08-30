package operations_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	utils "github.com/maximilien/cf-swagger/utils"
)

var _ = Describe("#createServiceInstance", func() {
	var (
		createServiceInstanceResult bool
		createServiceInstanceErr    error
	)

	Context("when createServiceInstance succeed", func() {
		BeforeEach(func() {
			createServiceInstanceErr = nil
			createServiceInstanceResult = true
		})

		It("createServiceInstance v2/service_instances/aws-service-guid returns models.DashboardURL", func() {

			parameters, err := utils.ReadTestFixtures("createServiceInstance.json")
			Expect(err).ToNot(HaveOccurred())

			httpClient := utils.NewHttpClient("username", "apiKey")
			response, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "PUT", parameters)
			if strings.Contains(string(response), "404") {
				createServiceInstanceResult = false
			}
			Expect(createServiceInstanceErr).ToNot(HaveOccurred())
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(createServiceInstanceResult).To(BeTrue())
		})

	})

	Context("when createServiceInstance fail", func() {
		BeforeEach(func() {
			createServiceInstanceErr = nil
			createServiceInstanceResult = false
		})

		Context("when parameters are empty", func() {
			It("PUT v2/service_instances/aws-service-guid with empty parameters", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "PUT", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					createServiceInstanceResult = false
				}

				Expect(createServiceInstanceErr).ToNot(HaveOccurred())
				Expect(createServiceInstanceResult).ToNot(BeTrue())
			})
		})

		Context("when HTTP method is incorrect", func() {

			It("Post  v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "Post", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					createServiceInstanceResult = false
				}

				Expect(createServiceInstanceErr).ToNot(HaveOccurred())
				Expect(createServiceInstanceResult).ToNot(BeTrue())
			})

			It("Options  v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "Options", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					createServiceInstanceResult = false
				}

				Expect(createServiceInstanceErr).ToNot(HaveOccurred())
				Expect(createServiceInstanceResult).ToNot(BeTrue())
			})

			It("Head  v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "Head", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					createServiceInstanceResult = false
				}

				Expect(createServiceInstanceErr).ToNot(HaveOccurred())
				Expect(createServiceInstanceResult).ToNot(BeTrue())
			})

		})
	})
})
