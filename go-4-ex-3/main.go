package main

import (
	"fmt"
	"math"
)

// computeQuadraticFormula calculates the roots of a quadratic equation ax^2 + bx + c = 0
func computeQuadraticFormula(a, b, c float64) (float64, float64, error) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, fmt.Errorf("the equation has no real roots")
	}
	sqrtDisc := math.Sqrt(discriminant)
	x1 := (-b + sqrtDisc) / (2 * a)
	x2 := (-b - sqrtDisc) / (2 * a)
	return x1, x2, nil
}

func main() {
	a, b, c := 1.0, -3.0, 2.0 // example: x^2 - 3x + 2 = 0
	x1, x2, err := computeQuadraticFormula(a, b, c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Roots: x1 = %.2f, x2 = %.2f\n", x1, x2)
}
