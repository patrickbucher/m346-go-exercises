package main

import "fmt"

func convertCelsiusToFahrenheit(c float64) float64 {
	return c*(9.0/5.0) + 32
}

func convertFahrenheitToCelsius(f float64) float64 {
	return (f - 32) * (5.0 / 9.0)
}

func main() {
	fmt.Println("Celsius to fahrenheit")
	fmt.Println(convertCelsiusToFahrenheit(37.5)) // 99.5
	fmt.Println(convertCelsiusToFahrenheit(-1.5)) // 29.3
	fmt.Println(convertCelsiusToFahrenheit(51))   // 123.8

	fmt.Println("fahrenheit to celsius")
	fmt.Println(convertFahrenheitToCelsius(99.5))  // 37.5
	fmt.Println(convertFahrenheitToCelsius(29.3))  // -1.5
	fmt.Println(convertFahrenheitToCelsius(123.8)) // 51
}
