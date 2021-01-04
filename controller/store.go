package controller

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Store struct {
	path string
}

func NewStore(path string) *Store {
	os.OpenFile(path,os.O_CREATE,0644)
	return &Store{path:path}
}

func (s *Store) GetTodoItems() []string {
	contents, err := ioutil.ReadFile(s.path)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(contents), "\n")
}

func (s *Store) SaveTodoItems(items ...string) {
	f, err := os.OpenFile(s.path, os.O_APPEND|os.O_WRONLY, 0600)
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
