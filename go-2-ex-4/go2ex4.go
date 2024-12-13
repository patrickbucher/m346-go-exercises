package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Name     string
	Students []Student
}

func main() {
	class1 := Class{
		Name: "INA23a",
		Students: []Student{
			{FirstName: "Hans", LastName: "Muster"},
			{FirstName: "Jan", LastName: "Gross"},
		},
	}
	class2 := Class{
		Name: "INA23c",
		Students: []Student{
			{FirstName: "Eva", LastName: "Zimmermann"},
			{FirstName: "Frank", LastName: "Müller"},
		},
	}
	class3 := Class{
		Name: "INA24a",
		Students: []Student{
			{FirstName: "Anna", LastName: "Klein"},
			{FirstName: "Sofua", LastName: "Gerber"},
		},
	}

	modules := map[string][]Class{
		"M117": {class1},
		"M346": {class1, class2},
		"M254": {class2, class1, class3},
	}

	for moduleName, classes := range modules {
		fmt.Println("Module:", moduleName)
		for _, class := range classes {
			fmt.Println("Class Name:", class.Name)
			fmt.Println("Students:")
			for _, student := range class.Students {
				fmt.Printf("- %s %s\n", student.FirstName, student.LastName)
			}
		}
		fmt.Println()
	}
}
