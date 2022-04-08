package lands

import (
	"errors"
	"math"
)

type Circle struct {
	Radius float64
}

func (circle Circle) CalculateArea() (float64, error) {
	hasRadius := circle.Radius > 0
	if (!hasRadius) {
		return 0, errors.New("radius cannot be negative or empty")
	}

	diameter := circle.Radius * 2

  return math.Pi * diameter, nil
}