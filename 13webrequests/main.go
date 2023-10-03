package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {

	fmt.Printf("This is how to use web requests in Go")
	responce, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Respnse type: %T\n", responce)

	defer responce.Body.Close() // this is a caller's responisbility to close the connection.

	databytes, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content: ", content)
}
