package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		printUsage()
		return
	}

	switch args[1] {
	case "list":
		fmt.Println("Listing todos...")
	case "add":
		if len(args) < 3 {
			fmt.Println("titleがありません")
			return
		}
		title := args[2]
		fmt.Printf("Adding todo: %s\n", title)
	default:
		fmt.Printf("Unknown command: %s\n\n", args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  go run . list")
	fmt.Println("  go run . add <title>")
}
