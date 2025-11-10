package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		panic(err)
	}
	defer eyesFile.Close()

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")


	
	diceFile, err := os.Create("dice.log")
	if err != nil {
		panic(err)
	}
	defer diceFile.Close()

	fmt.Fprintln(diceFile, "the dice was rolled at", when)
}