package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	fmt.Println(computeHypotenuse(5.1, 7.65))               // 9.19
	fmt.Println(computeHypotenuse(14.66, 11))               // 18.33
	fmt.Println(computeHypotenuse(52151.52152, 7512.16901)) // 52689.79
}
