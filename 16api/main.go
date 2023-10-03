// lets learn to make api calls with golang

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "https://example.com/api/users", nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Execute the HTTP request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the response body
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
