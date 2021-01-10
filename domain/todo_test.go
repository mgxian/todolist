package domain

import (
	"fmt"
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
	assertEqual(t, todo.Size(), 0)
}

func TestTodoAddTodoItem(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	assertEqual(t, todo.Size(), 1)
}

func TestListNotDoneTodoItems(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	anotherItem := NewTodoItem("need to be done item")
	anotherItem.Done()
	todo.Add(anotherItem)
	items := todo.NotDoneItems()
	assertEqual(t, len(items), 1)
}

func TestListDoneTodoItems(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	anotherItem := NewTodoItem("need to be done item")
	anotherItem.Done()
	todo.Add(anotherItem)
	items := todo.DoneItems()
	assertEqual(t, len(items), 1)
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

func TestDumpTodo(t *testing.T) {
	todo := NewTodo()
	item := NewTodoItem("a todo item")
	todo.Add(item)
	anotherItem := NewTodoItem("need to be done item")
	todo.Add(anotherItem)
	todo.Done(2)

	got := todo.Dump()
	want := fmt.Sprintf("false %s\ntrue %s\n", item.Name(), anotherItem.Name())
	assertEqual(t, got, want)
}

func TestNewTodoFromText(t *testing.T) {
	text := "false a new todo item\ntrue a done item\n"
	todo := NewTodoFromText(text)
	assertEqual(t, todo.Size(), 2)
}
