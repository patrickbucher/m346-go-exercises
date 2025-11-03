package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	firstName := "Lukas"
	lastName := "Traut"
	dayOfBirth := 14
	monthOfBirth := 11
	yearOfBirth := 2008
	numberOfSiblings := 1
	heightInMeters := 1.87
	zodiacSign := '\u264F' // Unicode für Skorpion 

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
