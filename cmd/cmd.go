package cmd

import (
	"fmt"
	"github.com/mgxian/todolist/controller"
	"github.com/mgxian/todolist/domain"
	"io"
	"strconv"
)

type Cmd struct {
	todoController *controller.TodoController
}

func (c *Cmd) Execute(args []string, w io.Writer) {
	cmdString := args[0]
	switch cmdString {
	case "add":
		item := args[1]
		index := c.todoController.Add(item)
		fmt.Fprintf(w, "1. %s\nItem %d added", item, index)
	case "done":
		if index, err := strconv.Atoi(args[1]); err == nil {
			fmt.Fprintf(w, "Item %d done", c.todoController.Done(index))
		}
	case "list":
		all := len(args) > 1 && args[1] == "--all"
		c.list(w, all)
	default:
		fmt.Fprintf(w, "unknown command: %s\n", cmdString)
	}
}

func (c *Cmd) list(w io.Writer, all bool) {
	var items []domain.TodoItem
	if all {
		items = c.todoController.TodoItems()
	} else {
		items = c.todoController.NotDoneTodoItems()
	}

	for i, item := range items {
		if item.IsDone() {
			fmt.Fprintf(w, "%d. [Done] %s\n", i+1, item.Name())
		} else {
			fmt.Fprintf(w, "%d. %s\n", i+1, item.Name())
		}
	}
	fmt.Fprintf(w, "\nTotal: %d items", len(items))
}

func NewCmd(todoController *controller.TodoController) *Cmd {
	return &Cmd{todoController: todoController}
}
