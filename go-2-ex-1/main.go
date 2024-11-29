package main

import "fmt"

type FullName struct {
	// TODO: add fields
	FirstName string
	Lastname  string
}

// TODO: declare a structure for birth date
type Birthdate struct {
	DayOfBirth   byte
	MonthOfBirth byte
	YearOfBirth  int16
}

type Profile struct {
	// TODO: embed full name and birth date information
	Name             FullName
	Birthday         Birthdate
	NumberOfSiblings byte
	ZodiacSign       rune
}

var myName = FullName{
	"Sirion",
	"Taiyasakda",
}

var myBirthday = Birthdate{
	01,
	02,
	1992,
}

func main() {
	var me = Profile{
		// TODO: set name and birth date information
		Name:             myName,
		Birthday:         myBirthday,
		NumberOfSiblings: 1,    // TODO: adjust
		ZodiacSign:       9810, // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings = 2
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
