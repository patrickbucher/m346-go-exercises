package main

import "fmt"

func main() {
	// Declare and initialize the variables
	firstName := "Nils"
	lastName := "Utiger"
	dayOfBirth := 15
	monthOfBirth := 6
	yearOfBirth := 1990
	numberOfSiblings := 2
	heightInMeters := 1.75
	zodiacSign := 'G'

	// Output the variables
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
