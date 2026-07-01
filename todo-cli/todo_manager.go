package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const todoFile = "todos.json"

type TodoManager struct {
	todos []Todo
}

func (tm *TodoManager) Add(title string) {
	todo := Todo{
		ID:    tm.nextID(),
		Title: title,
	}
	tm.todos = append(tm.todos, todo)
}

func (tm *TodoManager) Save() error {
	data, err := json.MarshalIndent(tm.todos, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(todoFile, data, 0644); err != nil {
		return err
	}

	return nil
}

func (tm TodoManager) List() {
	for _, todo := range tm.todos {
		fmt.Println(todo.display())
	}
}

func (tm *TodoManager) Load() error {
	existingData, err := os.ReadFile(todoFile)

	if os.IsNotExist(err) {
		return nil // If the file doesn't exist, we can just return without error
	}

	if err != nil {
		return err
	}

	if err := json.Unmarshal(existingData, &tm.todos); err != nil {
		return err
	}

	return nil
}

func (tm *TodoManager) MarkDone(id int) error {
	todo, err := tm.findByID(id)
	if err != nil {
		return err
	}

	todo.Done = true
	// Save the updated todos to the file
	return tm.Save()
}

func (tm *TodoManager) Delete(id int) error {
	for i, todo := range tm.todos {
		if todo.ID == id {
			tm.todos = append(tm.todos[:i], tm.todos[i+1:]...)
			return tm.Save() // Save the updated todos to the file
		}
	}
	return fmt.Errorf("Todo with ID %d not found", id)
}

func (t Todo) display() string {
	return fmt.Sprintf("ID: %d, Title: %s, Done: %t", t.ID, t.Title, t.Done)
}

func (tm *TodoManager) findByID(id int) (*Todo, error) {
	for i, todo := range tm.todos {
		if todo.ID == id {
			return &tm.todos[i], nil
		}
	}
	return nil, fmt.Errorf("Todo with ID %d not found", id)
}

func (tm *TodoManager) nextID() int {
	maxID := 0
	for _, todo := range tm.todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}
	return maxID + 1
}
