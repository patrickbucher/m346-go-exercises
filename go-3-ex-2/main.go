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
	// TODO: Replace if, else if branching with switch/case.
	// TODO: Define all 12 cases...
	switch {
	case (zodiacSign == Aries):
		fmt.Println("21.03. - 20.04")
	case (zodiacSign == Taurus):
		fmt.Println("21.04. - 21.05")
	case (zodiacSign == Gemini):
		fmt.Println("22.05. - 21.06")
	case (zodiacSign == Cancer):
		fmt.Println("22.06. - 22.07")
	case (zodiacSign == Leo):
		fmt.Println("23.07. - 23.08")
	case (zodiacSign == Virgo):
		fmt.Println("24.08. - 23.09")
	case (zodiacSign == Libra):
		fmt.Println("24.09. - 23.10")
	case (zodiacSign == Scorpius):
		fmt.Println("24.10. - 22.11")
	case (zodiacSign == Sagittarius):
		fmt.Println("23.11. - 21.12")
	case (zodiacSign == Capricornus):
		fmt.Println("22.12. - 20.01")
	case (zodiacSign == Aquarius):
		fmt.Println("21.01. - 19.02")
	case (zodiacSign == Pisces):
		fmt.Println("20.02. - 20.03")
	default:
		fmt.Println("No Zodiac Sign")
	}
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
}
