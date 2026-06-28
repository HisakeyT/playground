package main

import (
	"fmt"
	// "os"
)

func main() {
	todo1 := Todo{
		ID:    1,
		Title: "Buy groceries",
		Done:  false,
	}
	todo2 := Todo{
		ID:    2,
		Title: "Learn Go",
		Done:  false,
	}
	todo3 := Todo{
		ID:    3,
		Title: "Read a book",
		Done:  true,
	}

	todos := []Todo{}
	todos = append(todos, todo1, todo2, todo3)

	for _, todo := range todos {
		fmt.Printf("ID: %d, Title: %s, Done: %t\n", todo.ID, todo.Title, todo.Done)
	}
	//		args := os.Args
	//		if len(args) < 2 {
	//			printUsage()
	//			return
	//		}
	//
	//		switch args[1] {
	//		case "list":
	//			fmt.Println("Listing todos...")
	//		case "add":
	//			if len(args) < 3 {
	//				fmt.Println("titleがありません")
	//				return
	//			}
	//			title := args[2]
	//			fmt.Printf("Adding todo: %s\n", title)
	//		default:
	//			fmt.Printf("Unknown command: %s\n\n", args[1])
	//			printUsage()
	//		}
	//	}
	//
	//	func printUsage() {
	//		fmt.Println("Usage:")
	//		fmt.Println("  go run . list")
	//		fmt.Println("  go run . add <title>")
}
