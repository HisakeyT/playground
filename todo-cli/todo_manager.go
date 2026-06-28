package main

import (
	"encoding/json"
	"fmt"
)

type TodoManager struct {
	todos []Todo
}

func (tm *TodoManager) Add(title string) {
	todo := Todo{
		ID:    len(tm.todos) + 1,
		Title: title,
	}
	tm.todos = append(tm.todos, todo)
}

func (tm *TodoManager) Save() error {
	data, err := json.Marshal(tm.todos)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}

func (tm TodoManager) List() {
	for _, todo := range tm.todos {
		fmt.Println(todo.display())
	}
}

func (t Todo) display() string {
	return fmt.Sprintf("ID: %d, Title: %s, Done: %t", t.ID, t.Title, t.Done)
}
