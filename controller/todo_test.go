package controller

import (
	"testing"
)

func TestAddTodoItem(t *testing.T) {
	store := NewStore("./todo.txt")
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


func TestListNotDoneTodoItems(t *testing.T) {
	store := NewStore("./todo.txt")
	sut := NewTodoController(store)
	item := "a todo item"
	sut.Add(item)
	anotherItem := "another todo item"
	sut.Add(anotherItem)
	sut.Done(1)

	items:=sut.NotDoneTodoItems()
	if len(items) != 1 {
		t.Errorf("got %d,want %d",len(items),1)
	}
}