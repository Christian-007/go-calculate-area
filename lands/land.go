package lands

// genereate mock: mockgen -source=land.go -destination=mock_land.go
type Land interface {
	CalculateArea() (float64, error)
}

func CalculateTotalArea(lands []Land) (float64) {
	var result float64

	for _, land :=range lands {
		area, _ := land.CalculateArea()
		result += area
	}

	return result
}
