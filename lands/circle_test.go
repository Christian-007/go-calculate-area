package lands_test

import (
	"calculate-area/lands"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Circle", func() {
	var circle lands.Circle

	BeforeEach(func() {
		circle = lands.Circle{
			Radius: 2,
		}
	})

	It("should calculate the area of Circle", func() {
		result, _ := circle.CalculateArea()
		expected := 12.566370614359172
		Expect(result).To(Equal(expected))
	})

	It("should return error if radius is negative", func() {
		circle = lands.Circle{
			Radius: -1,
		}
		_, err := circle.CalculateArea()
		Expect(err).To(MatchError("radius cannot be negative or empty"))
	})
})
