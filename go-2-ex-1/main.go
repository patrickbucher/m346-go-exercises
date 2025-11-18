package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	dayOfBirth   byte
	monthOfBirth byte
	yearOfBirth  int16
}

type Profile struct {
	// TODO: embed full name and birth date information
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		FullName: FullName{
			firstName: "Alondra",
			lastName:  "Prado",
		},
		BirthDate: BirthDate{
			dayOfBirth:   23,
			monthOfBirth: 1,
			yearOfBirth:  2002,
		},
		NumberOfSiblings: 3,        // TODO: adjust
		ZodiacSign:       '\u2652', // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
