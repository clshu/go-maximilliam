package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("\nSaving the node failed")
		return
	}

	fmt.Println("\nSaving the node succeeded")
}

func getUserINput(prompt string) string {
	fmt.Printf("%v ", prompt)
	// Handle longer text input
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimRight(text, "\n\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserINput("Note title:")

	content := getUserINput("Note content:")

	return title, content

}
