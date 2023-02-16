package configs

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// TestConfigsFunctionality will test to make sure configs can get properly set
func TestConfigsFunctionality(t *testing.T) {

	RegisterFailHandler(Fail)  // registers the fail handler from ginkgo
	RunSpecs(t, "Error setup") // hands over control to the ginkgo testing framework
}

var _ = Describe("config functions", func() {})
