package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Introduction to Programming",
		117: "Data Structures",
		346: "Advanced Go Programming",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)

	modules[200] = "Machine Learning"

	modules[346] = "Concurrency in Go"

	fmt.Println(modules)
}
