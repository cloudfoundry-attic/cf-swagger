package operations_test

import (
	"fmt"


	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/maximilien/swagger-cf/generated_service_broker_test/cmd/utils"

	utils "github.com/maximilien/swagger-cf/utils"
)

var _ = Describe("#deprovisionServiceInstance", func() {
	var (
		deprovisionServiceInstanceData   string
		deprovisionServiceInstanceResult bool
		deprovisionServiceInstanceErr    error
	)

	Context("when deprovisionServiceInstance succeed", func() {
		BeforeEach(func() {
			err := createServiceInstance()
			Expect(err).ToNot(HaveOccurred())

			deprovisionServiceInstanceData = "deprovisionServiceInstance data"
			deprovisionServiceInstanceErr = nil
			deprovisionServiceInstanceResult = true
		})

		It("deprovisionServiceInstance with deprovisionServiceInstanceData and gets it back", func() {
			parameters, err := ReadTestFixtures("deprovisionServiceInstance.json")
			Expect(err).ToNot(HaveOccurred())
			fmt.Printf("======> parameters %#v\n", parameters.String())
			httpClient := utils.NewHttpClient("username", "apiKey")
			response, deprovisionServiceInstanceErr := httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "DELETE", parameters)
			fmt.Printf("======> response %#v\n", response)
			Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
			Expect(deprovisionServiceInstanceResult).To(BeTrue())
			Expect(deprovisionServiceInstanceData).To(Equal("deprovisionServiceInstance data"))
		})
	})

	Context("when deprovisionServiceInstance fail", func() {
		BeforeEach(func() {
			deprovisionServiceInstanceErr = nil
			deprovisionServiceInstanceResult = false
		})

		It("fails to execute deprovisionServiceInstance with deprovisionServiceInstanceData", func() {
			//Do deprovisionServiceInstance
			Expect(deprovisionServiceInstanceErr).ToNot(HaveOccurred())
			Expect(deprovisionServiceInstanceResult).To(BeFalse())
		})
	})
})

func createServiceInstance() error {
	parameters, err := ReadTestFixtures("createServiceInstance.json")
	Expect(err).ToNot(HaveOccurred())
	httpClient := utils.NewHttpClient("username", "apiKey")
	_, err = httpClient.DoRawHttpRequest("v2/service_instances/aws-service-guid", "PUT", parameters)
	return err
}
