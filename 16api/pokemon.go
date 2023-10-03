package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const myurl string = "http://pokeapi.co/api/v2/pokedex/kanto/"

func main() {
	fmt.Println("Pokedex API calls")

	res, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	resData, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(resData))
}
