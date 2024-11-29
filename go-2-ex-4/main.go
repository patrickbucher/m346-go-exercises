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
	modules := map[string]Class{
		"Math": {
			Students: []Student{
				{FirstName: "John", LastName: "Doe"},
				{FirstName: "Jane", LastName: "Smith"},
			},
		},
		"Science": {
			Students: []Student{
				{FirstName: "Alice", LastName: "Johnson"},
				{FirstName: "Bob", LastName: "Brown"},
			},
		},
	}

	for module, class := range modules {
		fmt.Println("Module:", module)
		for _, student := range class.Students {
			fmt.Printf("Student: %s %s\n", student.FirstName, student.LastName)
		}
	}
}
