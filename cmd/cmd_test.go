package cmd

import (
	"fmt"
	"github.com/mgxian/todolist/controller"
	"github.com/mgxian/todolist/repository"
	"reflect"
	"strings"
	"testing"
)

var path = "todolist.txt"

func TestAddTodoItem(t *testing.T) {
	aController := createTodoController()
	cmd := NewCmd(aController)
	cmdString := "add"
	item := "a todo item"
	args := []string{cmdString, item}

	var result strings.Builder
	cmd.Execute(args, &result)

	items := aController.TodoItems()
	assertEqual(t, len(items), 1)

	got := result.String()
	want := fmt.Sprintf("1. %s\nItem %d added", item, 1)
	assertEqual(t, got, want)
}

func TestDoneTodoItem(t *testing.T) {
	aController := createTodoController()
	item := "a todo item"
	aController.Add(item)
	anotherItem := "another todo item"
	aController.Add(anotherItem)

	cmd := NewCmd(aController)
	cmdString := "done"
	args := []string{cmdString, "2"}

	var result strings.Builder
	cmd.Execute(args, &result)

	items := aController.NotDoneTodoItems()
	assertEqual(t, len(items), 1)

	got := result.String()
	want := fmt.Sprintf("Item %d done", 2)
	assertEqual(t, got, want)
}

func TestListTodoItems(t *testing.T) {
	aController := createTodoController()
	item := "a todo item"
	aController.Add(item)
	anotherItem := "a another item"
	aController.Add(anotherItem)
	aController.Done(2)

	cmd := NewCmd(aController)
	cmdString := "list"
	args := []string{cmdString}

	var result strings.Builder
	cmd.Execute(args, &result)

	got := result.String()
	want := fmt.Sprintf("1. %s\n\nTotal: %d items", item, 1)
	assertEqual(t, got, want)
}

func TestListAllTodoItems(t *testing.T) {
	aController := createTodoController()
	item := "a todo item"
	aController.Add(item)
	anotherItem := "a another item"
	aController.Add(anotherItem)
	aController.Done(2)

	cmd := NewCmd(aController)
	cmdString := "list"
	args := []string{cmdString, "--all"}

	var result strings.Builder
	cmd.Execute(args, &result)

	got := result.String()
	want := fmt.Sprintf("1. %s\n2. [Done] %s\n\nTotal: %d items", item, anotherItem, 2)
	assertEqual(t, got, want)
}

func TestUnKonCmd(t *testing.T) {
	aController := createTodoController()
	cmd := NewCmd(aController)

	cmdString := "not-exists-cmd"
	args := []string{cmdString}

	var result strings.Builder
	cmd.Execute(args, &result)

	assertEqual(t, result.String(), "unknown command: not-exists-cmd\n")
}

func createTodoController() *controller.TodoController {
	repository.DeleteFile(path)
	store := repository.NewStore(path)
	aController := controller.NewTodoController(store)
	return aController
}

func assertEqual(t *testing.T, actual, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got %v, want %v", actual, expected)
	}
}
