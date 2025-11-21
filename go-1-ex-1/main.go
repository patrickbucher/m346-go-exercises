package main

import "fmt"

func main() {
	firstName := "Nino"
	lastName := "Meier"
	dayOfBirth := 12
	monthOfBirth := 4
	yearOfBirth := 2008
	numberOfSiblings := 1
	heightInMeters := 1.75
	zodiacSign := '\u2648'
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
