package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	a := 3.0
	b := 4.0
	hypotenuse := computeHypotenuse(a, b)
	fmt.Printf("The hypotenuse of a triangle with sides %.2f and %.2f is %.2f\n", a, b, hypotenuse)
}
