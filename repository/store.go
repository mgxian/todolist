package repository

import (
	"github.com/mgxian/todolist/domain"
	"io/ioutil"
	"os"
)

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func CreateFile(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func ClearFile(path string) {
	if err := os.Truncate(path, 0); err != nil {
		panic(err)
	}
}

func DeleteFile(path string) {
	if !IsFileExist(path) {
		return
	}

	if err := os.Remove(path); err != nil {
		panic(err)
	}
}
type GetTodo interface {
	GetTodo() *domain.Todo
}

type SaveTodo interface {
	SaveTodo(*domain.Todo)
}

type IStore interface {
	GetTodo
	SaveTodo
}
type Store struct {
	path string
}

func NewStore(path string) *Store {
	if !IsFileExist(path) {
		CreateFile(path)
	}
	return &Store{path: path}
}

func (s *Store) GetTodo() *domain.Todo {
	contents, err := ioutil.ReadFile(s.path)
	if err != nil {
		panic(err)
	}
	return domain.NewTodoFromText(string(contents))
}

func (s *Store) SaveTodo(todo *domain.Todo) {
	ClearFile(s.path)
	f, err := os.OpenFile(s.path, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(todo.Dump())
}
