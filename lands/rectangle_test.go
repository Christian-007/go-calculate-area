package lands_test

import (
	"calculate-area/lands"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rectangle", func() {
	var rectangle lands.Rectangle

	BeforeEach(func() {
		rectangle = lands.Rectangle{
			Width: 2,
			Length: 3,
		}
	})


	It("should calculate the area of Rectangle", func() {
		result, _ := rectangle.CalculateArea()
		expected := float64(6)
		Expect(result).To(Equal(expected))
	})

	It("should return error if length is empty", func() {
		rectangle = lands.Rectangle{
			Length: 0,
			Width: 2,
		}
		expected := MatchError("length cannot be negative or empty")

		_, err := rectangle.CalculateArea()

		Expect(err).To(expected)
	})

	It("should return error if width is empty", func() {
		rectangle = lands.Rectangle{
			Length: 2,
			Width: 0,
		}
		expected := MatchError("width cannot be negative or empty")

		_, err := rectangle.CalculateArea()

		Expect(err).To(expected)
	})
})
