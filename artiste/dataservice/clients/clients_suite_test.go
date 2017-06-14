package clients_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"gopkg.in/jarcoal/httpmock.v1"
	"testing"
)

func TestClients(t *testing.T) {

	RegisterFailHandler(Fail)

	BeforeSuite(func() {
		httpmock.Activate()
	})

	BeforeEach(func() {
		httpmock.Reset()
	})

	AfterSuite(func() {
		httpmock.DeactivateAndReset()
	})

	RunSpecs(t, "Clients Suite")
}
