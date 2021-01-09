package controller

import (
	"testing"
)

var path = "test-todolist.txt"

func TestAddTodoItem(t *testing.T) {
	DeleteFile(path)
	store := NewStore(path)
	sut := NewTodoController(store)

	item := "a todo item"
	sut.Add(item)

	items := store.GetTodoItems()
	if len(items) != 1 {
		t.Errorf("got %d,want %d", len(items), 1)
	}

	got := items[0]
	want := "false " + item
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDoneTodoItem(t *testing.T) {
	DeleteFile(path)
	store := NewStore(path)
	sut := NewTodoController(store)
	item := "a todo item"
	sut.Add(item)
	anotherItem := "another todo item"
	sut.Add(anotherItem)
	sut.Done(1)

	items := store.GetTodoItems()
	if len(items) != 2 {
		t.Errorf("got %d,want %d", len(items), 1)
	}

	got := items[1]
	want := "false " + anotherItem
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestListTodoItems(t *testing.T) {
	DeleteFile(path)
	store := NewStore(path)
	sut := NewTodoController(store)
	item := "a todo item"
	sut.Add(item)
	anotherItem := "another todo item"
	sut.Add(anotherItem)

	items := sut.TodoItems()
	if len(items) != 2 {
		t.Errorf("got %d,want %d", len(items), 2)
	}

	if items[0] != item {
		t.Errorf("got %q, want %q", items[0], item)
	}

	if items[1] != anotherItem {
		t.Errorf("got %q, want %q", items[1], anotherItem)
	}
}

func TestListNotDoneTodoItems(t *testing.T) {
	DeleteFile(path)
	store := NewStore(path)
	sut := NewTodoController(store)
	item := "a todo item"
	sut.Add(item)
	anotherItem := "another todo item"
	sut.Add(anotherItem)
	sut.Done(1)

	items := sut.NotDoneTodoItems()
	if len(items) != 1 {
		t.Errorf("got %d,want %d", len(items), 1)
	}

	if items[0] != anotherItem {
		t.Errorf("got %q, want %q", items[0], anotherItem)
	}
}
