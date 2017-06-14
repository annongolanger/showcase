package dataservice_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDataservice(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dataservice Suite")
}
