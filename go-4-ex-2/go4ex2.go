package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt((math.Pow(a, 2)) + (math.Pow(b, 2)))
}

func main() {
	fmt.Println(computeHypotenuse(3, 4))
}
