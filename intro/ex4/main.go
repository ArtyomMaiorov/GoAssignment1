package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func swap(x, y string) (string, string) {
	return y, x
}
func quotientRemainder(x int, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
func main() {
	fmt.Println("add result: ", add(1, 2))
	a, b := swap("benedict", "cumberbatch")
	fmt.Println("swap result: ", a, b)
	quotient, remainder := quotientRemainder(4, 2)
	fmt.Println("quotient and remainder: ", quotient, remainder)
}

// 1. by specifying multiple return types and listing them in the return statement of the function
// 2. it makes code easier to read, they are initialized automatically and make return statements simpler
// 3. you can use underscore "_"