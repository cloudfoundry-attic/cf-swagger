package operations_test

import (
	"bytes"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/maximilien/cf-swagger/utils"
)

var _ = Describe("#catalog", func() {
	var (
		catalogResult bool
		catalogErr    error
	)

	Context("when catalog succeed", func() {
		BeforeEach(func() {
			catalogErr = nil
			catalogResult = true
		})

		It("catalog v2/catalog returns models.CatalogServices", func() {

			httpClient := utils.NewHttpClient("username", "apiKey")
			response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "GET", new(bytes.Buffer))
			if strings.Contains(string(response), "404") {
				catalogResult = false
			}
			Expect(catalogErr).ToNot(HaveOccurred())
			Expect(string(response)).ToNot(ContainSubstring("Unprocessable Entity"))
			Expect(catalogResult).To(BeTrue())
		})

	})

	Context("when catalog fail", func() {
		BeforeEach(func() {
			catalogErr = nil
			catalogResult = false
		})

		Context("when HTTP method is incorrect", func() {

			It("Post  v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "Post", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("Put  v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "Put", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("Patch  v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "Patch", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("Options  v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "Options", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("Head  v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "Head", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

			It("Delete  v2/catalog fails with 404", func() {
				httpClient := utils.NewHttpClient("username", "apiKey")
				response, catalogErr := httpClient.DoRawHttpRequest("v2/catalog", "Delete", new(bytes.Buffer))
				if strings.Contains(string(response), "404") {
					catalogResult = false
				}

				Expect(catalogErr).ToNot(HaveOccurred())
				Expect(catalogResult).ToNot(BeTrue())
			})

		})
	})
})
