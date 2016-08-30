package operations_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/maximilien/cf-swagger/generated_service_broker_test/cmd/utils"
)

func TestTckSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "TCK Suite", []Reporter{NewTckReporter()})
}
