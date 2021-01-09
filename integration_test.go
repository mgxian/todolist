package todolist

import (
	cmd2 "github.com/mgxian/todolist/cmd"
	"github.com/mgxian/todolist/controller"
	"github.com/mgxian/todolist/repository"
	"strings"
	"testing"
)

var path = "todolist.txt"

func TestAddTodoItem(t *testing.T) {
	repository.DeleteFile(path)
	item := "a todo item"
	store := repository.NewStore(path)
	aController := controller.NewTodoController(store)
	var result strings.Builder
	cmd := cmd2.NewCmd(aController)
	cmdString := "add"
	args := []string{cmdString, item}
	cmd.Execute(args, &result)

	items := aController.TodoItems()
	if len(items) != 1 {
		t.Errorf("got %d,want %d", len(items), 1)
	}
}

func TestDoneTodoItem(t *testing.T) {
	repository.DeleteFile(path)
	item := "a todo item"
	store := repository.NewStore(path)
	aController := controller.NewTodoController(store)
	var result strings.Builder
	cmd := cmd2.NewCmd(aController)
	cmdString := "add"
	args := []string{cmdString, item}
	cmd.Execute(args, &result)

	cmdString = "done"
	args = []string{cmdString, "1"}
	cmd.Execute(args, &result)

	items := aController.NotDoneTodoItems()
	if len(items) != 0 {
		t.Errorf("got %d,want %d", len(items), 0)
	}
}

func TestListTodoItems(t *testing.T) {
	repository.DeleteFile(path)
	item := "a todo item"
	store := repository.NewStore(path)
	aController := controller.NewTodoController(store)
	aController.Add(item)
	anotherItem := "a another item"
	aController.Add(anotherItem)
	aController.Done(2)
	var result strings.Builder
	cmd := cmd2.NewCmd(aController)
	cmdString := "list"
	args := []string{cmdString}
	cmd.Execute(args, &result)
}
