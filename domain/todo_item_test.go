package domain

import "testing"

func TestCreateTodoItem(t *testing.T) {
	item := NewTodoItem("a new todo item")
	got := item.Name()
	want := "a new todo item"
	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func TestDoneTodoItem(t *testing.T) {
	item := NewTodoItem("a new todo item")
	item.Done()
	if !item.IsDone() {
		t.Errorf("todo item should be done, but not")
	}
}
