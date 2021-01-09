package domain

import (
	"reflect"
	"testing"
)

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

func TestDumpTodoItem(t *testing.T) {
	testCases := []struct {
		name string
		done bool
		want string
	}{
		{"a new todo item", false, "false a new todo item"},
		{"a new todo item", true, "true a new todo item"},
	}
	for _, tt := range testCases {
		name := tt.name
		item := NewTodoItem(name)
		if tt.done {
			item.Done()
		}
		got := item.Dump()
		want := tt.want
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
}

func TestCreateTodoItemFromText(t *testing.T) {
	testCases := []struct {
		text string
		done bool
		name string
	}{
		{"false a new todo item", false, "a new todo item"},
		{"true a done todo item", true, "a done todo item"},
	}
	for _, tt := range testCases {
		item := NewTodoItemFromText(tt.text)
		assertEqual(t, item.Name(), tt.name)
		assertEqual(t, item.IsDone(), tt.done)
	}
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
