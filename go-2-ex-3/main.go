package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Mathematics",
		117: "Programming Basics",
		346: "Go Programming",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)

	modules[205] = "Physics"

	modules[104] = "Advanced Mathematics"

	fmt.Println(modules)
}
