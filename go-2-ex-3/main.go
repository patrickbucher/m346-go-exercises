package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Informatik und Gesellschaft",
		117: "Datenbanken",
		346: "Go Programmierung",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)

	// TODO: add one
	modules[320] = "Netzwerke"

	// TODO: replace one
	modules[104] = "IT und Gesellschaft"

	fmt.Println(modules)
}
