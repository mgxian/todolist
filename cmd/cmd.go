package cmd

import (
	"github.com/mgxian/todolist/controller"
	"io"
	"strconv"
)

type Cmd struct {
	todoController *controller.TodoController
}

func (c *Cmd) Execute(args []string, w io.Writer) (result []string) {
	cmdString := args[0]
	switch cmdString {
	case "add":
		item := args[1]
		c.todoController.Add(item)
	case "done":
		if index, err := strconv.Atoi(args[1]); err == nil {
			c.todoController.Done(index)
		}
	case "list":
		if len(args) > 1 && args[1] == "--all" {
			c.todoController.TodoItems()
		}
		result = c.todoController.NotDoneTodoItems()
	}
	return
}

func NewCmd(todoController *controller.TodoController) *Cmd {
	return &Cmd{todoController: todoController}
}
