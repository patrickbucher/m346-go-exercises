package main

import (
	"fmt"
)

var firstName string = "lean"
var lastName string = "arnold"
var dayOfBirth int = 13
var monthOfBirth int = 05
var yearOfBirth int = 2008
var numberOfSiblings = 2
var heightInMeters = 1.87
var zodiacSign string = "taurus"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
