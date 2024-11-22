package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!

	// mit allem
	var firstName string = "Sirion"

	// ohne Datentyp, weil es eindeutlich ist
	var lastName = "Taiyasakda"
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)

	// mehrere Variable in eine Zeile deklarieren und initialisieren
	var dayOfBirth, monthOfBirth, yearOfBirth = 01, 02, 1992
	// dayOfBirth, monthOfBirth, yearOfBirth := 01, 02, 1992

	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)

	// := kurzschreibweise
	numberOfSiblings, heightInMeters, zodiacSign := 1, 1.57, 9810
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
