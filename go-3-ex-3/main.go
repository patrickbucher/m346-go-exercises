package main

const (
	Lower = 1
	Upper = 30
)

func main() {
	// TODO: Implement FizzBuzz using a for loop from Lower to Upper.
	// buzz ist wenn es durch 5 teilbar ist
	// fizz ist wenn es durch 3 teilbar ist
	// fizzbuzz ist wenn es durch 3 und 5 teilbar ist
	for n := Lower; n <= Upper; n++ {
		switch {
		case n%3 == 0 && n%5 == 0:
			println("FizzBuzz")
		case n%3 == 0:
			println("Fizz")
		case n%5 == 0:
			println("Buzz")
		default:
			println(n)
		}
	}
}
