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
	Name             FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Fabian",
			LastName:  "Scheuch",
		},
		Birth: BirthDate{
			Day:   1,
			Month: 1,
			Year:  1999,
		},
		NumberOfSiblings: 0,
		ZodiacSign:       '♋',
	}

	fmt.Println("Profile:", me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
