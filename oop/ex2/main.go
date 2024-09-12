package main

import "fmt"

type Employee struct {
	Name string
	ID string
}

type Manager struct{
	Employee
	Department string
}
func (e Employee) Work() string {
	return e.Name + " " + e.ID
}

func main() {
	m1 := Manager{
		Employee: Employee{
			Name: "Will",
			ID: "123098142",
		},
		Department: "FBI",
	}
	fmt.Println(m1.Work())
}

// 1. embedding allows a struct to have another struct in it which is similar to composition in traditional OOP. it is useful to reduce overload on a struct
// 2. Go allows you to use methods of embedded types on the outer structs
// 3. no
