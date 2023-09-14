package main

import "fmt"

func main() {
	fmt.Println("Learning functions in Go")

	res := add(10, 20)
	fmt.Println(res)

	res2 := product(10, 20)
	fmt.Println(res2)
}

	// function is a block of code which is used to perform a specific task.
	// function is a reusable code. we can call a function multiple times.

	// in go we should explicitly return values in a function compulsorily. 

	// syntax of a function
func add(a int, b int) int {
	return a + b;
}

func product(a int, b int) int {
	return a * b;
}