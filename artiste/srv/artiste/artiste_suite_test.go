package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/gexec"
	"os/exec"
	"testing"
)

func TestArtiste(t *testing.T) {
	RegisterFailHandler(Fail)

	var artisteProc *gexec.Session
	var pathToSelf = "github.com/benwaine/artistprof/artiste/srv/artiste"

	BeforeSuite(func() {
		pathToServer, err := gexec.Build(pathToSelf)
		Expect(err).NotTo(HaveOccurred())
		cmd := exec.Command(pathToServer, "-config=test/test_conf.json")

		artisteProc, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)

		Expect(err).NotTo(HaveOccurred())

		Consistently(artisteProc).ShouldNot(gexec.Exit(), "Artiste exited")

	})

	AfterSuite(func() {
		artisteProc.Kill()
		gexec.CleanupBuildArtifacts()
	})

	RunSpecs(t, "Artiste Suite")
}
