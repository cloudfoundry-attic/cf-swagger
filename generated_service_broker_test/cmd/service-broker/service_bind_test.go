package operations_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	utils "github.com/maximilien/cf-swagger/utils"
)

var _ = Describe("#serviceBind", func() {
	var (
		serviceBindResult bool
		serviceBindErr    error
	)

	Context("when serviceBind succeed", func() {
		BeforeEach(func() {
			serviceBindErr = nil
			serviceBindResult = true
		})

		FIt("serviceBind v2/service_instances/aws-service-guid/service_bindings/aws-service-binding returns models.BindingResponse", func() {

			parameters, err := utils.ReadTestFixtures("serviceBind.json")
			Expect(err).ToNot(HaveOccurred())

			httpClient := utils.NewHttpClient("username", "apiKey")
			response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "PUT", parameters)
			if strings.Contains(string(response), "404") {
				serviceBindResult = false
			}
			Expect(serviceBindErr).ToNot(HaveOccurred())
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(serviceBindResult).To(BeTrue())
		})

		FIt("serviceBind v2/service_instances/aws-service-guid/service_bindings/aws-service-binding with app_id", func() {
			parameters, err := utils.ReadTestFixtures("serviceBindWithAppid.json")
			Expect(err).ToNot(HaveOccurred())
			httpClient := utils.NewHttpClient("username", "apiKey")
			response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "PUT", parameters)
			Expect(response).ToNot(Equal(nil))
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(serviceBindErr).ToNot(HaveOccurred())
		})

	})

	Context("when serviceBind fail", func() {
		BeforeEach(func() {
			serviceBindErr = nil
			serviceBindResult = false
		})

		Context("when parameters are empty", func() {
			It("PUT v2/service_instances/aws-service-guid/service_bindings/aws-service-binding with empty parameters", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "PUT", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceBindResult = false
				}

				Expect(serviceBindErr).ToNot(HaveOccurred())
				Expect(serviceBindResult).ToNot(BeTrue())
			})
		})

		Context("when HTTP method is incorrect", func() {

			It("Post  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Post", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceBindResult = false
				}

				Expect(serviceBindErr).ToNot(HaveOccurred())
				Expect(serviceBindResult).ToNot(BeTrue())
			})

			It("Patch  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Patch", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceBindResult = false
				}

				Expect(serviceBindErr).ToNot(HaveOccurred())
				Expect(serviceBindResult).ToNot(BeTrue())
			})

			It("Options  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Options", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceBindResult = false
				}

				Expect(serviceBindErr).ToNot(HaveOccurred())
				Expect(serviceBindResult).ToNot(BeTrue())
			})

			It("Head  v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "Head", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					serviceBindResult = false
				}

				Expect(serviceBindErr).ToNot(HaveOccurred())
				Expect(serviceBindResult).ToNot(BeTrue())
			})

		})
	})
})
