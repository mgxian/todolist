package controller

import (
	"github.com/mgxian/todolist/domain"
	"github.com/mgxian/todolist/repository"
)

type TodoController struct {
	store *repository.Store
}

func (t *TodoController) Add(name string) int {
	todo := t.store.GetTodo()
	item := domain.NewTodoItem(name)
	todo.Add(item)
	t.store.SaveTodo(todo)
	return todo.Size()
}

func (t *TodoController) find(items []domain.TodoItem, i int) int {
	found := -1
	count := 0
	for index, item := range items {
		if item.IsDone() {
			continue
		}
		count++
		if count == i {
			found = index
		}
	}
	return found + 1
}

func (t *TodoController) Done(i int) int {
	todo := t.store.GetTodo()
	found := t.find(todo.Items(), i)
	if found == -1 {
		return -1
	}
	todo.Done(found)
	t.store.SaveTodo(todo)
	return found
}

func (t *TodoController) NotDoneTodoItems() []domain.TodoItem {
	todo := t.store.GetTodo()
	return todo.NotDoneItems()
}

func (t *TodoController) TodoItems() []domain.TodoItem {
	todo := t.store.GetTodo()
	return todo.Items()
}

func NewTodoController(store *repository.Store) *TodoController {
	return &TodoController{store: store}
}
