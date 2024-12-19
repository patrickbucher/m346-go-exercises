package main

import "fmt"

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	celsius := 25.0
	fahrenheit := convertCelsiusToFahrenheit(celsius)
	fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", celsius, fahrenheit)

	fahrenheit = 77.0
	celsius = convertFahrenheitToCelsius(fahrenheit)
	fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", fahrenheit, celsius)
}