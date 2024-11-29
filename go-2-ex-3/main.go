package main

import "fmt"

func main() {
	// TODO: create a map called "modules"
	modules := map[uint16]string{
		104: "Ka",
		117: "VonAllmen",
		346: "Cloud",
	}
	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)
	modules[114] = "Codierung"
	modules[346] = "Cloud computing"
	// TODO: delete one
	// TODO: add one
	// TODO: replace one
	fmt.Println(modules)
}
