package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type Date struct {
	Day   byte
	Month byte
	Year  uint16
}

type Profile struct {
	FullName         FullName
	BirthDate        Date
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		FullName: FullName{
			FirstName: "Eneas",
			LastName:  "Infanger",
		},
		BirthDate: Date{
			Day:   1,
			Month: 3,
			Year:  2008,
		},
		NumberOfSiblings: 2,
		ZodiacSign:       '\u2653',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
