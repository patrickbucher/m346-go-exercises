package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(6) + 1
	var when = time.Now()

	// Ausgabe der Augenzahl auf os.Stdout
	fmt.Fprintln(os.Stdout, "the dice shows", eyes, "eyes")

	// Ausgabe des Zeitpunkts auf os.Stderr
	fmt.Fprintln(os.Stderr, "the dice was rolled at", when)

	// TODO: how to write the output into eyes.txt and dice.log?
	// go run ex3/main.go > eyes.txt 2> dice.log
}
