package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Modul 104:",
		117: "Modul 117:",
		346: "Modul 346:",
}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 104) 

	modules[114] = "Modul 114"

	
	modules[117] = "Modul 254" 

	fmt.Println(modules)
}