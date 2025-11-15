package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
)

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   byte
	Month byte
	Year  uint16
}

type Person struct {
	FullName
	BirthDate
}

func outputWithZodiacSign(p Person) {
	var zodiacSign rune = '?'

	switch p.Month {
	case 1:
		if p.Day <= 20 {
			zodiacSign = Capricornus
		} else {
			zodiacSign = Aquarius
		}
	case 2:
		if p.Day <= 19 {
			zodiacSign = Aquarius
		} else {
			zodiacSign = Pisces
		}
	case 3:
		if p.Day <= 20 {
			zodiacSign = Pisces
		} else {
			zodiacSign = Aries
		}
	case 4:
		if p.Day <= 20 {
			zodiacSign = Aries
		} else {
			zodiacSign = Taurus
		}
	case 5:
		if p.Day <= 21 {
			zodiacSign = Taurus
		} else {
			zodiacSign = Gemini
		}
	case 6:
		if p.Day <= 21 {
			zodiacSign = Gemini
		} else {
			zodiacSign = Cancer
		}
	case 7:
		if p.Day <= 22 {
			zodiacSign = Cancer
		} else {
			zodiacSign = Leo
		}
	case 8:
		if p.Day <= 23 {
			zodiacSign = Leo
		} else {
			zodiacSign = Virgo
		}
	case 9:
		if p.Day <= 23 {
			zodiacSign = Virgo
		} else {
			zodiacSign = Libra
		}
	case 10:
		if p.Day <= 23 {
			zodiacSign = Libra
		} else {
			zodiacSign = Scorpius
		}
	case 11:
		if p.Day <= 22 {
			zodiacSign = Scorpius
		} else {
			zodiacSign = Sagittarius
		}
	case 12:
		if p.Day <= 21 {
			zodiacSign = Sagittarius
		} else {
			zodiacSign = Capricornus
		}
	}

	fmt.Printf("%s %s, born on %02d.%02d.%04d, has the zodiac sign %c.\n",
		p.FirstName, p.LastName, p.Day, p.Month, p.Year, zodiacSign)
}

func main() {
	grace := Person{FullName{"Grace", "Hopper"}, BirthDate{9, 12, 1906}}
	outputWithZodiacSign(grace)
}
