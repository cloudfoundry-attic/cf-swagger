package operations_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/maximilien/swagger-cf/utils"
)

var _ = Describe("#createServiceInstance", func() {
	var (
		createServiceInstanceData   string
		createServiceInstanceResult bool
		createServiceInstanceErr    error
	)

	Context("when createServiceInstance succeed", func() {
		BeforeEach(func() {
			createServiceInstanceData = "createServiceInstance data"
			createServiceInstanceErr = nil
			createServiceInstanceResult = true
		})

		FIt("createServiceInstance with createServiceInstanceData and gets it back", func() {
			wd, _ := os.Getwd()
	 		file, _:= ioutil.ReadFile(filepath.Join(wd, "..", "test_fixtures","createServiceInstance.json"))
			
			jsonParams, _ := json.Marshal(file)
			parameters := bytes.NewBuffer(jsonParams)
			httpClient := utils.NewHttpClient("username", "apiKey")
			_, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "PUT", parameters)
			Expect(createServiceInstanceErr).ToNot(HaveOccurred())
			Expect(createServiceInstanceResult).To(BeTrue())
			Expect(createServiceInstanceData).To(Equal("createServiceInstance data"))
		})
	})

	Context("when createServiceInstance fail", func() {
		BeforeEach(func() {
			createServiceInstanceErr = nil
			createServiceInstanceResult = false
		})

		Context("when parameters are incorrect", func() {
			It("PUT /v2/service_instances/aws-service-guid fails", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "PUT", new(bytes.Buffer))
				Expect(createServiceInstanceErr).To(HaveOccurred())

			})
		})

		Context("when HTTP method is incorrect", func() {
			It("GET /v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "GET", new(bytes.Buffer))
				Expect(createServiceInstanceErr).To(HaveOccurred())
			})

			It("POST /v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "POST", new(bytes.Buffer))
				Expect(createServiceInstanceErr).To(HaveOccurred())
			})

			It("DELETE /v2/service_instances/aws-service-guid fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "DELETE", new(bytes.Buffer))
				Expect(createServiceInstanceErr).To(HaveOccurred())
			})

			It("PATCH /v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				_, createServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "PATCH", new(bytes.Buffer))
				Expect(createServiceInstanceErr).ToNot(HaveOccurred())
			})
		})
	})
})
