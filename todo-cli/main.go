package main

func main() {
	manager := TodoManager{}

	manager.Add("Buy groceries")
	manager.Add("Learn Go")
	manager.Add("Read a book")

	manager.List()
}
