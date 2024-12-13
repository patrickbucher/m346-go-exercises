package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) []float64{
	return []float64{
		((-b + math.Pow(b, 2) - 4 * a * c) / (2 * a)),
		((-b - math.Pow(b, 2) - 4 * a * c) / (2 * a)),

	}
}

func main() {
	fmt.Println(computeQuadraticFormula(2, 3, 4))
	// TODO: call the function computeQuadraticFormula
}
