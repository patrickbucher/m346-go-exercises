package main

import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints float32, maxPoints float32) float32 {
	if gotPoints > maxPoints || gotPoints < 0 {
		fmt.Printf("Deine Punktzahl %v ist ungültig. \n", gotPoints)
	}

	note := (gotPoints/maxPoints)*5 + 1

	if note < 1 || note > 6 {
		fmt.Print("Die Note ist ungültig: ")
	}
	return note
}

func main() {
	// TODO: call the function computeGrade
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125
	fmt.Println(computeGrade(30, 28.0))   // 4.125

}
