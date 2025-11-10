package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	// Aufgabe
	var firstName string = "Selina"
	var lastName string = "Frey"
	var dayOfBirth int = 17
	var monthOfBirth int = 07
	var yearOfBirth int = 2007
	var numberOfSiblings int = 3
	heightInMeters := 1.55
	var zodiacSign = 'Cancer'
	// Aufgabe

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
