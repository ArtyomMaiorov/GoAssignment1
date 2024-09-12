package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("Positive")
	} else if number < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}

// syntax is different and in Go you can initialize the variable in the condition
