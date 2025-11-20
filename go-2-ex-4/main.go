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
    schueler1 := Student{
        FirstName: "Lukas",
        LastName:  "Traut",
    }
    schueler2 := Student{
        FirstName: "Rudi",
        LastName:  "Macher",
    }

    INA24bL := Class{
        Name:     "INA24bL",
        Students: []Student{schueler1, schueler2},
    }
    INA24aL := Class{
        Name:     "INA24aL",
        Students: []Student{schueler2},
    }

    modules := map[int][]Class{
        114: {INA24bL, INA24aL},
        999: {INA24aL},
        346: {INA24bL},
    }

    fmt.Println("Schüler in der Klasse INA24bL:", INA24bL.Students)
    fmt.Println("Schüler in der Klasse INA24aL:", INA24aL.Students)
    fmt.Println("Module:", modules)
}