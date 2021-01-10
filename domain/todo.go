package domain

import (
	"fmt"
	"strings"
)

type Todo struct {
	items []*TodoItem
}

func (t *Todo) Add(item *TodoItem) {
	t.items = append(t.items, item)
}

func (t *Todo) NotDoneItems() (result []TodoItem) {
	for _, item := range t.items {
		if item.IsDone() {
			continue
		}
		result = append(result, *item)
	}
	return
}

func (t *Todo) DoneItems() (result []TodoItem) {
	for _, item := range t.items {
		if item.IsDone() {
			result = append(result, *item)
		}
	}
	return
}

func (t *Todo) Done(i int) {
	t.items[i-1].Done()
}

func (t *Todo) Size() int {
	return len(t.items)
}

func (t *Todo) Dump() string {
	result := ""
	for _, item := range t.items {
		result += fmt.Sprintf("%s\n", item.Dump())
	}
	return result
}

func (t *Todo) Items() []TodoItem {
	result := make([]TodoItem, 0)
	for _, item := range t.items {
		result = append(result, *item)
	}
	return result
}

func NewTodo() *Todo {
	return &Todo{}
}

func NewTodoFromText(text string) *Todo {
	itemTexts := strings.Split(text, "\n")
	todo := NewTodo()
	for _, itemText := range itemTexts {
		if itemText != "" {
			todoItem := NewTodoItemFromText(itemText)
			todo.Add(todoItem)
		}
	}
	return todo
}
