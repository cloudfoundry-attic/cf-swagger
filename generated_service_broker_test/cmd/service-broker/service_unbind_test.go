package operations_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	utils "github.com/maximilien/cf-swagger/utils"
)

var _ = Describe("#serviceUnbind", func() {
	var (
		serviceUnbindResult bool
		serviceUnbindErr    error
	)

	Context("when serviceUnbind succeed", func() {
		BeforeEach(func() {
			serviceUnbindErr = nil
			serviceUnbindResult = true
		})

		It("serviceUnbind v2/service_instances/aws-service-guid/service_bindings/aws-service-binding returns models.Empty", func() {

			parameters, err := utils.ReadTestFixtures("serviceUnbind.json")
			Expect(err).ToNot(HaveOccurred())

			httpClient := utils.NewHttpClient("username", "apiKey")
			response, serviceUnbindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "DELETE", parameters)
			if strings.Contains(string(response), "404") {
				serviceUnbindResult = false
			}
			Expect(serviceUnbindErr).ToNot(HaveOccurred())
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(serviceUnbindResult).To(BeTrue())
		})

	})

	Context("when serviceUnbind fail", func() {
		BeforeEach(func() {
			serviceUnbindErr = nil
			serviceUnbindResult = false
		})

		Context("when parameters are empty", func() {
			It("DELETE v2/service_instances/aws-service-guid/service_bindings/aws-service-binding with empty parameters", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceUnbindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "DELETE", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceUnbindResult = false
				}

				Expect(serviceUnbindErr).ToNot(HaveOccurred())
				Expect(serviceUnbindResult).ToNot(BeTrue())
			})
		})

		Context("when HTTP method is incorrect", func() {

			It("Post  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceUnbindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Post", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceUnbindResult = false
				}

				Expect(serviceUnbindErr).ToNot(HaveOccurred())
				Expect(serviceUnbindResult).ToNot(BeTrue())
			})

			It("Patch  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceUnbindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Patch", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceUnbindResult = false
				}

				Expect(serviceUnbindErr).ToNot(HaveOccurred())
				Expect(serviceUnbindResult).ToNot(BeTrue())
			})

			It("Options  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceUnbindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Options", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceUnbindResult = false
				}

				Expect(serviceUnbindErr).ToNot(HaveOccurred())
				Expect(serviceUnbindResult).ToNot(BeTrue())
			})

			It("Head  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceUnbindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Head", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceUnbindResult = false
				}

				Expect(serviceUnbindErr).ToNot(HaveOccurred())
				Expect(serviceUnbindResult).ToNot(BeTrue())
			})

		})
	})
})
