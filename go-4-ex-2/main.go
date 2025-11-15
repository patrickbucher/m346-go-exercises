package main

import (
	"fmt"
	"math"
)

// computeHypotenuse calculates the hypotenuse given sides a and b
func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	sideA := 3.0
	sideB := 4.0

	hypotenuse := computeHypotenuse(sideA, sideB)
	fmt.Printf("The hypotenuse of a triangle with sides %.2f and %.2f is %.2f\n", sideA, sideB, hypotenuse)
}
