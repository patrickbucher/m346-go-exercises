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

type Module struct {
	ModuleNumber int
	Classes      []Class
}

func main() {
	studentsClassA := []Student{
		{"Benjamin", "Luis"},
		{"Julian", "Michael"},
		{"Greggy", "Mike"},
	}
	studentsClassB := []Student{
		{"Ioannis", "Kai"},
		{"Ürner", "Davide"},
		{"Saru", "Max"},
	}

	classA := Class{
		Name:     "D23a",
		Students: studentsClassA,
	}
	classB := Class{
		Name:     "D23b",
		Students: studentsClassB,
	}

	modules := []Module{
		{
			ModuleNumber: 346,
			Classes:      []Class{classA, classB},
		},
		{
			ModuleNumber: 104,
			Classes:      []Class{classA},
		},
		{
			ModuleNumber: 117,
			Classes:      []Class{classB},
		},
	}

	fmt.Println("Module:")
	for _, module := range modules {
		fmt.Printf("Modul %d beinhaltet:\n", module.ModuleNumber)
		for _, class := range module.Classes {
			fmt.Printf("  %s:\n", class.Name)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.FirstName, student.LastName)
			}
		}
	}
}
