package main

import "fmt"

type FullName struct {
	// TODO: add fields$
	FirstName string
	LastName string
}

// TODO: declare a structure for birth date
type BirthDate struct {
    DayOfBirth   byte
    MonthOfBirth byte
    YearOfBirth  int16
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
		FirstName:"Selina",
		LastName:"Frey",
		},
		BirthDate: BirthDate{
		DayOfBirth: 17,
		MonthOfBirth:7,
		YearOfBirth: 2007,
		},
		NumberOfSiblings: 3,   // TODO: adjust
		ZodiacSign:       '\u264B',

	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings = me.NumberOfSiblings + 1

	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
