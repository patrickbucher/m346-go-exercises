package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	result := make([]float64, 2)
	d := computeDiscriminant(a, b, c)

	if d < 0 {
		result[0] = 0
		return result
	} else if d == 0 {
		result[0] = (-b) / 2 * a
		return result
	} else {
		result[0] = (-b) + math.Sqrt(d)
		result[1] = (-b) - math.Sqrt(d)
		return result
	}
}

func computeDiscriminant(a float64, b float64, c float64) float64 {
	d := (math.Pow(2, b) - 4*a*c) // a = 3, b = 4, c = 1     d = 4
	return d
}

func main() {
	// TODO: call the function computeQuadraticFormula
	result := computeQuadraticFormula(3, 4, 1)
	fmt.Println(result)
}
