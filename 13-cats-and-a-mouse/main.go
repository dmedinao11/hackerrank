package main

import "math"

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	dx := math.Abs(float64(z - x))
	dy := math.Abs(float64(z - y))

	if dx == dy {
		return "Mouse C"
	} else if dx < dy {
		return "Cat A"
	} else {
		return "Cat B"
	}
}
