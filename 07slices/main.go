package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learning slices in Golang")
	// slices are used to store a sequence of elements of the same type. they are similar to arrays but they are dynamic in nature.

	// soemthing similar to the vectors in c++.

	var num []int // creating a slice of size 5 and capacity 10

	fmt.Printf("Type of num is %T: ", num)

	// append method of slices.
	num = append(num, 1, 2, 3) 
	fmt.Println("Printing num", num)

	// slice operator [start:end].  
	// start is inclusive and end is exclusive. 

	// syntax: var slice_name []data_type = make([]data_type, size, capacity).

	var list []int = make([]int, 5, 10) // creating a slice of size 5 and capacity 10
	fmt.Println(list) 

	sort.Ints(num) // sorting the slice
	
	// lets remove values from slice

	
}
