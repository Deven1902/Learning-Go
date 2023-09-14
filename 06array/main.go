package main

import "fmt"

func main() {
	fmt.Println("Learning arrays in Go")

	var num [5]int // declaring an array of size 5
	// how did the values get assigned to the array? when i didn't assign any value to the array.
	fmt.Println("Value of num is: ", num) // value of num is [0 0 0 0 0]

	var list [6]string

	list[0] = "Hello"
	list[1] = "World"
	list[2] = "How"
	// list [3]	= "are"
	list[4] = "you"
	fmt.Println("Value of list is: ", list) // value of list is [    ]

	fmt.Println("lenght of num is: ", len(num))   // lenght of num is:  5
	fmt.Println("lenght of list is: ", len(list)) // lenght of list is:  5

	// array is a fixed size data structure. once the size is defined, it cannot be changed. 
}
