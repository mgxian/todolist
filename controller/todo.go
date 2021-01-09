package controller

import (
	"github.com/mgxian/todolist/repository"
	"strings"
)

func IsItemDone(item string) bool {
	return strings.SplitN(item, " ", 2)[0] == "true"
}

type TodoController struct {
	store *repository.Store
}

func (t *TodoController) Add(s string) {
	items := t.store.GetTodoItems()
	s = "false " + s
	items = append(items, s)
	t.store.SaveTodoItems(items...)
}

func (t *TodoController) Done(i int) {
	items := t.store.GetTodoItems()
	needDoneItemIndex := -1
	count := 0
	for index, item := range items {
		if IsItemDone(item) {
			continue
		}
		count++
		if count == i {
			needDoneItemIndex = index
		}
	}
	item := strings.SplitN(items[needDoneItemIndex], " ", 2)[1]
	items[i-1] = "true " + item
	t.store.SaveTodoItems(items...)
}

func (t *TodoController) NotDoneTodoItems() (result []string) {
	items := t.store.GetTodoItems()
	for _, item := range items {
		data := strings.SplitN(item, " ", 2)
		isDone := data[0] == "true"
		if isDone {
			continue
		}
		result = append(result, item)
	}
	return
}

func (t *TodoController) TodoItems() (result []string) {
	items := t.store.GetTodoItems()
	return items
}

func NewTodoController(store *repository.Store) *TodoController {
	return &TodoController{store: store}
}
