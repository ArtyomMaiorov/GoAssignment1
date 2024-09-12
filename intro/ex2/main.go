package main

import "fmt"

func main() {
	var a int = 8
	var b bool = true
	c := 2.79
	d := "Str"
	fmt.Printf("a: %d, b: %t, c: %f, d: %s\n", a, b, c, d)
	fmt.Printf("Type of a: %T, Type of b: %T, Type of c: %T, Type of d: %T\n", a, b, c, d)
}

// := has implicit typing inside the function
// by using string formatting %T
// no, because Go is a statically typed language
