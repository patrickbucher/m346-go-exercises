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
	classA := Class{
		Students: []Student{
			{"Max", "Muster"},
			{"Tom", "Fischer"},
			{"Anna", "Meier"},
		},
	}
	classB := Class{
		Students: []Student{
			{"Sara", "Keller"},
			{"Tim", "Huber"},
			{"Nina", "Schmid"},
		},
	}
	modules := map[int][]Class{
		346: {classA, classB},
		117: {classA},
		210: {classB},
	}

	fmt.Println(classA)
	fmt.Println(classB)
	fmt.Println(modules)
}
