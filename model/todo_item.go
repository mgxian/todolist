package model

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

func NewTodoItem(name string) *TodoItem {
	return &TodoItem{name: name}
}
