package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome the app")
	fmt.Println("Pls rate our app between 1 to 5")

	read := bufio.NewReader(os.Stdin)

	input, _ := read.ReadString('\n')

	fmt.Println("thanks for ", input, "rating!!")

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated rating is", num+1)
	}
}
