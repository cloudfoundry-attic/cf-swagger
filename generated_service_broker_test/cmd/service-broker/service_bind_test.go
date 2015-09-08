package operations_test

import (
	"bytes"
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/maximilien/swagger-cf/generated_service_broker_test/cmd/utils"

	utils "github.com/maximilien/swagger-cf/utils"
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

		FIt("serviceBind with serviceBindData without app_id", func() {
			parameters, err := ReadTestFixtures("serviceBind.json")
			Expect(err).ToNot(HaveOccurred())
			fmt.Printf("parameters%#v",parameters.String())
			httpClient := utils.NewHttpClient("username", "apiKey")
			response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "PUT", parameters)
			Expect(response).ToNot(Equal(nil))
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(serviceBindErr).ToNot(HaveOccurred())
			Expect(serviceBindResult).To(BeTrue())

		})

		FIt("serviceBind with serviceBindData with app_id", func() {
			parameters, err := ReadTestFixtures("serviceBindWithAppid.json")
			Expect(err).ToNot(HaveOccurred())
			fmt.Printf("parameters%#v",parameters.String())
			httpClient := utils.NewHttpClient("username", "apiKey")
			response, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "PUT", parameters)
			Expect(response).ToNot(Equal(nil))
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(serviceBindErr).ToNot(HaveOccurred())
			Expect(serviceBindResult).To(BeTrue())

		})
	})

	Context("when serviceBind fail", func() {
		BeforeEach(func() {
			serviceBindErr = nil
			serviceBindResult = false

		})

		Context("when parameters are incorrect", func() {
			It("PUT v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails", func() {

				httpClient := utils.NewHttpClient("username", "apiKey")
				_, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "GET", new(bytes.Buffer))
				Expect(serviceBindErr).ToNot(HaveOccurred())

			})
		})

		Context("when HTTP method is incorrect", func() {
			It("GET v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "GET", new(bytes.Buffer))
				Expect(serviceBindErr).ToNot(HaveOccurred())
			})

			It("POST v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "POST", new(bytes.Buffer))
				Expect(serviceBindErr).ToNot(HaveOccurred())
			})

			It("PATCH v2/service_instances/aws-service-guid/service_bindings/aws-service-binding fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, serviceBindErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid/service_bindings/aws-service-binding", "PATCH", new(bytes.Buffer))
				Expect(serviceBindErr).ToNot(HaveOccurred())
			})
		})
	})
})
