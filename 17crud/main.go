package main

import (
	"fmt"
)

type Todo struct {
	ID        int
	Task      string
	Completed bool
}

var todos [10]Todo

func main() {
	// Create a new todo
	todo := Todo{
		ID:        1,
		Task:      "Finish this tutorial",
		Completed: false,
	}

	todos[1] = Todo{
		ID:        2,
		Task:      "watch movies",
		Completed: true,
	}

	todos[2] = Todo{
		ID:        3,
		Task:      "workout",
		Completed: true,
	}

	todos[3] = Todo{
		ID:        4,
		Task:      "abcd",
		Completed: false,
	}

	todos[4] = Todo{
		ID:        5,
		Task:      "efgh",
		Completed: true,
	}

	todos[0] = todo // Add the todo to the array

	todoByID := &todos[0] // Get a specific todo by ID

	todoByID.Completed = true // Update a todo

	todos[0] = Todo{} // Delete a todo

	fmt.Println(todo) // Print the todo

	for _, todo := range todos {
		fmt.Println(todo)
	}
}
