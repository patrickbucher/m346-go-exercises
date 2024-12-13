package main

import (
	"fmt"
	"math"
)



// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(sidea float64, sideb float64) float64{
	return math.Sqrt(math.Pow(sidea, 2) + math.Pow(sideb, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println(computeHypotenuse(3, 4))
}
