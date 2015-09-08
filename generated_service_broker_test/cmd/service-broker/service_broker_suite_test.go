package operations_test

import (
	"testing"

	. "github.com/maximilien/swagger-cf/generated_service_broker_test/cmd/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTckSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "TCK Suite", []Reporter{NewTckReporter()})
}
