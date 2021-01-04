package controller

import "strings"

type TodoController struct {
	store *Store
}

func (t *TodoController) Add(s string) {
	s = "false " + s
	t.store.SaveTodoItems(s)
}

func (t *TodoController) Done(i int) {
	items :=t.store.GetTodoItems()
	needDoneItem := items[i-1]
	item:=strings.SplitN(needDoneItem," ",2)[1]
	items[i-1] = "true "+item
	t.store.SaveTodoItems(items...)
}

func (t *TodoController) NotDoneTodoItems() (result []string) {
	items:=t.store.GetTodoItems()
	for _, item := range items {
		data:= strings.SplitN(item," ",2)
		isDone:=data[0]=="true"
		if isDone {
			continue
		}
		result = append(result,data[1])
	}
	return
}

func NewTodoController(store *Store) *TodoController {
	return &TodoController{store: store}
}
