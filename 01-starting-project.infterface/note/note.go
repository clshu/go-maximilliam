package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Note contains a note's info
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"context"`
	CreatedAt time.Time `json:"created_at"`
}

// Display displays struct Note
func (note Note) Display() {
	fmt.Printf("Your note titleed %v has the following content\n\n%v\n", note.Title, note.Content)
}

// Save content to the file with title as file name
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)
	fileName = fileName + ".json"

	json, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

// New is a contructur for Note struct
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Empty string. Invalid input")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
