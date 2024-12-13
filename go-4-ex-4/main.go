package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
// TODO: implement the function convertFahrenheitToCelsius

func convertCelsiusToFahrenheit(celsius float64) float64{
	return celsius * 9 / 5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64{
	return (fahrenheit -32) * 5 / 9
}


func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	// TODO: call the function convertFahrenheitToCelsius

	fmt.Println(convertCelsiusToFahrenheit(22))
	fmt.Println(convertFahrenheitToCelsius(100))
}
