package domain

import "strings"

type TodoItem struct {
	name   string
	status int
}

func (item TodoItem) Name() string {
	return item.name
}

const (
	Done = 1
)

func (item *TodoItem) Done() {
	item.status = Done
}

func (item TodoItem) IsDone() bool {
	return item.status == Done
}

func (item TodoItem) Dump() string {
	if item.IsDone() {
		return "true " + item.Name()
	}
	return "false " + item.Name()
}

func NewTodoItem(name string) *TodoItem {
	return &TodoItem{name: name}
}

func NewTodoItemFromText(text string) (item *TodoItem) {
	data := strings.SplitN(text, " ", 2)
	if len(data) != 2 {
		panic("can not create todo item from text: " + text)
	}
	item = NewTodoItem(data[1])
	if data[0] == "true" {
		item.Done()
	}
	return
}
