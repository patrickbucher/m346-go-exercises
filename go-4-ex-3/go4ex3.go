package main

import (
	"fmt"
	"math"
)

func computeDiscriminant(a float64, b float64, c float64) float64 {
	return math.Pow(b, 2) - (4 * a * c)
}

func computeQuadraticFormula(a float64, b float64, c float64) {
	if computeDiscriminant(a, b, c) < 0 {
		println("Keine Lösung")
	} else if computeDiscriminant(a, b, c) == 0 {
		println(-b / 2 * a)
	} else {
		println((-b+computeDiscriminant(a, b, c))/(2*a), " | ", (-b-computeDiscriminant(a, b, c))/(2*a))
	}

}

func main() {
	computeQuadraticFormula(3, 4, 1)
	computeQuadraticFormula(2, 4, 2)
	computeQuadraticFormula(3, 4, 2)

	/*	Test Quadratformel:
		Testdaten (3, 4, 1): x1 = 0.0; x2 = -1.33
		Testdaten (2, 4, 2): x = -4.0
		Testdaten (3, 4, 2): Keine Lösung
	*/

	fmt.Println(computeDiscriminant(3, 4, 1))
	fmt.Println(computeDiscriminant(2, 4, 2))
	fmt.Println(computeDiscriminant(3, 4, 2))

	/*	Test Diskriminante:
		Testdaten (3, 4, 1): D = 4
		Testdaten (2, 4, 2): D = 0
		Testdaten (3, 4, 2): D = -8
	*/
}
