package lands_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLands(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lands Suite")
}