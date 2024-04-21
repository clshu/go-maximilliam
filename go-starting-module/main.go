package main

import (
	"errors"
	"fmt"
)

func main() {

	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}

	result := fmt.Sprintf("title: %s\ncontent: %s\n", title, content)
	fmt.Print(result)
}

func getUserINput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Empty string. Invalid input")
	}

	return value, nil
}

func getNoteData() (string, string, error) {
	title, err := getUserINput("Note title: ")
	if err != nil {
		return "", "", err
	}
	content, err := getUserINput("Note content: ")
	if err != nil {
		return "", "", err
	}

	return title, content, nil

}
