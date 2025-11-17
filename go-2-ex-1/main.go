package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string	
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}


type Profile struct {
	FullName        FullName
	BirthDate      BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Ramon",
			LastName:  "Zielke",
		},
		BirthDate: BirthDate{
			Day:   30, 
			Month: 1,
			Year:  2009,
		},
		NumberOfSiblings: 4,
		ZodiacSign:       '\u2652',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	fmt.Println("Siblings After:", me.NumberOfSiblings+1)
}
