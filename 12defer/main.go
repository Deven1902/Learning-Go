package main

import "fmt"

func main() {

	// defer is used to ensure that a function call is performed later in a program's execution,
	// usually for purposes of cleanup.

	// defer is used to delay the execution of a statement until function exits.
	fmt.Println("Hello")

	defer fmt.Println("World")
	myDefer()
}

func myDefer() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}
