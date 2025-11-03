package main

import "fmt"

func main() {
	var firstName, lastName = "Eneas", "Infanger"
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	var dayOfBirth, monthOfBirth, yearOfBirth = 1, 3, 2008
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	var numberOfSiblings = 2
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	var heightInMeters = 1.9
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	var zodiacSign = '\u2653'
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
