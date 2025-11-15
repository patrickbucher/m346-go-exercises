package main

import "fmt"

// computeGrade receives a score and returns the corresponding grade
func computeGrade(score float64) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	scores := []float64{95, 82, 76, 61, 50}

	for _, score := range scores {
		grade := computeGrade(score)
		fmt.Printf("Score %.1f -> Grade %s\n", score, grade)
	}
}
