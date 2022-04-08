package lands

import "errors"

type Rectangle struct {
	Length float64
	Width float64
}

func (rectangle Rectangle) CalculateArea() (float64, error) {
	hasLength := rectangle.Length > 0
	if (!hasLength) {
		return 0, errors.New("length cannot be negative or empty")
	}

	hasWidth := rectangle.Width > 0
	if (!hasWidth) {
		return 0, errors.New("width cannot be negative or empty")
	}

  return rectangle.Length * rectangle.Width, nil
}
