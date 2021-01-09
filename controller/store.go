package controller

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

type Store struct {
	path string
}

func NewStore(path string) *Store {
	if !IsFileExist(path) {
		CreateFile(path)
	}
	return &Store{path: path}
}

func (s *Store) GetTodoItems() []string {
	contents, err := ioutil.ReadFile(s.path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(contents), "\n")
	return lines[:len(lines)-1]
}

func (s *Store) SaveTodoItems(items ...string) {
	ClearFile(s.path)
	f, err := os.OpenFile(s.path, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, item := range items {
		line := fmt.Sprintln(item)
		_, err := f.WriteString(line)
		if err != nil {
			panic(err)
		}
	}
}
