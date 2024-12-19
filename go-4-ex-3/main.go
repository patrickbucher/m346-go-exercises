package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) (float64, float64, error) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, fmt.Errorf("no real roots")
	}
	root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return root1, root2, nil
}

func main() {
	a, b, c := 1.0, -3.0, 2.0
	root1, root2, err := computeQuadraticFormula(a, b, c)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("The roots are: %f and %f\n", root1, root2)
	}
}
