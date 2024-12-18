package main

import (
	"fmt"
	"math"
)

type ShortSides struct {
	a float64
	b float64
}

func (s ShortSides) Hypotenuse() float64 {
	return computeHypotenuse(s.a, s.b)
}

func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func main() {
	fmt.Println(computeHypotenuse(5.1, 7.65))               // 9.19
	fmt.Println(computeHypotenuse(14.66, 11))               // 18.33
	fmt.Println(computeHypotenuse(52151.52152, 7512.16901)) // 52689.79

	fmt.Println()

	fmt.Println(ShortSides{5.1, 7.65}.Hypotenuse())               // 9.19
	fmt.Println(ShortSides{14.66, 11}.Hypotenuse())               // 18.33
	fmt.Println(ShortSides{52151.52152, 7512.16901}.Hypotenuse()) // 52689.79
}
