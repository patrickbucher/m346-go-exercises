package main
import "fmt"
// TODO: implement the function computeGrade
func computeGrade(e float64, m float64) float64{
	n := e/m * 5 + 1
	return n
}
func main() {
	// TODO: call the function computeGrade
	grade := computeGrade(18,20)
	fmt.Print(grade)
}
