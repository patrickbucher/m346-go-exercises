package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[int]string{
		104: "Datenmodell implementieren",
		117: "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// TODO: delete one
	delete(modules, 117)
	// TODO: add one
	modules[319] = "Applikationen entwerfen und implementieren"
	// TODO: replace one
	modules[104] = "Datenmodell implementieren und modellieren"

	fmt.Println(modules)
}
