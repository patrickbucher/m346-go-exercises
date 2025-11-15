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
	Name     FullName
	Birthday BirthDate
	NumberOfSiblings byte
	ZodiacSign       string
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Atieh",
			LastName:  "Rostamian",
		},
		Birthday: BirthDate{
			Day:   15,
			Month: 11,
			Year:  1999,
		},
		NumberOfSiblings: 0,
		ZodiacSign:       "Scorpio",
	}
	fmt.Println(me)
	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
