package main

import "fmt"

func main() {

	type Student struct {
		FirstName string
		LastName  string
	}
	
	type Class struct {
		Students []Student
	}
	
	modules := make(map[int]Class)

	student1 := Student{
		FirstName: "Ramon",
		LastName:  "Zielke",
	}

	student2 := Student{
		FirstName: "Leonardo",
		LastName:  "Ciccone",
	}

	student3 := Student{
		FirstName: "Ylli",
		LastName:  "Abazi",
	}

	INA24bL := Class{
		Students: []Student{student1, student2, student3},
	}

	modules[346] = INA24bL

	
	fmt.Println("Modules:", modules, "\nClasses:", INA24bL, "\nStudents:", student1, student2, student3)
}
