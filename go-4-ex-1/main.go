package main
import (
	"fmt"
)

// TODO: implement the function computeGrade
func computeGrade(yourPoints float64, maxPoints float64) float64 {
	return (yourPoints / maxPoints) * 5 + 1
}

func main() {
	fmt.Println(computeGrade(10, 12))
	// TODO: call the function computeGrade
}
