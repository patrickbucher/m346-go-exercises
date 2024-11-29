package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	fullName         FullName
	birthDate        BirthDate
	NumberOfSiblings byte
	ZodiacSign       string
}

func main() {
	var me = Profile{
		fullName: FullName{
			FirstName: "Yannick",
			LastName:  "Blatty",
		},
		birthDate: BirthDate{
			DayOfBirth:   28,
			MonthOfBirth: 03,
			YearOfBirth:  2007,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       "Aries",
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
