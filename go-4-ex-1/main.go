package main

import "fmt"

func computeGrade(score int, maxscore int) string {

	decimalGrade := float64(score) / float64(maxscore)
	swissGrade := 5*decimalGrade + 1

	return fmt.Sprintf("%.1f", swissGrade)
}

func main() {
	fmt.Println("Grade: ", computeGrade(20, 30))
	fmt.Println("Grade: ", computeGrade(18, 20))
}
