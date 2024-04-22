package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Todo contains a todo's info
type Todo struct {
	Text string `json:"text"`
}

// Display displays struct Todo
func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

// Save content to the file with title as file name
func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// New is a contructur for Todo struct
func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Empty string. Invalid input")
	}

	return Todo{
		Text: content,
	}, nil
}
