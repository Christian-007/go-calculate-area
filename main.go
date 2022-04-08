package main

import (
	"calculate-area/lands"
	"fmt"
)

func main() {
	rectangle := lands.Rectangle{
		Width: 2,
		Length: 3,
	}
	circle:= lands.Circle{
		Radius: 2,
	}

	landResult := lands.CalculateTotalArea([]lands.Land{rectangle, circle})
	fmt.Printf("Land Area: %f\n", landResult)
}