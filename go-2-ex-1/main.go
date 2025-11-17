package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	dayOfBirth   int
	monthOfBirth int
	yearOfBirth  int
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Lean",
			LastName:  "Arnold",
		},
		BirthDate: BirthDate{
			dayOfBirth:   13,
			monthOfBirth: 05,
			yearOfBirth:  2008,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u2649',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
