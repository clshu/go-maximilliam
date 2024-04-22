package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// Saver define save() to save struct info
type saver interface {
	Save() error
}
type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("Gen")

	todoText := getUserInput("Todo: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) {
	fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return data.Save()

}
func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	return nil
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
