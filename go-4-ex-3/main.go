package main

import (
	"fmt"
	"math"
)

func computeDiscriminant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
	var d = computeDiscriminant(a, b, c)

	if d > 0 {
		return []float64{
			(-b + math.Sqrt(d)) / (2 * a),
			(-b - math.Sqrt(d)) / (2 * a),
		}
	} else if d == 0 {
		return []float64{
			(-b + math.Sqrt(d)) / (2 * a),
		}
	} else {
		return []float64{}
	}
}

func main() {
	fmt.Println("Diskriminante")
	fmt.Println(computeDiscriminant(3, 4, 1)) // 4
	fmt.Println(computeDiscriminant(2, 4, 2)) // 0
	fmt.Println(computeDiscriminant(3, 4, 2)) // -8

	fmt.Println("Mitternachtsformel")
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // [-0.3333333333333333 -1]
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // [-1]
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // []
}
