package model

type Todo struct {
	items []*TodoItem
}

func (t *Todo) IsEmpty() bool {
	return len(t.items) == 0
}

func (t *Todo) Add(item *TodoItem) {
	t.items = append(t.items, item)
}

func (t *Todo) NotDoneItems() (result []*TodoItem) {
	for _, item := range t.items {
		if item.IsDone() {
			continue
		}
		result = append(result, item)
	}
	return
}

func (t *Todo) DoneItems() (result []*TodoItem) {
	for _, item := range t.items {
		if item.IsDone() {
			result = append(result, item)
		}
	}
	return
}

func (t *Todo) Done(i int) {
	items := t.NotDoneItems()
	if len(items) < i {
		return
	}
	items[i-1].Done()
}

func NewTodo() *Todo {
	return &Todo{}
}
