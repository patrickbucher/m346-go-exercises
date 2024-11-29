package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 2, 3, 5}

	fibs[2] = fibs[0] + fibs[1]
	// TODO: correct up to index 4 using direct element access

	fibs = append(fibs, 8) // TODO: replace 0 with the next Fibonacci number
	// TODO: compute three more Fibonacci numbers and append them
	fibs = append(fibs, 13)
	fibs = append(fibs, 21)
	fibs = append(fibs, 34)
	fmt.Println(fibs) // expected output: [1 1 2 3 5 8 13 21 34]
}
