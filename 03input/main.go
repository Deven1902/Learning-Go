package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a = "Hello World"

	fmt.Println(a)
	fmt.Println("Enter any number:- ")
	var reader = bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("You entered: ", input)
}
