package controller

import (
	"fmt"
	"github.com/mgxian/todolist/repository"
	"reflect"
	"testing"
)

var path = "test-todolist.txt"

func TestAddTodoItem(t *testing.T) {
	repository.DeleteFile(path)
	store := repository.NewStore(path)
	sut := NewTodoController(store)

	item := "a todo item"
	sut.Add(item)

	size := store.GetTodo().Size()
	assertEqual(t, size, 1)

	got := store.GetTodo().Items()[0]
	fmt.Println(got)
	assertEqual(t, got.Name(), item)
}

func TestDoneTodoItem(t *testing.T) {
	repository.DeleteFile(path)
	store := repository.NewStore(path)
	sut := NewTodoController(store)
	item := "a todo item"
	sut.Add(item)
	anotherItem := "another todo item"
	sut.Add(anotherItem)
	sut.Done(1)

	size := store.GetTodo().Size()
	assertEqual(t, size, 2)

	got := store.GetTodo().DoneItems()[0]
	assertEqual(t, got.Name(), item)
}

func TestListTodoItems(t *testing.T) {
	repository.DeleteFile(path)
	store := repository.NewStore(path)
	sut := NewTodoController(store)
	item := "a todo item"
	sut.Add(item)
	anotherItem := "another todo item"
	sut.Add(anotherItem)

	items := sut.TodoItems()

	assertEqual(t, len(items), 2)

	want := items[0]
	assertEqual(t, want.Name(), item)
	assertEqual(t, want.IsDone(), false)

	want = items[1]
	assertEqual(t, want.Name(), anotherItem)
	assertEqual(t, want.IsDone(), false)
}

func TestListNotDoneTodoItems(t *testing.T) {
	repository.DeleteFile(path)
	store := repository.NewStore(path)
	sut := NewTodoController(store)
	item := "a todo item"
	sut.Add(item)
	anotherItem := "another todo item"
	sut.Add(anotherItem)
	sut.Done(1)

	items := sut.NotDoneTodoItems()
	assertEqual(t, len(items), 1)

	want := items[0]
	assertEqual(t, want.Name(), anotherItem)
	assertEqual(t, want.IsDone(), false)
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
