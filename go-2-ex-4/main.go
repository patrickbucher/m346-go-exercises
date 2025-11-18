package main

import "fmt"

// TODO: declare a type for Student (with first and last name)
type Student struct {
	FirstName string
	LastName  string
}

// TODO: declare a type for Class (consisting of multiple students)
type Class struct {
	Name     string
	Students []Student
}

func main() {
	s1 := Student{FirstName: "Anna", LastName: "Müller"}
	s2 := Student{FirstName: "Clara", LastName: "Weber"}

	s3 := Student{FirstName: "David", LastName: "Meier"}
	s4 := Student{FirstName: "Finn", LastName: "Schmidt"}

	classInf23a := Class{
		Name:     "Inf-23A",
		Students: []Student{s1, s2},
	}

	classInf23b := Class{
		Name:     "Inf-23B",
		Students: []Student{s3, s4},
	}

	// TODO: declare a map of modules being attended by multiple classes
	modules := map[int][]Class{
		346: {classInf23a, classInf23b},
		117: {classInf23a},
		223: {classInf23b},
	}

	// TODO: output everything using fmt.Println()
	fmt.Println("--- Klassenverwaltung ---")
	fmt.Println("Klasse A:", classInf23a)
	fmt.Println("Klasse B:", classInf23b)
	fmt.Println("--- Modulbelegung ---")
	fmt.Println(modules)
}
