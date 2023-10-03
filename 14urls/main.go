package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.google.com"

func main() {
	fmt.Println("handling urls in Go")

	fmt.Println("url: ", myurl)

	// parsing an url

	res, _ := url.Parse(myurl)

	fmt.Println("Scheme: ", res.Scheme)
	fmt.Println("Host: ", res.Host)
	fmt.Println("Path: ", res.Path)
	fmt.Println("RawQuery: ", res.RawQuery)
	fmt.Println("Port: ", res.Port())
}
