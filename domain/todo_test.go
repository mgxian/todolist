package domain

import (
	"testing"
)

func TestCreateTodo(t *testing.T) {
	todo := NewTodo()
	if todo == nil {
		t.Errorf("expected todo not be nil")
	}
}

func TestNewTodoIsEmpty(t *testing.T) {
	todo := NewTodo()
	if todo.IsEmpty() != true {
		t.Errorf("expected new todo to be empty")
	}
}

func TestTodoAddTodoItem(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	if todo.IsEmpty() {
		t.Errorf("expected todo not be empty")
	}
}

func TestListNotDoneTodoItems(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	anotherItem := NewTodoItem("need to be done item")
	anotherItem.Done()
	todo.Add(anotherItem)
	items := todo.NotDoneItems()
	if len(items) != 1 {
		t.Errorf("got %d, want 1", len(items))
	}
}

func TestListDoneTodoItems(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	anotherItem := NewTodoItem("need to be done item")
	anotherItem.Done()
	todo.Add(anotherItem)
	items := todo.DoneItems()
	if len(items) != 1 {
		t.Errorf("got %d, want 1", len(items))
	}
}

func TestTodoDoneItem(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	anotherItem := NewTodoItem("need to be done item")
	todo.Add(anotherItem)

	todo.Done(2)
	items := todo.DoneItems()
	assertEqual(t, len(items), 1)
}

func TestNewTodoFromText(t *testing.T) {
	text := "false a new todo item\ntrue a done item\n"
	todo := NewTodoFromText(text)
	assertEqual(t, todo.Size(), 2)
}
