package main

import (
	"github.com/mgxian/todolist/cmd"
	"github.com/mgxian/todolist/controller"
	"github.com/mgxian/todolist/repository"
	"os"
)

func main() {
	path := "todolist.txt"
	args := os.Args[1:]
	store := repository.NewStore(path)
	aController := controller.NewTodoController(store)
	aCmd := cmd.NewCmd(aController)
	aCmd.Execute(args, os.Stdout)
}
