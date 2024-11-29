package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName = "Artur"
	var lastName = "Ferreira"
	var dayOfBirth = 12
	var monthOfBirth = 9
	var yearOfBirth = 2007
	var numberOfSiblings = 2
	var heightInMeters = 1.70
	var zodiacSign = "virgo"
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %s\n", zodiacSign)
}

// how to run: go run main.go

//Test commit
