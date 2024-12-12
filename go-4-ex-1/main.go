package main

import "fmt"

func computeGrade(getPoints float32, maxPoints float32) (float32, error) {
	if getPoints < 0 || maxPoints < 0 || getPoints > maxPoints {
		return -1, fmt.Errorf("erreichte punkte und maximale punkte müssen mehr als null sein und erreichte punkte nicht mehr als maximale punkte; erreicht: %f, max: %f", getPoints, maxPoints)
	}
	return getPoints/maxPoints*5 + 1, nil
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125
	fmt.Println(computeGrade(38, 42))     // 5,523...
	fmt.Println(computeGrade(9, 17))      // 3,647...
	fmt.Println(computeGrade(-9, 17))     // error
	fmt.Println(computeGrade(9, -17))     // error
	fmt.Println(computeGrade(9, 7))       // error
}
