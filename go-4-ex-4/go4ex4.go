package main

import "fmt"

func convertCelsiusToFahrenheit(celsius float32) float32 {
	return ((celsius * 9 / 5) + 32)
}

func convertFahrenheitToCelsius(fahrenheit float32) float32 {
	return ((fahrenheit - 32) * 5 / 9)
}

func main() {
	fmt.Println(convertCelsiusToFahrenheit(0))
	fmt.Println(convertCelsiusToFahrenheit(100))
	fmt.Println(convertCelsiusToFahrenheit(-50))

	/*	Test Celsius zu Fahrenheit:
		0° C = 32° F
		100° C = 212° F
		-50° C = -58° F
	*/

	fmt.Println()

	fmt.Println(convertFahrenheitToCelsius(32))
	fmt.Println(convertFahrenheitToCelsius(212))
	fmt.Println(convertFahrenheitToCelsius(-58))

	/*	Test Fahrenheit zu Celsius:
		32° F = 0° C
		212° F = 100° C
		-58° F = -50° C
	*/
}
