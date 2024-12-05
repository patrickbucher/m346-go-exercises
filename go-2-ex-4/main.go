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
	modules := make(map[uint16][]Class, 0)

	eneasStudent := Student{
		FirstName: "Eneas",
		LastName:  "Infanger",
	}

	timStudent := Student{
		FirstName: "Tim",
		LastName:  "Hüsler",
	}

	aaronStudent := Student{
		FirstName: "Aaron",
		LastName:  "Ettlin",
	}

	philStudent := Student{
		FirstName: "Phil",
		LastName:  "Hägeli",
	}

	arunStudent := Student{
		FirstName: "Arun",
		LastName:  "Aenishänslin",
	}

	class1 := Class{
		Students: []Student{
			eneasStudent,
			timStudent,
			aaronStudent,
		},
	}

	class2 := Class{
		Students: []Student{
			aaronStudent,
			philStudent,
			arunStudent,
		},
	}

	class3 := Class{
		Students: []Student{
			arunStudent,
			eneasStudent,
			timStudent,
		},
	}

	modules[123] = []Class{class1, class3}
	modules[346] = []Class{class1, class2, class3}
	modules[114] = []Class{class2}

	fmt.Println(modules)
}
