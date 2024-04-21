package main

import (
	"fmt"

	"example.com/go-project/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(userNote)
}

func getUserINput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}

func getNoteData() (string, string) {
	title := getUserINput("Note title: ")

	content := getUserINput("Note content: ")

	return title, content

}
