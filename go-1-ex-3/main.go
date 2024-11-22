package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Ensure randomness with a seed
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	// Use fmt.Fprintln instead of fmt.Println
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")
	fmt.Fprintln(os.Stdout, "the dice was rolled at", when)

	// Write to eyes.txt
	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating eyes.txt:", err)
		return
	}
	defer eyesFile.Close()

	fmt.Fprintln(eyesFile, eyes)

	// Write to dice.log
	logFile, err := os.Create("dice.log")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating dice.log:", err)
		return
	}
	defer logFile.Close()

	fmt.Fprintln(logFile, "the dice shows", eyes, "eyes")
	fmt.Fprintln(logFile, "the dice was rolled at", when)
}
