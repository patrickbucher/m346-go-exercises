package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	c := math.Sqrt(math.Pow(2, a) + math.Pow(2, b))
	return c
}

type ShortSides struct {
	a float64
	b float64
}

func main() {
	// TODO: call the function computeHypotenuse
	kathete := ShortSides{20, 41}
	result := computeHypotenuse(kathete.a, kathete.b)
	fmt.Println(result)
}
