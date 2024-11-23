package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")

	diceLogFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating dice.log:", err)
		return
	}
	defer diceLogFile.Close()

	fmt.Fprintln(diceLogFile, "the dice was rolled at", when)
}
