package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		pringUsage()
		return
	}

	manager := TodoManager{}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Error: Missing title for the todo")
			return
		}

		manager.Add(args[2])
		if err := manager.Save(); err != nil {
			fmt.Println("Error saving todos:", err)
		}
	case "list":
		manager.List()
	default:
		pringUsage()
		return
	}
}

func pringUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add <title>   - Add a new todo")
	fmt.Println("  list          - List all todos")
}
