package lands_test

import (
	"calculate-area/lands"

	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Land", func() {
	var mockCtrl *gomock.Controller
	var mockLand1 *lands.MockLand
	var mockLand2 *lands.MockLand
	var groupsOfLands []lands.Land

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockLand1 = lands.NewMockLand(mockCtrl)
		mockLand2 = lands.NewMockLand(mockCtrl)
		groupsOfLands = []lands.Land{mockLand1, mockLand2}
	})

	It("should calculate the total area of lands with 2 different lands", func() {
		mockLand1.EXPECT().CalculateArea().Return(25.00, nil)
		mockLand2.EXPECT().CalculateArea().Return(18.35, nil)
		expected := 43.35

		result := lands.CalculateTotalArea(groupsOfLands)

		Expect(result).To(Equal(expected))
	})

	It("should calculate the total area of lands with 2 of the same lands", func() {
		groupsOfLands = []lands.Land{mockLand1, mockLand1}
		mockLand1.EXPECT().CalculateArea().Return(20.00, nil).Times(2)
		expected := 40.00

		result := lands.CalculateTotalArea(groupsOfLands)

		Expect(result).To(Equal(expected))
	})
})
