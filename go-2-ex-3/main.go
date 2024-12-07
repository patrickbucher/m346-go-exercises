package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Einführung in die Programmierung",
		117: "Datenbanken und SQL",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104]) 
	fmt.Println("Modul 117:", modules[117]) 
	fmt.Println("Modul 346:", modules[346]) 

	delete(modules, 117)

	modules[202] = "Algorithmen und Datenstrukturen"

	modules[346] = "Moderne Cloud-Architekturen"

	fmt.Println(modules)
}
