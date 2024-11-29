package main

import "fmt"

func main() {
	// TODO: declare a type for Student (with first and last name)
	type Student struct {
		Firstname string
		Lastname  string
	}

	studentSiri := Student{
		"Siri",
		"On",
	}

	studentSirion := Student{
		"Sirion",
		"Taiyasakda",
	}

	studentTai := Student{
		"Tai",
		"Ya",
	}

	studentTaiya := Student{
		"Taiya",
		"Sakda",
	}

	var studenten = []Student{studentSiri, studentSirion}
	var studenten2 = []Student{studentTai, studentTaiya}

	// TODO: declare a type for Class (consisting of multiple students)
	type Class struct {
		Students []Student
	}

	var class1 = Class{
		studenten,
	}

	var class2 = Class{
		studenten2,
	}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int]Class{
		117: class1,
		118: class2,
	}

	// TODO: output everything using fmt.Println()
	fmt.Println(modules)
}
