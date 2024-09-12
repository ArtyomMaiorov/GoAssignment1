package main

import "fmt"

func main() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// you can write the for loop to be a while loop by removing the post statement
