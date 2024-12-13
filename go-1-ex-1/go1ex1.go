package main

import "fmt"

func main() {
	var firstName string = "Hans"
	var lastName string = "Meier"
	var dayOfBirth int = 4
	var monthOfBirth int = 7
	var yearOfBirth int = 1990
	var numberOfSiblings int = 2
	var heightInMeters float64 = 1.8
	var zodiacSign rune = '\u264F'
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
