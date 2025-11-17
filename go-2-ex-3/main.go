package main

import "fmt"

func main() {
	// TODO: create a map called "modules"

	modules := make(map[int]string)

	modules[104] = "Datenmodell implementieren"
	modules[117] = "Informatik- und Netzinfrastruktur für ein kleines Unternehmen realisieren"
	modules[346] = "Cloud Lösungen konzipieren und realisieren"

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 346)
	modules[320] = "Objektorientiert Programmieren"
	modules[117] = "replaced module 117"
	fmt.Println(modules)
}
