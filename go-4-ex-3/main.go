package main

import (
	"errors"
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula
func computeQuadraticFormula(a float64, b float64, c float64) (float64, float64, error) {

	d := math.Pow(b, 2) - 4*a*c

	switch {
	case d > 0:
		positiveValue := (-b + math.Sqrt(d)) / (2 * a)
		negativeValue := (-b - math.Sqrt(d)) / (2 * a)
		return positiveValue, negativeValue, nil
	case d == 0:
		x := -b / (2 * a)
		return x, x, nil
	case d < 0:
		return 0, 0, errors.New("Discriminante darf nicht negativ sein")
	}

	return 0, 0, errors.New("unerwarterer fall")
}

func main() {
	// Call the function with different cases
	// Case 1: Two distinct roots (d > 0)
	fmt.Println("Case 1:")
	root1, root2, err := computeQuadraticFormula(1, -3, 2) // Roots: 2, 1
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Roots: %.2f, %.2f\n", root1, root2)
	}

	// Case 2: One root (d == 0)
	fmt.Println("\nCase 2:")
	root1, root2, err = computeQuadraticFormula(1, 2, 1) // Root: -1
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Roots: %.2f, %.2f\n", root1, root2)
	}

	// Case 3: No real roots (d < 0)
	fmt.Println("\nCase 3:")
	root1, root2, err = computeQuadraticFormula(1, 2, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Roots: %.2f, %.2f\n", root1, root2)
	}

	// Case 4: Invalid input (a == 0)
	fmt.Println("\nCase 4:")
	root1, root2, err = computeQuadraticFormula(0, 2, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Roots: %.2f, %.2f\n", root1, root2)
	}
}
