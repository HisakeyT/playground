package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		pringUsage()
		return
	}

	manager := TodoManager{}
	if err := manager.Load(); err != nil {
		fmt.Println("Error loading todos:", err)
		return
	}

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
	case "done":
		if len(args) < 3 {
			fmt.Println("Error: Missing ID for the todo to mark as done")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			return
		}

		if err := manager.MarkDone(id); err != nil {
			fmt.Println("Error marking todo as done:", err)
		}
	case "delete":
		if len(args) < 3 {
			fmt.Println("Error: Missing ID for the todo to delete")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Error: ID must be a number")
			return
		}

		if err := manager.Delete(id); err != nil {
			fmt.Println("Error deleting todo:", err)
		}

		if err := manager.Save(); err != nil {
			fmt.Println("Error saving todos:", err)
		}
	default:
		pringUsage()
		return
	}
}

func pringUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add <title>   - Add a new todo")
	fmt.Println("  list          - List all todos")
	fmt.Println("  done <id>     - Mark a todo as done")
	fmt.Println("  delete <id>   - Delete a todo")
}
