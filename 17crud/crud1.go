package main

import (
	"fmt"
)

type Todo struct {
	ID          int
	Task        string
	Description string
	DueDate     string
}

var todos map[int]Todo

func main() {
	// Create a new TODO
	todo := Todo{
		ID:          1,
		Task:        "Finish this tutorial",
		Description: "Write a blog post about it",
		DueDate:     "2023-10-07",
	}

	// Add the TODO to the map
	todos[1] = todo

	// Get all TODOs
	allTodos := todos

	// Get a specific TODO by ID
	todoByID, ok := todos[1]
	if ok {
		fmt.Println(todoByID)
	} else {
		fmt.Println("TODO not found")
	}

	// Update a TODO
	todoByID.Task = "Finish this tutorial and publish it"

	// Delete a TODO
	delete(todos, 1)
}
