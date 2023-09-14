package main

import "fmt"

func main() {
	fmt.Println("Learning loops in Go")

	// synatx of for loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	// for loop without any condition will run forever, until you break it.

	for {
		fmt.Println("loop")
		break
	}
}
