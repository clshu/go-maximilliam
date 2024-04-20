package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println("Age:", *agePointer)
	fmt.Println("Age Pinter:", agePointer)
	adultYear := getAdultYear(age)
	fmt.Println("Adult Year:", adultYear)
}

func getAdultYear(age int) int {
	return age - 18
}
