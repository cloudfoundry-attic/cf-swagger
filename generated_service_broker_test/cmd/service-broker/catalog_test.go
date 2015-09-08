package operations_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/maximilien/swagger-cf/utils"
)

var _ = Describe("#catalog", func() {
	var (
		catalogData   string
		catalogResult bool
		catalogErr    error
	)

	Context("when catalog succeed", func() {
		BeforeEach(func() {
			catalogData = "catalog data"
			catalogErr = nil
			catalogResult = true
		})

		It("GET /v2/catalog returns catalogData", func() {
			
			httpClient := utils.NewHttpClient("username", "apiKey")
			response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "GET", new(bytes.Buffer))
			if response == nil {
				catalogResult = false
			}

			Expect(catalogErr).ToNot(HaveOccurred())
			Expect(catalogResult).To(BeTrue())
			Expect(catalogData).To(Equal("catalog data"))
		})
	})

	Context("when catalog fail", func() {
		BeforeEach(func() {
			catalogErr = nil
			catalogResult = false
		})

		Context("when parameters are incorrect", func() {
			It("GET /v2/catalog/#someId returns catalogData", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog/id", "GET", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
				Expect(catalogData).To(Equal("catalog data"))
			})
		})

		Context("when HTTP method is incorrect", func() {
			It("PUT /v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "PUT", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("POST /v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "PUT", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("DELETE /v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "PUT", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("PATCH /v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "PUT", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})
		})
	})
})
