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
	Name            FullName
	BirthDate       BirthDate
	NumberOfSiblings byte
	ZodiacSign      string
}

func main() {
	var me = Profile{
		Name: FullName{
			FirstName: "Benjamin",
			LastName:  "Maier",
		},
		BirthDate: BirthDate{
			Day:   8,
			Month: 5,
			Year:  2008,
		},
		NumberOfSiblings: 1,
		ZodiacSign:	"Stier",
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	me.NumberOfSiblings++
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
