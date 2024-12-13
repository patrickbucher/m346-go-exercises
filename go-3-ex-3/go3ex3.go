package main

import "fmt"

const (
	Lower = 1
	Upper = 30
)

func main() {
	for i := Lower; i < Upper; i++ {
		fmt.Print(i, ". ")
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
