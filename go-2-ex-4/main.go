package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

func main() {
	student1 := Student{"Atieh", "Rostamian"}
	student2 := Student{"Ali", "Ahmadi"}
	student3 := Student{"Sara", "Moradi"}
	student4 := Student{"Mina", "Shah"}

	classA := Class{
		Students: []Student{student1, student2},
	}
	classB := Class{
		Students: []Student{student3, student4},
	}

	modules := map[string][]Class{
		"Mathematics":      {classA},
		"Go Programming":   {classA, classB},
		"Physics":          {classB},
	}

	fmt.Println("Modules and Classes:")
	for module, classes := range modules {
		fmt.Println("Module:", module)
		for i, class := range classes {
			fmt.Printf("  Class %d:\n", i+1)
			for _, student := range class.Students {
				fmt.Printf("    %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
