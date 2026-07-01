package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	manager := TodoManager{}

	manager.Add("Test Todo")

	todo := manager.todos[0]

	if len(manager.todos) != 1 {
		t.Errorf("got %d, want 1", len(manager.todos))
	}

	if todo.ID != 1 {
		t.Errorf("got ID %d, want 1", todo.ID)
	}

	if todo.Title != "Test Todo" {
		t.Errorf("got title '%s', want 'Test Todo'", todo.Title)
	}

	if todo.Done {
		t.Errorf("got Done to be %t, want false", todo.Done)
	}
}

func TestDelete(t *testing.T) {
	manager := TodoManager{}
	manager.Add("Test Todo")
	manager.Add("Another Todo")

	if err := manager.Delete(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	todo := manager.todos[0]

	if len(manager.todos) != 1 {
		t.Errorf("got %v, want 1", len(manager.todos))
	}

	if todo.ID != 2 {
		t.Errorf("got ID %d, want 2", todo.ID)
	}

	if todo.Title != "Another Todo" {
		t.Errorf("got title '%s', want 'Another Todo'", todo.Title)
	}
}

func TestMarkDone(t *testing.T) {
	manager := TodoManager{}
	manager.Add("Test Todo")

	if err := manager.MarkDone(1); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	todo, err := manager.findByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !todo.Done {
		t.Errorf("got %t, want true", todo.Done)
	}
}
