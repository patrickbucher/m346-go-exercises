package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "This goes to STDOUT")
	fmt.Fprintln(os.Stderr, "This goes to STDERR")
}