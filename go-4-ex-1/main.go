package main

import "fmt"

func computeGrade(getPoints float32, maxPoints float32) float32 {
	return getPoints/maxPoints*5 + 1
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125
	fmt.Println(computeGrade(38, 42))     // 5,523...
	fmt.Println(computeGrade(9, 17))      // 3,647...
}
