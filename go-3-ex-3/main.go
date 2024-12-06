package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.
	for i := Lower; i < Upper+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("\nFizzBuzz")
			fmt.Println(i, "Number is devidable by 3 AND 5.")
		} else if i%3 == 0 {
			fmt.Println("\nFizz")
			fmt.Println(i, "Number is devidable by 3.")
		} else if i%5 == 0 {
			fmt.Println("\nBuzz")
			fmt.Println(i, "Number is devidable by 5.")
		} else {
			fmt.Println()
			fmt.Println(i, "Number IS  NOT devidable by 3 OR 5 OR both.")
		}
	}
}
