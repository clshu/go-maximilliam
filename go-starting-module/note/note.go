package note

import (
	"errors"
	"time"
)

// Note contains a note's info
type Note struct {
	title     string
	content   string
	createdAt time.Time
}

// New is a contructur for Note struct
func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Empty string. Invalid input")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
