package main

import "fmt"

type Person struct {
	Name string
	Age int
}
func (p Person) Greet() string {
	return "Hello, " + p.Name
}

func main() {
	p1 := Person{
		Name: "Artyom",
		Age: 22,
	}
	fmt.Println(p1.Greet())
}
// 1. by writing type <StructName> struct{}
// 2. methods can be called on struct instances and can access structs' fields
// 3. yes, for example with data types