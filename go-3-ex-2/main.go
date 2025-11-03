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

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)
	switch zodiacSign {
	case Aries:
		fmt.Println("21. März - 20. April")
	case Taurus:
		fmt.Println("21. April - 21. Mai")
	case Gemini:
		fmt.Println("22. Mai – 21. Juni")
	case Cancer:
		fmt.Println("22. Juni – 22. Juli")
	case Leo:
		fmt.Println("23. Juli – 22. August")
	case Virgo:
		fmt.Println("23. August – 22. September")
	case Libra:
		fmt.Println("23. September – 22. Oktober")
	case Scorpius:
		fmt.Println("23. Oktober – 22. November")
	case Sagittarius:
		fmt.Println("23. November – 20. Dezember")
	case Capricornus:
		fmt.Println("21. Dezember – 19. Januar")
	case Aquarius:
		fmt.Println("20. Januar – 18. Februar")
	case Pisces:
		fmt.Println("19. Februar – 20. März")
	default:
		fmt.Println("unknown zodiac sign")
	}
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}

	outputDateRange('😛')
}
