package main

import "fmt"

func main() {
	modules := map[uint16]string{
		104: "Modul 104 Informationen",
		117: "Modul 117 Informationen",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 104) // delete one

	modules[117] = "New Module 117 Information" // add one

	modules[346] = "New Module 346 Information xyz" // replace one

	fmt.Println(modules)
}
