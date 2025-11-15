package main

import "fmt"

// convertCelsiusToFahrenheit converts Celsius temperature to Fahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// convertFahrenheitToCelsius converts Fahrenheit temperature to Celsius
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	celsius := 25.0
	fahrenheit := 77.0

	// convert Celsius to Fahrenheit
	fmt.Printf("%.2f°C = %.2f°F\n", celsius, convertCelsiusToFahrenheit(celsius))

	// convert Fahrenheit to Celsius
	fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, convertFahrenheitToCelsius(fahrenheit))
}
