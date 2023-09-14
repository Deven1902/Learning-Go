package main

import "fmt"

func main() {
	fmt.Println("Learning pointers in Go")

	// pointers are used to store the address of a variable.
	// pointer act as references to our main variables. so that we can access the main variable from the pointer.

	// syntax: var pointer_name *data_type = &variable_name

	var ptr *int // declaring a pointer

	fmt.Println("Value of ptr is: ", ptr) // value of ptr is nil
	// created a pointer but not pointing to any variable, so it's value is nil.

	num := 10
	var ptr2 *int = &num // creating a pointer and pointing to num variable
						 // & = reference operator. it is used to get the address of a variable.
						 // * = dereference operator. it is used to get the value of a variable from the address.

	fmt.Println("Value of ptr2 is: ", ptr2) // value of ptr2 is the address of num variable
	fmt.Println("Value of num is: ", *ptr2) // value of num is 10

	*ptr2 = *ptr2 + 10 // changing the value of num variable using pointer
	fmt.Println("guess what the value is after update", num) 

	// here the value of num is changed using the pointer. here the operation is performed directly on the variable through the pointer.
}
