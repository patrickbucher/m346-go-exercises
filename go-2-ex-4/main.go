package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		First string
		Last  string
	}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class []Student

	// TODO: declare a map of modules being attended by multiple classes
	//this was solved with the help of chatgpt
	modules := map[string][]Class{
		"Modul 104": {
			{
				{First: "Anna", Last: "Muster"},
				{First: "Ben", Last: "Beispiel"},
			},
			{
				{First: "Clara", Last: "Test"},
				{First: "David", Last: "Demo"},
			},
		},
		"Modul 346": {
			{
				{First: "Eva", Last: "Practice"},
			},
		},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
