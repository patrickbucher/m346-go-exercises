package main

import "fmt"

func computeGrade(gotPoints float32, maxpoints float32) {
	var note float32 = ((gotPoints / maxpoints) * 5) + 1
	fmt.Println(note)
}

func main() {
	computeGrade(17.5, 28.0)
}
